package redisClient

import (
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
)

// New : Create a new redis client
func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}
