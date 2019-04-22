package sort

import (
	"fmt"
	"math"
	"reflect"
)

//
// Heap Sort
//
//                  10
//                  / \
//                 8   7
//                /\   /\
//               5  6 2  1
//              /\
//             3  4
//
// Array:
//
// +========================================+
// | 10 | 8 | 7 | 5 | 6 | 3 | 2 | 1 | 3 | 4 |
// +========================================+
//
// PARENT(i) = floor(i / 2)
// LEFT-CHILD(i) = 2 * i
// RIGHT-CHILD(i) = 2 * i + 1
// HEIGHT(n) = floor(lg(n))
//
// Maximum heap: A[PARENT[i]] >= A[i]
// Minimum heap: A[PARENT[i]] <= A[i]
//
// HeapSort(A) - O(nlg(n))
// 1) Build-Heap(A)
// 2) for i = A.length downto 2
// 3)   exchange A[1] and A[i]
// 4)   A.heap-size = A.heap-size - 1
// 5)   MAX-HEAPIFY(A, 1)
//

// HeapSort heap sort implementation
type HeapSort struct {
	data     interface{}
	compare  Compare
	heapSize int
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

// maxHeapify heapify the data into a max heap and its subtree in given position
func (h *HeapSort) maxHeapify(i int) {
	l := leftChild(i)
	r := rightChild(i)
	max := 0
	if l < h.heapSize && h.compare(l, i) == LargerThan {
		max = l
	} else {
		max = i
	}

	if r < h.heapSize && h.compare(r, max) == LargerThan {
		max = r
	}

	if max != i {
		swp := reflect.Swapper(h.data)
		swp(max, i)
		h.maxHeapify(max)
	}
}

func (h *HeapSort) buildHeap() {
	size := float64(h.heapSize)
	for i := int(math.Floor(size/2)) - 1; i >= 0; i-- {
		h.maxHeapify(i)
	}
}

// Sort sort given data with heap sort algorithm
func (h *HeapSort) Sort(data interface{}, cmp Compare) {
	typ := reflect.TypeOf(data)
	val := reflect.ValueOf(data)
	if typ.Kind() != reflect.Array && typ.Kind() != reflect.Slice {
		panic(fmt.Sprintf("data is a slice or array, got: %v", typ.Kind()))
	}

	length := val.Len()
	h.data = data
	h.compare = cmp
	h.heapSize = length

	h.buildHeap()
	for i := length - 1; i >= 0; i-- {
		// exchange A[0] and A[len - 1]
		swp := reflect.Swapper(h.data)
		swp(0, h.heapSize-1)
		h.heapSize--
		h.maxHeapify(0)
	}
}
