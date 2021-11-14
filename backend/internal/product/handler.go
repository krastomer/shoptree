package product

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service ProductService
}

var (
	ErrMsgIDProduct            = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgImageProduct         = fiber.NewError(fiber.StatusBadRequest, "Require image.")
	ErrMsgIDProductImage       = fiber.NewError(fiber.StatusBadRequest, "Product image 'ID' should postive integer.")
	ErrMsgProductIDNotFound    = fiber.NewError(fiber.StatusNotFound, "Product ID not found.")
	ErrMsgProductImageNotFound = fiber.NewError(fiber.StatusNotFound, "Product image not found.")
	ErrMsgCreateProductRequire = fiber.NewError(fiber.StatusBadRequest, "Product require Name, ScientificName, Price, Description and Status.")
	ErrMsgProductStatus        = fiber.NewError(fiber.StatusBadRequest, "Product Status should Unavailable, Available, Pending, Purchased.")
	ErrMsgCreateProductFailed  = fiber.NewError(fiber.StatusBadRequest, "Create Product failed.")
)

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &handler{service: service}

	router.Get("/:id", handler.getProductByID)
	router.Get("/:id/images", handler.getImagesProductID)
	router.Get("/images/:id", handler.getImageProductByID)

	router.Post("/", StaffMiddleware(), handler.createProduct)
	router.Post("/:id/images", StaffMiddleware(), handler.createImageProduct)
}

func (h *handler) getProductByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgIDProduct
	}

	response, err := h.service.GetProductByID(ctx, id)
	if err != nil {
		return ErrMsgProductIDNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *handler) getImageProductByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgIDProduct
	}
	path, err := h.service.GetImageProductByID(ctx, id)
	if err != nil {
		return ErrProductImageNotFound
	}

	return c.Status(fiber.StatusOK).SendFile(path)
}

func (h *handler) getImagesProductID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgIDProduct
	}

	response, err := h.service.GetImagesProductID(ctx, id)
	if err != nil {
		return ErrMsgProductImageNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *handler) createImageProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgIDProduct
	}

	file, err := c.FormFile("image")
	if err != nil {
		return ErrMsgImageProduct
	}

	request := &ImageProduct{
		ProductID: id,
		Image:     file,
	}

	err = h.service.CreateImageProduct(ctx, c, request)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add image successfully.",
	})
}

func (h *handler) createProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := &ProductRequest{}
	if err := c.BodyParser(request); err != nil {
		return ErrMsgCreateProductRequire
	}

	err := h.service.CreateProduct(ctx, request)
	if err != nil {
		return ErrMsgCreateProductFailed
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Create Product successfully.",
	})
}
