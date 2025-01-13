package routers

import (
	"github.com/easily-mistaken/OpinX-gateway/controllers"
	"github.com/gin-gonic/gin"
)

// Orders
func ordersRouter(rg *gin.RouterGroup) {
	router := rg.Group("/order")

	router.POST("/buy", func(c *gin.Context) {
		controllers.ForwardRequest(c, "/order/buy")
	})

	router.POST("/sell", func(c *gin.Context) {
		controllers.ForwardRequest(c, "/order/sell")
	})

	router.POST("/cancel", func(c *gin.Context) {
		controllers.ForwardRequest(c, "/order/cancel")
	})
}