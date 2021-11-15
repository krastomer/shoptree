package order

import (
	"context"
	"time"
)

type OrderPending struct {
	ID         int
	CustomerID int
	Status     string
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
	CreatePendingOrder(context.Context, *OrderPending) error
}

type OrderService interface {
	CreateOrder()
}

type OrderMessageQueue interface {
	CreateQueue()
}
