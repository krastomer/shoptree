package product

import (
	"context"
	"time"
)

type Product struct {
	ID          int                `dbq:"id" json:"id"`
	Name        string             `dbq:"name" json:"name"`
	Description string             `dbq:"description" json:"description"`
	Price       float32            `dbq:"price" json:"price"`
	CreatedAt   *time.Time         `dbq:"created_at" json:"-"`
	Categories  []*CategoryProduct `json:"catergory"`
	ImagesID    []int              `dbq:"-" json:"image_id"`

	Status string `dbq:"-" json:"status_product"`
}

type ProductPending struct {
	CustomerID int        `dbq:"customer_id"`
	ProductID  int        `dbq:"product_id"`
	CreatedAt  *time.Time `dbq:"created_at"`
}

type CategoryProduct struct {
	ID      int    `dbq:"id" json:"id"`
	Product int    `dbq:"product_id" json:"-"`
	Name    string `dbq:"name" json:"name"`
}

type ImageProduct struct {
	ID        int        `dbq:"id"`
	ProductID int        `dbq:"product_id"`
	ImagePath string     `dbq:"image_path"`
	CreatedAt *time.Time `dbq:"created_at"`
}

type ProductRepository interface {
	GetProductByID(context.Context, int) (*Product, error)
	GetImagesProductID(context.Context, int) ([]*ImageProduct, error)
	GetCategoriesProduct(context.Context, int) ([]*CategoryProduct, error)
	GetProductAvailableByID(context.Context, int) (*Product, error)
	GetProductPendingByID(context.Context, int) (*ProductPending, error)
}

type ProductService interface {
	GetProductByID(context.Context, int) (*Product, error)
}
