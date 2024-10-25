package controller

import (
	"errors"
	"net/http"
	"simple-social-app/db"
	"simple-social-app/dto"
	"simple-social-app/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Post struct {
}

func (p Post) Create(ctx *gin.Context) {
	var form dto.PostRequest
	user := ctx.MustGet("user").(dto.UserResponse)

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newPost := model.Post{
		UserId:  user.ID,
		Message: form.Message,
	}

	query := db.Conn.Preload("User").Create(&newPost)

	if err := query.Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"post": dto.PostResponse{
		ID:      newPost.ID,
		Message: newPost.Message,
		UserId:  newPost.UserId,
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Comment: []dto.CommentResponse{},
	}})

}
func (p Post) FindAll(ctx *gin.Context) {
	var posts []model.Post
	query := db.Conn.Preload("Comment.User").Preload(clause.Associations).Find(&posts)
	if err := query.Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	var result []dto.PostResponse

	for _, post := range posts {
		postResult := dto.PostResponse{
			ID:      post.ID,
			Message: post.Message,
			UserId:  post.UserId,
			User: dto.UserResponse{
				ID:        post.User.ID,
				Email:     post.User.Email,
				FirstName: post.User.FirstName,
				LastName:  post.User.LastName,
			},
		}

		var comments []dto.CommentResponse = []dto.CommentResponse{}
		for _, comment := range post.Comment {
			comments = append(comments, dto.CommentResponse{
				ID:      comment.ID,
				Message: comment.Message,
				UserId:  comment.UserId,
				User: dto.UserResponse{
					ID:        comment.User.ID,
					Email:     comment.User.Email,
					FirstName: comment.User.FirstName,
					LastName:  comment.User.LastName,
				},
			},
			)
		}
		postResult.Comment = comments
		result = append(result, postResult)
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": result})

}
func (p Post) Update(ctx *gin.Context) {
	postId := ctx.Param("postId")
	user := ctx.MustGet("user").(dto.UserResponse)

	var form dto.PostRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existPost model.Post

	if err := db.Conn.Preload("Comment.User").First(&existPost, postId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if existPost.UserId != user.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
		return
	}

	existPost.Message = form.Message

	db.Conn.Preload("User").Save(&existPost)

	var comments []dto.CommentResponse
	for _, comment := range existPost.Comment {
		comments = append(comments, dto.CommentResponse{
			ID:      comment.ID,
			Message: comment.Message,
			UserId:  comment.UserId,
			User: dto.UserResponse{
				ID:        comment.User.ID,
				Email:     comment.User.Email,
				FirstName: comment.User.FirstName,
				LastName:  comment.User.LastName,
			},
		},
		)
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": dto.PostResponse{
		ID:      existPost.ID,
		Message: existPost.Message,
		UserId:  existPost.UserId,
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Comment: comments,
	}})
}

func (p Post) Delete(ctx *gin.Context) {
	postId := ctx.Param("postId")
	user := ctx.MustGet("user").(dto.UserResponse)

	var existPost model.Post

	if err := db.Conn.First(&existPost, postId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if existPost.UserId != user.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
		return
	}

	db.Conn.Select("Comment").Delete(&model.Post{}, postId)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
