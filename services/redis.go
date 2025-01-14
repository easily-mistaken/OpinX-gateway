package redisclient

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Client     *redis.Client
	Subscriber *redis.Client
)

var ctx = context.Background()

func ConnectToRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	Subscriber = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Test connection
	if err := Client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Critical: Failed to connect to Redis: %v", err)
	}

	if err := Subscriber.Ping(ctx).Err(); err != nil {
		log.Fatalf("Critical: Failed to connect to Redis Subscriber: %v", err)
	}

	log.Println("Info: Successfully connected to Redis")
}

func PushToQueue(queueName, data string) {
	if err := Client.LPush(ctx, queueName, data).Err(); err != nil {
		log.Printf("Failed to push data to queue %s: %v", queueName, err)
		return
	}
	log.Printf("Data pushed to queue %s: %s", queueName, data)
}