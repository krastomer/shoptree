package product

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ImageProduct struct {
	ID        int                   `dbq:"id"`
	ProductID int                   `dbq:"product_id"`
	Image     *multipart.FileHeader `dbq:"-"`
	ImagePath string                `dbq:"image_path"`
	CreatedAt *time.Time            `dbq:"created_at"`
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
	ImagesID    []int `dbq:"-" json:"images_id"`
}

type CategoryProduct struct {
	ID      int    `dbq:"id" json:"id"`
	Product int    `dbq:"product_id" json:"-"`
	Name    string `dbq:"name" json:"name"`
}

type ProductRepository interface {
	GetProductByID(context.Context, int) (*Product, error)
	GetCategoriesProduct(context.Context, int) ([]*CategoryProduct, error)
	GetImagesProductID(context.Context, int) ([]*ImageProduct, error)

	GetImageProductByID(context.Context, int) (string, error)

	// CreateProduct(*ProductRequest) error
	CreateImageProduct(context.Context, *ImageProduct) error
}

type ProductService interface {
	GetProductByID(context.Context, int) (*Product, error)
	GetImageProductByID(context.Context, int) (string, error)
	GetImagesProductID(context.Context, int) ([]int, error)

	// AddProduct(*ProductRequest) error
	CreateImageProduct(context.Context, *fiber.Ctx, *ImageProduct) error
}
