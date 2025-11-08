package eviction

import (
	"time"

	"github.com/loki4514/in-memory.git/internal/storage"
)

func LeastRecentlyUsed(c *storage.Cache) {
	c.MU.Lock()
	defer c.MU.Unlock()

	if len(c.Store) <= c.Config.Cache.MaxSize {
		return
	}

	var leastRecentKey string
	var leastTime time.Time
	first := true

	for key, entry := range c.Store {
		if first || entry.LastAccessed.Before(leastTime) {
			leastTime = entry.LastAccessed
			leastRecentKey = key
			first = false
		}
	}

	// Remove the least recently used entry
	delete(c.Store, leastRecentKey)
}
