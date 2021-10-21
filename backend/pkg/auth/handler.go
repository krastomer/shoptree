package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrMsgFailedBodyParser    = fiber.NewError(fiber.StatusBadRequest, "Require Username and Password.")
	ErrMsgEmailInvalid        = fiber.NewError(fiber.StatusBadRequest, "Email invalid.")
	ErrMsgPasswordInvalid     = fiber.NewError(fiber.StatusBadRequest, "Password invalid.")
	ErrMsgUserNotFound        = fiber.NewError(fiber.StatusNotFound, "User not found.")
	ErrMsgInternalServerError = fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error.")
)

type authHandler struct {
	service AuthService
}

func NewAuthHandler(router fiber.Router, service AuthService) {
	handler := &authHandler{service: service}

	router.Post("/login", handler.login)
	router.Post("/logout", handler.logout)
	router.Post("/register", handler.register)
}

func (h *authHandler) login(c *fiber.Ctx) error {
	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	request := &loginRequest{}
	if err := c.BodyParser(request); err != nil || request.Username == "" || request.Password == "" {
		return ErrMsgFailedBodyParser
	}

	token, err := h.service.Login(request.Username, request.Password)
	if err != nil {
		switch err {
		case ErrEmailInvalid:
			return ErrMsgEmailInvalid
		case ErrPasswordInvalid:
			return ErrMsgPasswordInvalid
		case ErrUserNotFound:
			return ErrMsgUserNotFound
		default:
			return ErrMsgInternalServerError
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

func (h *authHandler) register(c *fiber.Ctx) error {
	return nil
}
