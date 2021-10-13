package entities

import "github.com/krastomer/shoptree/backend/internal/models"

type AuthService interface {
	LoginCustomer(string, string) (string, error)
}

type ProfileService interface {
	CreateProfile(*models.CustomerProfile) error
}

type ProductService interface {
	GetProductByID(id int) (*models.Product, error)
}
