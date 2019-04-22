package sort

// CompareResult compare function result
type CompareResult int

const (
	// LessThan less than
	LessThan CompareResult = -1
	// Equal two elements are equal
	Equal CompareResult = 0
	// LargerThan element left larger than element right
	LargerThan CompareResult = 1
)

// Compare compare function for given two index
//
// return -1 indicate element in i less than element in j
// return 1 indicate element in i larger than element in j
// return 0 indicate element in i equal to element in j
type Compare func(i, j int) CompareResult

// Sort sort given slice and return a new sorted slice
type Sort interface {
	Sort(data interface{}, cmp Compare)
}
