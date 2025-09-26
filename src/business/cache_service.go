package business

import (
	"encoding/json"
	"theatre-management-system/src/constants"
	"time"

	"github.com/patrickmn/go-cache"
)

// CacheService provides caching functionality for the application
type CacheService struct {
	cache *cache.Cache
}

// NewCacheService creates a new cache service
func NewCacheService() *CacheService {
	return &CacheService{
		cache: cache.New(
			time.Duration(constants.CacheDefaultExpiration)*time.Second,
			time.Duration(constants.CacheCleanupInterval)*time.Second,
		),
	}
}

// Set stores a value in the cache with the given key
func (cs *CacheService) Set(key string, value interface{}, duration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	cs.cache.Set(key, data, duration)
	return nil
}

// Get retrieves a value from the cache by key
func (cs *CacheService) Get(key string, dest interface{}) (bool, error) {
	data, found := cs.cache.Get(key)
	if !found {
		return false, nil
	}

	jsonData, ok := data.([]byte)
	if !ok {
		return false, nil
	}

	err := json.Unmarshal(jsonData, dest)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Delete removes a value from the cache
func (cs *CacheService) Delete(key string) {
	cs.cache.Delete(key)
}

// Clear removes all values from the cache
func (cs *CacheService) Clear() {
	cs.cache.Flush()
}

// GetCacheKey generates a cache key for different entities
func (cs *CacheService) GetCacheKey(entityType, operation string, params ...string) string {
	key := entityType + ":" + operation
	for _, param := range params {
		key += ":" + param
	}
	return key
}

// SetWithDefaultExpiration sets a value with the default cache expiration
func (cs *CacheService) SetWithDefaultExpiration(key string, value interface{}) error {
	return cs.Set(key, value, time.Duration(constants.CacheDefaultExpiration)*time.Second)
}

// GetStats returns cache statistics
func (cs *CacheService) GetStats() (int, int) {
	return cs.cache.ItemCount(), len(cs.cache.Items())
}
