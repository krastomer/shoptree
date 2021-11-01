package product

import (
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
)

// TODO: status -> enum
type ProductResponse struct {
	Product
	ImagesID []int `json:"images_id"`
}

type ProductImageRequest struct {
	ID    int
	Image *multipart.FileHeader
	Path  string
}

type Product struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	ScientificName string    `json:"scientific_name"`
	Price          float32   `json:"price"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	ImagesID       []int     `json:"images_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type ProductRepository interface {
	GetProductByID(int) (*Product, error)
	GetProductImagesID(int) ([]int, error)

	GetProductImageByID(int) (string, error)

	AddProductImagePath(*ProductImageRequest) error
}

type ProductService interface {
	GetProductByID(int) (*ProductResponse, error)
	GetProductImageByID(int) (string, error)

	AddProductImage(*fiber.Ctx, *ProductImageRequest) error
}
