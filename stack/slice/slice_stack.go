package slice

import "github.com/enix223/gokeeper/stack"

// StackImpl stack implementation with slice
type StackImpl[T any] struct {
	elemenets []T
	cap       int
}

// NewStack create a stack with given capacity
//
// If cap = 0, then unlimited capacity stack is created
func NewStack[T any](cap int) stack.Stack[T] {
	if cap < 0 {
		panic("cap should greater or equals 0")
	}
	s := new(StackImpl[T])
	s.elemenets = make([]T, 0, cap)
	s.cap = cap
	return s
}

// IsEmpty check stack is empty or not
func (s *StackImpl[T]) IsEmpty() bool {
	return len(s.elemenets) == 0
}

// IsFull check stack is full or not
func (s *StackImpl[T]) IsFull() bool {
	return s.cap > 0 && len(s.elemenets) == s.cap
}

// MakeEmpty clear the stack
func (s *StackImpl[T]) Clear() {
	// make the old elements gc
	s.elemenets = make([]T, 0, s.cap)
}

// Push push an element into stack
// if stack is full then return ErrStackFull
func (s *StackImpl[T]) Push(elem T) {
	if s.IsFull() {
		panic("stack is full")
	}

	s.elemenets = append(s.elemenets, elem)
}

// Peek return the element in the top of the stack, if stack is empty,
// then return panic
func (s *StackImpl[T]) Peek() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	return s.elemenets[s.head()]
}

// Pop return the element at the top of the stack, and remove it from the stack
// If stack is empty, then error = ErrStackEmpty
func (s *StackImpl[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	var head = s.head()
	elem := s.elemenets[head]
	s.elemenets = s.elemenets[:head]
	return elem
}

func (s *StackImpl[T]) head() int {
	return len(s.elemenets) - 1
}
