package product

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

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
