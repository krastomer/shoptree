package customer

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type UserToken struct {
	ID    int
	Email string
	Level string
}

func CustomerMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: jwtCustomerHandler,
		ErrorHandler:   errorHandler,
		SigningKey:     []byte(viper.GetString("JWT_SECRET")),
		SigningMethod:  "HS256",
		TokenLookup:    "cookie:jwt",
	})
}

func jwtCustomerHandler(c *fiber.Ctx) error {
	jwtData := c.Locals("user").(*jwt.Token)

	claims := jwtData.Claims.(jwt.MapClaims)

	token := &UserToken{
		ID:    int(claims["uid"].(float64)),
		Email: claims["user"].(string),
		Level: claims["aud"].(string),
	}

	if token.Level != "Customer" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "unauthorized",
			"message": "Your role can't access.",
		})
	}

	c.Locals("currentUser", token)

	return c.Next()
}

func errorHandler(c *fiber.Ctx, _ error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
		"status":  "unauthorized",
		"message": "Your role can't access.",
	})
}
