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

	router.Use(customerMiddleware())

	router.Get("/", handler.getCustomerProfile)
	router.Get("/addresses", handler.getAddresses)
	router.Post("/addresses", handler.createAddressCustomer)
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

func (h *handler) createAddressCustomer(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID
	defer cancel()

	request := &AddressRequest{}
	if err := c.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	err := h.service.CreateAddressCustomer(ctx, custID, request)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Create Address customer successfully.",
	})
}
