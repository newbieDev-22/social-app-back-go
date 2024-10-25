package db

import (
	"log"
	"os"
	"simple-social-app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func ConnectDB() {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(
		mysql.Open(databaseURL),
		&gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	Conn = db

}

func Migrate() {
	Conn.AutoMigrate(
		model.User{},
		model.Post{},
		model.Comment{},
	)
}
