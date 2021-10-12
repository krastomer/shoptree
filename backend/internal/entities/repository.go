package entities

import "github.com/krastomer/shoptree/backend/internal/models"

type CustomerRepo interface {
	GetCustomer(string) (*models.Customer, error)
}

type ItemRepo interface {
}
