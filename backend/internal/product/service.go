package product

import "errors"

type productService struct {
	repo ProductRepository
}

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductImageNotFound = errors.New("product image not found")
)

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProductByID(id int) (*ProductResponse, error) {
	product, err := s.repo.GetProductByID(id)

	if err != nil {
		return nil, ErrProductNotFound
	}

	images, _ := s.repo.GetProductImagesID(id)

	response := &ProductResponse{
		Product:  *product,
		ImagesID: images,
	}

	return response, nil
}

func (s *productService) GetProductImageByID(id int) (string, error) {
	path, err := s.repo.GetProductImageByID(id)
	if err != nil {
		return "", ErrProductImageNotFound
	}
	return path, nil
}
