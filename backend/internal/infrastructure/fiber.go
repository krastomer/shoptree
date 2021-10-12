package infrastructure

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/krastomer/shoptree/backend/internal/handlers"
	"github.com/krastomer/shoptree/backend/internal/repositories/mariadb"
	"github.com/krastomer/shoptree/backend/internal/services"
	"gorm.io/gorm"
)

var fiberConfig = fiber.Config{
	AppName:      "",
	Prefork:      true,
	ServerHeader: "Fiber",
}

func Run() {

	ch := make(chan int)

	var db *gorm.DB
	var err error
	go func() {
		for db == nil {
			db, err = connectToMariaDB()
			if err != nil {
				time.Sleep(5 * time.Second)
			} else {
				ch <- 1
			}
		}
	}()

	app := fiber.New(fiberConfig)
	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	<-ch

	custRepo := mariadb.NewCustomerRepo(db)

	authService := services.NewAuthService(custRepo)
	profileService := services.NewProfileService(custRepo)

	handlers.NewAuthHandler(v1.Group("/auth"), authService)
	handlers.NewProfileHandler(v1.Group("/profile"), profileService)

	app.Listen(":8080")
}
