package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
)

const (
	msgDataIncorrect = "Body not correct."
	msgPhoneInvalid  = "Phone existed."
)

type profileHandler struct {
	service entities.ProfileService
}

func NewProfileHandler(r fiber.Router, s entities.ProfileService) {
	handler := &profileHandler{service: s}

	r.Post("/register", handler.register)
}

func (h *profileHandler) register(c *fiber.Ctx) error {
	request := &models.CustomerProfile{}

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, msgDataIncorrect)
	}

	err := h.service.CreateProfile(request)
	if err != nil {
		if err == errors.ErrUserExisted {
			return fiber.NewError(fiber.StatusBadRequest, msgEmailInvalid)
		}
		if err == errors.ErrPhoneNumberInvalid {
			return fiber.NewError(fiber.StatusBadRequest, msgPhoneInvalid)
		}
		return fiber.ErrInternalServerError
	}
	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "Register successfully.",
	})
}
