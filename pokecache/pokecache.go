package pokecache

import (
	"sync"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry, ok := c.cache[key]

	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {

	if interval == 0 {
		interval = time.Second * 5 // default to 5 seconds
	}

	tick := time.NewTicker(interval)

	for range tick.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for key, entry := range c.cache {
		if now.Sub(entry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}
