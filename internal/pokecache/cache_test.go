package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5
	cases := []struct {
		key   string
		val[] byte
	}{
		{
			key: "https://testing.com/v1/3ae4",
			val: []byte("testing data"),
		},
		{
			key: "http://example.com/v3/3as",
			val: []byte("more testing data"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case: %d", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find the value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5
	delay := time.Millisecond * 10 + time.Second * interval
	cache := NewCache(interval)
	cache.Add("https://example.com", []byte("testdata"))

	if _, ok := cache.Get("https://example.com"); !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(delay)

	if _, ok := cache.Get("https://example.com"); ok {
		t.Errorf("expected to not find key")
		return
	}
}