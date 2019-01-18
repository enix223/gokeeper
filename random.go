package gokeeper

import (
	"math/rand"
	"reflect"
)

// Alphabet english alphabet
var Alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Digits digits
var Digits = []byte("0123456789")

// Random create a random slice from given slice
func Random(collection interface{}, length int) interface{} {
	v := reflect.ValueOf(collection)
	if kind := v.Kind(); kind != reflect.Slice {
		panic("collection should be slie")
	}

	res := reflect.MakeSlice(v.Type(), 0, length)
	for i := 0; i < length; i++ {
		vv := v.Index(rand.Intn(v.Len()))
		res = reflect.Append(res, vv)
	}
	return res.Interface()
}

// RandomString return a random string slice with `length` length
func RandomString(collection []string, length int) []string {
	res := Random(collection, length).([]string)
	return res
}

// RandomBytes return a random byte slice with `length` length
func RandomBytes(collection []byte, length int) []byte {
	res := Random(collection, length).([]byte)
	return res
}

// RandomAlphabetDigits return random string with `length` length,
// and possible values from alphabets
func RandomAlphabetDigits(length int) string {
	candidates := append(Alphabet, Digits...)
	res := Random(candidates, length).([]byte)
	return string(res)
}

// RandomDigits return random string with `length` length,
// and possible values from digits
func RandomDigits(length int) string {
	res := Random(Digits, length).([]byte)
	return string(res)
}
