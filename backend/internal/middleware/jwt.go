package middleware

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/krastomer/shoptree/backend/internal/models"
)

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

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: getDataFromJWT,
		ErrorHandler:   jwtError,
		SigningKey:     []byte("september"),
		SigningMethod:  "HS256",
		TokenLookup:    "cookie:jwt",
	})
}

func SoftJWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: getDataFromJWT,
		ErrorHandler:   setGeneralUser,
		SigningKey:     []byte("september"),
		SigningMethod:  "HS256",
		TokenLookup:    "cookie:jwt",
	})
}

func getDataFromJWT(c *fiber.Ctx) error {
	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	uid, err := strconv.ParseUint(claims["uid"].(string), 10, 32)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user := &models.User{
		ID:    uint32(uid),
		Email: claims["user"].(string),
		Level: strings.Split(claims["aud"].(string), "-")[1],
	}
	c.Locals("currentUser", user)
	return c.Next()
}

func setGeneralUser(c *fiber.Ctx, _ error) error {
	user := &models.User{
		Level: "GeneralUser",
	}
	c.Locals("currentUser", user)
	return c.Next()
}
