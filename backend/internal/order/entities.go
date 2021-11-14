package order

import (
	"context"
	"time"
)

type ProductPending struct {
	ID         int
	CustomerID int
	ExpiresAt  time.Time
}

type Order struct {
	ID         int
	CustomerID int
	AddressID  int
	PaymentID  int
	Status     string
	Review     string
	CreatedAt  *time.Time
}

type OrderRepository interface {
	CreatePendingProduct(context.Context, *ProductPending) error
}

type OrderService interface{}
