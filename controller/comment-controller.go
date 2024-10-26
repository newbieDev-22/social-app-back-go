package controller

import (
	"net/http"
	"simple-social-app/dto"
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
		postService    service.PostService
	}
)

func NewCommentController(commentService service.CommentService, postService service.PostService) CommentController {
	return &commentController{
		commentService: commentService,
		postService:    postService,
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

	existPost, err := c.postService.GetPostById(ctx.Request.Context(), postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_POST})
		ctx.Abort()
		return
	}

	newComment := dto.CommentCreate{
		UserId:  user.ID,
		Message: commentReq.Message,
		PostId:  existPost.ID,
	}

	result, err := c.commentService.CreateComment(ctx.Request.Context(), newComment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_CREATE_COMMENT})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"comment": result})
}

func (c *commentController) FindAll(ctx *gin.Context) {
	postId := ctx.Param("postId")

	_, err := c.postService.GetPostById(ctx.Request.Context(), postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_POST})
		ctx.Abort()
		return
	}

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

	existComment, err := c.commentService.GetCommentById(ctx.Request.Context(), commentId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_COMMENT})
		return
	}

	if existComment.UserId != user.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated"})
		return
	}

	existComment.Message = commentReq.Message

	result, err := c.commentService.UpdateComment(ctx.Request.Context(), existComment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_UPDATE_COMMENT})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"post": result})
}

func (c *commentController) Delete(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	user := ctx.MustGet("user").(dto.UserResponse)

	existComment, err := c.commentService.GetCommentById(ctx.Request.Context(), commentId)

	if err == nil {
		if existComment.UserId != user.ID {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": dto.MESSAGE_FAILED_UNAUTHORIZED})
			ctx.Abort()
			return
		}
		err = c.commentService.DeleteCommentById(ctx.Request.Context(), commentId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_DELETE_COMMENT})
			ctx.Abort()
			return
		}
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
