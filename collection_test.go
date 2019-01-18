package gokeeper

import "testing"

func TestIsIn(t *testing.T) {
	c1 := []byte{1, 2, 3, 4}
	c2 := []byte{1, 2, 3, 4, 5}

	if !IsIn(c1, c2) {
		t.Errorf(expectValue, true, false)
	}

	if IsIn(c2, c1) {
		t.Errorf(expectValue, false, true)
	}
}

func TestIsInWithDiffType(t *testing.T) {
	c1 := []byte{1, 2, 3, 4}
	c2 := []int{1, 2, 3, 4, 5}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf(expectValue, "panic", "nil")
		}
	}()

	IsIn(c1, c2)
}

func TestIsInWithWrongTypeC1(t *testing.T) {
	c1 := 1
	c2 := []int{1, 2, 3, 4, 5}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf(expectValue, "panic", "nil")
		}
	}()

	IsIn(c1, c2)
}

func TestIsInWithWrongTypeC2(t *testing.T) {
	c1 := []int{1, 2}
	c2 := 2

	defer func() {
		if r := recover(); r == nil {
			t.Errorf(expectValue, "panic", "nil")
		}
	}()

	IsIn(c1, c2)
}

func TestContainOK(t *testing.T) {
	c1 := []int{1, 2, 3, 4}

	if !Contain(c1, 1) {
		t.Errorf(expectValue, true, false)
	}

	if Contain(c1, 10) {
		t.Errorf(expectValue, false, true)
	}
}

func TestContainWithWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(expectValue, "panic", "nil")
		}
	}()

	Contain(2, 1)
}
