package db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	//redisHost := os.Getenv("REDIS_HOST")
	//redisPort := os.Getenv("REDIS_PORT")

	redisHost := "localhost"
	redisPort := "6379"
	//redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}
