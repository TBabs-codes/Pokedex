package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(5*time.Minute)
	if cache.storage == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(5*time.Minute)
	
	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 not found after add")
	}

	if string(actual) != "val1" {
		t.Error("value doesn't match")
	}
}

func TestReapLoopCache(t *testing.T) {
	cache := NewCache(5*time.Second)
	
	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 not found after add")
	}

	if string(actual) != "val1" {
		t.Error("value doesn't match")
	}
	time.Sleep(3*time.Second)
	
	actual, ok = cache.Get("key1")
	if !ok {
		t.Error("key1 not found after 3 seconds")
	}

	time.Sleep(2*time.Second)

	actual, ok = cache.Get("key1")
	if ok {
		t.Error("key1 found after 5 seconds")
	}
}