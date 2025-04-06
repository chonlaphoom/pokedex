package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

var cacheInstances []*Cache

func (c *Cache) Add(key string, val []byte) {
	c.reapLoop()

	c.mu.Lock()

	c.entries[key] = cacheEntry{
		createAt: time.Now(),
		val:      val,
	}

	defer c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.reapLoop()

	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.entries[key]; ok {
		fmt.Println("Cache hit")
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	for key, value := range c.entries {
		if time.Since(value.createAt) > c.interval {
			delete(c.entries, key)
		}
	}

	defer c.mu.Unlock()
}

func reapAllCache() {
	if len(cacheInstances) > 0 {
		for _, cache := range cacheInstances {
			cache.reapLoop()
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	reapAllCache()

	newCache := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	cacheInstances = append(cacheInstances, &newCache)
	return &newCache
}

func (c *Cache) PrintAddress() {
	fmt.Printf("%p", c)
}
