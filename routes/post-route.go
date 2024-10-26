package routes

import (
	"simple-social-app/controller"
	"simple-social-app/middleware"
	"simple-social-app/repository"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
)

func Post(route *gin.Engine, postController controller.PostController, jwtService service.JWTService, userRepository repository.UserRepository) {
	routes := route.Group("/posts")
	routes.Use(middleware.Authenticate(jwtService, userRepository))
	{
		routes.POST("/", postController.Create)
		routes.GET("/", postController.FindAll)
		routes.PATCH("/:postId", postController.Update)
		routes.DELETE("/:postId", postController.Delete)
	}

}
