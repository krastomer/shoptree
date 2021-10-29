package product

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrMsgFailedBodyParser  = fiber.NewError(fiber.StatusBadRequest, "Product Require `name`, `scientific_name`, `price`, `description`, `status`.")
	ErrMsgBadIDProduct      = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgIDProductNotFound = fiber.NewError(fiber.StatusNotFound, "Product not found.")
)

type productHandler struct {
	service ProductService
}

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &productHandler{service: service}

	router.Use(SoftJWTMiddleware())
	router.Get("/", handler.getProducts)
	router.Get("/:id", handler.getProductByID)
	router.Post("/", EmployeeMiddleware, handler.addProduct)
}

// TODO: edit middleware
func (h *productHandler) getProductByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return ErrMsgBadIDProduct
	}

	response, err := h.service.GetProductByID(uint32(id))
	if err != nil {
		return ErrMsgIDProductNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})

}

func (h *productHandler) getProducts(c *fiber.Ctx) error {
	return nil
}

func (h *productHandler) addProduct(c *fiber.Ctx) error {
	request := &Product{}

	if err := c.BodyParser(request); err != nil {
		return ErrMsgFailedBodyParser
	}

	_ = h.service.AddProduct(request)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add Product successfully",
	})
}
