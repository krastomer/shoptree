package profile

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: getDataFromJWT,
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

func getDataFromJWT(c *fiber.Ctx) error {
	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	uid, err := strconv.ParseUint(claims["uid"].(string), 10, 32)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	user := &User{
		ID:    uint32(uid),
		Email: claims["user"].(string),
		Level: strings.Split(claims["aud"].(string), "-")[1],
	}
	c.Locals("currentUser", user)
	return c.Next()
}
