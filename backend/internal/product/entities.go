package product

import (
	"context"
	"mime/multipart"
	"time"
)

type ProductImageRequest struct {
	ID    int
	Image *multipart.FileHeader
	Path  string
}

type ProductRequest struct {
	Name           string  `json:"name"`
	ScientificName string  `json:"scientific_name"`
	Price          float32 `json:"price"`
	Description    string  `json:"description"`
}

type Product struct {
	ID          int        `dbq:"id" json:"id"`
	Name        string     `dbq:"name" json:"name"`
	Description string     `dbq:"description" json:"description"`
	Price       float32    `dbq:"price" json:"price"`
	CreatedAt   *time.Time `dbq:"created_at" json:"-"`
	Categories  []*CategoryProduct
}

type CategoryProduct struct {
	ID      int    `dbq:"id" json:"id"`
	Product int    `dbq:"product_id" json:"-"`
	Name    string `dbq:"name" json:"name"`
}

type ProductRepository interface {
	GetProductByID(context.Context, int) (*Product, error)
	GetCategoriesProduct(context.Context, int) ([]*CategoryProduct, error)
	// GetProductImagesID(int) ([]int, error)

	// GetProductImageByID(int) (string, error)

	// CreateProduct(*ProductRequest) error
	// CreateProductImagePath(*ProductImageRequest) error
}

type ProductService interface {
	GetProductByID(context.Context, int) (*Product, error)
	// GetProductImageByID(int) (string, error)

	// AddProduct(*ProductRequest) error
	// AddProductImage(*fiber.Ctx, *ProductImageRequest) error
}
