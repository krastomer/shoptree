package services

import "github.com/krastomer/shoptree/backend/internal/entities"

type profileService struct {
}

func NewProfileService() entities.ProfileService {
	return &profileService{}
}
