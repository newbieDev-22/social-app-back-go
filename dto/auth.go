package dto

type UserCreateRequest struct {
	Email           string `form:"email" binding:"required"`
	FirstName       string `form:"firstName" binding:"required"`
	LastName        string `form:"lastName" binding:"required"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" binding:"required"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserLoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserLoginResponse struct {
	AccessToken string `json:"accessToken"`
}
