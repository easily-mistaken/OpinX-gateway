package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func orderbookRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/orderbook")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}