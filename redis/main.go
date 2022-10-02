package redis

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
)

var RDB *redis.Client
var Cache *cache.Cache

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	redisUrl := os.Getenv("REDIS_HOST")
	redisUser := os.Getenv("REDIS_USER")
	redisPass := os.Getenv("REDIS_PASS")

	RDB = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPass,
		Username: redisUser,
		DB:       0,
	})

	Cache = cache.New(&cache.Options{
		Redis:      RDB,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
}
