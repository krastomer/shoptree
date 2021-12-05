package search

import (
	"context"
	"errors"
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
	for _, p := range products {
		p.ImagesID, _ = s.repo.GetImageProductByID(ctx, p.ID)
	}

	if cat == "" {
		return products, nil
	}

	listProduct, err := s.repo.GetCategoriesProducts(ctx, cat)
	if err != nil {
		return nil, ErrSearchNotFound
	}

	var response []*Product

	for _, p := range products {
		for _, l := range listProduct {
			if l == p.ID {
				response = append(response, p)
			}
		}
	}

	return response, nil
}
