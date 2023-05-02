package infrastructure

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func redisSet(redisEndpoint string) error {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr: redisEndpoint, // Redis server address
		DB:   0,             // Redis database number
	})

	// Set a key-value pair
	err := client.Set(context.Background(), "name", "jay", 0).Err()
	if err != nil {
		return err
	}

	val, err := client.Get(context.Background(), "name").Result()
	if err != nil {
		return err
	}
	fmt.Println("key", val)
	return nil
}

// func main() {
// 	redisEndpoint := "myt-ca-1i62tyhzqdngv.cvvqf5.0001.apne1.cache.amazonaws.com:6379"
// 	redisSet(redisEndpoint)
// }
