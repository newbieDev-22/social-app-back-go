package controller

import (
	"net/http"
	"simple-social-app/dto"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		GetMe(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var userReq dto.UserCreateRequest

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_DATA_FROM_BODY})
		ctx.Abort()
		return
	}

	if userReq.Password != userReq.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_PASSWORD_NOT_MATCH})
		ctx.Abort()
		return
	}

	result, err := c.userService.Register(ctx.Request.Context(), userReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_REGISTER_USER})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *userController) Login(ctx *gin.Context) {
	var userReq dto.UserLoginRequest

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_GET_DATA_FROM_BODY})
		ctx.Abort()
		return
	}

	result, err := c.userService.Login(ctx.Request.Context(), userReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": dto.MESSAGE_FAILED_WRONG_EMAIL_OR_PASSWORD})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}
func (c *userController) GetMe(ctx *gin.Context) {
	user := ctx.MustGet("user").(dto.UserResponse)

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
