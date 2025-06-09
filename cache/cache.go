package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(addr, password string, ttl time.Duration) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &RedisCache{
		client: rdb,
		ttl:    ttl,
	}
}

func (r *RedisCache) Get(key string) (interface{}, bool) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil || err != nil {
		return nil, false
	}

	var result interface{}
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		return nil, false
	}
	return result, true
}

func (r *RedisCache) Set(key string, value interface{}) {
	ctx := context.Background()
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return
	}
	r.client.Set(ctx, key, jsonValue, r.ttl)
}
