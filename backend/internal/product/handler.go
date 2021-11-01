package product

import (
	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	service ProductService
}

var (
	ErrMsgIDProduct            = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgIDProductImage       = fiber.NewError(fiber.StatusBadRequest, "Product image 'ID' should postive integer.")
	ErrMsgProductIDNotFound    = fiber.NewError(fiber.StatusNotFound, "Product ID not found.")
	ErrMsgProductImageNotFound = fiber.NewError(fiber.StatusNotFound, "Product image not found.")
)

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &productHandler{service: service}

	router.Get("/:id", handler.getProductByID)
	router.Get("/images/:id", handler.getProductImageByID)

	router.Post("/:id/images", StaffMiddleware(), handler.addProductImage)
}

func (h *productHandler) getProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgIDProduct
	}

	response, err := h.service.GetProductByID(id)
	if err != nil {
		return ErrMsgProductIDNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *productHandler) getProductImageByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgIDProduct
	}
	path, err := h.service.GetProductImageByID(id)
	if err != nil {
		return ErrProductImageNotFound
	}

	return c.Status(fiber.StatusOK).SendFile(path)
}

func (h *productHandler) addProductImage(c *fiber.Ctx) error {
	return nil
}
