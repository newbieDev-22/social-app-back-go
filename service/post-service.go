package service

import (
	"context"
	"simple-social-app/dto"
	"simple-social-app/entity"
	"simple-social-app/repository"
	"strconv"
)

type (
	PostService interface {
		CreatePost(ctx context.Context, req dto.PostInput) (dto.PostResponse, error)
		UpdatePost(ctx context.Context, req dto.PostInput, postId string) (dto.PostResponse, error)
		GetPostAllPost(ctx context.Context) ([]dto.PostResponse, error)
		GetPostById(ctx context.Context, postId string) (entity.Post, error)
		DeletePostById(ctx context.Context, userId uint, postId string) error
	}

	postService struct {
		postRepo repository.PostRepository
	}
)

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{
		postRepo: postRepo,
	}
}

func (s *postService) CreatePost(ctx context.Context, req dto.PostInput) (dto.PostResponse, error) {

	post := entity.Post{
		UserId:  req.UserId,
		Message: req.Message,
	}

	createPost, err := s.postRepo.CreatePost(ctx, nil, post)
	if err != nil {
		return dto.PostResponse{}, err
	}

	newPost, err := s.postRepo.GetPostById(ctx, nil, strconv.FormatUint(uint64(createPost.ID), 10))
	if err != nil {
		return dto.PostResponse{}, err
	}

	return dto.PostResponse{
		ID:      newPost.ID,
		Message: newPost.Message,
		UserId:  newPost.UserId,
		User: dto.UserResponse{
			ID:        newPost.User.ID,
			Email:     newPost.User.Email,
			FirstName: newPost.User.FirstName,
			LastName:  newPost.User.LastName,
		},
		Comment: []dto.CommentResponse{},
	}, nil
}
func (s *postService) GetPostById(ctx context.Context, postId string) (entity.Post, error) {

	post, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return entity.Post{}, err
	}

	return post, nil
}
func (s *postService) UpdatePost(ctx context.Context, req dto.PostInput, postId string) (dto.PostResponse, error) {

	existPost, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return dto.PostResponse{}, err
	}

	if existPost.UserId != req.UserId {
		return dto.PostResponse{}, dto.ErrUnauthorized
	}

	existPost.Message = req.Message

	updatePost, err := s.postRepo.UpdatePost(ctx, nil, existPost)
	if err != nil {
		return dto.PostResponse{}, err
	}

	var comments []dto.CommentResponse
	for _, comment := range updatePost.Comment {
		comments = append(comments, dto.CommentResponse{
			ID:      comment.ID,
			Message: comment.Message,
			UserId:  comment.UserId,
			User: dto.UserResponse{
				ID:        comment.User.ID,
				Email:     comment.User.Email,
				FirstName: comment.User.FirstName,
				LastName:  comment.User.LastName,
			},
		},
		)
	}

	return dto.PostResponse{
		ID:      updatePost.ID,
		Message: updatePost.Message,
		UserId:  updatePost.UserId,
		User: dto.UserResponse{
			ID:        updatePost.User.ID,
			Email:     updatePost.User.Email,
			FirstName: updatePost.User.FirstName,
			LastName:  updatePost.User.LastName,
		},
		Comment: comments,
	}, nil
}

func (s *postService) GetPostAllPost(ctx context.Context) ([]dto.PostResponse, error) {

	posts, err := s.postRepo.GetAllPost(ctx, nil)
	if err != nil {
		return []dto.PostResponse{}, err
	}

	var result []dto.PostResponse

	for _, post := range posts {
		postResult := dto.PostResponse{
			ID:      post.ID,
			Message: post.Message,
			UserId:  post.UserId,
			User: dto.UserResponse{
				ID:        post.User.ID,
				Email:     post.User.Email,
				FirstName: post.User.FirstName,
				LastName:  post.User.LastName,
			},
		}

		var comments []dto.CommentResponse = []dto.CommentResponse{}
		for _, comment := range post.Comment {
			comments = append(comments, dto.CommentResponse{
				ID:      comment.ID,
				Message: comment.Message,
				UserId:  comment.UserId,
				User: dto.UserResponse{
					ID:        comment.User.ID,
					Email:     comment.User.Email,
					FirstName: comment.User.FirstName,
					LastName:  comment.User.LastName,
				},
			},
			)
		}
		postResult.Comment = comments
		result = append(result, postResult)
	}

	return result, nil
}

func (s *postService) DeletePostById(ctx context.Context, userId uint, postId string) error {
	existPost, err := s.postRepo.GetPostById(ctx, nil, postId)
	if err != nil {
		return err
	}

	if existPost.UserId != userId {
		return dto.ErrUnauthorized
	}

	err = s.postRepo.DeletePostById(ctx, nil, postId)

	return err
}
