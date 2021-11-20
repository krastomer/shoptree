package product

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
	CreatedAt      *time.Time `dbq:"created_at" json:"created_at"`
}

type ProductRepository interface {
	GetProducts(context.Context) ([]*Product, error)
	CreateProduct(context.Context, *Product) error
}

type ProductService interface {
	GetProducts(context.Context) ([]*Product, error)
	CreateProduct(context.Context, *Product) error
}
