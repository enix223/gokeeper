package collection

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

func TestGetMapValue(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": "2",
	}

	v := GetMapValue(m, "a", 2)
	val, ok := v.(int)
	if !ok {
		t.Fatalf("exp :%v, got: %v", true, ok)
	}

	exp := 1
	if val != exp {
		t.Fatalf("exp :%v, got: %v", exp, val)
	}

	v2 := GetMapValue(m, "c", 3)
	val2, ok := v2.(int)
	if !ok {
		t.Fatalf("exp :%v, got: %v", true, ok)
	}

	exp2 := 3
	if val2 != exp2 {
		t.Fatalf("exp :%v, got: %v", exp2, val2)
	}

	i1, i2 := 1, 2
	m2 := map[int]*int{
		1: &i1,
	}
	v3 := GetMapValue(m2, 1, &i2)
	val3, ok := v3.(*int)
	if !ok {
		t.Fatalf("exp :%v, got: %v", true, ok)
	}

	if val3 != &i1 {
		t.Fatalf("exp :%v, got: %v", &i1, val3)
	}

	v4 := GetMapValue(m2, 2, &i2)
	val4, ok := v4.(*int)
	if !ok {
		t.Fatalf("exp :%v, got: %v", true, ok)
	}

	if val4 != &i2 {
		t.Fatalf("exp :%v, got: %v", &i2, val4)
	}
}

func TestGetMapValueWithInvalidRightParam(t *testing.T) {
	defer func() {
		exp := "m must be a map"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	GetMapValue(1, 1, 1)
}

func TestKeyOfMapWithNormalMap(t *testing.T) {
	m := map[int]int{
		1: 2,
		2: 3,
		3: 4,
	}
	res := KeysOfMap(m)
	exp := []int{1, 2, 3}
	if !sliceElementEqual(res, exp) {
		t.Fatalf("exp: %v, got: %v", exp, res)
	}

	m2 := map[string]struct{}{
		"1": {},
		"2": {},
		"3": {},
	}
	res2 := KeysOfMap(m2)
	exp2 := []string{"1", "2", "3"}
	if !sliceElementEqual(res2, exp2) {
		t.Fatalf("exp: %v, got: %v", exp2, res2)
	}

	m3 := map[interface{}]interface{}{
		1:   1,
		"2": 2,
	}
	res3 := KeysOfMap(m3)
	exp3 := []interface{}{1, "2"}
	if !sliceElementEqual(res3, exp3) {
		t.Fatalf("exp: %v, got: %v", exp3, res3)
	}
}

func TestKeyOfMapWithNilMap(t *testing.T) {
	defer func() {
		exp := "m should not be nil"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	KeysOfMap(nil)
}

func TestKeyOfMapWithEmptyMap(t *testing.T) {
	var m map[int]string
	res := KeysOfMap(m)
	if res != nil {
		t.Fatalf("exp: nil, got: %v", res)
	}
}

func TestKeyOfMapWithNoMap(t *testing.T) {
	defer func() {
		exp := "m should be type of map"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	var m []string
	KeysOfMap(m)
}

func TestValueOfMapWithNormalMap(t *testing.T) {
	m := map[int]int{
		1: 2,
		2: 3,
		3: 4,
	}
	res := ValuesOfMap(m)
	exp := []int{2, 3, 4}
	if !sliceElementEqual(res, exp) {
		t.Fatalf("exp: %v, got: %v", exp, res)
	}

	type s struct {
		a int
	}
	s1 := s{a: 1}
	s2 := s{a: 2}
	s3 := s{a: 3}
	m2 := map[string]s{
		"1": s1,
		"2": s2,
		"3": s3,
	}
	res2 := ValuesOfMap(m2)
	exp2 := []s{s1, s2, s3}
	if !sliceElementEqual(res2, exp2) {
		t.Fatalf("exp: %v, got: %v", exp2, res2)
	}

	m3 := map[interface{}]interface{}{
		1:   1,
		"2": 2,
	}
	res3 := ValuesOfMap(m3)
	exp3 := []interface{}{1, 2}
	if !sliceElementEqual(res3, exp3) {
		t.Fatalf("exp: %v, got: %v", exp3, res3)
	}
}

func TestValueOfMapWithNilMap(t *testing.T) {
	defer func() {
		exp := "m should not be nil"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	ValuesOfMap(nil)
}

func TestValueOfMapWithEmptyMap(t *testing.T) {
	var m map[int]string
	res := ValuesOfMap(m)
	if res != nil {
		t.Fatalf("exp: nil, got: %v", res)
	}
}

func TestValueOfMapWithNoMap(t *testing.T) {
	defer func() {
		exp := "m should be type of map"
		if e := recover().(string); e != exp {
			t.Fatalf("exp: %v, got: %v", exp, e)
		}
	}()
	var m []string
	ValuesOfMap(m)
}

func sliceElementEqual(slice1, slice2 interface{}) bool {
	val1 := reflect.ValueOf(slice1)
	val2 := reflect.ValueOf(slice2)
	m := make(map[interface{}]struct{})
	if val1.Len() != val2.Len() {
		return false
	}
	for i := 0; i < val1.Len(); i++ {
		m[val1.Index(i).Interface()] = struct{}{}
	}
	for i := 0; i < val2.Len(); i++ {
		m[val2.Index(i).Interface()] = struct{}{}
	}
	return len(m) == val1.Len()
}
