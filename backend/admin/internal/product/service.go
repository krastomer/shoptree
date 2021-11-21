package product

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type service struct {
	repo ProductRepository
}

var (
	ErrProductNotFound       = errors.New("product not found")
	ErrCreateProductFailed   = errors.New("create product failed")
	ErrUpdateProductFailed   = errors.New("update product failed")
	ErrDeleteProductFailed   = errors.New("delete product failed")
	ErrAddImageProductFailed = errors.New("add image product failed")
	ErrProductImageNotFound  = errors.New("product image not found")
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

func (s *service) UpdateProduct(ctx context.Context, product *Product) error {
	old_product, err := s.repo.GetProductByID(ctx, product.ID)
	if err != nil {
		return ErrProductNotFound
	}

	s.cloneProduct(old_product, product)
	err = s.repo.UpdateProduct(ctx, product)
	if err != nil {
		return ErrUpdateProductFailed
	}
	return nil
}

func (s *service) cloneProduct(old_product, new_product *Product) {
	if new_product.Name == "" {
		new_product.Name = old_product.Name
	}
	if new_product.ScientificName == "" {
		new_product.ScientificName = old_product.ScientificName
	}
	if new_product.Description == "" {
		new_product.Description = old_product.Description
	}
	if new_product.Price == 0 {
		new_product.Price = old_product.Price
	}
}

func (s *service) DeleteProduct(ctx context.Context, id int) error {

	_, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return ErrProductNotFound
	}
	err = s.repo.DeleteProductByID(ctx, id)
	if err != nil {
		return ErrDeleteProductFailed
	}
	return nil
}

func (s *service) CreateImageProduct(ctx context.Context, c *fiber.Ctx, request *ImageProduct) error {
	uniqueId := uuid.New()
	request.ImagePath = fmt.Sprintf("%s/%s.jpg", viper.GetString("DIRECTORY_PRODUCT"), uniqueId)
	err := c.SaveFile(request.Image, request.ImagePath)
	if err != nil {
		return ErrAddImageProductFailed
	}

	err = s.repo.CreateImageProduct(ctx, request)
	if err != nil {
		return ErrAddImageProductFailed
	}
	return nil
}

func (s *service) GetImageProductByID(ctx context.Context, id int) (string, error) {
	path, err := s.repo.GetImageProductByID(ctx, id)
	if err != nil {
		return "", ErrProductImageNotFound
	}
	return path, nil
}
