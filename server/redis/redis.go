package redisClient

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
)

// New : Create a new redis client
func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // TODO : Change redis docker container to use the pass stored in .env
		DB:       0,  // use default DB
	})
}
