package stack

import (
	"errors"
)

var (
	// ErrStackInvalidSize error indicate invalid stack size
	ErrStackInvalidSize = errors.New("Invalid stack size")

	// ErrStackEmpty error indicate stack is empty
	ErrStackEmpty = errors.New("Stack is empty")

	// ErrStackFull error indicate stack is full
	ErrStackFull = errors.New("Stack is full")
)

// Stack is an interface for stack
type Stack interface {
	IsEmpty() bool
	IsFull() bool
	MakeEmpty()
	Push(elem interface{}) error
	Top() (interface{}, error)
	Pop() (interface{}, error)
}
