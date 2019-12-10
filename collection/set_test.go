package collection

import (
	"reflect"
	"testing"
)

func TestSetAdd(t *testing.T) {
	s := NewSet()

	s.Add(1)
	if _, ok := s[1]; !ok {
		t.Fatalf("exp %v, got %v", "exist", "not exist")
	}
}

func TestSetHas(t *testing.T) {
	s := NewSet()
	s[1] = struct{}{}

	if ok := s.Has(1); !ok {
		t.Fatalf("exp %v, got %v", "exist", "not exist")
	}

	if ok := s.Has(2); ok {
		t.Fatalf("exp %v, got %v", "not exist", "exist")
	}
}

func TestSetUnion(t *testing.T) {
	s1 := NewSet()
	s1.Add(1, 2, 3, 4, 5)
	s2 := NewSet()
	s2.Add(6, 7, 8, 9, 10)

	s3 := s1.Union(s2)
	exp := NewSet()
	exp.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if !reflect.DeepEqual(s3, exp) {
		t.Fatalf("exp %v, got %v", exp, s3)
	}
}

func TestSetIntercet(t *testing.T) {
	s1 := NewSet()
	s1.Add(1, 2, 3, 4, 5)
	s2 := NewSet()
	s2.Add(3, 7, 2, 9, 10)

	s3 := s1.Intersect(s2)
	exp := NewSet()
	exp.Add(2, 3)
	if !reflect.DeepEqual(s3, exp) {
		t.Fatalf("exp %v, got %v", exp, s3)
	}
}

func TestSetDiff(t *testing.T) {
	s1 := NewSet()
	s1.Add(1, 2, 3, 4, 5)
	s2 := NewSet()
	s2.Add(3, 7, 2, 9, 10)

	s3 := s1.Difference(s2)
	exp := NewSet()
	exp.Add(1, 4, 5)
	if !reflect.DeepEqual(s3, exp) {
		t.Fatalf("exp %v, got %v", exp, s3)
	}
}

func TestSetSize(t *testing.T) {
	s1 := NewSet()
	s1.Add(1, 2, 3, 4, 5)

	if s := s1.Size(); s != 5 {
		t.Fatalf("exp %v, got %v", 5, s)
	}
}

func TestSetRemove(t *testing.T) {
	s1 := NewSet()
	s1.Add(1, 2, 3, 4, 5)

	s1.Remove(3, 4, 5)
	exp := NewSet()
	exp.Add(1, 2)
	if !reflect.DeepEqual(s1, exp) {
		t.Fatalf("exp %v, got %v", exp, s1)
	}
}

func TestSetUnhashableValueAdd(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Fatalf("exp %v, got %v", "panic: runtime error: hash of unhashable type []string", "nil panic")
		}
	}()

	s1 := NewSet()
	k := []string{"1"}
	s1.Add(k)
}

func TestSetUnhashableValueRemove(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Fatalf("exp %v, got %v", "panic: runtime error: hash of unhashable type []string", "nil panic")
		}
	}()

	s1 := NewSet()
	k := []string{"1"}
	s1.Remove(k)
}

func TestSetUnhashableValueHas(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Fatalf("exp %v, got %v", "panic: runtime error: hash of unhashable type []string", "nil panic")
		}
	}()

	s1 := NewSet()
	k := []string{"1"}
	s1.Has(k)
}
