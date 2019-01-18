package gokeeper

import "testing"

func TestLRUCacheWithInvalidKey(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("should panic for unhashable key")
		}
	}()

	cache := NewLRUCache(0)
	cache.Set([]byte("a"), 1)
}

func TestLRUCacheSetZeroSize(t *testing.T) {
	cache := NewLRUCache(0)
	if cache.Size != defaultSize {
		t.Errorf(expectValue, defaultSize, cache.Size)
	}
}

func TestLRUCacheSetOneSize(t *testing.T) {
	size := 1
	cache := NewLRUCache(size)
	if cache.Size != size {
		t.Errorf(expectValue, size, cache.Size)
	}

	cache.Set("a", 1)
	cache.Set("b", 2)

	if len(cache.hashMap) != 1 {
		t.Errorf(expectValue, 1, len(cache.hashMap))
	}

	for k, v := range cache.hashMap {
		if k != "b" {
			t.Errorf(expectValue, "b", k)
		}

		if v.value != 2 {
			t.Errorf(expectValue, 1, v.value)
		}
	}

	if cache.head != cache.hashMap["b"] {
		t.Errorf(expectValue, "head is b", cache.head)
	}

	if cache.tail != cache.hashMap["b"] {
		t.Errorf(expectValue, "tail is b", cache.tail)
	}
}

func TestLRUCacheSetWithOneItem(t *testing.T) {
	size := 5
	cache := NewLRUCache(size)

	if cache.Size != size {
		t.Errorf(expectValue, size, cache.Size)
	}

	if cache.head != nil {
		t.Errorf(expectValue, nil, cache.head)
	}

	if cache.tail != nil {
		t.Errorf(expectValue, nil, cache.tail)
	}

	cache.Set("a", 1)

	if len(cache.hashMap) != 1 {
		t.Errorf(expectValue, 1, len(cache.hashMap))
	}

	for k, v := range cache.hashMap {
		if k != "a" {
			t.Errorf(expectValue, "a", k)
		}

		if v.value != 1 {
			t.Errorf(expectValue, 1, v.value)
		}
	}

	if cache.head != cache.hashMap["a"] {
		t.Errorf(expectValue, "head is a", cache.head)
	}

	if cache.tail != cache.hashMap["a"] {
		t.Errorf(expectValue, "tail is a", cache.tail)
	}

	// (head) b <-> a (tail)
	na := cache.hashMap["a"]
	if na.next != nil {
		t.Error(expectValue, nil, na.next)
	}
	if na.prev != nil {
		t.Error(expectValue, nil, na.prev)
	}
}

func TestLRUCacheSetWithTwoItems(t *testing.T) {
	size := 5
	cache := NewLRUCache(size)
	cache.Set("a", 1)
	cache.Set("b", 2)

	if len(cache.hashMap) != 2 {
		t.Errorf(expectValue, 2, len(cache.hashMap))
	}

	for k, v := range cache.hashMap {
		if k == "a" {
			if v.value != 1 {
				t.Errorf(expectValue, 1, v.value)
			}
		} else if k == "b" {
			if v.value != 2 {
				t.Errorf(expectValue, 2, v.value)
			}
		} else {
			t.Errorf(expectValue, "key should be a or b", k)
		}
	}

	if cache.head != cache.hashMap["b"] {
		t.Errorf(expectValue, "head is b", cache.head)
	}

	if cache.tail != cache.hashMap["a"] {
		t.Errorf(expectValue, "tail is a", cache.tail)
	}

	// (head) b <-> a (tail)
	na := cache.hashMap["a"]
	nb := cache.hashMap["b"]
	if na.next != nil {
		t.Error(expectValue, nil, na.next)
	}
	if na.prev != nb {
		t.Error(expectValue, nil, na.prev)
	}
	if nb.next != na {
		t.Error(expectValue, na, nb.next)
	}
	if nb.prev != nil {
		t.Error(expectValue, nil, nb.prev)
	}
}

func TestLRUCacheSetWithThreeItems(t *testing.T) {
	size := 5
	cache := NewLRUCache(size)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	if len(cache.hashMap) != 3 {
		t.Errorf(expectValue, 3, len(cache.hashMap))
	}

	for k, v := range cache.hashMap {
		if k == "a" {
			if v.value != 1 {
				t.Errorf(expectValue, 1, v.value)
			}
		} else if k == "b" {
			if v.value != 2 {
				t.Errorf(expectValue, 2, v.value)
			}
		} else if k == "c" {
			if v.value != 3 {
				t.Errorf(expectValue, 3, v.value)
			}
		} else {
			t.Errorf(expectValue, "key should be either of a, b, c", k)
		}
	}

	if cache.head != cache.hashMap["c"] {
		t.Errorf(expectValue, "head is c", cache.head)
	}

	if cache.tail != cache.hashMap["a"] {
		t.Errorf(expectValue, "tail is a", cache.tail)
	}

	// (head) c <-> b <-> a (tail)
	na := cache.hashMap["a"]
	nb := cache.hashMap["b"]
	nc := cache.hashMap["c"]
	if na.next != nil {
		t.Error(expectValue, nil, na.next)
	}
	if na.prev != nb {
		t.Error(expectValue, nil, na.prev)
	}
	if nb.next != na {
		t.Error(expectValue, na, nb.next)
	}
	if nb.prev != nc {
		t.Error(expectValue, nc, nb.prev)
	}
	if nc.next != nb {
		t.Error(expectValue, nb, nc.next)
	}
	if nc.prev != nil {
		t.Error(expectValue, nil, nc.prev)
	}
}

func TestLRUCacheSetWithExceedingSize(t *testing.T) {
	size := 3
	cache := NewLRUCache(size)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	cache.Set("d", 4)

	if len(cache.hashMap) != 3 {
		t.Errorf(expectValue, 3, len(cache.hashMap))
	}

	for k, v := range cache.hashMap {
		if k == "b" {
			if v.value != 2 {
				t.Errorf(expectValue, 2, v.value)
			}
		} else if k == "c" {
			if v.value != 3 {
				t.Errorf(expectValue, 3, v.value)
			}
		} else if k == "d" {
			if v.value != 4 {
				t.Errorf(expectValue, 4, v.value)
			}
		} else {
			t.Errorf(expectValue, "key should be either of b, c, d", k)
		}
	}

	if cache.head != cache.hashMap["d"] {
		t.Errorf(expectValue, "head is d", cache.head)
	}

	if cache.tail != cache.hashMap["b"] {
		t.Errorf(expectValue, "tail is b", cache.tail)
	}
}

func TestLRUCacheGetWithExistKey(t *testing.T) {
	size := 3
	cache := NewLRUCache(size)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	// now is (head) c <-> b <-> a (tail)
	if v, ok := cache.Get("a"); !ok || v != 1 {
		t.Errorf(expectValue, 1, v)
	}

	// now is (head) a <-> c <-> b (tail)
	val := cache.hashMap["a"]
	if val.prev != nil {
		t.Errorf(expectValue, nil, val.prev)
	}
	if val.next != cache.hashMap["c"] {
		t.Errorf(expectValue, "c", val.next)
	}

	if cache.head != cache.hashMap["a"] {
		t.Errorf(expectValue, "head is a", cache.head)
	}

	if cache.tail != cache.hashMap["b"] {
		t.Errorf(expectValue, "tail is b", cache.tail)
	}

	if v, ok := cache.Get("b"); !ok || v != 2 {
		t.Errorf(expectValue, 2, v)
	}

	// now is (head) b <-> a <-> c (tail)
	if cache.head != cache.hashMap["b"] {
		t.Errorf(expectValue, "head is b", cache.head)
	}

	if cache.tail != cache.hashMap["c"] {
		t.Errorf(expectValue, "tail is c", cache.tail)
	}

	if v, ok := cache.Get("c"); !ok || v != 3 {
		t.Errorf(expectValue, 3, v)
	}

	// now is (head) c <-> b <-> a (tail)
	if cache.head != cache.hashMap["c"] {
		t.Errorf(expectValue, "head is c", cache.head)
	}

	if cache.tail != cache.hashMap["a"] {
		t.Errorf(expectValue, "tail is a", cache.tail)
	}

	if v, ok := cache.Get("c"); !ok || v != 3 {
		t.Errorf(expectValue, 3, v)
	}

	// now is (head) c <-> b <-> a (tail)
	if cache.head != cache.hashMap["c"] {
		t.Errorf(expectValue, "head is c", cache.head)
	}

	if cache.tail != cache.hashMap["a"] {
		t.Errorf(expectValue, "tail is a", cache.tail)
	}

	// --------- Get node in the middle ---------

	if v, ok := cache.Get("b"); !ok || v != 2 {
		t.Errorf(expectValue, 2, v)
	}

	// now is (head) b <-> c <-> a (tail)
	if cache.head != cache.hashMap["b"] {
		t.Errorf(expectValue, "head is b", cache.head)
	}

	if cache.tail != cache.hashMap["a"] {
		t.Errorf(expectValue, "tail is a", cache.tail)
	}
}

func TestLRUCacheSetValueAgain(t *testing.T) {
	size := 5
	cache := NewLRUCache(size)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	// Set value in the head again
	cache.Set("c", 2)
	if cache.hashMap["c"].value != 2 {
		t.Errorf(expectValue, 2, cache.hashMap["c"].value)
	}

	if cache.head != cache.hashMap["c"] {
		t.Errorf(expectValue, "head is c", cache.head)
	}

	if cache.tail != cache.hashMap["a"] {
		t.Errorf(expectValue, "tail is a", cache.tail)
	}

	// Set value not in link list head
	cache.Set("b", 4)
	if cache.hashMap["b"].value != 4 {
		t.Errorf(expectValue, 4, cache.hashMap["4"].value)
	}

	if cache.head != cache.hashMap["b"] {
		t.Errorf(expectValue, "head is b", cache.head)
	}

	if cache.tail != cache.hashMap["a"] {
		t.Errorf(expectValue, "tail is a", cache.tail)
	}
}

func TestLRUCacheGetWithNonExistKey(t *testing.T) {
	size := 5
	cache := NewLRUCache(size)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	if v, ok := cache.Get("d"); v != nil || ok {
		t.Errorf(expectValue, nil, v)
	}
}
