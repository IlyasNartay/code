package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache struct holds a reference to the Redis client
type RedisCache struct {
	Client *redis.Client
}

// NewRedisCache creates a new instance of RedisCache
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		Client: client,
	}
}

// Get retrieves data from Redis cache by key
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil // Cache miss
		}
		return "", err
	}
	return val, nil
}

// Set sets data in Redis cache with an expiration time (TTL)
func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = r.Client.Set(ctx, key, data, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// Example usage:

// Initialize Redis client
// redisClient := redis.NewClient(&redis.Options{
// 	Addr:     "localhost:6379",
// 	Password: "", // no password set
// 	DB:       0,  // use default DB
// })

// // Initialize RedisCache
// cache := NewRedisCache(redisClient)

// // Example of getting data from cache
// ctx := context.Background()
// key := "example_key"
// data, err := cache.Get(ctx, key)
// if err != nil {
// 	fmt.Println("Error:", err)
// } else if data != "" {
// 	fmt.Println("Data retrieved from cache:", data)
// } else {
// 	fmt.Println("Cache miss")
// }

// // Example of setting data in cache
// value := "example_value"
// expiration := time.Hour
// err = cache.Set(ctx, key, value, expiration)
// if err != nil {
// 	fmt.Println("Error:", err)
// } else {
// 	fmt.Println("Data set in cache")
// }
