package auth

import (
	"context"
	"time"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Employee struct {
	ID          int        `dbq:"id" json:"-"`
	Name        string     `dbq:"name" json:"name"`
	Email       string     `dbq:"email" json:"username"`
	Password    string     `dbq:"password" json:"password"`
	PhoneNumber string     `dbq:"phone_number" json:"phone_number"`
	Level       string     `dbq:"level" json:"-"`
	CreatedAt   *time.Time `dbq:"created_at" json:"-"`
}

type AuthRepository interface {
	GetEmployeeByEmail(context.Context, string) (*Employee, error)
}

type AuthService interface {
	Login(context.Context, *UserRequest) (string, error)
}
