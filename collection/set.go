package collection

// Set set implemented by map
type Set map[interface{}]struct{}

// SetMapFunc map function
type SetMapFunc func(v interface{}) interface{}

// NewSet create a new set
func NewSet() Set {
	return make(map[interface{}]struct{})
}

// Size get the size of the set
func (s Set) Size() int {
	return len(s)
}

// Add element into set
func (s Set) Add(v ...interface{}) Set {
	for _, val := range v {
		s[val] = struct{}{}
	}
	return s
}

// Has check if element v in the set
func (s Set) Has(v interface{}) bool {
	_, ok := s[v]
	return ok
}

// Remove remove element(s) from the set
func (s Set) Remove(v ...interface{}) Set {
	for _, val := range v {
		delete(s, val)
	}
	return s
}

// Intersect calculate the intersection of current set against the target set
func (s Set) Intersect(a Set) Set {
	n := NewSet()

	for k := range s {
		if _, ok := a[k]; ok {
			n[k] = struct{}{}
		}
	}

	return n
}

// Union calculate the union of the current set and the target set
func (s Set) Union(a Set) Set {
	n := NewSet()

	for k := range s {
		n[k] = struct{}{}
	}

	for k := range a {
		n[k] = struct{}{}
	}

	return n
}

// Difference calculate the difference of the current set against the target set
func (s Set) Difference(a Set) Set {
	n := NewSet()

	for k := range s {
		if _, ok := a[k]; !ok {
			n[k] = struct{}{}
		}
	}

	return n
}

// Equal compare each element in the set s against set a, return true
// if all elements in set s equal to set a, and all element in set a equal
// to elements in s
func (s Set) Equal(a Set) bool {
	for k := range s {
		if _, ok := a[k]; !ok {
			return false
		}
	}

	for k := range a {
		if _, ok := s[k]; !ok {
			return false
		}
	}

	return true
}

// Map apply fn to each element in set s, and return a new set
func (s Set) Map(fn SetMapFunc) Set {
	ns := NewSet()

	for k := range s {
		ns[fn(k)] = struct{}{}
	}

	return ns
}

// Filter call fn on each element, and add it into new set when
// fn return true
func (s Set) Filter(fn FilterFunc) Set {
	ns := NewSet()

	for k := range s {
		if fn(k) {
			ns[k] = struct{}{}
		}
	}

	return ns
}
