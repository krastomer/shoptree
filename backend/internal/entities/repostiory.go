package entities

import "github.com/krastomer/shoptree/backend/internal/models"

type CustomerRepo interface {
	GetCustomerByEmail(string) (*models.Customer, error)
}

type EmployeeRepo interface {
	GetEmployeeByEmail(string) (*models.Employee, error)
}
