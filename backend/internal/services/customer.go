package services

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/models"
)

type customerService struct {
	repo entities.CustomerRepo
}

func NewCustomerService(repo entities.CustomerRepo) entities.CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) GetCustomer(id int) (*models.Customer, error) {
	return s.repo.GetCustomer(id)
}
