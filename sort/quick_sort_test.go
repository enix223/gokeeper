package sort

import (
	"reflect"
	"testing"
)

func TestQuickSortAsc(t *testing.T) {
	s := new(QuickSort)
	a := []int{8, 5, 3, 9, 4, 6, 7, 2}
	exp := []int{2, 3, 4, 5, 6, 7, 8, 9}
	s.Sort(a, func(i, j int) CompareResult {
		if a[i] == a[j] {
			return Equal
		} else if a[i] < a[j] {
			return LessThan
		}
		return LargerThan
	})
	if !reflect.DeepEqual(a, exp) {
		t.Errorf("exp %v, got %v", exp, a)
	}

	a = []int{7, 6, 5, 4, 3, 2, 1, 1}
	exp = []int{1, 1, 2, 3, 4, 5, 6, 7}
	s.Sort(a, func(i, j int) CompareResult {
		if a[i] == a[j] {
			return Equal
		} else if a[i] < a[j] {
			return LessThan
		}
		return LargerThan
	})
	if !reflect.DeepEqual(a, exp) {
		t.Errorf("exp %v, got %v", exp, a)
	}

	a = []int{1, 1, 1, 1, 1, 1, 1, 1}
	exp = []int{1, 1, 1, 1, 1, 1, 1, 1}
	s.Sort(a, func(i, j int) CompareResult {
		if a[i] == a[j] {
			return Equal
		} else if a[i] < a[j] {
			return LessThan
		}
		return LargerThan
	})
	if !reflect.DeepEqual(a, exp) {
		t.Errorf("exp %v, got %v", exp, a)
	}
}

func TestQuickSortDescAsc(t *testing.T) {
	s := new(QuickSort)
	a := []int{8, 5, 3, 9, 4, 6, 7, 2}
	exp := []int{9, 8, 7, 6, 5, 4, 3, 2}
	s.Sort(a, func(i, j int) CompareResult {
		if a[i] == a[j] {
			return Equal
		} else if a[i] > a[j] {
			return LessThan
		}
		return LargerThan
	})
	if !reflect.DeepEqual(a, exp) {
		t.Errorf("exp %v, got %v", exp, a)
	}

	a = []int{1, 1, 2, 3, 4, 5, 6, 7}
	exp = []int{7, 6, 5, 4, 3, 2, 1, 1}
	s.Sort(a, func(i, j int) CompareResult {
		if a[i] == a[j] {
			return Equal
		} else if a[i] > a[j] {
			return LessThan
		}
		return LargerThan
	})
	if !reflect.DeepEqual(a, exp) {
		t.Errorf("exp %v, got %v", exp, a)
	}

	a = []int{1, 1, 1, 1, 1, 1, 1, 1}
	exp = []int{1, 1, 1, 1, 1, 1, 1, 1}
	s.Sort(a, func(i, j int) CompareResult {
		if a[i] == a[j] {
			return Equal
		} else if a[i] > a[j] {
			return LessThan
		}
		return LargerThan
	})
	if !reflect.DeepEqual(a, exp) {
		t.Errorf("exp %v, got %v", exp, a)
	}
}

func TestQuickSortInvalidData(t *testing.T) {
	s := new(QuickSort)
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("exp panic, but no")
		}
	}()
	s.Sort(1, func(i, j int) CompareResult {
		return Equal
	})
}

func TestQuickSortWithString(t *testing.T) {
	s := new(QuickSort)
	a := []string{"z", "b", "f", "e", "k", "x", "c"}
	exp := []string{"b", "c", "e", "f", "k", "x", "z"}
	s.Sort(a, func(i, j int) CompareResult {
		if a[i] == a[j] {
			return Equal
		} else if a[i] < a[j] {
			return LessThan
		}
		return LargerThan
	})
	if !reflect.DeepEqual(a, exp) {
		t.Errorf("exp %v, got %v", exp, a)
	}
}

func TestQuickSortWithFloat(t *testing.T) {
	s := new(QuickSort)
	a := []float32{1.2, 0.9, 0, 5.9, 3.0, 4.2, 2}
	exp := []float32{0, 0.9, 1.2, 2, 3.0, 4.2, 5.9}
	s.Sort(a, func(i, j int) CompareResult {
		if a[i] == a[j] {
			return Equal
		} else if a[i] < a[j] {
			return LessThan
		}
		return LargerThan
	})
	if !reflect.DeepEqual(a, exp) {
		t.Errorf("exp %v, got %v", exp, a)
	}
}

type testStruct struct {
	val int
}

func TestQuickSortWithStruct(t *testing.T) {
	s := new(QuickSort)
	a := []*testStruct{
		&testStruct{val: 9},
		&testStruct{val: 3},
		&testStruct{val: 5},
		&testStruct{val: 7},
		&testStruct{val: 2},
		&testStruct{val: 1},
		&testStruct{val: 4},
	}
	exp := []*testStruct{
		&testStruct{val: 1},
		&testStruct{val: 2},
		&testStruct{val: 3},
		&testStruct{val: 4},
		&testStruct{val: 5},
		&testStruct{val: 7},
		&testStruct{val: 9},
	}
	s.Sort(a, func(i, j int) CompareResult {
		if a[i].val == a[j].val {
			return Equal
		} else if a[i].val < a[j].val {
			return LessThan
		}
		return LargerThan
	})
	if !reflect.DeepEqual(a, exp) {
		t.Errorf("exp %v, got %v", exp, a)
	}
}
