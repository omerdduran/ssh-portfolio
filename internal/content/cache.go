package content

import (
	"log"
	"sync"
	"time"
)

var (
	cacheMu   sync.RWMutex
	cached    *SiteContent
	cachedAt  time.Time
	defaultTTL = 5 * time.Minute
)

// Get returns cached content, refreshing if stale. Safe for concurrent use.
func Get() *SiteContent {
	cacheMu.RLock()
	if cached != nil && time.Since(cachedAt) < defaultTTL {
		defer cacheMu.RUnlock()
		return cached
	}
	cacheMu.RUnlock()

	cacheMu.Lock()
	defer cacheMu.Unlock()

	// Double-check after acquiring write lock
	if cached != nil && time.Since(cachedAt) < defaultTTL {
		return cached
	}

	sc, err := FetchAll()
	if err != nil {
		log.Printf("Content fetch failed: %v", err)
		if cached != nil {
			return cached // return stale data
		}
		// Fall back to embedded data
		sc = LoadFallback()
	}

	cached = sc
	cachedAt = time.Now()
	return cached
}

// Warm pre-populates the cache. Call at server startup.
func Warm() {
	_ = Get()
}
