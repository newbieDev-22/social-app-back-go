package dto

type PostRequest struct {
	Message string `form:"message" binding:"required"`
}

type PostInput struct {
	UserId  uint   `json:"userId"`
	Message string `json:"message"`
}

type PostResponse struct {
	ID      uint              `json:"id"`
	Message string            `json:"message"`
	UserId  uint              `json:"userId"`
	User    UserResponse      `json:"user"`
	Comment []CommentResponse `json:"comments"`
}
