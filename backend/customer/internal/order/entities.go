package order

type Order struct {
	ID         int
	CustomerID int
	AddressID  int
}

type OrderRepository interface{}

type OrderService interface{}
