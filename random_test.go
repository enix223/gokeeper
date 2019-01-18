package gokeeper

import (
	"math/rand"
	"testing"
	"time"
)

const expectValue = "expect: %v, got: %v"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestRandomString(t *testing.T) {
	res := Random(Alphabet, 2)
	if v, ok := res.([]byte); !ok {
		t.Fatal("result should be []byte type")
	} else {
		if len(v) != 2 {
			t.Fatalf(expectValue, 2, len(v))
		}

		if !IsIn(res, []byte(Alphabet)) {
			t.Fatalf(expectValue, "", res)
		}
	}
}

func TestRandomWithExceedLength(t *testing.T) {
	l := 100
	res := Random(Alphabet, l)
	if v, ok := res.([]byte); !ok {
		t.Fatal("result should be []byte type")
	} else {
		if len(v) != l {
			t.Fatalf(expectValue, l, len(v))
		}

		if !IsIn(res, []byte(Alphabet)) {
			t.Fatalf(expectValue, "", res)
		}
	}
}

func TestRandomSlice(t *testing.T) {
	candidates := []int{1, 2, 3, 4, 5}
	res := Random(candidates, 2)
	if v, ok := res.([]int); !ok {
		t.Fatal("result should be []int type")
	} else {
		if len(v) != 2 {
			t.Fatalf(expectValue, 2, len(v))
		}
	}

	if !IsIn(res, candidates) {
		t.Fatalf(expectValue, "", res)
	}
}

func TestRandomWithNonSlice(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The Random should panic")
		}
	}()

	Random(10, 1)
}

func TestRandomStringFunc(t *testing.T) {
	candidates := []string{"aa", "bb", "cc", "dd"}
	res := RandomString(candidates, 2)
	if len(res) != 2 {
		t.Errorf(expectValue, 2, len(res))
	}
	if !IsIn(res, candidates) {
		t.Fatalf(expectValue, "", res)
	}
}

func TestRandomByteFunc(t *testing.T) {
	res := RandomBytes(Alphabet, 2)
	if len(res) != 2 {
		t.Fatalf(expectValue, 2, len(res))
	}
	if !IsIn([]byte(res), Alphabet) {
		t.Fatalf(expectValue, "", res)
	}
}

func TestRandomAlphabetDigits(t *testing.T) {
	l := 3
	res := RandomAlphabetDigits(l)
	if len(res) != l {
		t.Fatalf(expectValue, l, len(res))
	}

	candidates := append(Alphabet, Digits...)
	if !IsIn([]byte(res), candidates) {
		t.Fatalf(expectValue, "", res)
	}
}

func TestRandomDigits(t *testing.T) {
	l := 3
	res := RandomDigits(l)
	if len(res) != l {
		t.Fatalf(expectValue, l, len(res))
	}

	if !IsIn([]byte(res), Digits) {
		t.Fatalf(expectValue, "", res)
	}
}
