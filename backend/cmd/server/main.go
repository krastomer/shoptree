package main

import (
	"github.com/krastomer/shoptree/backend/internal/config"
	"github.com/krastomer/shoptree/backend/internal/customer"
	"github.com/krastomer/shoptree/backend/internal/infrastructure"
)

func main() {
	config.LoadConfig(".")
	// infrastructure.Run()
	db := infrastructure.RunDB()

	custRepo := customer.NewCustomerRepository(db)
	// prodRepo := product.NewProductRepository(db)
}
