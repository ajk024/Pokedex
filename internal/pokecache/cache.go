package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry    map[string]cacheEntry //map key is url
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time //time when entry was created
	val       []byte    //represents the raw data we are caching
}

func NewCache(interval time.Duration) *Cache {
	//creates a new cache with a configurable interval (time.Duration)
	cache := &Cache{
		entry:    map[string]cacheEntry{},
		interval: interval,
		//mu field is zero-initialized
	}

	go cache.reapLoop() //runs in background
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	//fmt.Printf("Adding entry to pokeCache: %s\n", key)
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	//fmt.Printf("Getting entry from pokeCache: %s\n", key)
	//return bool is true if the entry was found else false
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entry[key]
	if !ok {
		return nil, ok //ok = false in this case
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	//fmt.Println("Starting reapLoop()")

	for range ticker.C { //C is a channel that receives a value every time the interval elapses
		c.mu.Lock()
		for key, entry := range c.entry {
			elapsed := time.Since(entry.createdAt)

			if elapsed > c.interval {
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}
