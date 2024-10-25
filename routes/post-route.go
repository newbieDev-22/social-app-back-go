package routes

import (
	"simple-social-app/controller"
	"simple-social-app/middleware"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
)

func PostRoute(r *gin.Engine) {
	postController := controller.Post{}
	postGroup := r.Group("/posts")
	postGroup.POST("/", middleware.Authenticate(service.NewJWTService()), postController.Create)
	postGroup.GET("/", middleware.Authenticate(service.NewJWTService()), postController.FindAll)
	postGroup.PATCH("/:postId", middleware.Authenticate(service.NewJWTService()), postController.Update)
	postGroup.DELETE("/:postId", middleware.Authenticate(service.NewJWTService()), postController.Delete)
}
