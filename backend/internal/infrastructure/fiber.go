package infrastructure

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/krastomer/shoptree/backend/internal/auth"
	"github.com/krastomer/shoptree/backend/internal/customer"
	"github.com/krastomer/shoptree/backend/internal/order"
	"github.com/krastomer/shoptree/backend/internal/product"
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

	rdb, err := connectToRedis()
	if err != nil {
		panic(err)
	}

	rbConn, err := connectToRabbitMQ()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	defer rdb.Close()
	defer rbConn.Close()

	_ = rdb

	app := fiber.New(fiberConfig)

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/dashboard", monitor.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	authRepo := auth.NewAuthRepository(db)
	custRepo := customer.NewCustomerRepository(db)
	prodRepo := product.NewProductRepository(db)
	// ordeRepo := order.NewOrderRepository(db)

	ordeMessage := order.NewOrderMessageQueue(rbConn)

	authService := auth.NewAuthService(authRepo)
	custService := customer.NewCustomerService(custRepo)
	prodService := product.NewProductService(prodRepo)
	ordeService := order.NewOrderService(ordeMessage)

	auth.NewAuthHandler(v1.Group("/auth"), authService)
	customer.NewCustomerHandler(v1.Group("/customers"), custService)
	product.NewProductHandler(v1.Group("/products"), prodService)
	order.NewOrderHandler(v1.Group("/orders"), ordeService)

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
