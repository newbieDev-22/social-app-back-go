package dto

type PostRequest struct {
	Message string `form:"message" binding:"required"`
}

type PostCreate struct {
	UserId  string `form:"userId" binding:"required"`
	Message string `form:"message" binding:"required"`
}

type PostResponse struct {
	ID      uint              `json:"id"`
	Message string            `json:"message"`
	UserId  uint              `json:"userId"`
	User    UserResponse      `json:"user"`
	Comment []CommentResponse `json:"comments"`
}
