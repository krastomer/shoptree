package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
)

const (
	msgFailedBodyParser    = "Require Username and Password."
	msgEmailInvalid        = "Email invalid."
	msgPasswordInvalid     = "Password invalid"
	msgUserNotFound        = "User not found."
	msgInternalServerError = "Internal Server Error."
)

type authHandler struct {
	service entities.AuthService
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(r fiber.Router, s entities.AuthService) {
	handler := &authHandler{service: s}

	r.Post("/login", handler.loginUser)
	r.Post("/logout", handler.logoutUser)
}

func (h *authHandler) loginUser(c *fiber.Ctx) error {
	request := &loginRequest{}
	if err := c.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, msgFailedBodyParser)
	}

	token, err := h.service.LoginCustomer(request.Username, request.Password)
	if err != nil {
		var status int
		var msg string
		switch err {
		case errors.ErrEmailInvalid:
			status = fiber.StatusBadRequest
			msg = msgEmailInvalid
		case errors.ErrEmailInvalid:
			status = fiber.StatusBadRequest
			msg = msgEmailInvalid
		case errors.ErrUserNotFound:
			status = fiber.StatusBadRequest
			msg = msgUserNotFound
		case errors.ErrPasswordInvalid:
			status = fiber.StatusBadRequest
			msg = msgPasswordInvalid
		default:
			status = fiber.StatusInternalServerError
			msg = msgInternalServerError
		}
		return fiber.NewError(status, msg)
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

func (h *authHandler) logoutUser(c *fiber.Ctx) error {
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
