package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func appRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}