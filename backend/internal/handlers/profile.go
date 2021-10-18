package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/middleware"
	"github.com/krastomer/shoptree/backend/internal/models"
)

const (
	msgUnavailableRole = "Can't access from this role"
	msgNotFoundAddress = "Not found address for user"
)

type profileHandler struct {
	service entities.ProfileService
}

func NewProfileHandler(r fiber.Router, s entities.ProfileService) {
	handler := &profileHandler{service: s}

	r.Use(middleware.JWTMiddleware())
	r.Get("", handler.getProfile)
	r.Get("/addresses", handler.getAddresses)
}

func (s *profileHandler) getProfile(c *fiber.Ctx) error {
	user := c.Locals("currentUser").(*models.User)
	if user.Level != "Customer" {
		return fiber.NewError(fiber.StatusUnauthorized, msgUnavailableRole)
	}
	custPro, err := s.service.GetProfile(user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, msgUserNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   custPro,
	})
}

func (s *profileHandler) getAddresses(c *fiber.Ctx) error {
	user := c.Locals("currentUser").(*models.User)
	if user.Level != "Customer" {
		return fiber.NewError(fiber.StatusUnauthorized, msgUnavailableRole)
	}
	addresses, err := s.service.GetAddresses(user.ID)
	if err != nil {
		switch err {
		case errors.ErrQueryNotFound:
			return fiber.NewError(fiber.StatusNotFound, msgUserNotFound)
		case errors.ErrNotFoundAddress:
			return fiber.NewError(fiber.StatusNotFound, msgNotFoundAddress)
		}
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   addresses,
	})
}
