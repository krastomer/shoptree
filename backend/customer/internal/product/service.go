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

func (s *service) GetProducts(ctx context.Context, custID int) ([]*ProductMinimal, error) {
	products, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, ErrProductNotFound
	}

	var response []*ProductMinimal
	for _, p := range products {
		images, _ := s.repo.GetImagesProductID(ctx, p.ID)
		r := &ProductMinimal{
			ID:     p.ID,
			Name:   p.Name,
			Price:  p.Price,
			Status: s.getStatusProduct(ctx, p.ID, custID),
		}
		if images != nil {
			r.ImageID = images[0].ID
		}
		response = append(response, r)
	}

	return response, nil
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

	product.Status = s.getStatusProduct(ctx, product.ID, custID)

	return product, nil
}

func (s *service) GetImageProductByID(ctx context.Context, id int) (string, error) {
	path, err := s.repo.GetImageProductByID(ctx, id)
	if err != nil {
		return "", ErrProductImageNotFound
	}
	return path, nil
}

func (s *service) getStatusProduct(ctx context.Context, productID, custID int) string {
	_, err := s.repo.GetProductAvailableByID(ctx, productID)

	if err == nil {
		return "Available"
	}

	owner, err := s.repo.GetProductPendingByID(ctx, productID)
	if err != nil {
		return "Purchased"
	}

	if custID == owner.CustomerID {
		return fmt.Sprintf("Pending(Owner), %s", owner.CreatedAt)
	}

	return fmt.Sprintf("Pending, %s", owner.CreatedAt)
}
