package order

import (
	"context"
	"time"
)

type Order struct {
	ID             int        `dbq:"id" json:"id"`
	CustomerID     int        `dbq:"customer_id" json:"-"`
	AddressID      int        `dbq:"address_id" json:"-"`
	PaymentID      int        `dbq:"payment_id" json:"-"`
	Status         string     `dbq:"status" json:"status"`
	Review         string     `dbq:"review" json:"-"`
	CreatedAt      *time.Time `dbq:"created_at" json:"created_at"`
	Products       []*Product `dbq:"-" json:"products"`
	AddressProfile *Address   `dbq:"-" json:"address"`
}

type Product struct {
	ID             int        `dbq:"id" json:"id"`
	Name           string     `dbq:"name" json:"name"`
	ScientificName string     `dbq:"scientific_name" json:"scientific_name"`
	Description    string     `dbq:"description" json:"description"`
	Price          float32    `dbq:"price" json:"price"`
	CreatedAt      *time.Time `dbq:"created_at" json:"-"`
	ImagePath      int        `dbq:"-" json:"image_path"`
}

type ProductPending struct {
	CustomerID int        `dbq:"customer_id"`
	ProductID  int        `dbq:"product_id"`
	CreatedAt  *time.Time `dbq:"created_at"`
}

type Address struct {
	ID          int        `json:"-" dbq:"id"`
	CustomerID  int        `json:"-" dbq:"customer_id"`
	CreatedAt   *time.Time `json:"-" dbq:"created_at"`
	Name        string     `json:"name" dbq:"name"`
	PhoneNumber string     `json:"phone_number" dbq:"phone_number"`
	AddressLine string     `json:"address_line" dbq:"address_line"`
	Country     string     `json:"country" dbq:"country"`
	State       string     `json:"state" dbq:"state"`
	City        string     `json:"city" dbq:"city"`
	District    string     `json:"district" dbq:"district"`
	PostalCode  string     `json:"postal_code" dbq:"postal_code"`
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
	GetImageProductByID(context.Context, int) (int, error)
	GetAddressesCustomer(context.Context, int) ([]*Address, error)
	GetAddressCustomerByID(context.Context, int) (*Address, error)
	UpdateAddressOrder(context.Context, int, int) error
	UpdateStatusOrder(context.Context, string, int) error
}

type OrderService interface {
	AddProductToCart(context.Context, int, int) error
	RemoveProductFromCart(context.Context, int, int) error
	GetProductOnCart(context.Context, int) ([]*Product, error)
	UpdateAddressOrder(context.Context, int, int) error
	GetCart(context.Context, int) (*Order, error)
	ConfirmOrder(context.Context, int) error
}
