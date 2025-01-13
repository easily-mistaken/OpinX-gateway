package routes

import (
	"github.com/easily-mistaken/OpinX-gateway/controllers"
	"github.com/gin-gonic/gin"
)

func balanceRoutes(route *gin.RouterGroup) {
	router := route.Group("/balance")


	router.GET("/inr", controllers.forwardRequest)

	router.GET("/inr/:userId", controllers.forwardRequest)

	router.GET("/stock", controllers.forwardRequest)

	router.GET("/stock/stockSymbol", controllers.forwardRequest)
}