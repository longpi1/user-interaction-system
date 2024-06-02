package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// Set sets a string value with an expiration time
func Set(key string, value interface{}, expiration time.Duration) error {
	serialized, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return GetClient().Set(key, serialized, expiration).Err()
}

// Get retrieves a string value
func Get(key string, dest interface{}) error {
	val, err := GetClient().Get(key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

// HSet sets a hash field to a value
func HSet(key, field string, value interface{}) error {
	return GetClient().HSet(key, field, value).Err()
}

// HGet retrieves a value from a hash field
func HGet(key, field string) (string, error) {
	return GetClient().HGet(key, field).Result()
}

// LPush pushes a value onto the head of a list
func LPush(key string, values ...interface{}) error {
	return GetClient().LPush(key, values...).Err()
}

// RPush pushes a value onto the tail of a list
func RPush(key string, values ...interface{}) error {
	return GetClient().RPush(key, values...).Err()
}

// LRange retrieves a range of elements from a list
func LRange(key string, start, stop int64) ([]string, error) {
	return GetClient().LRange(key, start, stop).Result()
}

// SAdd adds members to a set
func SAdd(key string, members ...interface{}) error {
	return GetClient().SAdd(key, members...).Err()
}

// SMembers retrieves all members of a set
func SMembers(key string) ([]string, error) {
	return GetClient().SMembers(key).Result()
}

// ZAdd adds a member to a sorted set with a score
func ZAdd(key string, members ...redis.Z) error {
	return GetClient().ZAdd(key, members...).Err()
}

// ZRange retrieves a range of members from a sorted set
func ZRange(key string, start, stop int64) ([]string, error) {
	return GetClient().ZRange(key, start, stop).Result()
}

// Del deletes keys
func Del(keys ...string) error {
	return GetClient().Del(keys...).Err()
}
