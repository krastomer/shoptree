package entities

import "github.com/krastomer/shoptree/backend/internal/models"

type AuthService interface {
	Login(string, string) (string, error)
	Register(*models.User) error
}

type ProfileService interface {
	GetProfile(id uint32) (*models.CustomerProfile, error)
}

type ProductService interface {
	GetProduct(id uint32) (*models.ProductResponse, error)
}
