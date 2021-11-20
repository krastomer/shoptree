package infrastructure

import (
	"log"
	"shoptree-backend-admin/internal/auth"
	"shoptree-backend-admin/internal/product"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func Run() {
	db, err := connectToMariaDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fiberConfig := fiber.Config{
		AppName:      "SHOPTREE API",
		Prefork:      true,
		ServerHeader: "Fiber",
		ErrorHandler: errorHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	app := fiber.New(fiberConfig)

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	authRepo := auth.NewAuthRepository(db)
	productRepo := product.NewProductRepository(db)

	authService := auth.NewAuthService(authRepo)
	productService := product.NewProductService(productRepo)

	auth.NewAuthHandler(v1.Group("/auth"), authService)
	product.NewProductHandler(v1.Group("/products", AdminMiddleware()), productService)

	log.Fatal(app.Listen(viper.GetString("APP_PORT")))
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(&fiber.Map{
		"status":  "error",
		"message": err.Error(),
	})
}
