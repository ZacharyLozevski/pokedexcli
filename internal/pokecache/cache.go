package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache        map[string]cacheEntry
	mux          *sync.RWMutex
	interval 	 time.Duration
}

type cacheEntry struct {
	createdAt    time.Time 
	val 		 []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache: 	  make(map[string]cacheEntry),
		mux: 	  &sync.RWMutex{},
		interval: interval,
	}
	go cache.reapLoop()
	return cache
} 

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}	
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	if entry, ok := c.cache[key]; ok {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	// create a new ticker that creates a channel that sends every interval time of the cache
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()

		// loop over the entries for the cache 
		for key, entry := range c.cache {
			// if the time elapsed since the entry was created was greater than the interval
			// delete the entry
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)	
			}
		}

		c.mux.Unlock()
	}
}