package customer

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type customerHandler struct {
	service CustomerService
}

var (
	ErrMsgCustomerRequestBody       = fiber.NewError(fiber.StatusBadRequest, "Require Name, Email, Password and PhoneNumber.")
	ErrMsgCustomerIDBody            = fiber.NewError(fiber.StatusBadRequest, "Require ID.")
	ErrMsgRegisterCustomerFailed    = fiber.NewError(fiber.StatusInternalServerError, "Registered Customer failed.")
	ErrMsgEmailUsed                 = fiber.NewError(fiber.StatusBadRequest, "Email used.")
	ErrMsgPhoneUsed                 = fiber.NewError(fiber.StatusBadRequest, "Phone used.")
	ErrMsgPhoneInvalid              = fiber.NewError(fiber.StatusBadRequest, "Phone invalid.")
	ErrMsgNameInvalid               = fiber.NewError(fiber.StatusBadRequest, "Name invalid.")
	ErrMsgPasswordInvalid           = fiber.NewError(fiber.StatusBadRequest, "Password invalid.")
	ErrMsgEmailInvalid              = fiber.NewError(fiber.StatusBadRequest, "Email invalid.")
	ErrMsgUnauthorizedID            = fiber.NewError(fiber.StatusUnauthorized, "You can't access to another ID.")
	ErrMsgAddressBody               = fiber.NewError(fiber.StatusBadRequest, "Require Name, PhoneNumber, AddressLine, Country, State, City, District and PostalCode.")
	ErrMsgAddressesCustomerNotFound = fiber.NewError(fiber.StatusNotFound, "Not found addresses.")
)

func NewCustomerHandler(router fiber.Router, service CustomerService) {
	handler := &customerHandler{service: service}

	router.Post("/", handler.registerCustomer)
	// 	router.Get("/", CustomerMiddleware(), handler.getCustomer)
	router.Get("/address", CustomerMiddleware(), handler.getAddressesCustomer)
	// 	router.Post("/address", CustomerMiddleware(), handler.addAddress)

	// 	router.Get("/orders", CustomerMiddleware(), handler.getOrders)
}

func (h *customerHandler) registerCustomer(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	request := &CustomerRequest{}

	if err := c.BodyParser(request); err != nil {
		return ErrMsgCustomerRequestBody
	}

	err := h.service.CreateNewCustomer(ctx, request)
	if err != nil {
		switch err {
		case ErrRegisterCustomerFailed:
			return ErrMsgRegisterCustomerFailed
		case ErrEmailUsed:
			return ErrMsgEmailUsed
		case ErrPhoneUsed:
			return ErrMsgPhoneUsed
		default:
			return fiber.ErrInternalServerError
		}
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Registered Customer successfully.",
	})
}

// func (h *customerHandler) getCustomer(c *fiber.Ctx) error {
// 	id := c.Locals("currentUser").(*UserToken).ID

// 	response, err := h.service.GetCustomer(id)
// 	if err != nil {
// 		return fiber.ErrInternalServerError
// 	}

// 	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 		"status": "success",
// 		"data":   response,
// 	})
// }

func (h *customerHandler) getAddressesCustomer(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	id := c.Locals("currentUser").(*UserToken).ID

	response, err := h.service.GetAddressesCustomer(ctx, id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if response == nil {
		return ErrMsgAddressesCustomerNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

// func (h *customerHandler) addAddress(c *fiber.Ctx) error {
// 	id := c.Locals("currentUser").(*UserToken).ID

// 	address := &Address{}
// 	if err := c.BodyParser(address); err != nil {
// 		return ErrMsgAddressBody
// 	}

// 	err := h.service.AddAddress(id, address)
// 	if err != nil {
// 		return nil
// 	}
// 	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 		"status":  "success",
// 		"message": "Add address successfully.",
// 	})
// }

// // TODO: Edit middleware
// // TODO: Add View
// func (h *customerHandler) getOrders(c *fiber.Ctx) error {
// 	id := c.Locals("currentUser").(*UserToken).ID

// 	response, err := h.service.GetOrders(id)
// 	if err != nil {
// 		return fiber.ErrInternalServerError
// 	}

// 	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 		"status": "success",
// 		"data":   response,
// 	})
// }
