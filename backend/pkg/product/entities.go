package product

import "time"

type Product struct {
	ID             uint32
	Name           string
	ScientificName string
	Price          float32
	Description    string
	Status         ProductStatusType
	CreatedAt      time.Time
}

type ProductStatusType string

const (
	Unavailable ProductStatusType = "Unavailable"
	Available   ProductStatusType = "Available"
	Pending     ProductStatusType = "Pending"
	Purchased   ProductStatusType = "Purchased"
)

type ProductRepository interface {
	GetProductByID(uint32) (*Product, error)
}

type ProductService interface {
	GetProductByID(uint32) (*Product, error)
}
