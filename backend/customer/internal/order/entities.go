package order

import "time"

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

type OrderRepository interface{}

type OrderService interface{}
