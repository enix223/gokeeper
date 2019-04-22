package cache

import "testing"

func TestLRUcachesWithInvalidKey(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("should panic for unhashable key")
		}
	}()

	caches := NewLRUCache(0)
	caches.Set([]byte("a"), 1)
}

func TestLRUcachesSetZeroSize(t *testing.T) {
	caches := NewLRUCache(0)
	if caches.Size != defaultSize {
		t.Errorf("expect: %v, got: %v", defaultSize, caches.Size)
	}
}

func TestLRUcachesSetOneSize(t *testing.T) {
	size := 1
	caches := NewLRUCache(size)
	if caches.Size != size {
		t.Errorf("expect: %v, got: %v", size, caches.Size)
	}

	caches.Set("a", 1)
	caches.Set("b", 2)

	if len(caches.hashMap) != 1 {
		t.Errorf("expect: %v, got: %v", 1, len(caches.hashMap))
	}

	for k, v := range caches.hashMap {
		if k != "b" {
			t.Errorf("expect: %v, got: %v", "b", k)
		}

		if v.value != 2 {
			t.Errorf("expect: %v, got: %v", 1, v.value)
		}
	}

	if caches.head != caches.hashMap["b"] {
		t.Errorf("expect: %v, got: %v", "head is b", caches.head)
	}

	if caches.tail != caches.hashMap["b"] {
		t.Errorf("expect: %v, got: %v", "tail is b", caches.tail)
	}
}

func TestLRUcachesSetWithOneItem(t *testing.T) {
	size := 5
	caches := NewLRUCache(size)

	if caches.Size != size {
		t.Errorf("expect: %v, got: %v", size, caches.Size)
	}

	if caches.head != nil {
		t.Errorf("expect: %v, got: %v", nil, caches.head)
	}

	if caches.tail != nil {
		t.Errorf("expect: %v, got: %v", nil, caches.tail)
	}

	caches.Set("a", 1)

	if len(caches.hashMap) != 1 {
		t.Errorf("expect: %v, got: %v", 1, len(caches.hashMap))
	}

	for k, v := range caches.hashMap {
		if k != "a" {
			t.Errorf("expect: %v, got: %v", "a", k)
		}

		if v.value != 1 {
			t.Errorf("expect: %v, got: %v", 1, v.value)
		}
	}

	if caches.head != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "head is a", caches.head)
	}

	if caches.tail != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "tail is a", caches.tail)
	}

	// (head) b <-> a (tail)
	na := caches.hashMap["a"]
	if na.next != nil {
		t.Errorf("expect: %v, got: %v", nil, na.next)
	}
	if na.prev != nil {
		t.Errorf("expect: %v, got: %v", nil, na.prev)
	}
}

func TestLRUcachesSetWithTwoItems(t *testing.T) {
	size := 5
	caches := NewLRUCache(size)
	caches.Set("a", 1)
	caches.Set("b", 2)

	if len(caches.hashMap) != 2 {
		t.Errorf("expect: %v, got: %v", 2, len(caches.hashMap))
	}

	for k, v := range caches.hashMap {
		if k == "a" {
			if v.value != 1 {
				t.Errorf("expect: %v, got: %v", 1, v.value)
			}
		} else if k == "b" {
			if v.value != 2 {
				t.Errorf("expect: %v, got: %v", 2, v.value)
			}
		} else {
			t.Errorf("expect: %v, got: %v", "key should be a or b", k)
		}
	}

	if caches.head != caches.hashMap["b"] {
		t.Errorf("expect: %v, got: %v", "head is b", caches.head)
	}

	if caches.tail != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "tail is a", caches.tail)
	}

	// (head) b <-> a (tail)
	na := caches.hashMap["a"]
	nb := caches.hashMap["b"]
	if na.next != nil {
		t.Errorf("expect: %v, got: %v", nil, na.next)
	}
	if na.prev != nb {
		t.Errorf("expect: %v, got: %v", nil, na.prev)
	}
	if nb.next != na {
		t.Errorf("expect: %v, got: %v", na, nb.next)
	}
	if nb.prev != nil {
		t.Errorf("expect: %v, got: %v", nil, nb.prev)
	}
}

func TestLRUcachesSetWithThreeItems(t *testing.T) {
	size := 5
	caches := NewLRUCache(size)
	caches.Set("a", 1)
	caches.Set("b", 2)
	caches.Set("c", 3)

	if len(caches.hashMap) != 3 {
		t.Errorf("expect: %v, got: %v", 3, len(caches.hashMap))
	}

	for k, v := range caches.hashMap {
		if k == "a" {
			if v.value != 1 {
				t.Errorf("expect: %v, got: %v", 1, v.value)
			}
		} else if k == "b" {
			if v.value != 2 {
				t.Errorf("expect: %v, got: %v", 2, v.value)
			}
		} else if k == "c" {
			if v.value != 3 {
				t.Errorf("expect: %v, got: %v", 3, v.value)
			}
		} else {
			t.Errorf("expect: %v, got: %v", "key should be either of a, b, c", k)
		}
	}

	if caches.head != caches.hashMap["c"] {
		t.Errorf("expect: %v, got: %v", "head is c", caches.head)
	}

	if caches.tail != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "tail is a", caches.tail)
	}

	// (head) c <-> b <-> a (tail)
	na := caches.hashMap["a"]
	nb := caches.hashMap["b"]
	nc := caches.hashMap["c"]
	if na.next != nil {
		t.Errorf("expect: %v, got: %v", nil, na.next)
	}
	if na.prev != nb {
		t.Errorf("expect: %v, got: %v", nil, na.prev)
	}
	if nb.next != na {
		t.Errorf("expect: %v, got: %v", na, nb.next)
	}
	if nb.prev != nc {
		t.Errorf("expect: %v, got: %v", nc, nb.prev)
	}
	if nc.next != nb {
		t.Errorf("expect: %v, got: %v", nb, nc.next)
	}
	if nc.prev != nil {
		t.Errorf("expect: %v, got: %v", nil, nc.prev)
	}
}

func TestLRUcachesSetWithExceedingSize(t *testing.T) {
	size := 3
	caches := NewLRUCache(size)
	caches.Set("a", 1)
	caches.Set("b", 2)
	caches.Set("c", 3)
	caches.Set("d", 4)

	if len(caches.hashMap) != 3 {
		t.Errorf("expect: %v, got: %v", 3, len(caches.hashMap))
	}

	for k, v := range caches.hashMap {
		if k == "b" {
			if v.value != 2 {
				t.Errorf("expect: %v, got: %v", 2, v.value)
			}
		} else if k == "c" {
			if v.value != 3 {
				t.Errorf("expect: %v, got: %v", 3, v.value)
			}
		} else if k == "d" {
			if v.value != 4 {
				t.Errorf("expect: %v, got: %v", 4, v.value)
			}
		} else {
			t.Errorf("expect: %v, got: %v", "key should be either of b, c, d", k)
		}
	}

	if caches.head != caches.hashMap["d"] {
		t.Errorf("expect: %v, got: %v", "head is d", caches.head)
	}

	if caches.tail != caches.hashMap["b"] {
		t.Errorf("expect: %v, got: %v", "tail is b", caches.tail)
	}
}

func TestLRUcachesGetWithExistKey(t *testing.T) {
	size := 3
	caches := NewLRUCache(size)
	caches.Set("a", 1)
	caches.Set("b", 2)
	caches.Set("c", 3)

	// now is (head) c <-> b <-> a (tail)
	if v, ok := caches.Get("a"); !ok || v != 1 {
		t.Errorf("expect: %v, got: %v", 1, v)
	}

	// now is (head) a <-> c <-> b (tail)
	val := caches.hashMap["a"]
	if val.prev != nil {
		t.Errorf("expect: %v, got: %v", nil, val.prev)
	}
	if val.next != caches.hashMap["c"] {
		t.Errorf("expect: %v, got: %v", "c", val.next)
	}

	if caches.head != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "head is a", caches.head)
	}

	if caches.tail != caches.hashMap["b"] {
		t.Errorf("expect: %v, got: %v", "tail is b", caches.tail)
	}

	if v, ok := caches.Get("b"); !ok || v != 2 {
		t.Errorf("expect: %v, got: %v", 2, v)
	}

	// now is (head) b <-> a <-> c (tail)
	if caches.head != caches.hashMap["b"] {
		t.Errorf("expect: %v, got: %v", "head is b", caches.head)
	}

	if caches.tail != caches.hashMap["c"] {
		t.Errorf("expect: %v, got: %v", "tail is c", caches.tail)
	}

	if v, ok := caches.Get("c"); !ok || v != 3 {
		t.Errorf("expect: %v, got: %v", 3, v)
	}

	// now is (head) c <-> b <-> a (tail)
	if caches.head != caches.hashMap["c"] {
		t.Errorf("expect: %v, got: %v", "head is c", caches.head)
	}

	if caches.tail != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "tail is a", caches.tail)
	}

	if v, ok := caches.Get("c"); !ok || v != 3 {
		t.Errorf("expect: %v, got: %v", 3, v)
	}

	// now is (head) c <-> b <-> a (tail)
	if caches.head != caches.hashMap["c"] {
		t.Errorf("expect: %v, got: %v", "head is c", caches.head)
	}

	if caches.tail != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "tail is a", caches.tail)
	}

	// --------- Get node in the middle ---------

	if v, ok := caches.Get("b"); !ok || v != 2 {
		t.Errorf("expect: %v, got: %v", 2, v)
	}

	// now is (head) b <-> c <-> a (tail)
	if caches.head != caches.hashMap["b"] {
		t.Errorf("expect: %v, got: %v", "head is b", caches.head)
	}

	if caches.tail != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "tail is a", caches.tail)
	}
}

func TestLRUcachesSetValueAgain(t *testing.T) {
	size := 5
	caches := NewLRUCache(size)
	caches.Set("a", 1)
	caches.Set("b", 2)
	caches.Set("c", 3)

	// Set value in the head again
	caches.Set("c", 2)
	if caches.hashMap["c"].value != 2 {
		t.Errorf("expect: %v, got: %v", 2, caches.hashMap["c"].value)
	}

	if caches.head != caches.hashMap["c"] {
		t.Errorf("expect: %v, got: %v", "head is c", caches.head)
	}

	if caches.tail != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "tail is a", caches.tail)
	}

	// Set value not in link list head
	caches.Set("b", 4)
	if caches.hashMap["b"].value != 4 {
		t.Errorf("expect: %v, got: %v", 4, caches.hashMap["4"].value)
	}

	if caches.head != caches.hashMap["b"] {
		t.Errorf("expect: %v, got: %v", "head is b", caches.head)
	}

	if caches.tail != caches.hashMap["a"] {
		t.Errorf("expect: %v, got: %v", "tail is a", caches.tail)
	}
}

func TestLRUcachesGetWithNonExistKey(t *testing.T) {
	size := 5
	caches := NewLRUCache(size)
	caches.Set("a", 1)
	caches.Set("b", 2)
	caches.Set("c", 3)

	if v, ok := caches.Get("d"); v != nil || ok {
		t.Errorf("expect: %v, got: %v", nil, v)
	}
}
