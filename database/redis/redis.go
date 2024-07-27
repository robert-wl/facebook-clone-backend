package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/yahkerobertkertasnya/TPAWebBack/helper"
)

var client *redis.Client

func GetInstance() *redis.Client {
	if client == nil {

		clientCurr := redis.NewClient(&redis.Options{
			Addr:     helper.GetDotENVVariable("REDIS_ADDRESS", "localhost:6379"),
			Password: "",
			DB:       0,
		})

		client = clientCurr
	}

	return client
}
