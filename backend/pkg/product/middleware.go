package product

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func SoftJWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: getDataFromJWT,
		ErrorHandler:   setGeneralUser,
		SigningKey:     []byte(viper.GetString("JWT_SECRET")),
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

	user := &User{
		ID:    uint32(uid),
		Email: claims["user"].(string),
		Level: strings.Split(claims["aud"].(string), "-")[1],
	}
	c.Locals("currentUser", user)
	return c.Next()
}

func setGeneralUser(c *fiber.Ctx, _ error) error {
	user := &User{
		Level: "GeneralUser",
	}
	c.Locals("currentUser", user)
	return c.Next()
}

func EmployeeMiddleware(c *fiber.Ctx) error {
	user := c.Locals("currentUser").(*User)

	if user.Level != "Admin" && user.Level != "Staff" && user.Level != "Deliver" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "unauthorized",
			"message": "Your role can't access.",
		})
	}

	return c.Next()
}
