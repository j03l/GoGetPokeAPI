package cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := newCache(interval)
			cache.Add(c.key, c.val, 0)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := newCache(baseTime)
	cache.OnAdd(func(val []byte) {
		fmt.Printf("ADD: %+v\n", val)
	})
	cache.OnEvicted(func(val []byte) {
		fmt.Printf("DELETE: %+v\n", val)
	})
	cache.Add("https://example.com", []byte("testdata"), baseTime)

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

var cache ICache = newCache()

func TestNewCache(t *testing.T) {
	assert.IsType(t, &Cache{}, cache)
}

func TestSet(t *testing.T) {
	cache.Set("key", "value")
	value, ok := cache.Get("key")
	assert.Equal(t, "value", value, "Unexpected value in cache")
	assert.True(t, ok)
}

func TestGet(t *testing.T) {
	cache.Set("key", "value")
	value, ok := cache.Get("key")
	assert.Equal(t, "value", value, "Unexpected value in cache")
	assert.True(t, ok)
}

func TestClear(t *testing.T) {
	cache.Set("key", "value")
	cache.Clear()
	value, ok := cache.Get("key")
	assert.Equal(t, nil, value, "Expected value to be nil")
	assert.False(t, ok)
}

func TestDelete(t *testing.T) {
	cache.Set("key", "value")
	cache.Delete("key")
	value, ok := cache.Get("key")
	assert.Equal(t, nil, value, "Expected value to be nil")
	assert.False(t, ok)
}

func TestSetExpiration(t *testing.T) {
	cache.SetExpiration(1 * time.Second)
	expiration := cache.Expiration()
	cache.Set("key", "value")
	time.Sleep(2 * time.Second)
	value, ok := cache.Get("key")
	assert.Equal(t, nil, value, "Expected value to be nil")
	assert.Equal(t, 1*time.Second, expiration, "Unexpected expiration time")
	assert.False(t, ok)
}

func TestGetActive(t *testing.T) {
	assert.True(t, cache.Active(), "Expected initial active value to be true")
}

func TestSetActive(t *testing.T) {
	cache.SetActive(false)
	assert.False(t, cache.Active(), "Expected active value to be false")
}
