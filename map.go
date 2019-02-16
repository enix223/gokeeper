package gokeeper

import (
	"strings"
)

// ExpandKeyMap expand the string map
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
func ExpandKeyMap(m map[string]interface{}, separator string) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range m {
		if len(k) > 1 {
			e := ExpandKey(k, separator, v)
			mergeMap(res, e)
		} else {
			// key contains no sep
			res[k] = v
		}
	}
	return res
}

// ExpandKey expand the `sep` separated key into embedded map
func ExpandKey(key, sep string, value interface{}) map[string]interface{} {
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

// mergeMap merge the r map into l map
func mergeMap(l, r map[string]interface{}) {
	for k, v := range r {
		l[k] = v
	}
}
