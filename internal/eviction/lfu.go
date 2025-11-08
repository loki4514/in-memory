package eviction

import "github.com/loki4514/in-memory.git/internal/storage"

func LeastFrequentlyUsed(c *storage.Cache) {
	c.MU.Lock()
	defer c.MU.Unlock()

	if len(c.Store) <= c.Config.Cache.MaxSize {
		return
	}

	var leastFreqKey string
	var minFreq int
	first := true
	for key, entry := range c.Store {
		if first || entry.Frequency < minFreq {
			leastFreqKey = key
			minFreq = entry.Frequency
			first = false
		}
	}

	delete(c.Store, leastFreqKey)
}
