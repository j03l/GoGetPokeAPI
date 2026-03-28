package cache

// Clear — wipes everything at once (you decided you didn't need this, but PokeGo uses it in tests for isolation).
// Runtime TTL change** — `SetExpiration()` lets you change the default TTL after creation.

import (
	"sync"
	"time"

	"github.com/j03l/GoGetPokeAPI/internal/logger"
)

const (
	defaultTTL = time.Minute * 15
)

type iCache interface {
	Add(key string, val []byte, ttt ...time.Duration)
	Get(key string) ([]byte, bool)
	SetActive(active bool)
	Active() bool
	OnAdd(callback func([]byte))
	OnEvicted(callback func([]byte))
}

type settings struct {
	defaultTTL time.Duration
	active     bool
}

type Cache struct {
	data     map[string]cacheEntry
	settings settings
	mu       sync.RWMutex

	onEvictedCallback func([]byte)
	onAddCallback     func([]byte)
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
	ttl       time.Duration
}

var CACHE *Cache
var log logger.ILogger = logger.LOG

func init() {
	CACHE = newCache()
	log.Info("Cache initialized")
}

func newCache(tickRate ...time.Duration) *Cache {
	var rate time.Duration
	if len(tickRate) > 0 {
		rate = tickRate[0]
	}
	if rate == 0 {
		rate = defaultTTL
	}
	cache := &Cache{
		data: make(map[string]cacheEntry),
		settings: settings{
			defaultTTL: rate,
			active:     true,
		},
	}
	go cache.reapLoop(rate)
	return cache
}

func (c *Cache) Add(key string, val []byte, ttl ...time.Duration) {
	var itemTTL time.Duration
	if len(ttl) > 0 {
		itemTTL = ttl[0]
	}
	c.mu.Lock()
	c.data[key] = cacheEntry{createdAt: time.Now(), val: val, ttl: itemTTL}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// Lazy expiration on Get isn't required for a static API like PokeAPI
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	if c.onAddCallback != nil {
		c.onAddCallback(entry.val)
	}
	return entry.val, true
}

func (c *Cache) deleteEntry(key string) ([]byte, bool) {
	c.mu.Lock()
	if ce, found := c.data[key]; found {
		delete(c.data, key)
		c.mu.Unlock()
		return ce.val, true
	}
	c.mu.Unlock()
	return nil, false
}

func (c *Cache) reapLoop(tickRate time.Duration) {
	tick := time.NewTicker(tickRate)
	var keys []string // keys to delete
	for {
		<-tick.C
		keys = nil // reset
		c.mu.RLock()
		for key, val := range c.data {
			curTTL := c.settings.defaultTTL
			if val.ttl != 0 {
				// extended item ttl
				curTTL = val.ttl
			}
			if time.Since(val.createdAt) > curTTL {
				keys = append(keys, key)
			}
		}
		c.mu.RUnlock()

		// delete all items older than ttl
		for _, key := range keys {
			ce, ok := c.deleteEntry(key)
			if ok && c.onEvictedCallback != nil {
				c.onEvictedCallback(ce)
			}
		}
	}
}

// For debug purposes
func (c *Cache) SetActive(active bool) {
	log.Warn("Cache active status changed", "active", active)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.settings.active = active
}

// Disables caching entirely without removing the cache. Useful for testing and debugging
func (c *Cache) Active() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.settings.active
}

// Wipes everything at once, for debugging and testing
func (c *Cache) Clear() {
	log.Warn("Cache cleared")
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[string]cacheEntry)
}

// Hooks
func (c *Cache) OnAdd(f func([]byte)) {
	c.onAddCallback = f
}

func (c *Cache) OnEvicted(f func([]byte)) {
	c.onEvictedCallback = f
}
