package search

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service SearchService
}

var (
	ErrMsgSearchNotFound = fiber.NewError(fiber.StatusNotFound, "Search Not Found.")
)

func NewSearchHandler(router fiber.Router, service SearchService) {
	handler := &handler{service: service}

	router.Get("/categories", handler.getCategories)
	router.Get("", handler.search)
}

func (h *handler) getCategories(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, _ := h.service.GetCategories(ctx)
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func (h *handler) search(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	productLike := c.Query("product")

	data, err := h.service.Search(ctx, "", productLike)
	if err != nil {
		return ErrMsgSearchNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   data,
	})
}
