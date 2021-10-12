package model

import "time"

type Customer struct {
	ID          int
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
}
