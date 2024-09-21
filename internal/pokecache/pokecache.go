package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
	ticker   *time.Ticker
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type timed interface {
	Add()
	Get() ([]byte, bool)
	reapLoop()
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	c := &Cache{
		interval: interval,
		entries:  make(map[string]cacheEntry),
		ticker:   ticker,
	}
	go c.reapLoop(ticker)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}
func (c *Cache) reapLoop(t *time.Ticker) {
	for range t.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) StopReapLoop() {
	if c.ticker != nil {
		c.ticker.Stop()
	}
}
