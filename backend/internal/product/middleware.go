package product

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

func StaffMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: jwtStaffHandler,
		ErrorHandler:   jwtError,
		SigningKey:     []byte(viper.GetString("JWT_SECRET")),
		SigningMethod:  "HS256",
		TokenLookup:    "cookie:jwt",
	})
}

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

func jwtStaffHandler(c *fiber.Ctx) error {
	jwtData := c.Locals("user").(*jwt.Token)

	claims := jwtData.Claims.(jwt.MapClaims)

	token := &UserToken{
		ID:    int(claims["uid"].(float64)),
		Email: claims["user"].(string),
		Level: claims["aud"].(string),
	}

	if token.Level != "Admin" && token.Level != "Staff" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "unauthorized",
			"message": "Your role can't access.",
		})
	}

	c.Locals("currentUser", token)

	return c.Next()
}
