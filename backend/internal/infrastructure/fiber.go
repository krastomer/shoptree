package infrastructure

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/auth"
)

var fiberConfig = fiber.Config{
	AppName: "SHOPTREE API",
	// Prefork:      true,
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

	authRepo := auth.NewAuthRepository(db)
	cust, _ := authRepo.GetCustomerByEmail("krastomer@gmail.com")
	fmt.Println(cust)
	empl, _ := authRepo.GetEmployeeByEmail("kasama.tsw@shoptree.com")
	fmt.Println(empl)

	// app := fiber.New(fiberConfig)

	// app.Use(logger.New())
	// app.Use(cors.New())
	// app.Use(recover.New())

	// app.Get("/dashboard", monitor.New())

	// api := app.Group("/api")
	// v1 := api.Group("/v1")

	// _ = v1
	// _ = db

	// log.Fatal(app.Listen(viper.GetString("APP_PORT")))
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
