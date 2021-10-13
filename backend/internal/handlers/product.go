package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/entities"
)

const (
	msgIDInvalid = "ID invalid"
)

type productHandler struct {
	service entities.ProductService
}

func NewProductHandler(r fiber.Router, s entities.ProductService) {
	handler := &productHandler{service: s}

	r.Get("/:id", handler.getProductByID)
}

func (h *productHandler) getProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, msgIDInvalid)
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, msgIDInvalid)
	}

	return c.JSON(&fiber.Map{
		"status": "success",
		"data":   product,
	})

}
