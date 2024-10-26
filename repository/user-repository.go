package repository

import (
	"context"
	"errors"
	"simple-social-app/entity"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		GetUserById(ctx context.Context, tx *gorm.DB, userId uint) (entity.User, error)
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error)
		CheckEmailInUse(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil

}

func (r *userRepository) GetUserById(ctx context.Context, tx *gorm.DB, userId uint) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}
	var user entity.User
	getUserQuery := tx.WithContext(ctx).Where("id = ?", userId).Take(&user)

	if err := getUserQuery.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, err
	}

	return user, nil

}

func (r *userRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User
	getUserQuery := tx.WithContext(ctx).Where("email = ?", email).Take(&user)

	if err := getUserQuery.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, err
	}
	return user, nil

}

func (r *userRepository) CheckEmailInUse(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User

	if err := tx.WithContext(ctx).Where("email = ?", email).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, false, err
	}
	return user, true, nil

}
