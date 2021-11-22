package order

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service OrderService
}

var (
	ErrMsgProductID = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
)

func NewOrderHandler(router fiber.Router, service OrderService) {
	handler := &handler{service: service}

	router.Post("/:productID", handler.addProductToCart)
	// router.Delete()
}

func (h *handler) addProductToCart(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	custID := c.Locals("currentUser")["ID"].(int)

	defer cancel()

	id, err := c.ParamsInt("productID")
	if err != nil {
		return ErrMsgProductID
	}

	err := h.service.AddProductToCart(ctx)

}
