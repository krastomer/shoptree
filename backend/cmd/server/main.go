package main

import (
	"github.com/krastomer/shoptree/backend/internal/config"
	"github.com/krastomer/shoptree/backend/internal/infrastructure"
)

func main() {
	// config, err := config.LoadConfig(".")
	// if err != nil {
	// 	panic(err)
	// }

	// infrastructure.Run(config)

	config.LoadConfig(".")
	infrastructure.Run()
}
