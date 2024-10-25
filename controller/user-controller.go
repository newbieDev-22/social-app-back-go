package controller

import (
	"errors"
	"net/http"
	"simple-social-app/db"
	"simple-social-app/dto"
	"simple-social-app/helpers"
	"simple-social-app/model"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
}

func (u User) Register(ctx *gin.Context) {
	var form dto.UserCreateRequest

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.Password != form.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password and confirm password is not match"})
		return
	}

	var existUser model.User
	query := db.Conn.Where("email = ?", form.Email).First(&existUser)
	if err := query.Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Email is already in use"})
		return
	}

	hashedPassword, _ := helpers.HashPassword(form.Password)

	newUser := model.User{
		Email:     form.Email,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Password:  hashedPassword,
	}

	if err := db.Conn.Create(&newUser).Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.UserResponse{
		ID:        newUser.ID,
		Email:     newUser.FirstName,
		FirstName: newUser.LastName,
		LastName:  newUser.LastName,
	})
}

func (u User) Login(ctx *gin.Context) {
	var form dto.UserLoginRequest
	var existUser model.User

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := db.Conn.Where("email = ?", form.Email).First(&existUser)
	if err := query.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or Password"})
		return
	}

	isMatch, _ := helpers.CheckPassword(existUser.Password, []byte(form.Password))

	if !isMatch {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or Password"})
		return
	}

	jwtService := service.NewJWTService()
	accessToken := jwtService.GenerateToken(existUser.ID)

	ctx.JSON(http.StatusOK, dto.UserLoginResponse{
		AccessToken: accessToken,
	})
}
func (u User) GetMe(ctx *gin.Context) {
	user := ctx.MustGet("user").(dto.UserResponse)

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
