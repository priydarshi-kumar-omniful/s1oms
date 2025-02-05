package utils

import (
	"context"
	"fmt"
	"time"
	"oms/constants"
	"github.com/omniful/go_commons/redis"
)

// ConnectToRedis establishes a Redis client connection
func ConnectToRedis() *redis.Client {
	// Define Redis config (adjust as needed)
	redisConfig := &redis.Config{
		ClusterMode: false,
		Hosts:       []string{constants.RedisPORT}, // Redis server address
		PoolSize:    10,
	}

	// Create Redis client
	redisClient := redis.NewClient(redisConfig)

	// Check if Redis client was created successfully
	if redisClient == nil {
		fmt.Println(" Failed to create Redis client")
		return nil
	}

	// Create a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Log the Redis connection attempt
	fmt.Println("Trying to connect to Redis...")

	// Use the correct Ping function
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println(" Error connecting to Redis:", err)
		return nil
	}

	fmt.Println(" Connected to Redis successfully!")
	return redisClient
}
