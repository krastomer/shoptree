package errors

import "errors"

// AuthService
var (
	ErrEmailInvalid        = errors.New("email invalid")
	ErrInternalServerError = errors.New("internal server error")
	ErrUserNotFound        = errors.New("user not found")
	ErrPasswordInvalid     = errors.New("password invalid")
)

// ProfileService
var (
	ErrPhoneNumberInvalid = errors.New("phone number invalid")
	ErrUserExisted        = errors.New("user existed")
)
