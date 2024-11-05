package main

import (
	"log"
	"os"
	"simple-social-app/config"
	"simple-social-app/controller"
	"simple-social-app/middleware"
	"simple-social-app/repository"
	"simple-social-app/routes"
	"simple-social-app/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)
	db = config.Migrate(db)

	var (
		jwtService        service.JWTService           = service.NewJWTService()
		userRepository    repository.UserRepository    = repository.NewUserRepository(db)
		userService       service.UserService          = service.NewUserService(userRepository, jwtService)
		userController    controller.UserController    = controller.NewUserController(userService)
		postRepository    repository.PostRepository    = repository.NewPostRepository(db)
		postService       service.PostService          = service.NewPostService(postRepository)
		postController    controller.PostController    = controller.NewPostController(postService)
		commentRepository repository.CommentRepository = repository.NewCommentRepository(db)
		commentService    service.CommentService       = service.NewCommentService(commentRepository, postRepository)
		commentController controller.CommentController = controller.NewCommentController(commentService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.User(server, userController, jwtService, userRepository)
	routes.Post(server, postController, jwtService, userRepository)
	routes.Comment(server, commentController, jwtService, userRepository)

	server.Run(":" + os.Getenv("PORT"))

}
