package config

import (
	"log"
	"os"
	"simple-social-app/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(
		mysql.Open(databaseURL),
		&gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db

}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		entity.User{},
		entity.Post{},
		entity.Comment{},
	)
	return db
}
