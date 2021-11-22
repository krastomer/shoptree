package order

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service OrderService
}

var (
	ErrMsgProductID                   = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgProductUnavailable          = fiber.NewError(fiber.StatusNotAcceptable, "Product unavailble.")
	ErrMsgDeleteProductFromCartFailed = fiber.NewError(fiber.StatusInternalServerError, "Delete Product from cart failed.")
	ErrMsgWrongOwnerProduct           = fiber.NewError(fiber.StatusBadRequest, "Wrong owner product.")
	ErrMsgBlankCart                   = fiber.NewError(fiber.StatusNotFound, "Blank Cart.")
)

func NewOrderHandler(router fiber.Router, service OrderService) {
	handler := &handler{service: service}

	router.Use(CustomerMiddleware())
	router.Get("/", handler.getOrderPending)
	router.Post("/:productID", handler.addProductToCart)
	router.Delete("/:productID", handler.removeProductFromCart)
}

func (h *handler) addProductToCart(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID

	defer cancel()

	productID, err := c.ParamsInt("productID")
	if err != nil {
		return ErrMsgProductID
	}

	err = h.service.AddProductToCart(ctx, custID, productID)
	if err != nil {
		if err == ErrProductUnavailable {
			return ErrMsgProductUnavailable
		}
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add product to cart successfully.",
	})
}

func (h *handler) removeProductFromCart(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID

	defer cancel()

	productID, err := c.ParamsInt("productID")
	if err != nil {
		return ErrMsgProductID
	}

	err = h.service.RemoveProductFromCart(ctx, custID, productID)
	if err != nil {
		if err == ErrWrongOwnerProduct {
			return ErrMsgWrongOwnerProduct
		}
		if err == ErrBlankCart {
			return ErrMsgBlankCart
		}
		return ErrMsgDeleteProductFromCartFailed
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Remove product from cart successfully.",
	})
}

func (h *handler) getOrderPending(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID

	defer cancel()

	data, err := h.service.GetProductOnCart(ctx, custID)

	if err != nil {
		return ErrMsgBlankCart
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   data,
	})
}
