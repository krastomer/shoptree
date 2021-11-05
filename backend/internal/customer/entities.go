package customer

import (
	"context"
	"time"
)

type CustomerRequest struct {
	Name        string `json:"name"`
	Email       string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type CustomerResponse struct {
	Name        string     `json:"name"`
	Email       string     `json:"username"`
	PhoneNumber string     `json:"phone_number"`
	Address     []*Address `json:"address"`
}

type Customer struct {
	ID          int
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
}

type Address struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	AddressLine string    `json:"address_line"`
	Country     string    `json:"country"`
	State       string    `json:"state"`
	City        string    `json:"city"`
	District    string    `json:"district"`
	PostalCode  string    `json:"postal_code"`
	CreatedAt   time.Time `json:"created_at"`
}

type Order struct {
	OrderResponse
	CustomerID int
}

type OrderResponse struct {
	ID              int
	AddressID       int
	PaymentEvidence string
	Status          string
	CreatedAt       time.Time
}

type CustomerRepository interface {
	CreateCustomer(context.Context, *CustomerRequest) error
	// CreateAddress(int, *Address) error
	// GetCustomerByEmail(string) (*Customer, error)
	// GetCustomerByID(int) (*Customer, error)
	// GetCustomerByPhone(string) (*Customer, error)
	// GetAddresses(int) ([]*Address, error)
	// GetInvoices(int) ([]*Order, error)
}

type CustomerService interface {
	// RegisterCustomer(*CustomerRequest) error
	// AddAddress(int, *Address) error
	// GetCustomer(int) (*CustomerResponse, error)
	// GetAddresses(int) ([]*Address, error)
	// GetOrders(int) ([]*OrderResponse, error)
}
