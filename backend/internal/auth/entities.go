package auth

import (
	"context"
	"time"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserToken struct {
	ID       int
	Email    string
	Password string
	Level    string
}

type Customer struct {
	ID          int
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   *time.Time
}

type Employee struct {
	ID          int
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Level       string
	CreatedAt   *time.Time
}

type AuthRepository interface {
	GetCustomerByEmail(context.Context, string) (*Customer, error)
	GetEmployeeByEmail(context.Context, string) (*Employee, error)
}

type AuthService interface {
	Login(context.Context, *UserRequest) (string, error)
}
