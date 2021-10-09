package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.DisableConsoleColor()
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
}

func NewHandler() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	v1 := api.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"hello": "world",
			})
		})
	}

	return router
}
