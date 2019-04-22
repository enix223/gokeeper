package factory

import (
	"errors"

	"github.com/enix223/gokeeper/stack"
	"github.com/enix223/gokeeper/stack/list"
	"github.com/enix223/gokeeper/stack/slice"
)

var (
	// ErrInvalidType error indicate invalid stack type
	ErrInvalidType = errors.New("Invalid stack type")
)

// CreateStack create a stack with given type
func CreateStack(typ string, size uint) (stack.Stack, error) {
	switch typ {
	case "list":
		return list.NewStack(size)
	case "slice":
		return slice.NewStack(size)
	default:
		return nil, ErrInvalidType
	}
}
