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
	ErrMsgAddressBody            = fiber.NewError(fiber.StatusBadRequest, "Require Name, PhoneNumber, AddressLine, Country, State, City, District and PostalCode.")
)

func NewCustomerHandler(router fiber.Router, service CustomerService) {
	handler := &customerHandler{service: service}

	router.Post("/", handler.registerCustomer)
	router.Get("/:id", CustomerMiddleware(), handler.getCustomer)
	router.Get("/:id/address", CustomerMiddleware(), handler.getAddresses)
	router.Post("/:id/address", CustomerMiddleware(), handler.addAddress)
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
	id, err := h.permissionCustomer(c)
	if err != nil {
		return err
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

func (h *customerHandler) addAddress(c *fiber.Ctx) error {
	id, err := h.permissionCustomer(c)
	if err != nil {
		return err
	}

	address := &Address{}
	if err := c.BodyParser(address); err != nil {
		return ErrMsgAddressBody
	}

	err = h.service.AddAddress(id, address)
	if err != nil {
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add address successfully.",
	})
}

func (h *customerHandler) permissionCustomer(c *fiber.Ctx) (int, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return 0, ErrMsgCustomerIDBody
	}

	if c.Locals("currentUser").(*UserToken).ID != id {
		return 0, ErrMsgUnauthorizedID
	}
	return id, nil
}
