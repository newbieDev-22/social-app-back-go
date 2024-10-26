package controller

import (
	"net/http"
	"simple-social-app/dto"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
)

type (
	PostController interface {
		Create(ctx *gin.Context)
		FindAll(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	postController struct {
		postService service.PostService
	}
)

func NewPostController(postService service.PostService) PostController {
	return &postController{
		postService: postService,
	}
}

func (c *postController) Create(ctx *gin.Context) {
	var postReq dto.PostRequest
	user := ctx.MustGet("user").(dto.UserResponse)

	if err := ctx.ShouldBindJSON(&postReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_DATA_FROM_BODY})
		ctx.Abort()
		return
	}

	newPost := dto.PostCreate{
		UserId:  user.ID,
		Message: postReq.Message,
	}

	result, err := c.postService.CreatePost(ctx.Request.Context(), newPost)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_CREATE_POST})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"post": result})
}

func (c *postController) FindAll(ctx *gin.Context) {
	result, err := c.postService.GetPostAllPost(ctx.Request.Context())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_POST})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": result})
}

func (c *postController) Update(ctx *gin.Context) {
	postId := ctx.Param("postId")
	user := ctx.MustGet("user").(dto.UserResponse)
	postReq := dto.PostRequest{}

	if err := ctx.ShouldBindJSON(&postReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_DATA_FROM_BODY})
		ctx.Abort()
		return
	}

	existPost, err := c.postService.GetPostById(ctx.Request.Context(), postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_UPDATE_POST})
		ctx.Abort()
		return
	}

	if existPost.UserId != user.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "dto.MESSAGE_FAILED_UNAUTHORIZED"})
		ctx.Abort()
		return
	}

	existPost.Message = postReq.Message

	result, err := c.postService.UpdatePost(ctx.Request.Context(), existPost)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_UPDATE_POST})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"post": result})
}

func (c *postController) Delete(ctx *gin.Context) {
	postId := ctx.Param("postId")
	user := ctx.MustGet("user").(dto.UserResponse)

	existPost, err := c.postService.GetPostById(ctx.Request.Context(), postId)

	if err == nil {
		if existPost.UserId != user.ID {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": dto.MESSAGE_FAILED_UNAUTHORIZED})
			ctx.Abort()
			return
		}
		err = c.postService.DeletePostById(ctx.Request.Context(), postId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_DELETE_POST})
			ctx.Abort()
			return
		}
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
