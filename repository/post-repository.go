package repository

import (
	"context"
	"errors"
	"simple-social-app/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	PostRepository interface {
		CreatePost(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error)
		UpdatePost(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error)
		DeletePostById(ctx context.Context, tx *gorm.DB, postId string) error
		GetAllPost(ctx context.Context, tx *gorm.DB) ([]entity.Post, error)
		GetPostById(ctx context.Context, tx *gorm.DB, postId string) (entity.Post, error)
	}

	postRepository struct {
		db *gorm.DB
	}
)

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) CreatePost(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&post).Error; err != nil {
		return entity.Post{}, err
	}
	return post, nil
}

func (r *postRepository) UpdatePost(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Preload("User").Save(&post).Error; err != nil {
		return entity.Post{}, err
	}
	return post, nil
}

func (r *postRepository) DeletePostById(ctx context.Context, tx *gorm.DB, postId string) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Select("Comment").Delete(&entity.Post{}, postId).Error; err != nil {
		return err
	}
	return nil
}

func (r *postRepository) GetAllPost(ctx context.Context, tx *gorm.DB) ([]entity.Post, error) {
	if tx == nil {
		tx = r.db
	}

	var posts []entity.Post
	if err := tx.WithContext(ctx).Order("created_at desc").Preload("Comment.User").Preload(clause.Associations).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) GetPostById(ctx context.Context, tx *gorm.DB, postId string) (entity.Post, error) {
	if tx == nil {
		tx = r.db
	}
	var post entity.Post

	if err := tx.WithContext(ctx).Preload("User").Preload("Comment.User").First(&post, postId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Post{}, err
	}

	return post, nil
}
