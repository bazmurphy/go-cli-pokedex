package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	cache := NewCache(500 * time.Millisecond)

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
	interval := 500 * time.Millisecond
	cache := NewCache(interval)

	keyOneValue := []byte("value-one")

	cache.Add("key1", keyOneValue)

	value, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 should exist in the cache")
	}
	if !bytes.Equal(keyOneValue, value) {
		t.Errorf("expected %v | got %v", keyOneValue, value)
	}

	// sleep so the reap loop will remove the (expired) entry from the cache
	wait := 100 * time.Millisecond
	time.Sleep(interval + wait)

	if _, ok := cache.Get("key1"); ok {
		t.Error("key1 should have been removed from the cache")
	}
}
