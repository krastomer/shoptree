package product

import "time"

type Product struct {
	ID             uint32            `json:"id"`
	Name           string            `json:"name"`
	ScientificName string            `json:"scientific_name"`
	Price          float32           `json:"price"`
	Description    string            `json:"description"`
	Status         ProductStatusType `json:"status"`
	CreatedAt      time.Time         `json:"created_at"`
}

type ProductStatusType string

const (
	Unavailable ProductStatusType = "Unavailable"
	Available   ProductStatusType = "Available"
	Pending     ProductStatusType = "Pending"
	Purchased   ProductStatusType = "Purchased"
)

type User struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Level       string `json:"level"`
}

type ProductRepository interface {
	AddProduct(*Product) error
	GetProductByID(uint32) (*Product, error)
	GetProducts() ([]*Product, error)
}

type ProductService interface {
	AddProduct(*Product) error
	GetProductByID(uint32) (*Product, error)
	GetProducts([]uint32) ([]*Product, error)
}
