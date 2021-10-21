package infrastructure

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/krastomer/shoptree/backend/pkg/auth"
	"github.com/krastomer/shoptree/backend/pkg/product"
	"github.com/krastomer/shoptree/backend/pkg/profile"
	"github.com/krastomer/shoptree/backend/pkg/repository/mariadb"
)

var fiberConfig = fiber.Config{
	AppName:      "SHOPTREE API",
	Prefork:      true,
	ServerHeader: "Fiber",
	ErrorHandler: errorHandler,
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
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

	authRepo := mariadb.NewAuthRepository(db)
	productRepo := mariadb.NewProductRepository(db)
	profileRepo := mariadb.NewProfileRepository(db)

	authService := auth.NewAuthService(authRepo)
	productService := product.NewProductService(productRepo)
	profileService := profile.NewProfileService(profileRepo)

	auth.NewAuthHandler(v1.Group("/auth"), authService)
	product.NewProductHandler(v1.Group("/products"), productService)
	profile.NewProfileHandler(v1.Group("/profile"), profileService)

	log.Fatal(app.Listen(":8080"))
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
