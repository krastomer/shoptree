package product

import (
	"context"
	"errors"
)

type service struct {
	repo ProductRepository
}

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrCreateProductFailed = errors.New("create product failed")
)

func NewProductService(repo ProductRepository) ProductService {
	return &service{repo: repo}
}

func (s *service) GetProducts(ctx context.Context) ([]*Product, error) {
	products, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, ErrProductNotFound
	}
	return products, nil
}

func (s *service) CreateProduct(ctx context.Context, product *Product) error {
	err := s.repo.CreateProduct(ctx, product)
	if err != nil {
		return ErrCreateProductFailed
	}
	return nil
}
