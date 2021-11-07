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
	CreatedAt   *time.Time
}

type Address struct {
	CustomerID int
	AddressResponse
}

type AddressResponse struct {
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
	GetCustomerByEmail(context.Context, string) (*Customer, error)
	GetCustomerByID(context.Context, int) (*Customer, error)
	GetCustomerByPhone(context.Context, string) (*Customer, error)
	// CreateAddress(int, *Address) error
	GetAddressesCustomer(context.Context, int) ([]*Address, error)
	// GetInvoices(int) ([]*Order, error)
}

type CustomerService interface {
	CreateNewCustomer(context.Context, *CustomerRequest) error
	// AddAddress(int, *Address) error
	// GetCustomer(int) (*CustomerResponse, error)
	GetAddressesCustomer(context.Context, int) ([]*AddressResponse, error)
	// GetOrders(int) ([]*OrderResponse, error)
}
