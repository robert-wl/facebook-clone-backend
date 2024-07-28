package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/yahkerobertkertasnya/facebook-clone-backend/database"
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

func (r *RedisAdapter) getGeneralizedType(value interface{}) string {
	dataType := reflect.TypeOf(value)

	switch dataType.Kind() {
	case reflect.Int64:
		return "int"
	case reflect.Int32:
		return "int"
	case reflect.Int16:
		return "int"
	case reflect.Int8:
		return "int"
	case reflect.Int:
		return "int"
	case reflect.Uint64:
		return "uint"
	case reflect.Uint32:
		return "uint"
	case reflect.Uint16:
		return "uint"
	case reflect.Uint8:
		return "uint"
	case reflect.Uint:
		return "uint"
	case reflect.Float64:
		return "float"
	case reflect.Float32:
		return "float"
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "string"
	default:
		return "unknown"
	}
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

func (r *RedisAdapter) scanKeys(key string) ([]string, error) {
	var keys []string

	keys, _, err := r.client.Scan(ctx, 0, key, 0).Result()

	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return keys, nil
	}

	return keys, nil
}

func (r *RedisAdapter) Set(value interface{}, cacheKeys []string) error {
	dataType := r.getGeneralizedType(reflect.TypeOf(value))

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

	return r.client.Set(ctx, key, jsonByte, 10*60*time.Minute).Err()
}

func (r *RedisAdapter) Get(keys []string, dest interface{}) error {
	dataType := r.getGeneralizedType(reflect.TypeOf(dest))

	var cacheKey string
	for _, key := range keys {
		cacheKey += fmt.Sprintf("*%v", key)
	}

	key := fmt.Sprintf("%s:%s*", dataType, cacheKey)
	key = strings.Replace(key, "[]", "Array", -1)

	keys, err := r.scanKeys(key)

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

	keys, err := r.scanKeys(key)

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

func (r *RedisAdapter) GetOrSet(keys []string, dest interface{}, callback func() (interface{}, error), cacheKeys []string) error {
	if errG := r.Get(keys, dest); errG == nil && dest != nil {
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

func (r *RedisAdapter) Del(obj interface{}, deleteVal []string) error {
	if obj == nil {
		return nil
	}

	dataType := r.getGeneralizedType(reflect.TypeOf(obj))

	for _, val := range deleteVal {
		key := fmt.Sprintf("%s:*%s*", dataType, val)

		keys, err := r.scanKeys(key)

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

	return nil
}
