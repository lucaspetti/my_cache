package cache

import (
	"os"
	"log"

	"github.com/go-redis/redis/v8"
)

const password = ""
const db       = 0

func NewRedisClient() *redis.Client {
	address := os.Getenv("REDIS_ADDRESS")
	log.Println("Connecting to redis on ", address)

	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
}
