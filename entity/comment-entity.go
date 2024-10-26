package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Message string `gorm:"type:text;not null" json:"message"`
	UserId  uint   `gorm:"not null" json:"userId"`
	User    User   `json:"user"`
	PostId  uint   `gorm:"not null" json:"postId"`
}
