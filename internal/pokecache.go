package internal

import (
	"sync"
	"time"
)

type CacheEntry struct {
	val []byte
	createdAt time.Time
}


type Cache struct {
	mu *sync.RWMutex
	store map[string]CacheEntry
}


func NewCache(dur time.Duration) Cache {
	ch := Cache{
		mu: &sync.RWMutex{},
		store: map[string]CacheEntry{},
	}
	go ch.reapLoop(dur)
	return ch
}


func (c *Cache) Add (key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = CacheEntry{
		val: val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get (key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.store[key]
	if ok {
		return v.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	tk := time.NewTicker(interval)

	for range tk.C {
		c.mu.Lock()
		for k, v := range c.store {
			if time.Now().Add(interval).After(v.createdAt) {
				delete(c.store, k)
			}
		}
		c.mu.Unlock()
	}
}
