package list

import "github.com/enix223/gokeeper/stack"

type node struct {
	next  *node
	value interface{}
}

// StackImpl stack implementation with linked list
type StackImpl struct {
	head        *node
	size        uint
	currentSize uint
}

// NewStack create a stack, if size equal 0, then error = ErrStackInvalidSize
func NewStack(size uint) (stack.Stack, error) {
	if size == 0 {
		return nil, stack.ErrStackInvalidSize
	}

	a := new(StackImpl)
	a.head = nil
	a.size = size
	a.currentSize = 0
	return a, nil
}

// IsEmpty check stack is empty or not
func (a *StackImpl) IsEmpty() bool {
	return a.head == nil
}

// IsFull check stack is full or not
func (a *StackImpl) IsFull() bool {
	return a.currentSize == a.size
}

// MakeEmpty clear stack
func (a *StackImpl) MakeEmpty() {
	a.head = nil
	a.currentSize = 0
}

// Push push an element into the stack.
// if stack is full, then error = ErrStackFull
func (a *StackImpl) Push(elem interface{}) error {
	if a.IsFull() {
		return stack.ErrStackFull
	}

	n := new(node)
	n.next = a.head
	n.value = elem
	a.head = n
	a.currentSize++
	return nil
}

// Top returns the element at the top of the stack
// If stack is empty, then error = ErrStackEmpty
func (a *StackImpl) Top() (interface{}, error) {
	if a.IsEmpty() {
		return nil, stack.ErrStackEmpty
	}

	return a.head.value, nil
}

// Pop returns the element at the top of the stack, and then remove it from stack
// If stack is empty, then error = ErrStackEmpty
func (a *StackImpl) Pop() (interface{}, error) {
	if a.IsEmpty() {
		return nil, stack.ErrStackEmpty
	}

	elem := a.head.value
	a.head = a.head.next
	a.currentSize--
	return elem, nil
}
