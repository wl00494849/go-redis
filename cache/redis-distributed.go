package cache

import (
	"encoding/json"
	"fmt"
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

func (cache *redisDistributedCache) Set(key string, value map[string]string) {
	client := cache.getClient()

	exist, _ := json.Marshal(cache.Get(key))
	val, _ := json.Marshal(value)

	if string(val) == string(exist) {
		client.Expire(ctx, key, 300)
		fmt.Println("key reset time")
	} else if string(exist) == "{}" {
		client.Set(ctx, key, val, cache.expires)
		fmt.Println("key reset")
	} else {
		fmt.Println("key exist")
	}
}

func (cache *redisDistributedCache) Get(key string) *map[string]string {
	client := cache.getClient()
	user := make(map[string]string)

	val, _ := client.Get(ctx, key).Result()
	json.Unmarshal([]byte(val), &user)
	return &user
}

func (cache *redisDistributedCache) Push(key string, value *[]map[string]string) {

}

func (cache *redisDistributedCache) Lrange(key string, start int64, stop int64) *[]map[string]string {
	data := make([]map[string]string, 0)
	return &data
}
