package routers

import (
	"github.com/easily-mistaken/OpinX-gateway/controllers"
	"github.com/gin-gonic/gin"
)

// Balances
func balanceRouter(route *gin.RouterGroup) {
	router := route.Group("/balances")


	router.GET("/inr", func (c *gin.Context){
		controllers.ForwardRequest(c, "balances/inr")
	})

	router.GET("/inr/:userId", func (c *gin.Context){
		controllers.ForwardRequest(c, "/balances/inr/:uderId")
	})

	router.GET("/stock", func (c *gin.Context){
		controllers.ForwardRequest(c, "/balances/stock")
	})

	router.GET("/stock/stockSymbol", func (c *gin.Context){
		controllers.ForwardRequest(c, "balances/stock/stockSymbol")
	})
}