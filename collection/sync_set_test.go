package collection

import (
	"reflect"
	"sync"
	"testing"
)

func TestSyncSetAdd(t *testing.T) {
	s := NewSyncSet()

	s.Add(1)
	if _, ok := s.set[1]; !ok {
		t.Fatalf("exp %v, got %v", "exist", "not exist")
	}
}

func TestSyncSetHas(t *testing.T) {
	s := NewSyncSet()
	s.set[1] = struct{}{}

	if ok := s.Has(1); !ok {
		t.Fatalf("exp %v, got %v", "exist", "not exist")
	}

	if ok := s.Has(2); ok {
		t.Fatalf("exp %v, got %v", "not exist", "exist")
	}
}

func TestSyncSetUnion(t *testing.T) {
	s1 := NewSyncSet()
	s1.Add(1, 2, 3, 4, 5)
	s2 := NewSyncSet()
	s2.Add(6, 7, 8, 9, 10)

	s3 := s1.Union(s2)
	exp := NewSyncSet()
	exp.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if !reflect.DeepEqual(s3, exp) {
		t.Fatalf("exp %v, got %v", exp, s3)
	}
}

func TestSyncSetIntercet(t *testing.T) {
	s1 := NewSyncSet()
	s1.Add(1, 2, 3, 4, 5)
	s2 := NewSyncSet()
	s2.Add(3, 7, 2, 9, 10)

	s3 := s1.Intersect(s2)
	exp := NewSyncSet()
	exp.Add(2, 3)
	if !reflect.DeepEqual(s3, exp) {
		t.Fatalf("exp %v, got %v", exp, s3)
	}
}

func TestSyncSetDiff(t *testing.T) {
	s1 := NewSyncSet()
	s1.Add(1, 2, 3, 4, 5)
	s2 := NewSyncSet()
	s2.Add(3, 7, 2, 9, 10)

	s3 := s1.Difference(s2)
	exp := NewSyncSet()
	exp.Add(1, 4, 5)
	if !reflect.DeepEqual(s3, exp) {
		t.Fatalf("exp %v, got %v", exp, s3)
	}
}

func TestSyncSetSize(t *testing.T) {
	s1 := NewSyncSet()
	s1.Add(1, 2, 3, 4, 5)

	if s := s1.Size(); s != 5 {
		t.Fatalf("exp %v, got %v", 5, s)
	}
}

func TestSyncSetRemove(t *testing.T) {
	s1 := NewSyncSet()
	s1.Add(1, 2, 3, 4, 5)

	s1.Remove(3, 4, 5)
	exp := NewSyncSet()
	exp.Add(1, 2)
	if !reflect.DeepEqual(s1, exp) {
		t.Fatalf("exp %v, got %v", exp, s1)
	}
}

func TestSyncSetUnhashableValueAdd(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Fatalf("exp %v, got %v", "panic: runtime error: hash of unhashable type []string", "nil panic")
		}
	}()

	s1 := NewSyncSet()
	k := []string{"1"}
	s1.Add(k)
}

func TestSyncSetUnhashableValueRemove(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Fatalf("exp %v, got %v", "panic: runtime error: hash of unhashable type []string", "nil panic")
		}
	}()

	s1 := NewSyncSet()
	k := []string{"1"}
	s1.Remove(k)
}

func TestSyncSetUnhashableValueHas(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Fatalf("exp %v, got %v", "panic: runtime error: hash of unhashable type []string", "nil panic")
		}
	}()

	s1 := NewSyncSet()
	k := []string{"1"}
	s1.Has(k)
}

func TestSyncSetEqual(t *testing.T) {
	s1 := NewSyncSet()
	s1.Add("1", "2", "3")
	s2 := NewSyncSet()
	s2.Add("1", "2", "3")

	if !s1.Equal(s2) {
		t.Errorf("exp %v, got %v", "equal", "not equal")
	}

	s3 := NewSyncSet()
	s3.Add("1", "3")
	if s1.Equal(s3) {
		t.Errorf("exp %v, got %v", "not equal", "equal")
	}

	s4 := NewSyncSet()
	s4.Add("1", "2", "3", "4")
	if s1.Equal(s4) {
		t.Errorf("exp %v, got %v", "not equal", "equal")
	}

	s5 := NewSyncSet()
	s5.Add(1, 2, 3)
	if s1.Equal(s5) {
		t.Errorf("exp %v, got %v", "not equal", "equal")
	}
}

func TestSyncSetMap(t *testing.T) {
	s1 := NewSyncSet()
	s1.Add("1", "2", "3")

	s0 := NewSyncSet()
	s0.Add("11", "22", "33")

	s2 := s1.Map(func(v interface{}) interface{} {
		return v.(string) + v.(string)
	})
	if !s0.Equal(s2) {
		t.Errorf("exp %v, got %v", s0, s2)
	}
}

func TestSyncSetFilter(t *testing.T) {
	s1 := NewSyncSet()
	s1.Add("1", "2", "33")

	s0 := NewSyncSet()
	s0.Add("1", "2")

	s2 := s1.Filter(func(v interface{}) bool {
		return len(v.(string)) == 1
	})
	if !s0.Equal(s2) {
		t.Errorf("exp %v, got %v", s0, s2)
	}
}

func TestSyncSetCurrentAdd(t *testing.T) {
	s1 := NewSyncSet()

	var wg sync.WaitGroup
	go func() {
		for i := 0; i < 1000; i++ {
			s1.Add(i)
		}
		wg.Done()
	}()
	wg.Add(1)

	go func() {
		for i := 1000; i < 2000; i++ {
			s1.Add(i)
		}
		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()

	exp := NewSyncSet()
	for i := 0; i < 2000; i++ {
		exp.Add(i)
	}

	if !exp.Equal(s1) {
		t.Errorf("exp %v, got %v", exp, s1)
	}
}
