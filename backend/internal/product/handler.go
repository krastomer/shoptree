package product

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	service ProductService
}

var (
	ErrMsgIDProduct            = fiber.NewError(fiber.StatusBadRequest, "Product 'ID' should postive integer.")
	ErrMsgImageProduct         = fiber.NewError(fiber.StatusBadRequest, "Require image.")
	ErrMsgIDProductImage       = fiber.NewError(fiber.StatusBadRequest, "Product image 'ID' should postive integer.")
	ErrMsgProductIDNotFound    = fiber.NewError(fiber.StatusNotFound, "Product ID not found.")
	ErrMsgProductImageNotFound = fiber.NewError(fiber.StatusNotFound, "Product image not found.")
	ErrMsgAddProductRequire    = fiber.NewError(fiber.StatusBadRequest, "Product require Name, ScientificName, Price, Description and Status.")
	ErrMsgProductStatus        = fiber.NewError(fiber.StatusBadRequest, "Product Status should Unavailable, Available, Pending, Purchased.")
)

func NewProductHandler(router fiber.Router, service ProductService) {
	handler := &productHandler{service: service}

	router.Get("/:id", handler.getProductByID)
	router.Get("/images/:id", handler.getImageProductByID)

	// router.Post("/", StaffMiddleware(), handler.addProduct)
	router.Post("/:id/images", StaffMiddleware(), handler.createImageProduct)
}

func (h *productHandler) getProductByID(c *fiber.Ctx) error {
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

func (h *productHandler) getImageProductByID(c *fiber.Ctx) error {
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

// func (h *productHandler) addProduct(c *fiber.Ctx) error {
// 	request := &ProductRequest{}

// 	if err := c.BodyParser(request); err != nil {
// 		return ErrMsgAddProductRequire
// 	}

// 	err := h.service.AddProduct(request)
// 	if err != nil {
// 		if err == ErrProductStatus {
// 			return ErrMsgProductStatus
// 		}
// 		return fiber.ErrInternalServerError
// 	}

// 	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
// 		"status":  "success",
// 		"message": "Add product successfully.",
// 	})
// }

func (h *productHandler) createImageProduct(c *fiber.Ctx) error {
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
		return ErrMsgProductIDNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Add image successfully.",
	})
}
