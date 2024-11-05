package dto

import "errors"

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY      = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER           = "failed create user"
	MESSAGE_FAILED_GET_USER_TOKEN          = "failed get user token"
	MESSAGE_FAILED_TOKEN_NOT_VALID         = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND         = "token not found"
	MESSAGE_FAILED_GET_USER                = "failed get user"
	MESSAGE_FAILED_LOGIN                   = "failed login"
	MESSAGE_FAILED_WRONG_EMAIL_OR_PASSWORD = "wrong email or password"
	MESSAGE_FAILED_PASSWORD_NOT_MATCH      = "password and confirm password is not match"
	MESSAGE_FAILED_UNAUTHORIZED            = "unauthorized"
	MESSAGE_FAILED_CREATE_POST             = "failed create post"
	MESSAGE_FAILED_GET_POST                = "failed get post"
	MESSAGE_FAILED_UPDATE_POST             = "failed update post"
	MESSAGE_FAILED_DELETE_POST             = "failed delete post"
	MESSAGE_FAILED_CREATE_COMMENT          = "failed create post"
	MESSAGE_FAILED_GET_COMMENT             = "failed get post"
	MESSAGE_FAILED_UPDATE_COMMENT          = "failed update post"
	MESSAGE_FAILED_DELETE_COMMENT          = "failed delete post"
	MESSAGE_FAILED_FORGOT_PARAMS           = "forgot params"
)

var (
	ErrCreateUser         = errors.New("failed to create user")
	ErrGetUserById        = errors.New("failed to get user by id")
	ErrGetUserByEmail     = errors.New("failed to get user by email")
	ErrEmailAlreadyExists = errors.New("email already exist")
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailNotFound      = errors.New("email not found")
	ErrPasswordNotMatch   = errors.New("password not match")
	ErrEmailOrPassword    = errors.New("wrong email or password")
	ErrTokenInvalid       = errors.New("token invalid")
	ErrTokenExpired       = errors.New("token expired")
	ErrUnauthorized       = errors.New("Unauthorized")
	ErrCannotFindPost     = errors.New("cannot find post")
)
