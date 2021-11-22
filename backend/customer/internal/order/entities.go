package order

import (
	"context"
	"time"
)

type Order struct {
	ID         int        `dbq:"id" json:"id"`
	CustomerID int        `dbq:"customer_id"`
	AddressID  int        `dbq:"address_id"`
	PaymentID  int        `dbq:"payment_id"`
	Status     string     `dbq:"status"`
	Review     string     `dbq:"review"`
	CreatedAt  *time.Time `dbq:"created_at"`
}

type Product struct {
	ID             int        `dbq:"id" json:"id"`
	Name           string     `dbq:"name" json:"name"`
	ScientificName string     `dbq:"scientific_name" json:"scientific_name"`
	Description    string     `dbq:"description" json:"description"`
	Price          float32    `dbq:"price" json:"price"`
	CreatedAt      *time.Time `dbq:"created_at" json:"-"`
}

type ProductPending struct {
	CustomerID int        `dbq:"customer_id"`
	ProductID  int        `dbq:"product_id"`
	CreatedAt  *time.Time `dbq:"created_at"`
}

type CurrentUserIDType string

const CurrentUserID CurrentUserIDType = "currentUserID"

type OrderRepository interface {
	CreateOrderPending(context.Context, int) error
	GetOrderPendingByCustomerID(context.Context, int) (*Order, error)
	AddProductToOrder(context.Context, int, int) error
	GetAvailableProductByID(context.Context, int) (*Product, error)
	DeleteProductFromOrder(context.Context, int) error
	GetProductPendingByCustomerID(context.Context, int) ([]*ProductPending, error)
	GetProductByID(context.Context, int) (*Product, error)
}

type OrderService interface {
	AddProductToCart(context.Context, int, int) error
	RemoveProductFromCart(context.Context, int, int) error
	GetProductOnCart(context.Context, int) ([]*Product, error)
}
