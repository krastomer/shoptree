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
	ImagesID       []int      `dbq:"-" json:"image_id"`
	Status         string     `dbq:"-" json:"status_product"`
}

type CategoryProduct struct {
	ID      int    `dbq:"id" json:"id"`
	Product int    `dbq:"product_id" json:"-"`
	Name    string `dbq:"name" json:"name"`
}

type SearchRepository interface {
	GetCategoriesProduct(context.Context) ([]*CategoryProduct, error)
	GetProductsLike(context.Context, string) ([]*Product, error)
}
type SearchService interface {
	GetCategories(context.Context) ([]*CategoryProduct, error)
	Search(context.Context, string, string) ([]*Product, error)
}
