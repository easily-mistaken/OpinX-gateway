package config

import (
	"os"
	"time"
)

type Config struct {
    Redis struct {
        QueueName string
        Addr     string
        Password string
        DB       int
        Timeout  time.Duration
    }
    Server struct {
        Port string
    }
}

var AppConfig = loadConfig()

func loadConfig() Config {
    var cfg Config
    
    cfg.Redis.QueueName = getEnvOrDefault("REDIS_QUEUE_NAME", "redis-server")
    cfg.Redis.Addr = getEnvOrDefault("REDIS_ADDR", "localhost:6379")
    cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")
    cfg.Redis.Timeout = 30 * time.Second
    cfg.Server.Port = getEnvOrDefault("PORT", "4000")
    
    return cfg
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
} 