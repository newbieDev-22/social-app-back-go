package service

import (
	"simple-social-app/db"
	"simple-social-app/model"
)

type AuthService struct {
}

func (a AuthService) FindUserByEmail(user model.User, email string) model.User {
	db.Conn.Where("email = ?", email).First(&user)
	return user
}
