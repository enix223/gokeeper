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

func TestFilterSliceInt(b *testing.T) {
	arr := []int{1, 2, 3}
	res := FilterSlice(arr, func(v interface{}) bool {
		return v.(int) == 2 || v.(int) == 1
	})

	exp := []int{1, 2}
	if !reflect.DeepEqual(res, exp) {
		b.Fatalf("exp: %v, got: %v", exp, res)
	}
}

func TestFilterSliceInvalidType(t *testing.T) {
	arr := [3]int{1, 2, 3}

	defer func() {
		if err := recover(); err == nil {
			t.Fatal("should be panic")
		}
	}()

	FilterSlice(arr, func(v interface{}) bool {
		return v.(int) == 3
	})
}

func TestFilterSliceEmpty(b *testing.T) {
	arr := []int{1, 2, 3}
	res := FilterSlice(arr, func(v interface{}) bool {
		return v.(int) == 100
	})

	exp := []int{}
	if !reflect.DeepEqual(res, exp) {
		b.Fatalf("exp: %v, got: %v", exp, res)
	}
}

func TestFilterSliceStruct(b *testing.T) {
	type S struct {
		Name string
	}
	arr1 := []S{
		{Name: "abc"},
		{Name: "bbc"},
		{Name: "bbq"},
	}
	res1 := FilterSlice(arr1, func(v interface{}) bool {
		return v.(S).Name == "bbc"
	})

	exp1 := []S{{Name: "bbc"}}
	if !reflect.DeepEqual(res1, exp1) {
		b.Fatalf("exp: %v, got: %v", exp1, res1)
	}

	arr2 := []*S{
		{Name: "abc"},
		{Name: "bbc"},
		{Name: "bbq"},
	}
	res2 := FilterSlice(arr2, func(v interface{}) bool {
		return v.(*S).Name == "bbc"
	})

	exp2 := []*S{arr2[1]}
	if !reflect.DeepEqual(res2, exp2) {
		b.Fatalf("exp: %v, got: %v", exp2, res2)
	}

	arr3 := []*S{
		{Name: "abc"},
		{Name: "bbc"},
		{Name: "bbq"},
	}
	res3 := FilterSlice(&arr3, func(v interface{}) bool {
		return v.(*S).Name == "bbc"
	})

	exp3 := []*S{arr3[1]}
	if !reflect.DeepEqual(res3, exp3) {
		b.Fatalf("exp: %v, got: %v", exp3, res3)
	}
}

func TestAnyWithSlice(t *testing.T) {
	a := []int{1, 2, 3, 4}

	res := Any(a, func(item interface{}) bool {
		return item.(int) > 5
	})

	if res {
		t.Fatal("exp: false, got true")
	}

	res = Any(a, func(item interface{}) bool {
		return item.(int) >= 4
	})

	if !res {
		t.Fatal("exp: true, got false")
	}

	// ptr
	res = Any(&a, func(item interface{}) bool {
		return item.(int) >= 4
	})

	if !res {
		t.Fatal("exp: true, got false")
	}
}

func TestAnyWithInvalidType(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("expect panic, but got nil")
		}
	}()

	Any(1, func(item interface{}) bool {
		return true
	})
}

func TestAnyWithDiffElementType(t *testing.T) {
	a := []interface{}{1, nil, 12.0}

	res := Any(a, func(item interface{}) bool {
		return item != nil
	})

	if !res {
		t.Fatal("exp: true, got false")
	}

	a = []interface{}{nil, nil, nil}

	res = Any(a, func(item interface{}) bool {
		return item != nil
	})

	if res {
		t.Fatal("exp: false, got true")
	}
}

func TestAllWithSlice(t *testing.T) {
	a := []int{1, 2, 3, 4}

	res := All(a, func(item interface{}) bool {
		return item.(int) > 5
	})

	if res {
		t.Fatal("exp: false, got true")
	}

	res = All(a, func(item interface{}) bool {
		return item.(int) >= 1
	})

	if !res {
		t.Fatal("exp: true, got false")
	}

	// ptr
	res = All(&a, func(item interface{}) bool {
		return item.(int) >= 1
	})

	if !res {
		t.Fatal("exp: true, got false")
	}
}

func TestAllWithInvalidType(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("expect panic, but got nil")
		}
	}()

	All(1, func(item interface{}) bool {
		return true
	})
}

func TestAllWithDiffElementType(t *testing.T) {
	a := []interface{}{1, nil, 12.0}

	res := All(a, func(item interface{}) bool {
		return item != nil
	})

	if res {
		t.Fatal("exp: false, got true")
	}

	a = []interface{}{1, "bool", true}

	res = Any(a, func(item interface{}) bool {
		return item != nil
	})

	if !res {
		t.Fatal("exp: true, got false")
	}
}

func TestAnyNil(t *testing.T) {
	a := []interface{}{1, nil, 12.0}
	res := AnyNil(a)
	if !res {
		t.Fatal("exp: true, got false")
	}

	a = []interface{}{1, nil, nil}
	res = AnyNil(a)
	if !res {
		t.Fatal("exp: true, got false")
	}

	a = []interface{}{1, true, "bool"}
	res = AnyNil(a)
	if res {
		t.Fatal("exp: false, got true")
	}
}

func TestAnyNotNil(t *testing.T) {
	a := []interface{}{1, nil, 12.0}
	res := AnyNotNil(a)
	if !res {
		t.Fatal("exp: true, got false")
	}

	a = []interface{}{1, nil, nil}
	res = AnyNotNil(a)
	if !res {
		t.Fatal("exp: true, got false")
	}

	a = []interface{}{1, true, "bool"}
	res = AnyNotNil(a)
	if !res {
		t.Fatal("exp: true, got false")
	}

	a = []interface{}{nil, nil, nil}
	res = AnyNotNil(a)
	if res {
		t.Fatal("exp: false, got true")
	}
}

func TestAllNil(t *testing.T) {
	a := []interface{}{1, nil, 12.0}

	res := AllNil(a)
	if res {
		t.Fatal("exp: false, got true")
	}

	a = []interface{}{1, nil, nil}
	res = AllNil(a)
	if res {
		t.Fatal("exp: false, got true")
	}

	a = []interface{}{nil, nil, nil}
	res = AllNil(a)
	if !res {
		t.Fatal("exp: true, got false")
	}
}

func TestAllNotNil(t *testing.T) {
	a := []interface{}{1, nil, 12.0}

	res := AllNotNil(a)
	if res {
		t.Fatal("exp: false, got true")
	}

	a = []interface{}{1, nil, nil}
	res = AllNotNil(a)
	if res {
		t.Fatal("exp: false, got true")
	}

	a = []interface{}{nil, nil, nil}
	res = AllNotNil(a)
	if res {
		t.Fatal("exp: false, got true")
	}

	a = []interface{}{1, "a", true}
	res = AllNotNil(a)
	if !res {
		t.Fatal("exp: true, got false")
	}
}
