package customer

import "time"

type CustomerRequest struct {
	Name        string `json:"name"`
	Email       string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type Customer struct {
	ID          int
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	BagLevel    int
	CreatedAt   time.Time
}

type CustomerRepository interface {
	CreateCustomer(*Customer) error
	GetCustomerByEmail(string) (*Customer, error)
	GetCustomerByPhone(string) (*Customer, error)
}

type CustomerService interface {
	RegisterCustomer(*CustomerRequest) error
}
