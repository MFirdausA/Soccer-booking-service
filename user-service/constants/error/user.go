package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordIncorrect    = errors.New("password is incorrect")
	ErrUsernameExist       = errors.New("username already exists")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
	ErrEmailExist		  = errors.New("email already exists")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
	ErrEmailExist,
}