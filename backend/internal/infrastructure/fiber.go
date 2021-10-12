package infrastructure

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/krastomer/shoptree/backend/internal/repositories/mariadb"
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

	app.Get("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi((c.Params("id")))
		fmt.Println(id)
		cust, err := custRepo.GetCustomer(id)
		if err != nil {
			return c.SendString("error")
		}
		return c.JSON(cust)
	})

	app.Listen("127.0.0.1:8080")
}
