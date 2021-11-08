package product

import (
	"context"
	"errors"
)

type productService struct {
	repo ProductRepository
}

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductImageNotFound = errors.New("product image not found")
	ErrAddProductFailed     = errors.New("add product failed")
	ErrProductStatus        = errors.New("product status not match")
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
	// images, _ := s.repo.GetProductImagesID(id)

	// response := &ProductResponse{
	// 	Product:  *product,
	// 	ImagesID: images,
	// }

	return product, nil
}

// func (s *productService) GetProductImageByID(id int) (string, error) {
// 	path, err := s.repo.GetProductImageByID(id)
// 	if err != nil {
// 		return "", ErrProductImageNotFound
// 	}
// 	return path, nil
// }

// func (s *productService) AddProduct(request *ProductRequest) error {
// 	if !request.StatusValid() {
// 		return ErrProductStatus
// 	}

// 	err := s.repo.CreateProduct(request)
// 	if err != nil {
// 		return ErrAddProductFailed
// 	}
// 	return nil
// }

// func (s *productService) AddProductImage(c *fiber.Ctx, request *ProductImageRequest) error {
// 	uniqueId := uuid.New()
// 	request.Path = fmt.Sprintf("%s/%s.jpg", viper.GetString("DIRECTORY_PRODUCT"), uniqueId)

// 	err := s.repo.CreateProductImagePath(request)

// 	if err != nil {
// 		return ErrProductNotFound
// 	}

// 	c.SaveFile(request.Image, request.Path)
// 	return nil
// }
