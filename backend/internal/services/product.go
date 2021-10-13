package services

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/models"
)

type productService struct {
	repo entities.ProductRepo
}

func NewProductService(repo entities.ProductRepo) entities.ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProductByID(id int) (*models.Product, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
