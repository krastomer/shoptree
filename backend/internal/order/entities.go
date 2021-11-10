package order

import "time"

type Order struct {
	ID         int
	CustomerID int
	AddressID  int
	PaymentID  int
	Status     string
	Review     string
	CreatedAt  *time.Time
}

type OrderRepository interface{}

type OrderService interface{}
