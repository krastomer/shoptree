package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/entities"
)

type productHandler struct {
	service entities.ProductService
}

func NewProductHandler(r fiber.Router, s entities.ProductService) {
	handler := &productHandler{service: s}

	r.Get("/:id", handler.getProduct)
	r.Post("/")
}

func (s *productHandler) getProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "ID sould integer.")
	}
	res, err := s.service.GetProduct(uint32(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "ID not found")
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   res,
	})
}
