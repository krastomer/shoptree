package models

import "time"

type CustomerProfile struct {
	Name        string `json:"name"`
	Email       string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type Customer struct {
	CustomerProfile
	ID        int
	CreatedAt time.Time
}
