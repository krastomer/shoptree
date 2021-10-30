package product

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrAddProductFailed     = errors.New("add product failed")
	ErrProductImageNotFound = errors.New("product image not found")
)

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProductByID(id uint32) (*Product, error) {
	product, err := s.repo.GetProductByID(id)

	if err != nil {
		return nil, ErrProductNotFound
	}
	images, _ := s.repo.GetProductImagesID(id)
	product.ImagesID = images

	return product, nil
}

func (s *productService) GetProducts(id []uint32) ([]*Product, error) {
	products, err := s.repo.GetProducts()
	return products, err
}

func (s *productService) AddProduct(product *Product) error {
	err := s.repo.AddProduct(product)
	if err != nil {
		return ErrAddProductFailed
	}
	return nil
}

func (s *productService) AddProductImage(id uint32, c *fiber.Ctx, file *multipart.FileHeader) error {
	uniqueId := uuid.New()
	path := fmt.Sprintf("%s/%s.jpg", viper.GetString("DIRECTORY_PRODUCT"), uniqueId)
	err := s.repo.AddProductImage(id, path)

	if err != nil {
		return ErrProductNotFound
	}

	c.SaveFile(file, path)
	return nil
}

func (s *productService) GetProductImageByID(id uint32) (string, error) {
	path, err := s.repo.GetProductImageByID(id)
	if err != nil {
		return "", ErrProductImageNotFound
	}
	return path, nil
}
