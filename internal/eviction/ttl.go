package eviction

import (
	"time"

	"github.com/loki4514/in-memory.git/internal/storage"
)

func ExpiredAt(c *storage.Cache) {
	c.MU.Lock()
	defer c.MU.Unlock()

	var expiredKeys []string

	// Collect all expired keys
	for key, entry := range c.Store {
		if entry.ExpiresAt.Before(time.Now()) {
			expiredKeys = append(expiredKeys, key)
		}
	}

	// Delete all expired keys
	for _, key := range expiredKeys {
		delete(c.Store, key)
	}
}
