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
	ErrMsgOrderNotFound   = fiber.NewError(fiber.StatusNotFound, "Order not found.")
	ErrMsgNeedIDAddress   = fiber.NewError(fiber.StatusBadRequest, "Need ID Address.")
	ErrMsgNeedOrderID     = fiber.NewError(fiber.StatusBadRequest, "Need OrderID.")
)

func NewCustomerHandler(router fiber.Router, service CustomerService) {
	handler := &handler{service: service}

	router.Use(customerMiddleware())

	router.Get("/", handler.getCustomerProfile)
	router.Get("/addresses", handler.getAddresses)
	router.Post("/addresses", handler.createAddressCustomer)
	router.Get("/addresses/:id", handler.getAddressByID)
	router.Delete("/addresses/:id", handler.deleteAddressByID)
	router.Get("/orders/:id/image", handler.getPaymentImageByOrderID)
	router.Get("/orders", handler.getOrderCustomer)
}

func (h *handler) getOrderCustomer(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID

	defer cancel()

	data, err := h.service.GetOrdersCustomer(ctx, custID)
	if err != nil {
		return ErrMsgOrderNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   data,
	})
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

func (h *handler) getAddressByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID
	defer cancel()

	addressID, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgNeedIDAddress
	}

	response, err := h.service.GetAddressCustomerByID(ctx, custID, addressID)
	if err != nil {
		if err == ErrAddressNotFound {
			return ErrMsgNotFoundAddress
		}
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *handler) deleteAddressByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID
	defer cancel()

	addressID, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgNeedIDAddress
	}

	err = h.service.DeleteAddressCustomer(ctx, custID, addressID)
	if err != nil {
		if err == ErrAddressNotFound {
			return ErrMsgNotFoundAddress
		}
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Delete Address Customer successfully.",
	})
}

func (h *handler) getPaymentImageByOrderID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID
	defer cancel()

	orderID, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgNeedOrderID
	}

	path, err := h.service.GetPaymentSlip(ctx, custID, orderID)
	if err != nil {
		return ErrMsgOrderNotFound
	}

	return c.Status(fiber.StatusOK).SendFile(path)
}
