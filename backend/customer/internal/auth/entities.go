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
	ID          int        `dbq:"id" json:"-"`
	Name        string     `dbq:"name" json:"name"`
	Email       string     `dbq:"email" json:"username"`
	Password    string     `dbq:"password" json:"password"`
	PhoneNumber string     `dbq:"phone_number" json:"phone_number"`
	CreatedAt   *time.Time `dbq:"created_at" json:"-"`
}

type AuthRepository interface {
	GetCustomerByEmail(context.Context, string) (*Customer, error)
	GetCustomerByPhone(context.Context, string) (*Customer, error)
	CreateCustomer(context.Context, *Customer) error
}

type AuthService interface {
	Login(context.Context, *UserRequest) (string, error)
	Register(context.Context, *Customer) error
}
