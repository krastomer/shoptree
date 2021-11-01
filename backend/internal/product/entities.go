package product

import "time"

// TODO: status -> enum
type ProductResponse struct {
	Product
	ImagesID []int `json:"images_id"`
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
}

type ProductService interface {
	GetProductByID(int) (*ProductResponse, error)
}
