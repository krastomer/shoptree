package entities

import "github.com/krastomer/shoptree/backend/internal/models"

type CustomerRepo interface {
	// CreateCustomer(models.Customer) error
	GetCustomer(int) (*models.Customer, error)
}

type ItemRepo interface {
}
