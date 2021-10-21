package infrastructure

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/krastomer/shoptree/backend/pkg/auth"
	"github.com/krastomer/shoptree/backend/pkg/repository/mariadb"
)

var fiberConfig = fiber.Config{
	AppName:      "SHOPTREE API",
	Prefork:      true,
	ServerHeader: "Fiber",
	ErrorHandler: errorHandler,
}

func Run() {
	db, err := connectToMariaDB()
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiberConfig)

	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	authRepo := mariadb.NewAuthRepository(db)

	authService := auth.NewAuthService(authRepo)

	auth.NewAuthHandler(v1.Group("/auth"), authService)

	log.Fatal(app.Listen("127.0.0.1:8080"))
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(&fiber.Map{
		"status":  "fail",
		"message": err.Error(),
	})
}
