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
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get server port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Initialize Redis connection
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := 0

	// Initialize Redis singleton
	redisClient := redisclient.GetInstance(redisAddr, redisPassword, redisDB)
	defer redisClient.Close()

	// Initialize Gin server
	r := gin.Default()

	// Health check route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	routers.Run(r)

	// Start the server
	fmt.Printf("Server running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Println("Error running server:", err)
	}
}
