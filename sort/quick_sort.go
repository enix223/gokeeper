package sort

import (
	"fmt"
	"reflect"
)

// QuickSort
//
// 1. Pick a pivot from the array, and partition the array into two parts:
//    The elements in left part are all smaller or equal to the pivot number,
//    elements in the right part are all greater than the pivot number
// 2. Use divide and conquer to apply the quick sort algorithm to the left part and right part
//
// =====================================
// Algorithm
// =====================================
//
// QUICKSORT(A, p, r)
// 1. if p < r
// 2.   q = PARTITION(A, p, r)
// 3.   QUICKSORT(A, p, q - 1)
// 4.   QUICKSORT(A, q + 1, r)
//
// PARTITION(A, p, r)
// 1. x = A[r]
// 2. i = p - 1
// 3. for j = p to r - 1
// 4.   if A[j] <= x
// 5.     i = i + 1
// 6.     exchange A[i] and A[j]
// 7. exchange A[i + 1] and A[r]
// 8. return i + 1
//
// =====================================
// Example:
// =====================================
// A = | 1 | 5 | 9 | 7 | 4 | 3 | 2 | 8 |
//
// Partition(A, 1, 8):
//
//    i  j                           r
// 1)  | 1 | 5 | 9 | 7 | 8 | 3 | 2 | 4 |
//       i   j                       r
// 2)  | 1 | 5 | 9 | 7 | 8 | 3 | 2 | 4 |
//       i       j                   r
// 3)  | 1 | 5 | 9 | 7 | 8 | 3 | 2 | 4 |
//       i           j               r
// 3)  | 1 | 5 | 9 | 7 | 8 | 3 | 2 | 4 |
//       i               j           r
// 4)  | 1 | 5 | 9 | 7 | 8 | 3 | 2 | 4 |
//       i                   j       r
// 5)  | 1 | 5 | 9 | 7 | 8 | 3 | 2 | 4 |
//           i                   j   r
// 6)  | 1 | 3 | 9 | 7 | 8 | 5 | 2 | 4 |
//               i               j   r
// 7)  | 1 | 3 | 2 | 7 | 8 | 5 | 9 | 4 |
//               i               j   r
// 8)  | 1 | 3 | 2 | 4 | 8 | 5 | 9 | 7 |
//

// QuickSort quick sort implementation
type QuickSort struct {
	data interface{}
	cmp  Compare
}

// Sort sort given slice/array with quick sort
func (s *QuickSort) Sort(data interface{}, cmp Compare) {
	typ := reflect.TypeOf(data)
	val := reflect.ValueOf(data)
	if typ.Kind() != reflect.Array && typ.Kind() != reflect.Slice {
		panic(fmt.Sprintf("data is a slice or array, got: %v", typ.Kind()))
	}

	s.cmp = cmp
	s.data = data
	s.internalSort(0, val.Len()-1)
}

func (s *QuickSort) internalSort(p, r int) {
	if p < r {
		q := s.partition(p, r)
		s.internalSort(p, q-1)
		s.internalSort(q+1, r)
	}
}

func (s *QuickSort) partition(p, r int) int {
	i := p - 1
	swp := reflect.Swapper(s.data)

	for j := p; j <= r-1; j++ {
		if s.cmp(j, r) == LessThan || s.cmp(j, r) == Equal {
			i++
			swp(i, j)
		}
	}
	swp(i+1, r)
	return i + 1
}
