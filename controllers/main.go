package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type QueueData struct {
	ID       string                 `json:"_id"`
	Endpoint string                 `json:"endpoint"`
	Req      map[string]interface{} `json:"req"`
}

var (
	ctx         = context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Subscriber = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	QueueName = "redis-server" 
)

func PushToQueue(queueName, data string) error {
	log.Printf("Pushing to queue: %s", data)
	return RedisClient.LPush(ctx, queueName, data).Err()
}

func ForwardRequest(c *gin.Context, endpoint string) {
	
	payload := QueueData{
		ID:       uuid.New().String(),
		Endpoint: endpoint,
		Req: map[string]interface{}{
			"body":   c.Request.PostForm,
			"params": c.Params,
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	channel := payload.ID
	sub := Subscriber.Subscribe(ctx, channel)
	defer sub.Close()

	
	if err := PushToQueue(QueueName, string(payloadBytes)); err != nil {
		log.Printf("Failed to push to queue: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push to queue"})
		return
	}

	go func() {
		for msg := range sub.Channel() {
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
			return
		}
	}()

	// Timeout if no response is received
	select {
	case <-time.After(30 * time.Second):
		c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
		return
	}
}
