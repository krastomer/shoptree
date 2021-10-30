package product

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrMsgFailedBodyParser      = fiber.NewError(fiber.StatusBadRequest, "Product Require `name`, `scientific_name`, `price`, `description`, `status`.")
	ErrMsgBadIDProduct          = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgProductIDNotFound     = fiber.NewError(fiber.StatusNotFound, "Product not found.")
	ErrMsgMissingImage          = fiber.NewError(fiber.StatusBadRequest, "Missing Product Image.")
	ErrMsgProductImagesNotFound = fiber.NewError(fiber.StatusNotFound, "Product image not found.")
)

type productHandler struct {
	service ProductService
}

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &productHandler{service: service}

	router.Use(SoftJWTMiddleware())
	router.Get("/", handler.getProducts)
	router.Get("/:id", handler.getProductByID)
	router.Get("/image/:id", handler.getProductImageByID)

	router.Post("/", EmployeeMiddleware, handler.addProduct)
	router.Post("/:id/image", EmployeeMiddleware, handler.addProductImage)
}

// TODO: edit middleware
func (h *productHandler) getProductByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return ErrMsgBadIDProduct
	}

	response, err := h.service.GetProductByID(uint32(id))
	if err != nil {
		return ErrMsgProductIDNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})

}

func (h *productHandler) getProducts(c *fiber.Ctx) error {
	type productsRequest struct {
		ID []uint32 `json:"id"`
	}

	request := &productsRequest{}
	_ = c.BodyParser(request)

	response, err := h.service.GetProducts(request.ID)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *productHandler) addProduct(c *fiber.Ctx) error {
	request := &Product{}

	if err := c.BodyParser(request); err != nil {
		return ErrMsgFailedBodyParser
	}

	_ = h.service.AddProduct(request)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add Product successfully.",
	})
}

func (h *productHandler) addProductImage(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return ErrMsgBadIDProduct
	}
	file, err := c.FormFile("image")
	if err != nil {
		return ErrMsgMissingImage
	}

	err = h.service.AddProductImage(uint32(id), c, file)

	if err != nil {
		return ErrMsgBadIDProduct
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add image successfully.",
	})
}

func (h *productHandler) getProductImageByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return ErrMsgBadIDProduct
	}
	path, err := h.service.GetProductImageByID(uint32(id))
	if err != nil {
		return ErrMsgProductImagesNotFound
	}

	return c.Status(fiber.StatusOK).SendFile(path)
}
