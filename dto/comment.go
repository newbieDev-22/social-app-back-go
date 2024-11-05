package dto

type CommentRequest struct {
	Message string `form:"message" binding:"required"`
}

type CommentInput struct {
	UserId  uint   `json:"userId"`
	Message string `json:"message"`
}

type CommentResponse struct {
	ID      uint         `json:"id"`
	Message string       `json:"message"`
	UserId  uint         `json:"userId"`
	PostId  uint         `json:"postId"`
	User    UserResponse `json:"user"`
}
