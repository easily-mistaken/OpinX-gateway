package main

import (
	"fmt"
	"os"

	"github.com/easily-mistaken/OpinX-gateway/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err  := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	r  :=gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Group("/api", routers.getRoutes)

	err = r.Run(":" + port)
	if err != nil {
		fmt.Println("Error running server")
	}

	fmt.Println("Server running on port", port)
}