package routers

import (
	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine) {
	getRoutes(r)
}

func getRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	appRoutes(v1)
	balanceRoutes(v1)
	orderRoutes(v1)
	orderbookRoutes(v1)

	// v2 := router.Group("/v2")
	// comingSoon(v2)x
}
