package customer

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service CustomerService
}

var (
	ErrMsgNotFoundAddress = fiber.NewError(fiber.StatusNotFound, "Not found address.")
)

func NewCustomerHandler(router fiber.Router, service CustomerService) {
	handler := &handler{service: service}

	router.Use(CustomerMiddleware())

	router.Get("/", handler.getCustomerProfile)
	router.Get("/addresses", handler.getAddresses)
}

func (h *handler) getAddresses(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID

	defer cancel()

	data, err := h.service.GetAddresses(ctx, custID)
	if err != nil {
		return ErrMsgNotFoundAddress
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func (h *handler) getCustomerProfile(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID

	defer cancel()

	data, err := h.service.GetCustomerProfile(ctx, custID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   data,
	})
}
