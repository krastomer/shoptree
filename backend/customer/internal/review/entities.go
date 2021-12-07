package review

import (
	"context"
	"time"
)

type Order struct {
	ID         int        `dbq:"id" json:"id"`
	CustomerID int        `dbq:"customer_id" json:"-"`
	AddressID  int        `dbq:"address_id" json:"-"`
	Status     string     `dbq:"status" json:"status"`
	Review     string     `dbq:"review" json:"-"`
	CreatedAt  *time.Time `dbq:"created_at" json:"created_at"`
}

type Review struct {
	No     int    `json:"no"`
	Review string `json:"review"`
	Star   int    `json:"star"`
}

type ReviewRequest struct {
	Message string `json:"message"`
	Star    int    `json:"star"`
}

type ReviewRepository interface {
	GetOrdersDoneWithReview(context.Context) ([]*Order, error)
	GetOrdersDoneCustomer(context.Context, int) ([]*Order, error)
	UpdateOrderReview(context.Context, int, string) error
}

type ReviewService interface {
	GetReviews(context.Context) ([]*Review, error)
	GetOrdersDoneCustomer(context.Context, int) ([]*Order, error)
	UpdateOrderReview(context.Context, int, int, string, int) error
}
