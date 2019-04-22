package list

import (
	"testing"

	"github.com/enix223/gokeeper/stack"
)

func TestStackCreateWithErr(t *testing.T) {
	_, err := NewStack(0)
	if err == nil {
		t.Errorf("exp error = nil, but: %v", err)
	}
}

func TestStackEmptyCheck(t *testing.T) {
	s, err := NewStack(10)
	if err != nil {
		t.Fatalf("exp stack create success, but error: %v", err)
	}

	exp := true
	got := s.IsEmpty()
	if !got {
		t.Errorf("exp: %v, got: %v", exp, got)
	}
}

func TestStackFullCheck(t *testing.T) {
	s, err := NewStack(1)
	if err != nil {
		t.Fatalf("exp stack create success, but error: %v", err)
	}

	err = s.Push(1)
	if err != nil {
		t.Fatalf("exp: nil, got: %v", err)
	}

	exp := true
	got := s.IsFull()
	if !got {
		t.Errorf("exp: %v, got: %v", exp, got)
	}
}

func TestStackMakeEmpty(t *testing.T) {
	s, err := NewStack(10)
	if err != nil {
		t.Fatalf("exp stack create success, but error: %v", err)
	}

	err = s.Push(1)
	if err != nil {
		t.Fatalf("exp: nil, got: %v", err)
	}

	exp := false
	got := s.IsEmpty()
	if got {
		t.Errorf("exp: %v, got: %v", exp, got)
	}

	s.MakeEmpty()
	exp = true
	got = s.IsEmpty()
	if !got {
		t.Errorf("exp: %v, got: %v", exp, got)
	}
}

func TestStackPush(t *testing.T) {
	s, err := NewStack(1)
	if err != nil {
		t.Fatalf("exp stack create success, but error: %v", err)
	}

	err = s.Push(1)
	if err != nil {
		t.Fatalf("exp: nil, got: %v", err)
	}

	err = s.Push(2)
	if err == nil {
		t.Fatalf("exp: %v, got: %v", stack.ErrStackFull, err)
	}
}

func TestStackPop(t *testing.T) {
	s, err := NewStack(1)
	if err != nil {
		t.Fatalf("exp stack create success, but error: %v", err)
	}

	err = s.Push(1)
	if err != nil {
		t.Fatalf("exp: nil, got: %v", err)
	}

	elem, err := s.Pop()
	if err != nil {
		t.Fatalf("exp: %v, got: %v", nil, err)
	}

	v := elem.(int)
	if v != 1 {
		t.Errorf("exp: %v, got: %v", 1, v)
	}

	_, err = s.Pop()
	if err == nil {
		t.Errorf("exp: %v, got: %v", stack.ErrStackEmpty, err)
	}
}

func TestStackTop(t *testing.T) {
	s, err := NewStack(1)
	if err != nil {
		t.Fatalf("exp stack create success, but error: %v", err)
	}

	err = s.Push(1)
	if err != nil {
		t.Fatalf("exp: nil, got: %v", err)
	}

	elem, err := s.Top()
	if err != nil {
		t.Fatalf("exp: %v, got: %v", nil, err)
	}

	v := elem.(int)
	if v != 1 {
		t.Errorf("exp: %v, got: %v", 1, v)
	}

	empty := s.IsEmpty()
	if empty {
		t.Errorf("exp: %v, got: %v", false, empty)
	}

	_, err = s.Pop()
	if err != nil {
		t.Fatalf("exp: nil, got: %v", err)
	}

	_, err = s.Top()
	if err == nil {
		t.Fatalf("exp: %v, got: %v", stack.ErrStackEmpty, err)
	}
}
