package search

import (
	"context"
	"time"
)

type Product struct {
	ID             int        `dbq:"id" json:"id"`
	Name           string     `dbq:"name" json:"name"`
	ScientificName string     `dbq:"scientific_name" json:"scientific_name"`
	Description    string     `dbq:"description" json:"description"`
	Price          float32    `dbq:"price" json:"price"`
	CreatedAt      *time.Time `dbq:"created_at" json:"-"`
	ImageID        int        `dbq:"-" json:"image_id"`
	Status         string     `dbq:"-" json:"status_product"`
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

type SearchRepository interface {
	GetCategoriesProduct(context.Context) ([]*CategoryProduct, error)
	GetProductsLike(context.Context, string) ([]*Product, error)
	GetImageProductByID(context.Context, int) (int, error)
	GetProductAvailableByID(context.Context, int) (*Product, error)
	GetProductPendingByID(context.Context, int) (*ProductPending, error)
}
type SearchService interface {
	GetCategories(context.Context) ([]*CategoryProduct, error)
	Search(context.Context, string, string) ([]*Product, error)
}
