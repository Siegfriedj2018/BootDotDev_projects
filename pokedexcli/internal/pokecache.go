package internal

import (
	"log"
	"sync"
	"time"
)


type Cache struct {
	mu       sync.Mutex
	cash     map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	log.Println("Creating Cash...")
	c := &Cache{
		mu:       sync.Mutex{},
		cash:     make(map[string]cacheEntry),
		interval: interval,
	}
	
	// this is a blocking go routine
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	log.Println("adding cash")
	
	c.cash[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	log.Print("getting cash")
	data, exists := c.cash[key]

	return data.val, exists
}

func (c *Cache) reapLoop() {
	
	cleanUpTick := time.Tick(c.interval)

	for range cleanUpTick {
		c.mu.Lock()
		for key, val := range c.cash {
			duration := time.Since(val.createdAt)
			if duration >= c.interval {
				delete(c.cash, key)
			}
		}
		c.mu.Unlock()
	}
}
