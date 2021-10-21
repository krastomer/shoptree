package product

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrMsgBadIDProduct      = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgIDProductNotFound = fiber.NewError(fiber.StatusNotFound, "Product not found.")
)

type productHandler struct {
	service ProductService
}

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &productHandler{service: service}

	router.Get("/:id", handler.getProductByID)
}

func (h *productHandler) getProductByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return ErrMsgBadIDProduct
	}

	response, err := h.service.GetProductByID(uint32(id))
	if err != nil {
		return ErrMsgIDProductNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})

}
