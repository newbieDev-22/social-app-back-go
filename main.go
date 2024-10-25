package main

import (
	"log"
	"os"
	"simple-social-app/db"
	"simple-social-app/middleware"
	"simple-social-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectDB()
	db.Migrate()
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	// routes
	routes.AuthRoute(server)
	routes.PostRoute(server)
	routes.CommentRoute(server)

	server.Run(":" + os.Getenv("PORT"))

}
