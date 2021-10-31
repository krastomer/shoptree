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
)

func NewAuthHandler(router fiber.Router, service AuthService) {
	handler := &authHandler{service: service}

	router.Post("/login", handler.login)
}

func (h *authHandler) login(c *fiber.Ctx) error {
	request := &UserRequest{}

	if err := c.BodyParser(request); err != nil {
		return fiber.ErrInternalServerError
	}

	token, err := h.service.Login(request)
	if err != nil {
		switch err {
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
