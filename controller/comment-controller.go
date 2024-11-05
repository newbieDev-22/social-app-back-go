package controller

import (
	"net/http"
	"simple-social-app/dto"
	"simple-social-app/entity"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
)

type (
	CommentController interface {
		Create(ctx *gin.Context)
		FindAll(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	commentController struct {
		commentService service.CommentService
	}
)

func NewCommentController(commentService service.CommentService) CommentController {
	return &commentController{
		commentService: commentService,
	}
}

func (c *commentController) Create(ctx *gin.Context) {
	var commentReq dto.CommentRequest
	postId := ctx.Param("postId")
	user := ctx.MustGet("user").(dto.UserResponse)

	if err := ctx.ShouldBindJSON(&commentReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_DATA_FROM_BODY})
		ctx.Abort()
		return
	}

	newComment := entity.Comment{
		UserId:  user.ID,
		Message: commentReq.Message,
	}

	result, err := c.commentService.CreateComment(ctx.Request.Context(), newComment, postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"comment": result})
}

func (c *commentController) FindAll(ctx *gin.Context) {
	postId := ctx.Param("postId")

	result, err := c.commentService.GetAllComment(ctx.Request.Context(), postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_COMMENT})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"comments": result})
}

func (c *commentController) Update(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	user := ctx.MustGet("user").(dto.UserResponse)

	var commentReq dto.CommentRequest
	if err := ctx.ShouldBind(&commentReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_DATA_FROM_BODY})
		return
	}

	updateComment := entity.Comment{
		UserId:  user.ID,
		Message: commentReq.Message,
	}

	result, err := c.commentService.UpdateComment(ctx.Request.Context(), updateComment, commentId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_UPDATE_COMMENT})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"post": result})
}

func (c *commentController) Delete(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	user := ctx.MustGet("user").(dto.UserResponse)

	err := c.commentService.DeleteCommentById(ctx.Request.Context(), user.ID, commentId)
	if err != nil && err == dto.ErrUnauthorized {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
