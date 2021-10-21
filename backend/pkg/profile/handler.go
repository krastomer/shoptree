package profile

import "github.com/gofiber/fiber/v2"

var (
	ErrMsgUnavailableRole  = fiber.NewError(fiber.StatusUnauthorized, "Can't access from this role.")
	ErrMsgCustomerNotFound = fiber.NewError(fiber.StatusNotFound, "User not found.")
)

type profileHandler struct {
	service ProfileService
}

func NewProfileHandler(router fiber.Router, service ProfileService) {
	handler := &profileHandler{service: service}

	router.Use(JWTMiddleware())

	router.Get("/", handler.getProfile)
}

func (h *profileHandler) getProfile(c *fiber.Ctx) error {
	user := c.Locals("currentUser").(*User)
	if user.Level != "Customer" {
		return ErrMsgUnavailableRole
	}

	response, err := h.service.GetProfileCustomer(user.ID)
	if err != nil {
		return ErrMsgCustomerNotFound
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   response,
	})
}
