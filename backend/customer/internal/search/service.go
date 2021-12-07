package search

import (
	"context"
	"errors"
	"fmt"
)

type service struct {
	repo SearchRepository
}

var (
	ErrSearchNotFound = errors.New("search not found")
)

func NewSearchService(repo SearchRepository) SearchService {
	return &service{repo: repo}
}

func (s *service) GetCategories(ctx context.Context) ([]*CategoryProduct, error) {
	cat, _ := s.repo.GetCategoriesProduct(ctx)
	return cat, nil
}

func (s *service) Search(ctx context.Context, cat, data string) ([]*Product, error) {
	products, err := s.repo.GetProductsLike(ctx, data)
	if err != nil {
		return nil, ErrSearchNotFound
	}
	for _, product := range products {
		product.ImageID, _ = s.repo.GetImageProductByID(ctx, product.ID)
		_, err = s.repo.GetProductAvailableByID(ctx, product.ID)

		if err == nil {
			product.Status = "Available"
			continue
		}

		owner, err := s.repo.GetProductPendingByID(ctx, product.ID)
		if err != nil {
			product.Status = "Purchased"
			continue
		}

		product.Status = fmt.Sprintf("Pending, %s", owner.CreatedAt)
	}

	return products, nil
}
