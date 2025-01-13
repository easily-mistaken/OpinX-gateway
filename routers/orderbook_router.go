package routers

import (
	"github.com/easily-mistaken/OpinX-gateway/controllers"
	"github.com/gin-gonic/gin"
)

// Orderbook
func orderbookRouter(rg *gin.RouterGroup) {
	router := rg.Group("/orderbook")

	router.GET("/", func (c *gin.Context){
		controllers.ForwardRequest(c, "orderbook")
	})

	router.GET("/:stockSymbol", func (c *gin.Context){
		controllers.ForwardRequest(c, "/orderbook/:stockSymbol")
	})
}