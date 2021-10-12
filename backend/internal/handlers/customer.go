package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/krastomer/shoptree/backend/internal/entities"
)

type customerHandler struct {
	service entities.CustomerService
}

func NewCustomerHandler(router fiber.Router, service entities.CustomerService) {
	handler := &customerHandler{service: service}

	router.Get(":id", handler.GetCustomer)
}

func (h *customerHandler) GetCustomer(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	cust, _ := h.service.GetCustomer(id)
	return c.JSON(cust)
}
