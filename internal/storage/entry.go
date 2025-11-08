package storage

import (
	"time"
)

func (c *Cache) Set(key string, value interface{}, ttl time.Time) {
	c.MU.Lock()
	defer c.MU.Unlock()

	if c.Store == nil {
		c.Store = make(map[string]*CacheEntry) // initialize if needed
	}

	c.Store[key] = &CacheEntry{
		Value:     value,
		ExpiresAt: ttl,
	}
}

func (c *Cache) Get(key string) (*CacheEntry, bool) {
	c.MU.RLock()
	defer c.MU.RUnlock()

	cache, ok := c.Store[key]
	if ok {
		// check expiry if needed
		if !cache.ExpiresAt.IsZero() && cache.ExpiresAt.Before(time.Now()) {
			return nil, false
		}
		return cache, true
	}

	return nil, false // return nil pointer when key not found
}

func (c *Cache) GetAndUpdate(key string) (*CacheEntry, bool) {
	c.MU.Lock()
	defer c.MU.Unlock()

	cache, ok := c.Store[key]
	if ok {
		// check if expired
		if cache.ExpiresAt.IsZero() || cache.ExpiresAt.After(time.Now()) {
			cache.Frequency++
			cache.LastAccessed = time.Now()
			return cache, true
		}

		// expired -> remove from cache
		delete(c.Store, key)
	}

	return nil, false // return nil pointer when not found or expired
}

func (c *Cache) CleanupExpired() { // uppercase first letter
	c.MU.Lock()
	defer c.MU.Unlock()
	now := time.Now()
	for key, entry := range c.Store {
		if !entry.ExpiresAt.IsZero() && entry.ExpiresAt.Before(now) {
			delete(c.Store, key)
		}
	}
}
