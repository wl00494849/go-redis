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
	fmt.Println(cache.host)
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *model.User) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	errCheck(err)

	if data := client.Get(ctx, key); data == nil {
		client.Set(ctx, key, json, cache.expires*time.Second)
	}
}

func (cache *redisCache) Get(key string) *model.User {
	client := cache.getClient()
	user := &model.User{}

	val, _ := client.Get(ctx, key).Result()
	json.Unmarshal([]byte(val), user)
	return user
}

func (cache *redisCache) Push(key string, value *model.User) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	errCheck(err)

	client.LPush(ctx, key, json)
}

func (cache *redisCache) Lrange(key string, start int64, stop int64) *[]model.User {
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
