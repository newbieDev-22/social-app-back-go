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
)

type Comment struct {
}

func (c Comment) Create(ctx *gin.Context) {

	var form dto.CommentRequest
	postId := ctx.Param("postId")
	user := ctx.MustGet("user").(dto.UserResponse)

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existPost := model.Post{}

	existPostQuery := db.Conn.First(&existPost, postId)
	if err := existPostQuery.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newComment := model.Comment{
		UserId:  user.ID,
		Message: form.Message,
		PostId:  existPost.ID,
	}

	newCommentQuery := db.Conn.Preload("User").Create(&newComment)

	if err := newCommentQuery.Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"comment": dto.CommentResponse{
		ID:      newComment.ID,
		Message: newComment.Message,
		UserId:  newComment.UserId,
		PostId:  newComment.PostId,
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}})
}

func (c Comment) FindAll(ctx *gin.Context) {
	postId := ctx.Param("postId")

	existPost := model.Post{}

	existPostQuery := db.Conn.First(&existPost, postId)
	if err := existPostQuery.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var comments []model.Comment
	query := db.Conn.Preload("User").Find(&comments)
	if err := query.Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	var result []dto.CommentResponse

	for _, comment := range comments {
		result = append(result, dto.CommentResponse{
			ID:      comment.ID,
			Message: comment.Message,
			UserId:  comment.UserId,
			PostId:  comment.PostId,
			User: dto.UserResponse{
				ID:        comment.User.ID,
				Email:     comment.User.Email,
				FirstName: comment.User.FirstName,
				LastName:  comment.User.LastName,
			},
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"comments": result})
}

func (c Comment) Update(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	user := ctx.MustGet("user").(dto.UserResponse)

	var form dto.CommentRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existComment model.Comment

	if err := db.Conn.First(&existComment, commentId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if existComment.UserId != user.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
		return
	}

	existComment.Message = form.Message
	db.Conn.Preload("User").Save(&existComment)

	ctx.JSON(http.StatusOK, gin.H{"comment": dto.CommentResponse{
		ID:      existComment.ID,
		Message: existComment.Message,
		UserId:  existComment.UserId,
		PostId:  existComment.PostId,
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}})
}

func (c Comment) Delete(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	user := ctx.MustGet("user").(dto.UserResponse)

	var existComment model.Comment

	if err := db.Conn.First(&existComment, commentId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if existComment.UserId != user.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
		return
	}

	db.Conn.Unscoped().Delete(&model.Comment{}, commentId)
	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
