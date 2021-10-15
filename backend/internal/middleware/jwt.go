package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/krastomer/shoptree/backend/internal/models"
)

// JWT error message.
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT!",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
		"status":  "error",
		"message": "Invalid or expired JWT!",
	})
}

// Guards a specific endpoint in the API.
func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: getDataFromJWT,
		ErrorHandler:   jwtError,
		SigningKey:     []byte("september"),
		SigningMethod:  "HS256",
		TokenLookup:    "cookie:jwt",
	})
}

// Gets user data (their ID) from the JWT middleware. Should be executed after calling 'JWTMiddleware()'.
func getDataFromJWT(c *fiber.Ctx) error {
	// // Get userID from the previous route.
	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	uid, err := strconv.ParseUint(claims["uid"].(string), 10, 32)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user := &models.User{
		ID:    uint32(uid),
		Email: claims["user"].(string),
		Level: claims["aud"].(string),
	}
	c.Locals("currentUser", user)
	return c.Next()
}
