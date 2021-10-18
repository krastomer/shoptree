package errors

import "errors"

// Global
var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotAuthorized       = errors.New("not authorized")
)

// AuthService
var (
	ErrNotFoundUser   = errors.New("not found user")
	ErrEmailInvalid   = errors.New("email invalid")
	ErrPasswordInvlid = errors.New("password invalid")
	ErrEmailUsed      = errors.New("email used")
	ErrPhoneUsed      = errors.New("password used")
)

// ProductService
var (
	ErrNotFoundProduct = errors.New("not found product")
)

// ProfileService
var (
	ErrNotFoundAddress = errors.New("not found address")
)
