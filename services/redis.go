package redisclient

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	instance *RedisClient
	once     sync.Once
)

type RedisClient struct {
	Client     *redis.Client
	Subscriber *redis.Client
	ctx        context.Context
}

// GetInstance returns a singleton instance of RedisClient
func GetInstance(addr, password string, db int) *RedisClient {
	once.Do(func() {
		instance = NewRedisClient(addr, password, db)
	})
	return instance
}

// NewRedisClient initializes and returns a new instance of RedisClient
func NewRedisClient(addr, password string, db int) *RedisClient {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	subscriber := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Test connections
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Critical: Failed to connect to Redis Client: %v", err)
	}
	if err := subscriber.Ping(ctx).Err(); err != nil {
		log.Fatalf("Critical: Failed to connect to Redis Subscriber: %v", err)
	}

	log.Println("Info: Successfully connected to Redis")

	return &RedisClient{
		Client:     client,
		Subscriber: subscriber,
		ctx:        ctx,
	}
}

func (r *RedisClient) PushToQueue(queueName, data string) error {
	err := r.Client.LPush(r.ctx, queueName, data).Err()
	if err != nil {
		log.Printf("Failed to push data to queue %s: %v", queueName, err)
		return err
	}
	log.Printf("Data pushed to queue %s: %s", queueName, data)
	return nil
}

func (r *RedisClient) Subscribe(channel string) *redis.PubSub {
	return r.Subscriber.Subscribe(r.ctx, channel)
}

func (r *RedisClient) GetContext() context.Context {
	return r.ctx
}

func (r *RedisClient) Close() {
	if err := r.Client.Close(); err != nil {
		log.Printf("Warning: Failed to close Redis Client: %v", err)
	}
	if err := r.Subscriber.Close(); err != nil {
		log.Printf("Warning: Failed to close Redis Subscriber: %v", err)
	}
	log.Println("Info: Redis connections closed")
}

func (r *RedisClient) Unsubscribe(pubsub *redis.PubSub, channel string) error {
	return pubsub.Unsubscribe(r.ctx, channel)
}

func (r *RedisClient) HealthCheck() error {
	return r.Client.Ping(r.ctx).Err()
}

func (r *RedisClient) CleanupSubscription(pubsub *redis.PubSub, channel string) {
	if err := r.Unsubscribe(pubsub, channel); err != nil {
		log.Printf("Warning: Failed to unsubscribe from channel %s: %v", channel, err)
	}
	if err := pubsub.Close(); err != nil {
		log.Printf("Warning: Failed to close pubsub for channel %s: %v", channel, err)
	}
}
