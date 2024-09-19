package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type timed interface {
	Add() error
	Get() ([]byte, bool) // NOTE: should this have an error? or is the bool the error state (found/not found)?
	reapLoop() error
}

func NewCache(interval time.Duration) {

}

func (c *Cache) Add(key string, val []byte) error {
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	return []byte{}, false
}
func (c *Cache) reapLoop() error {
	return nil
}
