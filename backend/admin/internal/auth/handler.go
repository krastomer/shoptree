package auth

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service AuthService
}

var (
	ErrMsgPasswordIncorrect      = fiber.NewError(fiber.StatusBadRequest, "Password incorrect.")
	ErrMsgUserRequestBody        = fiber.NewError(fiber.StatusBadRequest, "Require Username and Password.")
	ErrMsgEmailInvalid           = fiber.NewError(fiber.StatusBadRequest, "Email invalid.")
	ErrMsgEmailNotFound          = fiber.NewError(fiber.StatusNotFound, "Email not found.")
	ErrMsgPasswordInvalid        = fiber.NewError(fiber.StatusBadRequest, "Password invalid.")
	ErrMsgCustomerRequestBody    = fiber.NewError(fiber.StatusBadRequest, "Require Name, Email, Password and PhoneNumber.")
	ErrMsgRegisterCustomerFailed = fiber.NewError(fiber.StatusInternalServerError, "Registered Customer failed.")
	ErrMsgEmailUsed              = fiber.NewError(fiber.StatusBadRequest, "Email used.")
	ErrMsgPhoneUsed              = fiber.NewError(fiber.StatusBadRequest, "Phone used.")
	ErrMsgPhoneInvalid           = fiber.NewError(fiber.StatusBadRequest, "Phone invalid.")
)

func NewAuthHandler(router fiber.Router, service AuthService) {
	handler := &handler{service: service}

	router.Post("/login", handler.login)
	router.Post("/logout", handler.logout)
}

func (h *handler) login(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := &UserRequest{}

	if err := c.BodyParser(request); err != nil {
		return ErrMsgUserRequestBody
	}

	token, err := h.service.Login(ctx, request)
	if err != nil {
		switch err {
		case ErrEmailIncorrect:
			return ErrMsgEmailNotFound
		case ErrEmailInvalid:
			return ErrMsgEmailInvalid
		case ErrPasswordInvalid:
			return ErrMsgPasswordInvalid
		case ErrPasswordIncorrect:
			return ErrMsgPasswordIncorrect
		default:
			return fiber.ErrInternalServerError
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  token,
	})
}

func (h *handler) logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "loggedOut",
		Path:     "/",
		Expires:  time.Now().Add(10 * time.Second),
		Secure:   false,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Logged out successfully.",
	})
}
