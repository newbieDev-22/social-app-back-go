package routes

import (
	"simple-social-app/controller"
	"simple-social-app/middleware"
	"simple-social-app/repository"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService, userRepository repository.UserRepository) {
	routes := route.Group("/auth")
	{
		routes.POST("/register", userController.Register)
		routes.POST("/login", userController.Login)
		routes.GET("/get-me", middleware.Authenticate(jwtService, userRepository), userController.GetMe)
	}

}
