package product

import "errors"

var (
	ErrProductNotFound  = errors.New("product not found")
	ErrAddProductFailed = errors.New("add product failed")
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
