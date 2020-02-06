package redis

import (
	"github.com/go-redis/redis/v7"
)

var (
	Client *redis.Client
)

// CreateClient creates a redis connection
func CreateClient() *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return Client
}

// CloseClient closes an open redis connection
func CloseClient() {
	Client.Close()
}
