package product

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service ProductService
}

var (
	ErrMsgProductNotFound = fiber.NewError(fiber.StatusNotFound, "Product Not found.")
	ErrMsgProductBody     = fiber.NewError(fiber.StatusBadRequest, "Product need name, scientific_name, description, price.")
)

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &handler{service: service}

	router.Get("/", handler.getProducts)
	router.Post("/", handler.createProduct)
}

func (h *handler) getProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response, err := h.service.GetProducts(ctx)
	if err != nil {
		return ErrMsgProductNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *handler) createProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	request := &Product{}

	if err := c.BodyParser(request); err != nil {
		return ErrMsgProductBody
	}

	err := h.service.CreateProduct(ctx, request)
	if err != nil {
		return ErrCreateProductFailed
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Create Product successfully",
	})
}
