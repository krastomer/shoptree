package entities

import "github.com/krastomer/shoptree/backend/internal/models"

type CustomerRepo interface {
	CreateCustomer(*models.CustomerProfile) error
	GetCustomerByEmail(string) (*models.Customer, error)
	GetCustomerByPhone(string) (*models.Customer, error)
	UpdateCustomerPassword(string) error
}

type AuthRepo interface {
	GetCustomerByEmail(string) (*models.Customer, error)
}

type ProductRepo interface {
	GetProductByID(int) (*models.Product, error)
}
