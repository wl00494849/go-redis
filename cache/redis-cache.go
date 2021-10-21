package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-redis/model"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisBasicCache struct {
	host    string
	db      int
	expires time.Duration
}

var ctx = context.Background()

func NewBasicRedisCache(host string, db int, exp time.Duration) *redisBasicCache {
	return &redisBasicCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisBasicCache) getClient() *redis.Client {
	fmt.Println(cache.host)
	return redis.NewClient(&redis.Options{
		Addr: cache.host,
		DB:   cache.db,
	})
}

func (cache *redisBasicCache) Set(key string, value *model.User) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	errCheck(err)

	client.Set(ctx, key, json, cache.expires*time.Second)
}

func (cache *redisBasicCache) Get(key string) *model.User {
	client := cache.getClient()
	user := &model.User{}

	val, _ := client.Get(ctx, key).Result()
	json.Unmarshal([]byte(val), user)
	return user
}

func (cache *redisBasicCache) Push(key string, value *[]model.User) {
	client := cache.getClient()
	for _, user := range *value {
		json, err := json.Marshal(user)
		errCheck(err)
		client.LPush(ctx, key, json)
	}
}

func (cache *redisBasicCache) Lrange(key string, start int64, stop int64) *[]model.User {
	client := cache.getClient()
	users := make([]model.User, 0)
	jsonString, err := client.LRange(ctx, key, start, stop).Result()
	errCheck(err)

	for _, str := range jsonString {
		user := &model.User{}
		json.Unmarshal([]byte(str), &user)
		users = append(users, *user)
	}

	return &users
}
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
