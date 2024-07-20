package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheMap: map[string]cacheEntry{},
		mu:       &sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Time{},
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if v, ok := c.cacheMap[key]; ok {
		return v.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for tick := range ticker.C {
		for key, value := range c.cacheMap {
			if tick.Sub(value.createdAt) > 5 {
				c.mu.Lock()
				delete(c.cacheMap, key)
				c.mu.Unlock()
			}
		}
	}
}
