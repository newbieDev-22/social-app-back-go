package routes

import (
	"simple-social-app/controller"
	"simple-social-app/middleware"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.Engine) {

	authController := controller.User{}
	authGroup := r.Group("/auth")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)
	authGroup.GET("/get-me", middleware.Authenticate(service.NewJWTService()), authController.GetMe)
}
