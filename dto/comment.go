package dto

type CommentRequest struct {
	Message string `form:"message" binding:"required"`
}

type CommentCreate struct {
	UserId  uint   `form:"userId" binding:"required"`
	Message string `form:"message" binding:"required"`
	PostId  uint   `form:"postId" binding:"required"`
}

type CommentResponse struct {
	ID      uint         `json:"id"`
	Message string       `json:"message"`
	UserId  uint         `json:"userId"`
	PostId  uint         `json:"postId"`
	User    UserResponse `json:"user"`
}
