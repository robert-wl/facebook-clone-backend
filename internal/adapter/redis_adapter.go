package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/internal/database"
	"reflect"
	"strings"
	"time"
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

func (r *RedisAdapter) getGeneralizedType(value interface{}) reflect.Type {
	dataType := reflect.TypeOf(value)

	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}

	if strings.HasPrefix(dataType.String(), "int") {
		return reflect.TypeOf(int64(0))
	}
	if strings.HasPrefix(dataType.String(), "uint") {
		return reflect.TypeOf(uint(0))
	}
	if strings.HasPrefix(dataType.String(), "float") {
		return reflect.TypeOf(float64(0))
	}
	return dataType

}

func (r *RedisAdapter) getValueByAttr(value interface{}, attr string) interface{} {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		if reflect.ValueOf(value).Len() == 0 {
			return nil
		}
		return reflect.ValueOf(value).Index(0).FieldByName(attr).Interface()
	}
	return reflect.ValueOf(value).Elem().FieldByName(attr).Interface()
}

func (r *RedisAdapter) scanKeys(key string, cursor uint64) ([]string, error) {
	var keys []string

	keys, newCursor, err := r.client.Scan(ctx, cursor, key, 0).Result()

	if err != nil {
		return nil, err
	}

	if newCursor != 0 {
		newKeys, err := r.scanKeys(key, newCursor)
		if err != nil {
			return nil, err
		}
		keys = append(keys, newKeys...)
	}

	return keys, nil
}

func (r *RedisAdapter) Set(value interface{}, cacheKeys []string, duration time.Duration) error {
	dataType := r.getGeneralizedType(value)

	var cacheKey string
	for _, key := range cacheKeys {
		cacheKey += fmt.Sprintf(":%v", key)
	}

	key := fmt.Sprintf("%s%s", dataType, cacheKey)
	key = strings.Replace(key, "[]", "Array", -1)
	jsonByte, err := json.Marshal(value)

	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, jsonByte, duration).Err()
}

func (r *RedisAdapter) Get(keys []string, dest interface{}) error {
	dataType := r.getGeneralizedType(dest)

	var cacheKey string
	for _, key := range keys {
		cacheKey += fmt.Sprintf("*%v", key)
	}

	key := fmt.Sprintf("%s:%s*", dataType, cacheKey)
	key = strings.Replace(key, "[]", "Array", -1)

	if key[0] == '*' {
		key = key[1:]
	}

	keys, err := r.scanKeys(key, 0)

	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return fmt.Errorf("NO KEYS")
	}

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

	keys, err := r.scanKeys(key, 0)

	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return fmt.Errorf("NO KEYS")
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

func (r *RedisAdapter) GetOrSet(keys []string, dest interface{}, callback func() (interface{}, error), duration time.Duration) error {
	if errG := r.Get(keys, dest); errG == nil && dest != nil {
		return nil
	}

	value, err := callback()

	if err != nil {
		return err
	}

	if err := r.Set(value, keys, duration); err != nil {
		return err
	}

	if err := r.Get(keys, dest); err != nil && dest != nil {
		return err
	}

	return nil
}

func (r *RedisAdapter) DelType(obj interface{}, deleteVal []string) error {
	if obj == nil {
		return nil
	}

	go func() {
		dataType := r.getGeneralizedType(obj)

		for _, val := range deleteVal {
			key := fmt.Sprintf("%s:*%s*", dataType, val)

			keys, err := r.scanKeys(key, 0)

			if err != nil {
				continue
			}

			if len(keys) == 0 {
				continue
			}

			for _, key := range keys {
				if err := r.client.Del(ctx, key).Err(); err != nil {
					continue
				}
			}
		}

	}()

	return nil
}

func (r *RedisAdapter) DelAllType(obj interface{}) error {
	if obj == nil {
		return nil
	}

	go func() {
		dataType := r.getGeneralizedType(obj)

		key := fmt.Sprintf("%s:*", dataType)

		keys, err := r.scanKeys(key, 0)

		if err != nil {
			return
		}

		if len(keys) == 0 {
			return
		}

		for _, key := range keys {
			if err := r.client.Del(ctx, key).Err(); err != nil {
				continue
			}
		}
	}()

	return nil
}

func (r *RedisAdapter) Del(keys []string) error {

	go func() {

		for _, val := range keys {
			key := fmt.Sprintf("*%s*", val)

			keys, err := r.scanKeys(key, 0)

			if err != nil {
				continue
			}

			if len(keys) == 0 {
				continue
			}

			for _, key := range keys {
				if err := r.client.Del(ctx, key).Err(); err != nil {
					continue
				}
			}
		}

	}()

	return nil
}

func (r *RedisAdapter) DelAll() error {
	keys, err := r.scanKeys("*", 0)

	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	for _, key := range keys {
		if err := r.client.Del(ctx, key).Err(); err != nil {
			continue
		}
	}

	return nil
}
