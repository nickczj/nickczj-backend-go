package cache

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Cache provides a simple caching interface with memcached backend
// and local in-memory fallback for development
type Cache struct {
	client     *memcache.Client
	local      map[string]cacheEntry
	localMu    sync.RWMutex
	expiration int32 // seconds
}

type cacheEntry struct {
	data      []byte
	expiresAt time.Time
}

var instance *Cache

// Init initializes the cache based on environment
func Init() {
	if viper.GetString("APP_ENV") == "prod" {
		servers := viper.GetString("memcached.servers")
		if servers == "" {
			log.Error("Memcached servers not configured")
			initLocalCache()
			return
		}

		serverList := strings.Split(servers, ",")
		client := memcache.New(serverList...)
		client.Timeout = 100 * time.Millisecond

		// Test connection
		err := client.Ping()
		if err != nil {
			log.Error("Error connecting to memcached: ", err)
			initLocalCache()
			return
		}

		instance = &Cache{
			client:     client,
			expiration: 3600, // 1 hour default
		}
		log.Info("Initialized memcached: ", servers)
	} else {
		initLocalCache()
	}
}

func initLocalCache() {
	instance = &Cache{
		local:      make(map[string]cacheEntry),
		expiration: 600, // 10 minutes for local
	}
	log.Info("Initialized local in-memory cache")
}

// Get retrieves a value from cache
func Get(key string, dest any) error {
	if instance == nil {
		return memcache.ErrCacheMiss
	}

	if instance.client != nil {
		item, err := instance.client.Get(key)
		if err != nil {
			return err
		}
		return json.Unmarshal(item.Value, dest)
	}

	// Local cache
	instance.localMu.RLock()
	entry, ok := instance.local[key]
	instance.localMu.RUnlock()

	if !ok || time.Now().After(entry.expiresAt) {
		if ok {
			// Clean up expired entry
			instance.localMu.Lock()
			delete(instance.local, key)
			instance.localMu.Unlock()
		}
		return memcache.ErrCacheMiss
	}

	return json.Unmarshal(entry.data, dest)
}

// Set stores a value in cache
func Set(key string, value any) error {
	if instance == nil {
		return nil
	}

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if instance.client != nil {
		return instance.client.Set(&memcache.Item{
			Key:        key,
			Value:      data,
			Expiration: instance.expiration,
		})
	}

	// Local cache
	instance.localMu.Lock()
	instance.local[key] = cacheEntry{
		data:      data,
		expiresAt: time.Now().Add(time.Duration(instance.expiration) * time.Second),
	}
	instance.localMu.Unlock()

	return nil
}

// GetOrSet retrieves from cache or calls the loader function on miss
func GetOrSet[T any](key string, loader func() (T, error)) (T, error) {
	var result T

	// Try cache first
	err := Get(key, &result)
	if err == nil {
		return result, nil
	}

	// Cache miss - call loader
	if err == memcache.ErrCacheMiss {
		log.Debug("Cache miss: ", key)
	}

	result, err = loader()
	if err != nil {
		return result, err
	}

	// Store in cache (ignore errors)
	if setErr := Set(key, result); setErr != nil {
		log.Warn("Failed to cache result: ", setErr)
	}

	return result, nil
}

// GenerateKey creates a cache key from multiple parts
func GenerateKey(args ...string) string {
	return strings.Join(args, "-")
}

// IsAvailable returns true if cache is initialized
func IsAvailable() bool {
	return instance != nil
}
