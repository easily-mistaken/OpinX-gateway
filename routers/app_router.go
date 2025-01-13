package routers

import (
	"github.com/easily-mistaken/OpinX-gateway/controllers"
	"github.com/gin-gonic/gin"
)

func appRouter(rg *gin.RouterGroup) {
	router := rg.Group("/")

	// Create User
	router.POST("/user/create/:userId", func(c *gin.Context) {
		controllers.ForwardRequest(c, "/user/create/:userId")
	})

	// Create Symbol
	router.POST("/symbol/create/:stockSymbol", func(c *gin.Context) {
		controllers.ForwardRequest(c, "/symbol/create/:stockSymbol")
	})

	// Onramp Money
	router.POST("/onramp/inr", func(c *gin.Context) {
		controllers.ForwardRequest(c, "/onramp/inr")
	})

	// Mint Tokens
	router.POST("/trade/mint", func(c *gin.Context) {
		controllers.ForwardRequest(c, "/trade/mint")
	})

	// Reset database
	router.POST("/reset", func(c *gin.Context) {
		controllers.ForwardRequest(c, "/reset")
	})
}