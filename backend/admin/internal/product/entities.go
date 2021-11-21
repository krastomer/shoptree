package product

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID             int        `dbq:"id" json:"id"`
	Name           string     `dbq:"name" json:"name"`
	ScientificName string     `dbq:"scientific_name" json:"scientific_name"`
	Description    string     `dbq:"description" json:"description"`
	Price          float32    `dbq:"price" json:"price"`
	CreatedAt      *time.Time `dbq:"created_at" json:"created_at"`
}

type ImageProduct struct {
	ID        int                   `dbq:"id"`
	ProductID int                   `dbq:"product_id"`
	Image     *multipart.FileHeader `dbq:"-"`
	ImagePath string                `dbq:"image_path"`
	CreatedAt *time.Time            `dbq:"created_at"`
}

type ProductRepository interface {
	GetProducts(context.Context) ([]*Product, error)
	GetProductByID(context.Context, int) (*Product, error)
	CreateProduct(context.Context, *Product) error
	UpdateProduct(context.Context, *Product) error
	DeleteProductByID(context.Context, int) error
	CreateImageProduct(context.Context, *ImageProduct) error
}

type ProductService interface {
	GetProducts(context.Context) ([]*Product, error)
	CreateProduct(context.Context, *Product) error
	UpdateProduct(context.Context, *Product) error
	DeleteProduct(context.Context, int) error
	CreateImageProduct(context.Context, *fiber.Ctx, *ImageProduct) error
}
