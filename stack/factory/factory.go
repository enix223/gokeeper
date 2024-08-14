package factory

import (
	"github.com/enix223/gokeeper/stack"
	"github.com/enix223/gokeeper/stack/list"
	"github.com/enix223/gokeeper/stack/slice"
)

// CreateStack create a stack with given type
func CreateStack[T any](typ string, size int) stack.Stack[T] {
	switch typ {
	case "list":
		return list.NewStack[T](size)
	case "slice":
		return slice.NewStack[T](size)
	default:
		panic("invalid stack type")
	}
}
