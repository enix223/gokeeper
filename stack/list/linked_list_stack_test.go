package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackCreateWithErr(t *testing.T) {
	assert.PanicsWithValue(t, "size should not be 0", func() {
		NewStack[int](0)
	})
}

func TestStackEmptyCheck(t *testing.T) {
	var s = NewStack[int](10)
	assert.True(t, s.IsEmpty())
}

func TestStackFullCheck(t *testing.T) {
	var s = NewStack[int](1)
	s.Push(1)
	assert.True(t, s.IsFull())
}

func TestStackMakeEmpty(t *testing.T) {
	var s = NewStack[int](10)
	s.Push(1)
	assert.False(t, s.IsEmpty())

	s.MakeEmpty()
	assert.True(t, s.IsEmpty())
}

func TestStackPush(t *testing.T) {
	var s = NewStack[int](1)
	s.Push(1)
	assert.PanicsWithValue(t, "stack is full", func() {
		s.Push(2)
	})
}

func TestStackPop(t *testing.T) {
	var s = NewStack[int](1)
	s.Push(1)
	var elem = s.Pop()
	assert.Equal(t, 1, elem)

	assert.PanicsWithValue(t, "stack is empty", func() {
		s.Pop()
	})
}

func TestStackTop(t *testing.T) {
	var s = NewStack[int](1)
	assert.PanicsWithValue(t, "stack is empty", func() {
		s.Top()
	})

	s.Push(1)
	var elem = s.Top()
	assert.Equal(t, 1, elem)
	assert.False(t, s.IsEmpty())
	s.Pop()
	assert.PanicsWithValue(t, "stack is empty", func() {
		s.Pop()
	})
}
