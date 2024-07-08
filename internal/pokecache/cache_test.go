package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	cache := NewCache(3 * time.Second)

	keyOneValue := []byte("value-one")
	keyTwoValue := []byte("value-two")

	cache.Add("key1", keyOneValue)
	cache.Add("key2", keyTwoValue)

	value, ok := cache.Get("key1")
	if !ok {
		t.Errorf("key1 should exist in the cache")
	}
	if !bytes.Equal(keyOneValue, value) {
		t.Errorf("expected %v | got %v", keyOneValue, value)
	}

	value, ok = cache.Get("key2")
	if !ok {
		t.Errorf("key2 should exist in the cache")
	}
	if !bytes.Equal(keyTwoValue, value) {
		t.Errorf("expected %v | got %v", keyOneValue, value)
	}
}

func TestCacheReapLoop(t *testing.T) {
	interval := 100 * time.Millisecond
	cache := NewCache(interval)

	cache.Add("key1", []byte("value-one"))
	cache.Add("key2", []byte("value-two"))

	time.Sleep(interval + time.Second)

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should have been removed from the cache")
	}
	if _, ok := cache.Get("key2"); ok {
		t.Error("key2 should have been removed from the cache")
	}
}
