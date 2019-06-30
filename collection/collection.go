package collection

import "reflect"

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