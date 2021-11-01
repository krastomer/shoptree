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

type ProductRequest struct {
	Name           string            `json:"name"`
	ScientificName string            `json:"scientific_name"`
	Price          float32           `json:"price"`
	Description    string            `json:"description"`
	Status         ProductStatusType `json:"status"`
}

type Product struct {
	ProductRequest
	ID        int       `json:"id"`
	ImagesID  []int     `json:"images_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductStatusType int

const (
	Undefined ProductStatusType = iota
	Unavailable
	Available
	Pending
	Purchased
)

func (t ProductStatusType) String() string {
	switch t {
	case Unavailable:
		return "Unavailable"
	case Available:
		return "Available"
	case Pending:
		return "Pending"
	case Purchased:
		return "Purchased"
	default:
		return "Undefined"
	}
}

type ProductRepository interface {
	GetProductByID(int) (*Product, error)
	GetProductImagesID(int) ([]int, error)

	GetProductImageByID(int) (string, error)

	CreateProduct(*ProductRequest) error
	CreateProductImagePath(*ProductImageRequest) error
}

type ProductService interface {
	GetProductByID(int) (*ProductResponse, error)
	GetProductImageByID(int) (string, error)

	AddProduct(*ProductRequest) error
	AddProductImage(*fiber.Ctx, *ProductImageRequest) error
}
