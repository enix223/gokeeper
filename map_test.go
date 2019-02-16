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
