package customer

import (
	"github.com/gofiber/fiber/v2"
)

type customerHandler struct {
	service CustomerService
}

var (
	ErrMsgCustomerRequestBody    = fiber.NewError(fiber.StatusBadRequest, "Require Name, Email, Password and PhoneNumber.")
	ErrMsgRegisterCustomerFailed = fiber.NewError(fiber.StatusInternalServerError, "Registered Customer failed.")
	ErrMsgEmailUsed              = fiber.NewError(fiber.StatusBadRequest, "Email used.")
	ErrMsgPhoneUsed              = fiber.NewError(fiber.StatusBadRequest, "Phone used.")
	ErrMsgPhoneInvalid           = fiber.NewError(fiber.StatusBadRequest, "Phone invalid.")
	ErrMsgNameInvalid            = fiber.NewError(fiber.StatusBadRequest, "Name invalid.")
	ErrMsgPasswordInvalid        = fiber.NewError(fiber.StatusBadRequest, "Password invalid.")
	ErrMsgEmailInvalid           = fiber.NewError(fiber.StatusBadRequest, "Email invalid.")
)

func NewCustomerHandler(router fiber.Router, service CustomerService) {
	handler := &customerHandler{service: service}

	router.Post("/", handler.registerCustomer)
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
