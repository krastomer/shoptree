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
	ErrMsgMissingAddressID            = fiber.NewError(fiber.StatusBadRequest, "Missing AddressID")
	ErrMsgAddressNotFound             = fiber.NewError(fiber.StatusNotFound, "Address not found.")
	ErrMsgMissingImagePayment         = fiber.NewError(fiber.StatusBadRequest, "Missing image payment.")
	ErrMsgAddImageProductFailed       = fiber.NewError(fiber.StatusInternalServerError, "Add product image failed.")
)

func NewOrderHandler(router fiber.Router, service OrderService) {
	handler := &handler{service: service}

	router.Use(customerMiddleware())
	router.Get("/", handler.getOrderPending)
	router.Get("/products", handler.getProductOrderPending)
	router.Post("/products/:productID", handler.addProductToCart)
	router.Delete("/products/:productID", handler.removeProductFromCart)
	router.Patch("/address/:addressID", handler.updateAddress)
	router.Post("/complete", handler.confirmOrder)
	router.Post("/payment", handler.addPaymentSlip)
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

func (h *handler) getProductOrderPending(c *fiber.Ctx) error {
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

func (h *handler) updateAddress(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID

	defer cancel()

	addressID, err := c.ParamsInt("addressID")
	if err != nil {
		return ErrMsgMissingAddressID
	}

	err = h.service.UpdateAddressOrder(ctx, custID, addressID)
	switch err {
	case nil:
		break
	case ErrBlankCart:
		return ErrMsgBlankCart
	case ErrWrongOwnerProduct:
		return ErrMsgWrongOwnerProduct
	case ErrAddressNotFound:
		return ErrMsgAddressNotFound
	default:
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Update order successfully.",
	})
}

func (h *handler) getOrderPending(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID

	defer cancel()

	data, err := h.service.GetCart(ctx, custID)

	if err != nil {
		return ErrMsgBlankCart
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func (h *handler) confirmOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID

	defer cancel()

	err := h.service.ConfirmOrder(ctx, custID)

	if err != nil {
		return ErrMsgBlankCart
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Confirm order successfully.",
	})
}

func (h *handler) addPaymentSlip(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser").(*UserToken).ID
	defer cancel()

	file, err := c.FormFile("image")
	if err != nil {
		return ErrMsgMissingImagePayment
	}

	request := &Payment{
		Image: file,
	}

	err = h.service.SendPayment(ctx, c, request, custID)

	if err != nil {
		if err == ErrBlankCart {
			return ErrMsgBlankCart
		}
		return ErrMsgAddImageProductFailed
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add image successfully.",
	})
}
