package pokicache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheData map[string]cacheEntry
	mtx       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {

	c := Cache{
		cacheData: make(map[string]cacheEntry),
		mtx:       &sync.Mutex{},
	}

	go c.clearCache(interval)

	return c

}

func (c *Cache) Add(key string, val []byte) {

	c.mtx.Lock()
	c.cacheData[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
	c.mtx.Unlock()

}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mtx.Lock()
	data, ok := c.cacheData[key]
	c.mtx.Unlock()

	return data.val, ok

}

func (c *Cache) clearCache(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {

	diff := time.Now().UTC().Add(interval)

	for k, v := range c.cacheData {
		if v.createdAt.Before(diff) {
			c.mtx.Lock()
			delete(c.cacheData, k)
			c.mtx.Unlock()
		}
	}
}
