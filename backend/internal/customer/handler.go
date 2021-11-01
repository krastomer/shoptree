package customer

import (
	"github.com/gofiber/fiber/v2"
)

type customerHandler struct {
	service CustomerService
}

var (
	ErrMsgCustomerRequestBody    = fiber.NewError(fiber.StatusBadRequest, "Require Name, Email, Password and PhoneNumber.")
	ErrMsgCustomerIDBody         = fiber.NewError(fiber.StatusBadRequest, "Require ID.")
	ErrMsgRegisterCustomerFailed = fiber.NewError(fiber.StatusInternalServerError, "Registered Customer failed.")
	ErrMsgEmailUsed              = fiber.NewError(fiber.StatusBadRequest, "Email used.")
	ErrMsgPhoneUsed              = fiber.NewError(fiber.StatusBadRequest, "Phone used.")
	ErrMsgPhoneInvalid           = fiber.NewError(fiber.StatusBadRequest, "Phone invalid.")
	ErrMsgNameInvalid            = fiber.NewError(fiber.StatusBadRequest, "Name invalid.")
	ErrMsgPasswordInvalid        = fiber.NewError(fiber.StatusBadRequest, "Password invalid.")
	ErrMsgEmailInvalid           = fiber.NewError(fiber.StatusBadRequest, "Email invalid.")
	ErrMsgUnauthorizedID         = fiber.NewError(fiber.StatusUnauthorized, "You can't access to another ID.")
)

func NewCustomerHandler(router fiber.Router, service CustomerService) {
	handler := &customerHandler{service: service}

	router.Post("/", handler.registerCustomer)
	router.Get("/:id", CustomerMiddleware(), handler.getCustomer)
	router.Get("/:id/address", CustomerMiddleware(), handler.getAddresses)
}

func (h *customerHandler) registerCustomer(c *fiber.Ctx) error {
	request := &CustomerRequest{}

	if err := c.BodyParser(request); err != nil {
		return ErrMsgCustomerRequestBody
	}

	err := h.service.RegisterCustomer(request)
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

func (h *customerHandler) getCustomer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgCustomerIDBody
	}

	if c.Locals("currentUser").(*UserToken).ID != id {
		return ErrMsgUnauthorizedID
	}

	response, err := h.service.GetCustomer(id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *customerHandler) getAddresses(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgCustomerIDBody
	}

	if c.Locals("currentUser").(*UserToken).ID != id {
		return ErrMsgUnauthorizedID
	}

	response, err := h.service.GetAddresses(id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}
