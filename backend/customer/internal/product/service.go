package product

import (
	"context"
	"errors"
	"fmt"
)

type service struct {
	repo ProductRepository
}

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductImageNotFound = errors.New("product image not found")
)

func NewProductService(repo ProductRepository) ProductService {
	return &service{repo: repo}
}

func (s *service) GetProductByID(ctx context.Context, id int, custID int) (*Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)

	if err != nil {
		return nil, ErrProductNotFound
	}

	categories, _ := s.repo.GetCategoriesProduct(ctx, id)
	product.Categories = categories

	images, _ := s.repo.GetImagesProductID(ctx, id)
	for _, image := range images {
		product.ImagesID = append(product.ImagesID, image.ID)
	}

	_, err = s.repo.GetProductAvailableByID(ctx, product.ID)

	if err == nil {
		product.Status = "Available"
		return product, nil
	}

	owner, err := s.repo.GetProductPendingByID(ctx, id)
	if err != nil {
		product.Status = "Purchased"
		return product, nil
	}

	if custID == owner.CustomerID {
		product.Status = fmt.Sprintf("Pending(Owner), %s", owner.CreatedAt)
		return product, nil
	}

	product.Status = fmt.Sprintf("Pending, %s", owner.CreatedAt)
	return product, nil
}

func (s *service) GetImageProductByID(ctx context.Context, id int) (string, error) {
	path, err := s.repo.GetImageProductByID(ctx, id)
	if err != nil {
		return "", ErrProductImageNotFound
	}
	return path, nil
}
