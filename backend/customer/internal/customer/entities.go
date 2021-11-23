package customer

import (
	"context"
	"time"
)

type Address struct {
	ID          int        `json:"id" dbq:"id"`
	CustomerID  int        `json:"-" dbq:"customer_id"`
	CreatedAt   *time.Time `json:"created_at" dbq:"created_at"`
	Name        string     `json:"name" dbq:"name"`
	PhoneNumber string     `json:"phone_number" dbq:"phone_number"`
	AddressLine string     `json:"address_line" dbq:"address_line"`
	Country     string     `json:"country" dbq:"country"`
	State       string     `json:"state" dbq:"state"`
	City        string     `json:"city" dbq:"city"`
	District    string     `json:"district" dbq:"district"`
	PostalCode  string     `json:"postal_code" dbq:"postal_code"`
}

type AddressRequest struct {
	Name        string `json:"name" dbq:"name"`
	PhoneNumber string `json:"phone_number" dbq:"phone_number"`
	AddressLine string `json:"address_line" dbq:"address_line"`
	Country     string `json:"country" dbq:"country"`
	State       string `json:"state" dbq:"state"`
	City        string `json:"city" dbq:"city"`
	District    string `json:"district" dbq:"district"`
	PostalCode  string `json:"postal_code" dbq:"postal_code"`
}

type Customer struct {
	ID          int        `dbq:"id" json:"-"`
	Name        string     `dbq:"name" json:"name"`
	Email       string     `dbq:"email" json:"username"`
	Password    string     `dbq:"password" json:"-"`
	PhoneNumber string     `dbq:"phone_number" json:"phone_number"`
	CreatedAt   *time.Time `dbq:"created_at" json:"-"`
	Address     []*Address `dbq:"-" json:"address"`
}

type CustomerRepository interface {
	GetCustomerByID(context.Context, int) (*Customer, error)
	GetAddressesCustomer(context.Context, int) ([]*Address, error)
	CreateAddressCustomer(context.Context, *Address) error
}

type CustomerService interface {
	GetCustomerProfile(context.Context, int) (*Customer, error)
	GetAddresses(context.Context, int) ([]*Address, error)
	CreateAddressCustomer(context.Context, int, *AddressRequest) error
}
