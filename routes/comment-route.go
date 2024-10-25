package routes

import (
	"simple-social-app/controller"
	"simple-social-app/middleware"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
)

func CommentRoute(r *gin.Engine) {
	commentController := controller.Comment{}
	commentGroup := r.Group("/comments")
	commentGroup.POST("/post/:postId", middleware.Authenticate(service.NewJWTService()), commentController.Create)
	commentGroup.GET("/post/:postId", middleware.Authenticate(service.NewJWTService()), commentController.FindAll)
	commentGroup.PATCH("/:commentId", middleware.Authenticate(service.NewJWTService()), commentController.Update)
	commentGroup.DELETE("/:commentId", middleware.Authenticate(service.NewJWTService()), commentController.Delete)
}
