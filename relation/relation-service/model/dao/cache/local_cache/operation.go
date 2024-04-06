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
	return GetClient().Set(key, serialized)
}

// Get gets a value from bigcache.
func Get(key string, dest interface{}) error {
	entry, err := GetClient().Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(entry, dest)
}

// Delete delete value from bigcache.
func Delete(key string) error {
	err := GetClient().Delete(key)
	return err
}
