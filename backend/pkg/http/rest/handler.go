package rest

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func init() {
	// gin.DisableConsoleColor()
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
}

func NewHandler() *fiber.App {

	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/login", func(c *fiber.Ctx) error {
		type loginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		type jwtClaims struct {
			UserID string `json:"uid"`
			User   string `json:"user"`
			jwt.StandardClaims
		}

		request := &loginRequest{}
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
			os.Getenv("API_USERID"),
			os.Getenv("API_USERNAME"),
			jwt.StandardClaims{
				Audience:  "docker-mariadb-clean-arch-users",
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				Issuer:    "docker-mariadb-clean-arch",
			},
		})
		signedToken, err := token.SignedString([]byte("september"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    signedToken,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			Secure:   false,
			HTTPOnly: true,
		})

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "success",
			"token":  signedToken,
		})
	})

	// router := gin.Default()
	// api := router.Group("/api")

	// v1 := api.Group("/v1")
	// {
	// 	v1.GET("/", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, map[string]interface{}{
	// 			"hello": "world",
	// 		})
	// 	})
	// }

	return app
}
