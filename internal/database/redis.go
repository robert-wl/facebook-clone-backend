package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/utils"
)

var ctx = context.Background()
var redisClient *redis.Client

func GetRedisInstance() *redis.Client {
	if redisClient == nil {

		host := utils.GetDotENVVariable("REDIS_HOST", "localhost")
		port := utils.GetDotENVVariable("REDIS_PORT", "6379")
		password := utils.GetDotENVVariable("REDIS_PASSWORD", "")

		addr := fmt.Sprintf("%s:%s", host, port)

		fmt.Println("Connecting to Redis at", addr)
		clientCurr := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		})

		if err := clientCurr.Ping(ctx).Err(); err != nil {
			panic(err)
		}

		redisClient = clientCurr
	}

	return redisClient
}
