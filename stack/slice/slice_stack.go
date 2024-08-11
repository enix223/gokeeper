package slice

import "github.com/enix223/gokeeper/stack"

// StackImpl stack implementation with slice
type StackImpl[T any] struct {
	elemenets []T
	head      int
}

// NewStack create a stack with given size
func NewStack[T any](size uint) stack.Stack[T] {
	if size == 0 {
		panic("size should not be 0")
	}

	s := new(StackImpl[T])
	s.elemenets = make([]T, 0, size)
	s.head = -1
	return s
}

// IsEmpty check stack is empty or not
func (s *StackImpl[T]) IsEmpty() bool {
	return s.head == -1
}

// IsFull check stack is full or not
func (s *StackImpl[T]) IsFull() bool {
	return s.head == cap(s.elemenets)-1
}

// MakeEmpty clear the stack
func (s *StackImpl[T]) MakeEmpty() {
	s.head = -1
}

// Push push an element into stack
// if stack is full then return ErrStackFull
func (s *StackImpl[T]) Push(elem T) {
	if s.IsFull() {
		panic("stack is full")
	}

	s.elemenets = append(s.elemenets, elem)
	s.head++
}

// Top return the element in the top of the stack, if stack is empty,
// then return ErrStackEmpty
func (s *StackImpl[T]) Top() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	return s.elemenets[s.head]
}

// Pop return the element at the top of the stack, and remove it from the stack
// If stack is empty, then error = ErrStackEmpty
func (s *StackImpl[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	elem := s.elemenets[s.head]
	s.head--
	return elem
}
