package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/krastomer/shoptree/backend/internal/auth"
	"github.com/krastomer/shoptree/backend/internal/customer"
	"github.com/spf13/viper"
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
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/dashboard", monitor.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	authRepo := auth.NewAuthRepository(db)
	custRepo := customer.NewCustomerRepository(db)
	// prodRepo := product.NewProductRepository(db)

	authService := auth.NewAuthService(authRepo)
	custService := customer.NewCustomerService(custRepo)
	// prodService := product.NewProductService(prodRepo)

	auth.NewAuthHandler(v1.Group("/auth"), authService)
	customer.NewCustomerHandler(v1.Group("/customers"), custService)
	// product.NewProductHandler(v1.Group("/products"), prodService)

	log.Fatal(app.Listen(viper.GetString("APP_PORT")))
}

func RunDB() {
	db, err := connectToMariaDB()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// newCust := &customer.CustomerRequest{
	// 	Name:        "test test",
	// 	Email:       "test@example.com",
	// 	Password:    "Test1234",
	// 	PhoneNumber: "0000000000",
	// }

	repo := customer.NewCustomerRepository(db)
	// err = repo.CreateCustomer(ctx, newCust)

	cust, _ := repo.GetCustomerByEmail(ctx, "krastomer@gmail.com")
	fmt.Println(cust)
	cust, _ = repo.GetCustomerByPhone(ctx, "028702739")

	fmt.Println(cust)

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
