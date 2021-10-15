package models

import "time"

type Address struct {
	ID          uint32    `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	AddressLine string    `json:"adress_line"`
	Country     string    `json:"country"`
	State       string    `json:"state"`
	City        string    `json:"city"`
	District    string    `json:"district"`
	PostalCode  uint32    `json:"postal_code"`
	CreatedAt   time.Time `json:"created_at"`
}
