package main

import (
	"shoptree-backend-admin/internal/config"
	"shoptree-backend-admin/internal/infrastructure"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	infrastructure.Run()
}
