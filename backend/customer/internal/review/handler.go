package review

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service ReviewService
}

var (
	ErrMsgNotFoundOrders = fiber.NewError(fiber.StatusNotFound, "No have orders with Review.")
)

func NewReviewHandler(router fiber.Router, service ReviewService) {
	handler := &handler{service: service}

	router.Get("/", handler.getReviews)
}

func (h *handler) getReviews(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response, err := h.service.GetReviews(ctx)
	if err != nil {
		return ErrMsgNotFoundOrders
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}
