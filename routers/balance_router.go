package routers

import (
	"github.com/easily-mistaken/OpinX-gateway/controllers"
	"github.com/gin-gonic/gin"
)

func balanceRoutes(route *gin.RouterGroup) {
	router := route.Group("/balance")


	router.GET("/inr", controllers.ForwardRequest)

	router.GET("/inr/:userId", controllers.ForwardRequest)

	router.GET("/stock", controllers.ForwardRequest)

	router.GET("/stock/stockSymbol", controllers.ForwardRequest)
}