package gokeeper

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
