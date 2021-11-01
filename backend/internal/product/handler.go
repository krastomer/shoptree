package product

import "github.com/gofiber/fiber/v2"

type productHandler struct {
	service ProductService
}

var (
	ErrMsgIDProduct         = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgProductIDNotFound = fiber.NewError(fiber.StatusNotFound, "Product ID not found.")
)

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &productHandler{service: service}

	router.Get("/:id", handler.getProductByID)
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
