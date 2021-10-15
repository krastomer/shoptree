package services

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
)

type productService struct {
	repo entities.ProductRepo
}

func NewProductService(repo entities.ProductRepo) entities.ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProduct(id uint32) (*models.ProductResponse, error) {
	product, err := s.repo.GetProduct(id)

	if err != nil {
		return nil, errors.ErrNotFoundProduct
	}

	prodRes := &models.ProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		ScientificName: product.ScientificName,
		Price:          product.Price,
		Description:    product.Description,
		Status:         product.Status,
	}

	return prodRes, nil
}
