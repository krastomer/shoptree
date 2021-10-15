package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/middleware"
)

type profileHandler struct {
	service entities.ProfileService
}

func NewProfileHandler(r fiber.Router, s entities.ProfileService) {
	handler := &profileHandler{service: s}

	r.Use(middleware.JWTMiddleware())
	r.Get("", handler.getProfile)
}

func (s *profileHandler) getProfile(c *fiber.Ctx) error {
	return c.JSON(c.Locals("currentUser"))
}
