package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-redis/model"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

var ctx = context.Background()

func NewRedisCache(host string, db int, exp time.Duration) *redisCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *model.User) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	fmt.Println(cache.host)
	fmt.Println(cache.db)

	client.Set(ctx, key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *model.User {
	client := cache.getClient()
	user := &model.User{}

	val, _ := client.Get(ctx, key).Result()
	json.Unmarshal([]byte(val), user)
	return user
}
