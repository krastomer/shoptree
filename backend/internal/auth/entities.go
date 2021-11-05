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
	ID          int        `mapstructure:"id"`
	Name        string     `mapstructure:"name"`
	Email       string     `mapstructure:"email"`
	Password    string     `mapstructure:"password"`
	PhoneNumber string     `mapstructure:"phone_number"`
	CreatedAt   *time.Time `mapstructure:"created_at"`
}

type Employee struct {
	ID          int        `mapstructure:"id"`
	Name        string     `mapstructure:"name"`
	Email       string     `mapstructure:"email"`
	Password    string     `mapstructure:"password"`
	PhoneNumber string     `mapstructure:"phone_number"`
	Level       string     `mapstructure:"level"`
	CreatedAt   *time.Time `mapstructure:"created_at"`
}

type AuthRepository interface {
	GetCustomerByEmail(context.Context, string) (*Customer, error)
	GetEmployeeByEmail(context.Context, string) (*Employee, error)
}

type AuthService interface {
	Login(context.Context, *UserRequest) (string, error)
}
