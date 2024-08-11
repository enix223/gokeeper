package list

import "github.com/enix223/gokeeper/stack"

type node[T any] struct {
	next  *node[T]
	value T
}

// StackImpl stack implementation with linked list
type StackImpl[T any] struct {
	head        *node[T]
	size        uint
	currentSize uint
}

// NewStack create a stack, if size equal 0, then error = ErrStackInvalidSize
func NewStack[T any](size uint) stack.Stack[T] {
	if size == 0 {
		panic("size should not be 0")
	}

	a := new(StackImpl[T])
	a.head = nil
	a.size = size
	a.currentSize = 0
	return a
}

// IsEmpty check stack is empty or not
func (a *StackImpl[T]) IsEmpty() bool {
	return a.head == nil
}

// IsFull check stack is full or not
func (a *StackImpl[T]) IsFull() bool {
	return a.currentSize == a.size
}

// MakeEmpty clear stack
func (a *StackImpl[T]) MakeEmpty() {
	a.head = nil
	a.currentSize = 0
}

// Push push an element into the stack.
// if stack is full, then error = ErrStackFull
func (a *StackImpl[T]) Push(elem T) {
	if a.IsFull() {
		panic("stack is full")
	}

	n := new(node[T])
	n.next = a.head
	n.value = elem
	a.head = n
	a.currentSize++
}

// Top returns the element at the top of the stack
// If stack is empty, then error = ErrStackEmpty
func (a *StackImpl[T]) Top() T {
	if a.IsEmpty() {
		panic("stack is empty")
	}

	return a.head.value
}

// Pop returns the element at the top of the stack, and then remove it from stack
// If stack is empty, then error = ErrStackEmpty
func (a *StackImpl[T]) Pop() T {
	if a.IsEmpty() {
		panic("stack is empty")
	}

	elem := a.head.value
	a.head = a.head.next
	a.currentSize--
	return elem
}
