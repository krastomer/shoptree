package product

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service ProductService
}

var (
	ErrMsgProductID         = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgProductIDNotFound = fiber.NewError(fiber.StatusNotFound, "Product ID not found.")
)

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &handler{service: service}

	router.Get("/:id", softCustomerMiddleware(), handler.getProductByID)
	router.Get("/images/:id", handler.getImageProductByID)
}

func (h *handler) getProductByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	var custID int
	if c.Locals("currentUser") != nil {
		cust := c.Locals("currentUser").(*UserToken)
		if cust.Level == "Customer" {
			custID = cust.ID
		}
	}

	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgProductID
	}

	response, err := h.service.GetProductByID(ctx, id, custID)
	if err != nil {
		return ErrMsgProductIDNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *handler) getImageProductByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgProductID
	}
	path, err := h.service.GetImageProductByID(ctx, id)
	if err != nil {
		return ErrProductImageNotFound
	}

	return c.Status(fiber.StatusOK).SendFile(path)
}
