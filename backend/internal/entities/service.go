package entities

import "github.com/krastomer/shoptree/backend/internal/models"

type CustomerService interface {
	GetCustomer(int) (*models.Customer, error)
}
