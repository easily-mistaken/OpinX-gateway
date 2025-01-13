package routers

import (
	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine) {
	getRoutes(r)
}

func getRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")

	appRouter(v1)
	balanceRouter(v1)
	ordersRouter(v1)
	orderbookRouter(v1)

	// v2 := router.Group("/v2")
	// comingSoon(v2)x
}
