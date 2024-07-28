package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/database"
	"reflect"
	"strings"
)

var ctx = context.Background()

type RedisAdapter struct {
	client *redis.Client
}

func NewRedisCacheAdapter() *RedisAdapter {
	rdb := database.GetRedisInstance()

	return &RedisAdapter{
		client: rdb,
	}
}

func (r *RedisAdapter) Set(value interface{}, cachedAttributes []string) error {
	dataType := reflect.TypeOf(value)

	var cacheKey string
	for _, attr := range cachedAttributes {

		var attrValue interface{}
		if reflect.TypeOf(value).Kind() == reflect.Slice {
			if reflect.ValueOf(value).Len() == 0 {
				continue
			}
			attrValue = reflect.ValueOf(value).Index(0).FieldByName(attr).Interface()
		} else {
			attrValue = reflect.ValueOf(value).Elem().FieldByName(attr).Interface()
		}
		cacheKey += fmt.Sprintf(":%v", attrValue)
	}
	key := fmt.Sprintf("%s%s", dataType, cacheKey)
	key = strings.Replace(key, "[]", "Array", -1)

	jsonByte, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, jsonByte, 0).Err()
}

func (r *RedisAdapter) Get(keys []string, dest interface{}) error {
	dataType := reflect.TypeOf(dest)

	var cacheKey string
	for _, key := range keys {
		cacheKey += fmt.Sprintf("*%v", key)
	}

	key := fmt.Sprintf("%s:%s*", dataType, cacheKey)
	key = strings.Replace(key, "[]", "Array", -1)

	fmt.Println("KEY", key)

	keys, _, err := r.client.Scan(ctx, 0, key, 0).Result()
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	fmt.Println("GET KEY", keys[0])

	val, err := r.client.Get(ctx, keys[0]).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(val), dest); err != nil {
		return err
	}

	return nil
}

func (r *RedisAdapter) Gets(keys []string, dest []interface{}) error {
	dataType := reflect.TypeOf(dest)

	var cacheKey string
	for _, key := range keys {
		cacheKey += fmt.Sprintf("*%v", key)
	}

	key := fmt.Sprintf("%s:%s*", dataType, cacheKey)

	keys, _, err := r.client.Scan(ctx, 0, key, 0).Result()
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	val, err := r.client.Get(ctx, keys[0]).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(val), &dest); err != nil {
		return err
	}

	return nil
}

func (r *RedisAdapter) GetOrSet(keys []string, dest interface{}, cacheKeys []string, callback func() (interface{}, error)) error {
	fmt.Println("GETTING FROM CACHE", keys)
	if err := r.Get(keys, dest); err == nil && (dest != nil || reflect.ValueOf(dest).Len() > 0) {
		fmt.Println("RETURNING")
		return nil
	}

	fmt.Println("GETTING FROM DB", keys)
	value, err := callback()

	if err != nil {
		return err
	}

	if err := r.Set(value, cacheKeys); err != nil {
		return err
	}

	if err := r.Get(keys, dest); err != nil {
		return err
	}

	return nil
}

func (r *RedisAdapter) GetsOrSet(keys []string, dest []interface{}, cacheKeys []string, callback func() (interface{}, error)) error {
	if err := r.Gets(keys, dest); err == nil {
		return nil
	}

	value, err := callback()

	if err != nil {
		return err
	}

	if err := r.Set(value, cacheKeys); err != nil {
		return err
	}

	if err := r.Get(keys, dest); err != nil {
		return err
	}

	return nil
}

func (r *RedisAdapter) Del(key string) error {
	return r.client.Del(ctx, key).Err()
}
