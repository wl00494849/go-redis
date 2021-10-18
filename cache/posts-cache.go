package cache

import "go-redis/model"

type PostCache struct {
	mc ManagerCache
}
type ManagerCache interface {
	Set(key string, value *model.User)
	Get(key string) *model.User
	Push(key string, value *model.User)
	Lrange(key string, start int64, stop int64) *[]model.User
}

func NewPostCache(m ManagerCache) *PostCache {
	return &PostCache{mc: m}
}

func (p PostCache) Set(key string, value *model.User) {
	p.mc.Set(key, value)
}

func (p PostCache) Get(key string) *model.User {
	return p.mc.Get(key)
}

func (p PostCache) Push(key string, value *model.User) {
	p.mc.Push(key, value)
}

func (p PostCache) Lrange(key string, start int64, stop int64) *[]model.User {
	return p.mc.Lrange(key, start, stop)
}
