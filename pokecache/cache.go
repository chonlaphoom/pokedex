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
	mu       *sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createAt: time.Now(),
		val:      val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.entries[key]; ok {
		fmt.Println("Cache hit")
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		// WARN: lock mutex outside this section can cause deadlock
		// because for range through channel will wait until channel is closed
		c.mu.Lock()
		for key, value := range c.entries {
			if time.Since(value.createAt) > c.interval {
				delete(c.entries, key)
				fmt.Println("cache delete: ", key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		mu:       &sync.Mutex{},
	}

	go newCache.reapLoop()
	return &newCache
}

func (c *Cache) PrintAddress() {
	fmt.Printf("%p", c)
}
