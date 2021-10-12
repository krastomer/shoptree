package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/krastomer/shoptree/backend/internal/handlers"
	"github.com/krastomer/shoptree/backend/internal/repositories/mariadb"
	"github.com/krastomer/shoptree/backend/internal/services"
)

var fiberConfig = fiber.Config{
	AppName:      "",
	Prefork:      true,
	ServerHeader: "Fiber",
}

func Run() {

	db, err := connectToMariaDB()
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiberConfig)
	app.Use(logger.New())
	app.Use(recover.New())

	custRepo := mariadb.NewCustomerRepo(db)
	custService := services.NewCustomerService(custRepo)

	handlers.NewCustomerHandler(app.Group("/api/v1/auth"), custService)

	app.Listen("127.0.0.1:8080")
}
