package sort

import (
	"reflect"
	"testing"
)

func TestSortInvalidType(t *testing.T) {
	hs := new(HeapSort)
	data := map[int]int{}
	defer func() {
		e := recover()
		if e == nil {
			t.Fatalf("exp panic, but got: nil")
		}
	}()

	hs.Sort(data, func(i, j int) CompareResult {
		return Equal
	})
}

func TestSortDesc(t *testing.T) {
	hs := new(HeapSort)
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	hs.Sort(data, func(i, j int) CompareResult {
		if data[i] > data[j] {
			return LessThan
		} else if data[i] == data[j] {
			return Equal
		} else {
			return LargerThan
		}
	})

	exp := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	if !reflect.DeepEqual(data, exp) {
		t.Fatalf("exp: %v, got: %v", exp, data)
	}
}

func TestSortAsc(t *testing.T) {
	hs := new(HeapSort)
	data := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	hs.Sort(data, func(i, j int) CompareResult {
		if data[i] < data[j] {
			return LessThan
		} else if data[i] == data[j] {
			return Equal
		} else {
			return LargerThan
		}
	})

	exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(data, exp) {
		t.Fatalf("exp: %v, got: %v", exp, data)
	}
}
