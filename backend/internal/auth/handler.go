package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service AuthService
}

var (
	ErrMsgPasswordIncorrect = fiber.NewError(fiber.StatusBadRequest, "Password incorrect.")
	ErrMsgUserRequestBody   = fiber.NewError(fiber.StatusBadRequest, "Require Username and Password.")
	ErrMsgEmailInvalid      = fiber.NewError(fiber.StatusBadRequest, "Email invalid.")
	ErrMsgEmailNotFound     = fiber.NewError(fiber.StatusNotFound, "Email not found.")
	ErrMsgPasswordInvalid   = fiber.NewError(fiber.StatusBadRequest, "Password invalid.")
)

func NewAuthHandler(router fiber.Router, service AuthService) {
	handler := &authHandler{service: service}

	router.Post("/login", handler.login)
	router.Post("/logout", handler.logout)
}

func (h *authHandler) login(c *fiber.Ctx) error {
	request := &UserRequest{}

	if err := c.BodyParser(request); err != nil {
		return ErrMsgUserRequestBody
	}

	token, err := h.service.Login(request)
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

func (h *authHandler) logout(c *fiber.Ctx) error {
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
