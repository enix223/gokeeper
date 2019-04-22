package slice

import "github.com/enix223/gokeeper/stack"

// StackImpl stack implementation with slice
type StackImpl struct {
	elemenets []interface{}
	head      int
}

// NewStack create a stack with given size
func NewStack(size uint) (stack.Stack, error) {
	if size == 0 {
		return nil, stack.ErrStackInvalidSize
	}

	s := new(StackImpl)
	s.elemenets = make([]interface{}, 0, size)
	s.head = -1
	return s, nil
}

// IsEmpty check stack is empty or not
func (s *StackImpl) IsEmpty() bool {
	return s.head == -1
}

// IsFull check stack is full or not
func (s *StackImpl) IsFull() bool {
	return s.head == cap(s.elemenets)-1
}

// MakeEmpty clear the stack
func (s *StackImpl) MakeEmpty() {
	s.head = -1
}

// Push push an element into stack
// if stack is full then return ErrStackFull
func (s *StackImpl) Push(elem interface{}) error {
	if s.IsFull() {
		return stack.ErrStackFull
	}

	s.elemenets = append(s.elemenets, elem)
	s.head++
	return nil
}

// Top return the element in the top of the stack, if stack is empty,
// then return ErrStackEmpty
func (s *StackImpl) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, stack.ErrStackEmpty
	}

	return s.elemenets[s.head], nil
}

// Pop return the element at the top of the stack, and remove it from the stack
// If stack is empty, then error = ErrStackEmpty
func (s *StackImpl) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, stack.ErrStackEmpty
	}

	elem := s.elemenets[s.head]
	s.head--
	return elem, nil
}
