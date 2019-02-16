package gokeeper

import (
	"reflect"
	"testing"
)

func TestExpandKeyWithDotSep(t *testing.T) {
	k := "a.b.c.d"
	m := ExpandKey(k, ".", 1)
	exp := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": map[string]interface{}{
					"d": 1,
				},
			},
		},
	}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("exp: %v, got: %v", exp, m)
	}
}

func TestExpandKeyWithSlashSep(t *testing.T) {
	k := "a/b/c/d"
	m := ExpandKey(k, "/", 1)
	exp := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": map[string]interface{}{
					"d": 1,
				},
			},
		},
	}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("exp: %v, got: %v", exp, m)
	}
}

func TestExpandKeyWithoutSep(t *testing.T) {
	k := "a.b.c.d"
	m := ExpandKey(k, "/", 1)
	exp := map[string]interface{}{
		"a.b.c.d": 1,
	}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("exp: %v, got: %v", exp, m)
	}
}

func TestExpandKeyMap(t *testing.T) {
	m := ExpandKeyMap(
		map[string]interface{}{
			"a.b.c.d": 1,
			"e.f":     true,
			"g":       "3",
		},
		".",
	)
	exp := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": map[string]interface{}{
					"d": 1,
				},
			},
		},
		"e": map[string]interface{}{
			"f": true,
		},
		"g": "3",
	}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("exp: %v, got: %v", exp, m)
	}
}

func TestExpandKeyMapWithoutSep(t *testing.T) {
	o := map[string]interface{}{
		"a.b.c.d": 1,
		"e.f":     true,
		"g":       "3",
	}
	m := ExpandKeyMap(o, "/")
	if !reflect.DeepEqual(m, o) {
		t.Fatalf("exp: %v, got: %v", o, m)
	}
}

func TestExtractKeyMap(t *testing.T) {
	o := map[string]interface{}{
		"a.b.c": 1,
		"a.b.d": true,
	}
	m := ExtractKeyMap(o, "a.b")
	exp := map[string]interface{}{
		"c": 1,
		"d": true,
	}
	if !reflect.DeepEqual(m, exp) {
		t.Fatalf("exp: %v, got: %v", exp, m)
	}
}

func TestExtractKeyMapWithNoPrefixMatch(t *testing.T) {
	o := map[string]interface{}{
		"a.b.c": 1,
		"a.b.d": true,
	}
	m := ExtractKeyMap(o, "x.y")
	if v, ok := m.(map[string]interface{}); !(len(v) == 0 && ok) {
		t.Fatalf("exp: %v, got: %v", "empty map", m)
	}
}
