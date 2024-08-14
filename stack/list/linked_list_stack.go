package list

import "github.com/enix223/gokeeper/stack"

type node[T any] struct {
	next  *node[T]
	value T
}

// StackImpl stack implementation with linked list
type StackImpl[T any] struct {
	head       *node[T]
	cap        int
	currentcap int
}

// NewStack create a stack with linked list
//
// If cap = 0, then unlimited capacity stack is created
func NewStack[T any](cap int) stack.Stack[T] {
	if cap < 0 {
		panic("cap should greater or equals 0")
	}
	a := new(StackImpl[T])
	a.head = nil
	a.cap = cap
	a.currentcap = 0
	return a
}

// IsEmpty check stack is empty or not
func (a *StackImpl[T]) IsEmpty() bool {
	return a.head == nil
}

// IsFull check stack is full or not
func (a *StackImpl[T]) IsFull() bool {
	return a.currentcap == a.cap
}

// MakeEmpty clear stack
func (a *StackImpl[T]) Clear() {
	a.head = nil
	a.currentcap = 0
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
	a.currentcap++
}

// Top returns the element at the top of the stack
// If stack is empty, then error = ErrStackEmpty
func (a *StackImpl[T]) Peek() T {
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
	a.currentcap--
	return elem
}
