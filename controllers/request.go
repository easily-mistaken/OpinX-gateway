package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ForwardRequest(c *gin.Context) {
	// forward request to redis queue and wait for the response

	// Waiting for the response to get published on unique channel

	// Push request to the queue

	c.JSON(http.StatusOK, gin.H{
		"message": "Request forwarded to queue",
	})
}
