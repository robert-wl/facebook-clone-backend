package database

import (
	"github.com/redis/go-redis/v9"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/utils"
)

var redisClient *redis.Client

func GetRedisInstance() *redis.Client {
	if redisClient == nil {

		clientCurr := redis.NewClient(&redis.Options{
			Addr:     utils.GetDotENVVariable("REDIS_ADDRESS", "localhost:6379"),
			Password: "",
			DB:       0,
		})

		redisClient = clientCurr
	}

	return redisClient
}
