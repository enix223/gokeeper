package utils

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func isInBytes(col1 []byte, col2 []byte) bool {
	for _, i := range col1 {
		found := false
		for _, j := range col2 {
			if j == i {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func isInStrings(col1 []string, col2 []string) bool {
	for _, i := range col1 {
		found := false
		for _, j := range col2 {
			if j == i {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func isInInt(col1 []int, col2 []int) bool {
	for _, i := range col1 {
		found := false
		for _, j := range col2 {
			if j == i {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestRandomString(t *testing.T) {
	res := Random(Alphabet, 2)
	if v, ok := res.([]byte); !ok {
		t.Fatal("result should be []byte type")
	} else {
		if len(v) != 2 {
			t.Fatalf("exp: %v, got: %v", 2, len(v))
		}

		if !isInBytes(v, Alphabet) {
			t.Fatalf("exp: %v, got: %v", "characters in alphabets", v)
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
			t.Fatalf("exp: %v, got: %v", l, len(v))
		}

		if !isInBytes(v, Alphabet) {
			t.Fatalf("exp: %v, got: %v", "characters in alphabets", v)
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
			t.Fatalf("exp: %v, got: %v", 2, len(v))
		}

		if !isInInt(v, candidates) {
			t.Fatalf("exp: %v, got: %v", "", res)
		}
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
		t.Errorf("exp: %v, got: %v", 2, len(res))
	}
	if !isInStrings(res, candidates) {
		t.Fatalf("exp: %v, got: %v", "characters in alphabets", res)
	}
}

func TestRandomByteFunc(t *testing.T) {
	res := RandomBytes(Alphabet, 2)
	if len(res) != 2 {
		t.Fatalf("exp: %v, got: %v", 2, len(res))
	}
	if !isInBytes([]byte(res), Alphabet) {
		t.Fatalf("exp: %v, got: %v", "", res)
	}
}

func TestRandomAlphabetDigits(t *testing.T) {
	l := 3
	res := RandomAlphabetDigits(l)
	if len(res) != l {
		t.Fatalf("exp: %v, got: %v", l, len(res))
	}

	candidates := append(Alphabet, Digits...)
	if !isInBytes([]byte(res), candidates) {
		t.Fatalf("exp: %v, got: %v", "", res)
	}
}

func TestRandomDigits(t *testing.T) {
	l := 3
	res := RandomDigits(l)
	if len(res) != l {
		t.Fatalf("exp: %v, got: %v", l, len(res))
	}

	if !isInBytes([]byte(res), Digits) {
		t.Fatalf("exp: %v, got: %v", "", res)
	}
}
