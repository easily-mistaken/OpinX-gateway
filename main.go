package main

import (
	"fmt"
	"os"

	"github.com/easily-mistaken/OpinX-gateway/routers"
	redisclient "github.com/easily-mistaken/OpinX-gateway/services"
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

	redisclient.ConnectToRedis()

	r  :=gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	routers.Run(r) 

	err = r.Run(":" + port)
	if err != nil {
		fmt.Println("Error running server")
	}   

	fmt.Println("Server running on port", port)
}