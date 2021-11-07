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
	ID          int        `dbq:"id"`
	Name        string     `dbq:"name"`
	Email       string     `dbq:"email"`
	Password    string     `dbq:"password"`
	PhoneNumber string     `dbq:"phone_number"`
	CreatedAt   *time.Time `dbq:"created_at"`
}

type Employee struct {
	ID          int        `dbq:"id"`
	Name        string     `dbq:"name"`
	Email       string     `dbq:"email"`
	Password    string     `dbq:"password"`
	PhoneNumber string     `dbq:"phone_number"`
	Level       string     `dbq:"level"`
	CreatedAt   *time.Time `dbq:"created_at"`
}

type AuthRepository interface {
	GetCustomerByEmail(context.Context, string) (*Customer, error)
	GetEmployeeByEmail(context.Context, string) (*Employee, error)
}

type AuthService interface {
	Login(context.Context, *UserRequest) (string, error)
}
