package gokeeper

import (
	"reflect"
	"testing"
)

func TestExpandKeyMapWithInvalidKeyType(t *testing.T) {
	o := map[int]interface{}{
		1: 1,
		2: true,
	}
	defer func() {
		exp := "key of map m should be type of string"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	ExpandKeyMap(o, ".")
}

func TestExpandKeyMapWithInvalidMap(t *testing.T) {
	defer func() {
		exp := "m must be a map"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	ExpandKeyMap(1, "a.b")
}

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

func TestExtractKeyMapWithInvalidKeyType(t *testing.T) {
	o := map[int]interface{}{
		1: 1,
		2: true,
	}
	defer func() {
		exp := "key of map m should be type of string"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	ExtractKeyMap(o, "a.b")
}

func TestExtractKeyMapWithInvalidMap(t *testing.T) {
	defer func() {
		exp := "m must be a map"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	ExtractKeyMap(1, "a.b")
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

func TestMergeMap(t *testing.T) {
	l := map[string]interface{}{
		"a": 1,
		"b": true,
	}
	r := map[string]interface{}{
		"c": 2,
		"d": "d",
	}
	ll := map[string]interface{}{}
	for k, v := range l {
		ll[k] = v
	}
	mergeMap(ll, r)
	exp := map[string]interface{}{
		"a": 1,
		"b": true,
		"c": 2,
		"d": "d",
	}
	if !reflect.DeepEqual(exp, ll) {
		t.Fatalf("exp: %v, got: %v", exp, ll)
	}
}

func TestMergeMapWithInvalidLeftParam(t *testing.T) {
	defer func() {
		exp := "l must be a map"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	mergeMap(1, map[string]string{})
}

func TestMergeMapWithInvalidRightParam(t *testing.T) {
	defer func() {
		exp := "r must be a map"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	mergeMap(map[string]string{}, 2)
}
