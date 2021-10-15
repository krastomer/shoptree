package entities

import "github.com/krastomer/shoptree/backend/internal/models"

type CustomerRepo interface {
	GetCustomerByEmail(string) (*models.Customer, error)
}

type EmployeeRepo interface {
	GetEmployeeByEmail(string) (*models.Employee, error)
}

type CustomerProfileRepo interface {
	GetCustomerByID(uint32) (*models.Customer, error)
	GetAddresses(uint32) ([]*models.Address, error)
}

type ProductRepo interface {
	GetProduct(uint32) (*models.Product, error)
}
