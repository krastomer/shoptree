package infrastructure

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

func adminMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: jwtAdminHandler,
		ErrorHandler:   unauthHandler,
		SigningKey:     []byte(viper.GetString("JWT_SECRET")),
		SigningMethod:  "HS256",
		TokenLookup:    "cookie:jwt",
	})
}

func jwtAdminHandler(c *fiber.Ctx) error {
	jwtData := c.Locals("user").(*jwt.Token)

	claims := jwtData.Claims.(jwt.MapClaims)

	token := &UserToken{
		ID:    int(claims["uid"].(float64)),
		Email: claims["user"].(string),
		Level: claims["aud"].(string),
	}

	if (token.Level != "Admin") && (token.Level != "Deliver") && (token.Level != "Staff") {
		return unauthHandler(c, nil)
	}

	c.Locals("currentUser", token)

	return c.Next()
}

func unauthHandler(c *fiber.Ctx, _ error) error {
	return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized from this user.")
}
