package cache

import "go-redis/model"

type PostCache interface {
	Set(key string, value *model.User)
	Get(key string) *model.User
}
