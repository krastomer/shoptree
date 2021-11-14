package order

import "github.com/gofiber/fiber/v2"

type handler struct {
	service OrderService
}

func NewOrderHandler(router fiber.Router, service OrderService) {
	handler := &handler{service: service}

	router.Get("/", handler.NewQueue)
}

func (h *handler) NewQueue(c *fiber.Ctx) error {
	return nil
}
