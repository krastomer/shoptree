package product

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service ProductService
}

var (
	ErrMsgIDProduct         = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgProductIDNotFound = fiber.NewError(fiber.StatusNotFound, "Product ID not found.")
)

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &handler{service: service}

	router.Get("/:id", CustomerMiddleware(), handler.getProductByID)
	// router.Get("/:id/images", handler.getImagesProductID)
	// router.Get("/images/:id", handler.getImageProductByID)
}

func (h *handler) getProductByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	if c.Locals("currentUser") != nil {
		ctx = context.WithValue(ctx, "currentUserID", c.Locals("currentUser").(*UserToken).ID)
	}

	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgIDProduct
	}

	response, err := h.service.GetProductByID(ctx, id)
	if err != nil {
		return ErrMsgProductIDNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
	return nil
}
