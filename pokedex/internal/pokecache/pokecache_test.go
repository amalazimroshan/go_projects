package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		}, {
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s value dosen't match %s", cas.inputKey, cas.inputVal)
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyone := "key1"
	cache.Add(keyone, []byte("val1"))
	time.Sleep(interval / 2)

	_, ok := cache.Get(keyone)
	if !ok {
		t.Errorf("%s shouldn't have been reaped", keyone)
	}

	time.Sleep(interval + time.Millisecond)
	_, ok = cache.Get(keyone)
	if ok {
		t.Errorf("%s should have been reaped", keyone)
	}

}
