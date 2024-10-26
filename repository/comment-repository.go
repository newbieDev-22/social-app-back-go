package repository

import (
	"context"
	"errors"
	"simple-social-app/entity"

	"gorm.io/gorm"
)

type (
	CommentRepository interface {
		CreateComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error)
		UpdateComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error)
		DeleteCommentById(ctx context.Context, tx *gorm.DB, commentId string) error
		GetCommentById(ctx context.Context, tx *gorm.DB, commentId string) (entity.Comment, error)
		GetAllComment(ctx context.Context, tx *gorm.DB, postId string) ([]entity.Comment, error)
	}

	commentRepository struct {
		db *gorm.DB
	}
)

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) CreateComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Preload("User").Create(&comment).Error; err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}

func (r *commentRepository) UpdateComment(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Preload("User").Save(&comment).Error; err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}

func (r *commentRepository) DeleteCommentById(ctx context.Context, tx *gorm.DB, commentId string) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Select("Comment").Delete(&entity.Comment{}, commentId).Error; err != nil {
		return err
	}
	return nil
}

func (r *commentRepository) GetAllComment(ctx context.Context, tx *gorm.DB, postId string) ([]entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}

	var comments []entity.Comment
	if err := tx.WithContext(ctx).Where("post_id = ?", postId).Preload("User").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) GetCommentById(ctx context.Context, tx *gorm.DB, commentId string) (entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}
	var comment entity.Comment

	if err := tx.WithContext(ctx).Preload("User").First(&comment, commentId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Comment{}, err
	}
	return comment, nil
}
