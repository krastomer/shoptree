package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
)

const (
	msgUserExisted = "User existed."
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
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": msgUserExisted,
		})
	}

	err := h.service.CreateProfile(request)
	if err != nil {
		if err == errors.ErrUserExisted {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  "fail",
				"message": msgUserExisted,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": msgInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Register successfully.",
	})
}
