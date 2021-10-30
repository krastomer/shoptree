package profile

import "time"

type User struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Level       string `json:"level"`
}

type Customer struct {
	ID          uint32
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	BagLevel    uint8
	CreatedAt   time.Time
}

type CustomerProfile struct {
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	Addresses   []*Address `json:"addresses"`
	CreatedAt   time.Time  `json:"created_at"`
}

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

type ProfileRepository interface {
	GetCustomerByID(uint32) (*Customer, error)
	GetAddresses(uint32) ([]*Address, error)
}

type ProfileService interface {
	GetProfileCustomer(uint32) (*CustomerProfile, error)
}
