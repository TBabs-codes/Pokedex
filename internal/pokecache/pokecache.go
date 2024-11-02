package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	storage map[string]cacheEntry
	safe    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		storage: make(map[string]cacheEntry),
	}

	go newCache.reapLoop(interval)
	return newCache
}

func (c Cache) Add(key string, val []byte) {
	c.safe.Lock()
	defer c.safe.Unlock()
	c.storage[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.safe.Lock()
	defer c.safe.Unlock()
	if entry, ok := c.storage[key]; ok {
		return entry.val, true
	}

	return []byte{}, false
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.safe.Lock()
		for key, cache := range c.storage {
			if time.Since(cache.createdAt) > interval {
				delete(c.storage, key)
			}
		}
		c.safe.Unlock()
	}
}
