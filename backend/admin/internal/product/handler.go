package product

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service ProductService
}

var (
	ErrMsgProductNotFound       = fiber.NewError(fiber.StatusNotFound, "Product Not found.")
	ErrMsgProductBody           = fiber.NewError(fiber.StatusBadRequest, "Product need name, scientific_name, description, price.")
	ErrMsgIDProduct             = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgCreateProductFailed   = fiber.NewError(fiber.StatusInternalServerError, "Create Product failed.")
	ErrMsgUpdateProductFailed   = fiber.NewError(fiber.StatusInternalServerError, "Update Product failed.")
	ErrMsgDeleteProductFailed   = fiber.NewError(fiber.StatusInternalServerError, "Delete Product failed.")
	ErrMsgImageProduct          = fiber.NewError(fiber.StatusBadRequest, "Require image.")
	ErrMsgAddImageProductFailed = fiber.NewError(fiber.StatusInternalServerError, "Add Image Product failed.")
)

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &handler{service: service}

	router.Get("/", handler.getProducts)
	router.Post("/", handler.createProduct)
	router.Patch("/:id", handler.updateProduct)
	router.Delete("/:id", handler.deleteProduct)
	router.Post("/:id/images", handler.addImageProduct)
}

func (h *handler) getProducts(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response, err := h.service.GetProducts(ctx)
	if err != nil {
		return ErrMsgProductNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *handler) createProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	request := &Product{}

	if err := c.BodyParser(request); err != nil {
		return ErrMsgProductBody
	}

	err := h.service.CreateProduct(ctx, request)
	if err != nil {
		return ErrMsgCreateProductFailed
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Create Product successfully",
	})
}

func (h *handler) updateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := &Product{}
	if err := c.BodyParser(request); err != nil {
		return ErrMsgProductBody
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgIDProduct
	}
	request.ID = id

	err = h.service.UpdateProduct(ctx, request)

	if err != nil {
		if err == ErrProductNotFound {
			return ErrMsgProductNotFound
		}
		return ErrMsgUpdateProductFailed
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Update Product successfully",
	})
}

func (h *handler) deleteProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return ErrMsgIDProduct
	}

	err = h.service.DeleteProduct(ctx, id)

	if err != nil {
		if err == ErrProductNotFound {
			return ErrMsgProductNotFound
		}
		return ErrMsgDeleteProductFailed
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Delete Product successfully",
	})
}

func (h *handler) addImageProduct(c *fiber.Ctx) error {
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
		return ErrMsgAddImageProductFailed
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add image successfully.",
	})
}
