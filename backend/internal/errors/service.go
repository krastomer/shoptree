package errors

import "errors"

var (
	ErrEmailInvalid        = errors.New("email invalid")
	ErrInternalServerError = errors.New("internal server error")
	ErrUserNotFound        = errors.New("user not found")
	ErrPasswordInvalid     = errors.New("password invalid")
)
