package infrastructure

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/krastomer/shoptree/backend/internal/handlers"
	"github.com/krastomer/shoptree/backend/internal/repositories/mariadb"
	"github.com/krastomer/shoptree/backend/internal/services"
)

var fiberConfig = fiber.Config{
	AppName:      "SHOPTREE API",
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

	app.Get("/dashboard", monitor.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	custRepo := mariadb.NewCustomerRepo(db)
	custProRepo := mariadb.NewCustomerProfileRepo(db)
	emplRepo := mariadb.NewEmployeeRepo(db)

	authService := services.NewAuthService(custRepo, emplRepo)
	profileService := services.NewProfileService(custProRepo)

	handlers.NewAuthHandler(v1.Group("/auth"), authService)
	handlers.NewProfileHandler(v1.Group("/profile"), profileService)

	log.Fatal(app.Listen("127.0.0.1:8080"))

}
