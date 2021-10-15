package errors

import "errors"

// Global
var (
	ErrInternalServerError = errors.New("internal server error")
)

// AuthService
var (
	ErrNotFoundUser   = errors.New("not found user")
	ErrEmailInvalid   = errors.New("email invalid")
	ErrPasswordInvlid = errors.New("password invalid")
)

// ProductService
var (
	ErrNotFoundProduct = errors.New("not found product")
)
