package routes

import (
	"simple-social-app/controller"
	"simple-social-app/middleware"
	"simple-social-app/repository"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
)

func Comment(route *gin.Engine, commentController controller.PostController, jwtService service.JWTService, userRepository repository.UserRepository) {
	routes := route.Group("/comments")
	routes.Use(middleware.Authenticate(jwtService, userRepository))
	{
		routes.POST("/post/:postId", commentController.Create)
		routes.GET("/post/:postId", commentController.FindAll)
		routes.PATCH("/:commentId", commentController.Update)
		routes.DELETE("/:commentId", commentController.Delete)
	}
}
