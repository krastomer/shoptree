package review

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service ReviewService
}

var (
	ErrMsgNotFoundOrders   = fiber.NewError(fiber.StatusNotFound, "No have orders with Review.")
	ErrMsgBadReviewRequest = fiber.NewError(fiber.StatusNotAcceptable, "We need star in range 0-5 and message.")
)

func NewReviewHandler(router fiber.Router, service ReviewService) {
	handler := &handler{service: service}

	router.Get("/", handler.getReviews)
	router.Get("/customer", customerMiddleware(), handler.getOrdersDoneCustomer)
	router.Post("/:id", customerMiddleware(), handler.updateOrderReview)
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

func (h *handler) getOrdersDoneCustomer(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID
	defer cancel()

	response, err := h.service.GetOrdersDoneCustomer(ctx, custID)
	if err != nil {
		return ErrMsgNotFoundOrders
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *handler) updateOrderReview(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID
	defer cancel()

	orderID, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgBadReviewRequest
	}

	var request ReviewRequest
	if err := c.BodyParser(&request); err != nil {
		return ErrMsgBadReviewRequest
	}

	err = h.service.UpdateOrderReview(ctx, orderID, custID, request.Message, request.Star)
	if err != nil {
		if err == ErrDataReviewBad {
			return ErrMsgBadReviewRequest
		}
		if err == ErrNotFoundOrders {
			return ErrMsgNotFoundOrders
		}
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Update Review successfully.",
	})
}
