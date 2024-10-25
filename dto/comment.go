package dto

type CommentRequest struct {
	Message string `form:"message" binding:"required"`
}

type CommentCreate struct {
	UserId  string `form:"userId" binding:"required"`
	PostId  string `form:"postId" binding:"required"`
	Message string `form:"message" binding:"required"`
}

type CommentResponse struct {
	ID      uint         `json:"id"`
	Message string       `json:"message"`
	UserId  uint         `json:"userId"`
	PostId  uint         `json:"postId"`
	User    UserResponse `json:"user"`
}
