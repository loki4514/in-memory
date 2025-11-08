package storage

import (
	"sync"
	"time"

	"github.com/loki4514/in-memory.git/internal/config"
)

type CacheEntry struct {
	Value        interface{}
	ExpiresAt    time.Time
	LastAccessed time.Time
	Frequency    int
}

type Cache struct {
	MU     sync.RWMutex
	Store  map[string]*CacheEntry
	Config *config.Config
}
