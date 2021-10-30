package main

import (
	"github.com/krastomer/shoptree/backend/pkg/config"
	"github.com/krastomer/shoptree/backend/pkg/infrastructure"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	_ = config

	infrastructure.Run(config)
}
