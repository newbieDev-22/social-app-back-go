package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;type:varchar(100);not null" json:"email"`
	FirstName string `gorm:"varchar(100);not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(100);not null" json:"last_name"`
	Password  string `gorm:"type:varchar(100);not null" json:"password"`
}
