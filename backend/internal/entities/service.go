package entities

import "github.com/krastomer/shoptree/backend/internal/models"

type AuthService interface {
	Login(string, string) (string, error)
}

type ProfileService interface {
	GetProfile(id uint32) (*models.CustomerProfile, error)
}
