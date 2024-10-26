package service

import (
	"context"
	"simple-social-app/dto"
	"simple-social-app/entity"
	"simple-social-app/helpers"
	"simple-social-app/repository"
)

type (
	UserService interface {
		Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		GetUserById(ctx context.Context, userId uint) (dto.UserResponse, error)
		GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error)
		Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
	}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *userService) Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {

	_, isExistEmail, _ := s.userRepo.CheckEmailInUse(ctx, nil, req.Email)

	if isExistEmail {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	hashedPassword, _ := helpers.HashPassword(req.Password)

	newUser := entity.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  hashedPassword,
	}

	userRegister, err := s.userRepo.CreateUser(ctx, nil, newUser)

	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:        userRegister.ID,
		Email:     userRegister.FirstName,
		FirstName: userRegister.LastName,
		LastName:  userRegister.LastName,
	}, nil

}

func (s *userService) GetUserById(ctx context.Context, userId uint) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	return dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, nil, email)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserByEmail
	}

	return dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (s *userService) Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	user, isExistEmail, err := s.userRepo.CheckEmailInUse(ctx, nil, req.Email)
	if err != nil || !isExistEmail {
		return dto.UserLoginResponse{}, dto.ErrEmailOrPassword
	}

	isMatch, err := helpers.CheckPassword(user.Password, []byte(req.Password))

	if err != nil || !isMatch {
		return dto.UserLoginResponse{}, dto.ErrEmailOrPassword
	}

	accessToken := s.jwtService.GenerateToken(user.ID)

	return dto.UserLoginResponse{
		AccessToken: accessToken,
	}, nil
}
