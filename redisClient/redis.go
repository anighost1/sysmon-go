package redisClient

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Client *redis.Client

// Init Redis with pooling + retry
func InitRedis() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       0,

		// Connection Pool Settings
		PoolSize:     10, // max connections
		MinIdleConns: 2,  // keep idle connections ready
		MaxRetries:   3,  // retry failed commands

		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// Retry connection
	var err error
	for i := 0; i < 5; i++ {
		_, err = Client.Ping(Ctx).Result()
		if err == nil {
			slog.Info("Connected to Redis")
			return nil
		}

		slog.Warn("Retrying Redis connection...", "attempt", i+1, "error", err)
		time.Sleep(2 * time.Second)
	}

	return err
}

// Helper: Env fallback
func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

// Graceful Shutdown
func CloseRedis() {
	if Client != nil {
		err := Client.Close()
		if err != nil {
			slog.Error("Error closing Redis", "error", err)
		} else {
			slog.Info("Redis connection closed")
		}
	}
}

// Listen for shutdown signals (CTRL+C, Docker stop, etc.)
func HandleGracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	slog.Info("Shutting down application...")

	CloseRedis()

	os.Exit(0)
}
