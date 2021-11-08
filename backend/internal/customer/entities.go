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
	ID          int        `dbq:"id"`
	Name        string     `dbq:"name"`
	Email       string     `dbq:"email"`
	Password    string     `dbq:"password"`
	PhoneNumber string     `dbq:"phone_number"`
	CreatedAt   *time.Time `dbq:"created_at"`
}

type Address struct {
	ID          int        `json:"id" dbq:"id"`
	CustomerID  int        `json:"-" dbq:"customer_id"`
	Name        string     `json:"name" dbq:"name"`
	PhoneNumber string     `json:"phone_number" dbq:"phone_number"`
	AddressLine string     `json:"address_line" dbq:"address_line"`
	Country     string     `json:"country" dbq:"country"`
	State       string     `json:"state" dbq:"state"`
	City        string     `json:"city" dbq:"city"`
	District    string     `json:"district" dbq:"district"`
	PostalCode  string     `json:"postal_code" dbq:"postal_code"`
	CreatedAt   *time.Time `json:"created_at" dbq:"created_at"`
}

type CustomerRepository interface {
	CreateCustomer(context.Context, *CustomerRequest) error
	GetCustomerByEmail(context.Context, string) (*Customer, error)
	GetCustomerByID(context.Context, int) (*Customer, error)
	GetCustomerByPhone(context.Context, string) (*Customer, error)
	CreateAddressCustomer(context.Context, *Address) error
	GetAddressesCustomer(context.Context, int) ([]*Address, error)
}

type CustomerService interface {
	CreateNewCustomer(context.Context, *CustomerRequest) error
	// AddAddress(int, *Address) error
	// GetCustomer(int) (*CustomerResponse, error)
	GetAddressesCustomer(context.Context, int) ([]*Address, error)
	CreateAddressCustomer(context.Context, *Address) error
}
