package collection

import (
	"reflect"
)

// FilterFunc filter function
type FilterFunc func(v interface{}) bool

// IsIn test all elements in collection1 are in collection2
func IsIn(collection1, collection2 interface{}) bool {
	v1 := reflect.ValueOf(collection1)
	if v1.Kind() != reflect.Slice {
		panic("collection1 should be a slice")
	}

	v2 := reflect.ValueOf(collection2)
	if v2.Kind() != reflect.Slice {
		panic("collection2 should be a slice")
	}

	if v1.Type() != v2.Type() {
		panic("collection1 and collection2 should be with the same type")
	}

	for i := 0; i < v1.Len(); i++ {
		if !Contain(collection2, v1.Index(i).Interface()) {
			return false
		}
	}

	return true
}

// Contain test if `item` in `collection`
func Contain(collection interface{}, item interface{}) bool {
	v := reflect.ValueOf(collection)
	if v.Kind() != reflect.Slice {
		panic("collection should be a slice")
	}

	for i := 0; i < v.Len(); i++ {
		val := v.Index(i)
		if reflect.DeepEqual(val.Interface(), item) {
			return true
		}
	}

	return false
}

// MapFunc convert one instance to another instance
type MapFunc func(idx int, item interface{}) interface{}

// SliceMap iterate from the slice, and apply mapFunc to each item
// and combine the map result to a new slice
func SliceMap(slice interface{}, mapFunc MapFunc) interface{} {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic("slice should be a slice")
	}

	if val.Len() == 0 {
		return nil
	}

	first := mapFunc(0, val.Index(0).Interface())
	firstVal := reflect.ValueOf(first)
	res := reflect.MakeSlice(reflect.SliceOf(firstVal.Type()), 0, val.Len())
	res = reflect.Append(res, firstVal)

	for i := 1; i < val.Len(); i++ {
		nItem := mapFunc(i, val.Index(i).Interface())
		nItemVal := reflect.ValueOf(nItem)
		res = reflect.Append(res, nItemVal)
	}

	return res.Interface()
}

// FilterSlice the given slice with given iterator function
func FilterSlice(slice interface{}, iterator FilterFunc) interface{} {
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)

	slicePtr := t.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Slice

	if t.Kind() != reflect.Slice && !slicePtr {
		panic("slice should be slice")
	}

	if slicePtr {
		v = v.Elem()
		t = v.Type()
	}

	r := reflect.MakeSlice(t, 0, 0)

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		if iterator(item.Interface()) {
			r = reflect.Append(r, item)
		}
	}

	return r.Interface()
}

// TestFunc bool test iterator function
type TestFunc func(item interface{}) bool

// Any enumerate each element in given slice, and call the iterator function
// to check if it is true, if any of the iterator function return
// value is true, then return true, otherwise return false
func Any(slice interface{}, iterator TestFunc) bool {
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)

	slicePtr := t.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Slice

	if t.Kind() != reflect.Slice && !slicePtr {
		panic("slice should be slice")
	}

	if slicePtr {
		v = v.Elem()
		t = v.Type()
	}

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		if iterator(item.Interface()) {
			return true
		}
	}

	return false
}

// AnyNil enumerate each element in the slice to check if it is nil
// if any of the element is nil, return true, otherwise return false
func AnyNil(slice interface{}) bool {
	return Any(slice, func(item interface{}) bool {
		return item == nil
	})
}

// All enumerate each element in given slice, and call the iterator function
// to check if it is true, if any of the iterator function return
// value is true, then return true, otherwise return false
func All(slice interface{}, iterator TestFunc) bool {
	t := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)

	slicePtr := t.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Slice

	if t.Kind() != reflect.Slice && !slicePtr {
		panic("slice should be slice")
	}

	if slicePtr {
		v = v.Elem()
		t = v.Type()
	}

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		v := iterator(item.Interface())
		if !v {
			return false
		}
	}

	return true
}

// AllNil enumerate each element in the slice to check if it is nil
// if all of them are nil, return true, otherwise return false
func AllNil(slice interface{}) bool {
	return All(slice, func(item interface{}) bool {
		return item == nil
	})
}
