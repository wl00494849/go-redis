package cache

import (
	"encoding/json"
	"fmt"
	"go-redis/model"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisPwdCache struct {
	host    string
	db      int
	expires time.Duration
	pwd     string
}

func NewPwdRedisCache(host string, db int, pwd string, exp time.Duration) *redisPwdCache {
	return &redisPwdCache{
		host:    host,
		db:      db,
		expires: exp,
		pwd:     pwd,
	}
}

func (cache *redisPwdCache) getClient() *redis.Client {
	fmt.Println(cache.host)
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		DB:       cache.db,
		Password: cache.pwd,
	})
}

func (cache *redisPwdCache) Set(key string, value *model.User) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	errCheck(err)

	client.Set(ctx, key, json, cache.expires*time.Second)
}

func (cache *redisPwdCache) Get(key string) *model.User {
	client := cache.getClient()
	user := &model.User{}

	val, _ := client.Get(ctx, key).Result()
	json.Unmarshal([]byte(val), user)
	return user
}

func (cache *redisPwdCache) Push(key string, value *model.User) {
	client := cache.getClient()
	value.Id = value.Id + 2
	json, err := json.Marshal(value)
	errCheck(err)

	client.LPush(ctx, key, json)
}

func (cache *redisPwdCache) Lrange(key string, start int64, stop int64) *[]model.User {
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
