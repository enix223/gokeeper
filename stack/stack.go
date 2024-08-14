package stack

// Stack is an interface for stack
type Stack[T any] interface {
	IsEmpty() bool
	IsFull() bool
	Clear()
	Push(elem T)
	Peek() T
	Pop() T
}
