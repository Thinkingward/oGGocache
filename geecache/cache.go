package geecache

import (
	"geecache/lru"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

// 封装add
func (c *cache) add(key string, value *ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		if lru.DefaultMaxBytes > c.cacheBytes {
			c.lru = lru.New(lru.DefaultMaxBytes, nil)
		} else {
			c.lru = lru.New(c.cacheBytes, nil)
		}
	}
	c.lru.Add(key, value, value.Expire())
}

// 封装get
func (c *cache) get(key string) (value *ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(*ByteView), ok
	}
	return
}
