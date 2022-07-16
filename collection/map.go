package collection

import (
	"reflect"
	"strings"
)

// GetMapValue get value from map, or return defaultVal when key not exist
func GetMapValue(m interface{}, key interface{}, defaultVal interface{}) interface{} {
	val := reflect.ValueOf(m)
	k := reflect.ValueOf(key)
	if val.Kind() != reflect.Map {
		panic("m must be a map")
	}

	value := val.MapIndex(k)
	if !value.IsValid() {
		return defaultVal
	}

	return value.Interface()
}

// ExtractKeyMap extract the key, value pair with same key prefix, and output the key value as a new map
// m should be a map, and the key type should be string or it will panic
// eg.,
//   m := map[string]interface{}{
//     "data.abc": 1,
//     "data.bbc": "hello",
//   }
// After extracted, the output is:
//   o := map[string]interface{}{
// 		"abc": 1,
// 		"bbc": "hello",
// 	 }
func ExtractKeyMap(m interface{}, prefix string) interface{} {
	typ := reflect.TypeOf(m)
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		panic("m must be a map")
	}
	if typ.Key().Kind() != reflect.String {
		panic("key of map m should be type of string")
	}

	res := reflect.MakeMap(reflect.MapOf(typ.Key(), typ.Elem()))
	for _, k := range val.MapKeys() {
		ks := k.String()
		if strings.HasPrefix(ks, prefix) {
			ks = ks[len(prefix)+1:]
			res.SetMapIndex(reflect.ValueOf(ks), val.MapIndex(k))
		}
	}
	return res.Interface()
}

// ExpandKeyMap expand the string map
// m should be a map, and the key type should be string or it will panic
// Eg., we have a map like this:
//   m := map[string]interface{}{
//     "data.abc": 1,
//     "data.bbc": "hello",
//     "data.bbq.name": "foo"
//   }
// Expand the dot separated key into emmbedding map
//   m := map[string]interface{}{
//      "data": {
//          "abc": 1,
//          "bbc": "hello",
//          "bbq": {
//              "name": "foo"
//          }
//      }
//   }
func ExpandKeyMap(m interface{}, separator string) interface{} {
	typ := reflect.TypeOf(m)
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		panic("m must be a map")
	}
	if typ.Key().Kind() != reflect.String {
		panic("key of map m should be type of string")
	}

	res := reflect.MakeMap(reflect.MapOf(typ.Key(), typ.Elem()))
	for _, k := range val.MapKeys() {
		ks := k.String()
		if len(ks) > 1 {
			e := ExpandKey(ks, separator, val.MapIndex(k).Interface())
			mergeMap(res.Interface(), e)
		} else {
			// key contains no sep
			res.SetMapIndex(k, val.MapIndex(k))
		}
	}
	return res.Interface()
}

// ExpandKey expand the `sep` separated key into embedded map
func ExpandKey(key, sep string, value interface{}) interface{} {
	i := strings.Index(key, sep)
	if i == -1 {
		return map[string]interface{}{
			key: value,
		}
	}

	ckey, skey := key[:i], key[i+1:]
	return map[string]interface{}{
		ckey: ExpandKey(skey, sep, value),
	}
}

// Get keys of the map
// Panic if m is not a map or m is nil, if m size is 0, then nil is return
func KeysOfMap(m interface{}) interface{} {
	if m == nil {
		panic("m should not be nil")
	}
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		panic("m should be type of map")
	}
	if val.Len() == 0 {
		return nil
	}

	keys := val.MapKeys()
	res := reflect.MakeSlice(reflect.SliceOf(keys[0].Type()), 0, len(keys))
	res = reflect.Append(res, keys...)
	return res.Interface()
}

// Get values of the map
// Panic if m is not a map or m is nil, if m size is 0, then nil is return
func ValuesOfMap(m interface{}) []interface{} {
	if m == nil {
		panic("m should not be nil")
	}
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		panic("m should be type of map")
	}
	if val.Len() == 0 {
		return nil
	}

	res := make([]interface{}, 0, val.Len())
	iter := val.MapRange()
	for iter.Next() {
		res = append(res, iter.Value().Interface())
	}
	return res
}

// mergeMap merge the r map into l map
func mergeMap(l, r interface{}) {
	lval := reflect.ValueOf(l)
	rval := reflect.ValueOf(r)
	if lval.Kind() != reflect.Map {
		panic("l must be a map")
	}
	if rval.Kind() != reflect.Map {
		panic("r must be a map")
	}

	for _, k := range rval.MapKeys() {
		lval.SetMapIndex(k, rval.MapIndex(k))
	}
}
