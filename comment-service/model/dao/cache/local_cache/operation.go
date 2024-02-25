package localcache

import (
	"encoding/json"
	"time"
)

// Set sets a value in bigcache.
func Set(key string, value interface{}, duration time.Duration) error {
	serialized, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return localCache.Set(key, serialized)
}

// Get gets a value from bigcache.
func Get(key string, dest interface{}) error {
	entry, err := localCache.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(entry, dest)
}
