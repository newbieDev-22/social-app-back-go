package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Message string    `gorm:"type:text;not null"`
	UserId  uint      `gorm:"not null"`
	User    User      `json:"user"`
	Comment []Comment `gorm:"foreignKey:PostId"`
}
