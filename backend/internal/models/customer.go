package models

import "time"

type Customer struct {
	ID          uint32
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
}

type CustomerProfile struct {
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	Addresses   []*Address `json:"addresses"`
	CreatedAt   time.Time  `json:"created_at"`
}
