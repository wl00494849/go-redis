package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisDistributedCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewDistributedRedisCache(host string, db int, exp time.Duration) *redisDistributedCache {
	return &redisDistributedCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisDistributedCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cache.host,
		DB:   cache.db,
	})
}

func (cache *redisDistributedCache) Set(key string, value *map[string]string) {

}

func (cache *redisDistributedCache) Get(key string) *map[string]string {
	client := cache.getClient()
	user := &map[string]string{}

	val, _ := client.Get(ctx, key).Result()
	json.Unmarshal([]byte(val), user)
	return user
}
