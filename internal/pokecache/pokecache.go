package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Cache    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		Cache:    map[string]cacheEntry{},
		mu:       sync.Mutex{},
		interval: interval,
	}

	go c.reapLoop(interval)

	return &c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, value := range c.Cache {
			if time.Since(value.createdAt) > interval {
				delete(c.Cache, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	// Lock the map so that no other functions can write while adding new entries
	c.mu.Lock()
	c.Cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	value, ok := c.Cache[key]
	c.mu.Unlock()
	if !ok {
		return []byte{}, false
	}

	return value.val, ok
}
