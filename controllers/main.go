package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/easily-mistaken/OpinX-gateway/config"
	redisclient "github.com/easily-mistaken/OpinX-gateway/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Add constants for timeouts and error messages
const (
	RequestTimeout = 30 * time.Second
	ErrInternalServer = "Internal server error"
	ErrRequestTimeout = "Request timed out"
)

// Add structured logging
type RequestLog struct {
	ID        string    `json:"id"`
	Endpoint  string    `json:"endpoint"`
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error,omitempty"`
}

type QueueData struct {
	ID        string                 `json:"_id"`
	Endpoint  string                 `json:"endpoint"`
	Timestamp time.Time             `json:"timestamp"`
	Req struct {
		Body   interface{}       `json:"body"`
		Params gin.Params        `json:"params"`
		Query  map[string][]string `json:"query"`  // Changed from map[string]string
	} `json:"req"`
}

// ForwardRequest handles forwarding requests to the queue and waiting for a response
func ForwardRequest(c *gin.Context, endpoint string) {
	redisClient := redisclient.GetInstance("", "", 0) // Get the singleton instance

	// Create payload
	payload := QueueData{
		ID:       uuid.New().String(),
		Endpoint: endpoint,
		Req: struct {
			Body   interface{}       `json:"body"`
			Params gin.Params        `json:"params"`
			Query  map[string][]string `json:"query"`  // Changed from map[string]string
		}{
			Body:   c.Request.PostForm,
			Params: c.Params,
			Query:  c.Request.URL.Query(),
		},
	}

	// Serialize payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Subscribe before pushing to queue
	pubSub := redisClient.Subscribe(payload.ID)
	defer func() {
		if err := redisClient.Unsubscribe(pubSub, payload.ID); err != nil {
			log.Printf("Warning: Failed to unsubscribe from channel %s: %v", payload.ID, err)
		}
		pubSub.Close()
	}()

	// Push to queue after subscription is ready
	if err := redisClient.PushToQueue(config.AppConfig.Redis.QueueName, string(payloadBytes)); err != nil {
		log.Printf("Failed to push to queue: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push to queue"})
		return
	}

	// Create a context that's cancelled when the client disconnects
	requestCtx, cancel := context.WithTimeout(c.Request.Context(), RequestTimeout)
	defer cancel()

	ch := pubSub.Channel()
	select {
	case msg := <-ch:
		var response struct {
			StatusCode int         `json:"statusCode"`
			Data       interface{} `json:"data"`
		}
		if err := json.Unmarshal([]byte(msg.Payload), &response); err != nil {
			log.Printf("Failed to unmarshal response: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		c.JSON(response.StatusCode, response.Data)

	case <-requestCtx.Done():
		if requestCtx.Err() == context.Canceled {
			log.Printf("Request cancelled by client: %s", payload.ID)
			return
		}
		c.JSON(http.StatusGatewayTimeout, gin.H{"error": ErrRequestTimeout})
	}
}
