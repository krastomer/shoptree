package product

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type productService struct {
	repo ProductRepository
}

var (
	ErrProductNotFound          = errors.New("product not found")
	ErrProductImageNotFound     = errors.New("product image not found")
	ErrCreateProductFailed      = errors.New("add product failed")
	ErrCreateImageProductFailed = errors.New("add image product failed")
	ErrProductStatus            = errors.New("product status not match")
)

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProductByID(ctx context.Context, id int) (*Product, error) {
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

	return product, nil
}

func (s *productService) GetImageProductByID(ctx context.Context, id int) (string, error) {
	path, err := s.repo.GetImageProductByID(ctx, id)
	if err != nil {
		return "", ErrProductImageNotFound
	}
	return path, nil
}

func (s *productService) GetImagesProductID(ctx context.Context, id int) ([]int, error) {
	imagesID, err := s.repo.GetImagesProductID(ctx, id)
	if err != nil {
		return nil, ErrProductImageNotFound
	}
	var response []int
	for _, imageID := range imagesID {
		response = append(response, imageID.ID)
	}
	return response, nil
}

func (s *productService) CreateImageProduct(ctx context.Context, c *fiber.Ctx, request *ImageProduct) error {
	uniqueId := uuid.New()
	request.ImagePath = fmt.Sprintf("%s/%s.jpg", viper.GetString("DIRECTORY_PRODUCT"), uniqueId)
	err := c.SaveFile(request.Image, request.ImagePath)
	if err != nil {
		return ErrCreateImageProductFailed
	}

	err = s.repo.CreateImageProduct(ctx, request)
	if err != nil {
		return ErrCreateImageProductFailed
	}

	return nil
}

func (s *productService) CreateProduct(ctx context.Context, prod *ProductRequest) error {
	err := s.repo.CreateProduct(ctx, prod)
	if err != nil {
		return ErrCreateProductFailed
	}

	return nil
}
