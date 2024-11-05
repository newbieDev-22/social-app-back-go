package service

import (
	"context"
	"simple-social-app/dto"
	"simple-social-app/entity"
	"simple-social-app/repository"
	"strconv"
)

type (
	CommentService interface {
		CreateComment(ctx context.Context, req dto.CommentInput, postId string) (dto.CommentResponse, error)
		UpdateComment(ctx context.Context, req dto.CommentInput, commentId string) (dto.CommentResponse, error)
		GetCommentById(ctx context.Context, commentId string) (entity.Comment, error)
		GetAllComment(ctx context.Context, postId string) ([]dto.CommentResponse, error)
		DeleteCommentById(ctx context.Context, userId uint, commentId string) error
	}

	commentService struct {
		commentRepo repository.CommentRepository
		postRepo    repository.PostRepository
	}
)

func NewCommentService(commentRepo repository.CommentRepository, postRepo repository.PostRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

func (s *commentService) CreateComment(ctx context.Context, req dto.CommentInput, postId string) (dto.CommentResponse, error) {

	existPost, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	comment := entity.Comment{
		UserId:  req.UserId,
		Message: req.Message,
		PostId:  existPost.ID,
	}

	createComment, err := s.commentRepo.CreateComment(ctx, nil, comment)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	newComment, err := s.commentRepo.GetCommentById(ctx, nil, strconv.FormatUint(uint64(createComment.ID), 10))
	if err != nil {
		return dto.CommentResponse{}, err
	}

	return dto.CommentResponse{
		ID:      newComment.ID,
		Message: newComment.Message,
		UserId:  newComment.UserId,
		PostId:  newComment.PostId,
		User: dto.UserResponse{
			ID:        newComment.User.ID,
			Email:     newComment.User.Email,
			FirstName: newComment.User.FirstName,
			LastName:  newComment.User.LastName,
		},
	}, nil
}

func (s *commentService) UpdateComment(ctx context.Context, req dto.CommentInput, commentId string) (dto.CommentResponse, error) {

	existComment, err := s.commentRepo.GetCommentById(ctx, nil, commentId)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	if existComment.UserId != req.UserId {
		return dto.CommentResponse{}, dto.ErrUnauthorized
	}

	existComment.Message = req.Message

	updateComment, err := s.commentRepo.UpdateComment(ctx, nil, existComment)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	return dto.CommentResponse{
		ID:      updateComment.ID,
		Message: updateComment.Message,
		UserId:  updateComment.UserId,
		PostId:  updateComment.PostId,
		User: dto.UserResponse{
			ID:        updateComment.User.ID,
			Email:     updateComment.User.Email,
			FirstName: updateComment.User.FirstName,
			LastName:  updateComment.User.LastName,
		},
	}, nil
}

func (s *commentService) GetCommentById(ctx context.Context, commentId string) (entity.Comment, error) {

	comment, err := s.commentRepo.GetCommentById(ctx, nil, commentId)
	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (s *commentService) GetAllComment(ctx context.Context, postId string) ([]dto.CommentResponse, error) {

	_, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return []dto.CommentResponse{}, err
	}

	comments, err := s.commentRepo.GetAllComment(ctx, nil, postId)
	if err != nil {
		return []dto.CommentResponse{}, err
	}

	var result []dto.CommentResponse

	for _, comment := range comments {
		result = append(result, dto.CommentResponse{
			ID:      comment.ID,
			Message: comment.Message,
			UserId:  comment.UserId,
			PostId:  comment.PostId,
			User: dto.UserResponse{
				ID:        comment.User.ID,
				Email:     comment.User.Email,
				FirstName: comment.User.FirstName,
				LastName:  comment.User.LastName,
			},
		})
	}

	return result, nil
}

func (s *commentService) DeleteCommentById(ctx context.Context, userId uint, commentId string) error {
	existComment, err := s.commentRepo.GetCommentById(ctx, nil, commentId)
	if err != nil {
		return err
	}

	if existComment.UserId != userId {
		return dto.ErrUnauthorized
	}
	err = s.commentRepo.DeleteCommentById(ctx, nil, commentId)

	return err
}
