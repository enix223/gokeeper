package collection

import (
	"reflect"
	"testing"
)

func TestIsIn(t *testing.T) {
	c1 := []byte{1, 2, 3, 4}
	c2 := []byte{1, 2, 3, 4, 5}

	if !IsIn(c1, c2) {
		t.Errorf("expect: %v, got: %v", true, false)
	}

	if IsIn(c2, c1) {
		t.Errorf("expect: %v, got: %v", false, true)
	}
}

func TestIsInWithDiffType(t *testing.T) {
	c1 := []byte{1, 2, 3, 4}
	c2 := []int{1, 2, 3, 4, 5}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect: %v, got: %v", "panic", "nil")
		}
	}()

	IsIn(c1, c2)
}

func TestIsInWithWrongTypeC1(t *testing.T) {
	c1 := 1
	c2 := []int{1, 2, 3, 4, 5}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect: %v, got: %v", "panic", "nil")
		}
	}()

	IsIn(c1, c2)
}

func TestIsInWithWrongTypeC2(t *testing.T) {
	c1 := []int{1, 2}
	c2 := 2

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect: %v, got: %v", "panic", "nil")
		}
	}()

	IsIn(c1, c2)
}

func TestContainOK(t *testing.T) {
	c1 := []int{1, 2, 3, 4}

	if !Contain(c1, 1) {
		t.Errorf("expect: %v, got: %v", true, false)
	}

	if Contain(c1, 10) {
		t.Errorf("expect: %v, got: %v", false, true)
	}
}

func TestContainWithWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect: %v, got: %v", "panic", "nil")
		}
	}()

	Contain(2, 1)
}

func TestSliceMapWithWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expect: %v, got: %v", "panic", "nil")
		}
	}()

	SliceMap(2, func(i int, item interface{}) interface{} {
		return i
	})
}

func TestSliceMapWithEmptyMap(t *testing.T) {
	m := []string{}
	m1 := SliceMap(m, func(i int, item interface{}) interface{} {
		return i
	})

	if m1 != nil {
		t.Fatalf("exp: %v, got: %v", nil, m1)
	}
}

func TestSliceMap(t *testing.T) {
	m := []string{"a", "b"}
	m1 := SliceMap(m, func(i int, item interface{}) interface{} {
		return i
	})

	exp := []int{0, 1}
	if !reflect.DeepEqual(exp, m1) {
		t.Fatalf("exp: %v, got: %v", exp, m1)
	}
}

func BenchmarkSliceMap(b *testing.B) {
	m := []string{"a", "b"}
	for i := 0; i < b.N; i++ {
		m1 := SliceMap(m, func(i int, item interface{}) interface{} {
			return i
		})

		exp := []int{0, 1}
		if !reflect.DeepEqual(exp, m1) {
			b.Fatalf("exp: %v, got: %v", exp, m1)
		}
	}
}