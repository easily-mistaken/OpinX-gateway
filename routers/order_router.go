package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func orderRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/order")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}