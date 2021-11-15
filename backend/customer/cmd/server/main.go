package main

import (
	"shoptree-backend-customer/internal/config"
	"shoptree-backend-customer/internal/infrastructure"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	infrastructure.Run()
}
