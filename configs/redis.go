package configs

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Context = context.Background()

func InitRedis() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		// Password: "P@$$w0rd##2143",
		DB:       0,
	})

	_, pingError := redisClient.Ping(Context).Result()
	if pingError != nil {
		panic("Failed to connect to Redis")
	}

	RedisClient = redisClient
	fmt.Println("Connected to Redis")
}
