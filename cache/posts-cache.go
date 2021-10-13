package cache

import "go-redis/model"

type PostCache interface {
	Set(key string, value *model.User)
	Get(key string) *model.User
	Push(key string, value *model.User)
	Lrange(key string, start int64, stop int64) *[]model.User
}
