package cache

type PostCache struct {
	mc ManagerCache
}
type ManagerCache interface {
	Set(key string, value map[string]string)
	Get(key string) *map[string]string
	Push(key string, value *[]map[string]string)
	Lrange(key string, start int64, stop int64) *[]map[string]string
}

func NewPostCache(m ManagerCache) *PostCache {
	return &PostCache{mc: m}
}

func (p *PostCache) Set(key string, value map[string]string) {
	p.mc.Set(key, value)
}

func (p *PostCache) Get(key string) *map[string]string {
	return p.mc.Get(key)
}

func (p *PostCache) Push(key string, value *[]map[string]string) {
	p.mc.Push(key, value)
}

func (p *PostCache) Lrange(key string, start int64, stop int64) *[]map[string]string {
	return p.mc.Lrange(key, start, stop)
}
