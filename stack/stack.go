package stack

// Stack is an interface for stack
type Stack[T any] interface {
	IsEmpty() bool
	IsFull() bool
	MakeEmpty()
	Push(elem T)
	Top() T
	Pop() T
}
