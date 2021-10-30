package product

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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
	router.Get("/:id/image", handler.getProductImages)
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
		return ErrMsgIDProductNotFound
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
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["image"]
		for _, file := range files {
			if err := c.SaveFile(file, fmt.Sprintf("%s/%s", viper.GetString("DIRECTORY_PRODUCT"), file.Filename)); err != nil {
				return err
			}
		}
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add image successfully.",
	})
}

func (h *productHandler) getProductImages(c *fiber.Ctx) error {
	return nil
}
