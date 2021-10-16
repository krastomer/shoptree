package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
)

const (
	msgFailedBodyParser    = "Require Username and Password."
	msgEmailInvalid        = "Email invalid."
	msgPasswordInvalid     = "Password invalid."
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
	r.Post("/register", handler.registerUser)
}

func (h *authHandler) loginUser(c *fiber.Ctx) error {
	request := &loginRequest{}
	if err := c.BodyParser(request); err != nil || request.Username == "" || request.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, msgFailedBodyParser)
	}

	token, err := h.service.Login(request.Username, request.Password)
	if err != nil {
		status := fiber.StatusInternalServerError
		msg := msgInternalServerError
		switch err {
		case errors.ErrEmailInvalid:
			status = fiber.StatusBadRequest
			msg = msgEmailInvalid
		case errors.ErrNotFoundUser:
			status = fiber.StatusNotFound
			msg = msgUserNotFound
		case errors.ErrPasswordInvlid:
			status = fiber.StatusBadRequest
			msg = msgPasswordInvalid
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

func (h *authHandler) registerUser(c *fiber.Ctx) error {
	request := &models.User{}
	if err := c.BodyParser(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, msgFailedBodyParser)
	}
	err := h.service.Register(request)
	if err != nil {
		switch err {
		case errors.ErrNotAuthorized:
			return fiber.NewError(fiber.StatusUnauthorized, msgUnavailableRole)
		default:
			return fiber.NewError(fiber.StatusInternalServerError, msgInternalServerError)
		}
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Registered successfully.",
	})
}
