package collection

import "sync"

// SyncSet thread safe set
type SyncSet struct {
	mu  sync.RWMutex
	set map[interface{}]struct{}
}

// NewSyncSet create a new SyncSet
func NewSyncSet() *SyncSet {
	return &SyncSet{
		set: make(map[interface{}]struct{}),
	}
}

// Size get the size of the SyncSet
func (s *SyncSet) Size() int {
	return len(s.set)
}

// Add element into SyncSet
func (s *SyncSet) Add(v ...interface{}) *SyncSet {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, val := range v {
		s.set[val] = struct{}{}
	}
	return s
}

// Has check if element v in the SyncSet
func (s *SyncSet) Has(v interface{}) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.set[v]
	return ok
}

// Remove remove element(s) from the SyncSet
func (s *SyncSet) Remove(v ...interface{}) *SyncSet {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, val := range v {
		if _, ok := s.set[val]; ok {
			delete(s.set, val)
		}
	}
	return s
}

// Intersect calculate the intersection of current SyncSet against the target SyncSet
func (s *SyncSet) Intersect(a *SyncSet) *SyncSet {
	n := NewSyncSet()

	s.mu.RLock()
	defer s.mu.RUnlock()

	for k := range s.set {
		if _, ok := a.set[k]; ok {
			n.set[k] = struct{}{}
		}
	}

	return n
}

// Union calculate the union of the current SyncSet and the target SyncSet
func (s *SyncSet) Union(a *SyncSet) *SyncSet {
	n := NewSyncSet()

	s.mu.RLock()
	defer s.mu.RUnlock()

	for k := range s.set {
		n.set[k] = struct{}{}
	}

	for k := range a.set {
		n.set[k] = struct{}{}
	}

	return n
}

// Difference calculate the difference of the current SyncSet against the target SyncSet
func (s *SyncSet) Difference(a *SyncSet) *SyncSet {
	n := NewSyncSet()

	s.mu.RLock()
	defer s.mu.RUnlock()

	for k := range s.set {
		if _, ok := a.set[k]; !ok {
			n.set[k] = struct{}{}
		}
	}

	return n
}

// Equal compare each element in the SyncSet s against SyncSet a, return true
// if all elements in SyncSet s equal to SyncSet a, and all element in SyncSet a equal
// to elements in s
func (s *SyncSet) Equal(a *SyncSet) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for k := range s.set {
		if _, ok := a.set[k]; !ok {
			return false
		}
	}

	for k := range a.set {
		if _, ok := s.set[k]; !ok {
			return false
		}
	}

	return true
}

// Map apply fn to each element in SyncSet s, and return a new SyncSet
func (s *SyncSet) Map(fn SetMapFunc) *SyncSet {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ns := NewSyncSet()

	for k := range s.set {
		ns.set[fn(k)] = struct{}{}
	}

	return ns
}

// Filter call fn on each element, and add it into new SyncSet when
// fn return true
func (s *SyncSet) Filter(fn FilterFunc) *SyncSet {
	ns := NewSyncSet()

	s.mu.RLock()
	defer s.mu.RUnlock()

	for k := range s.set {
		if fn(k) {
			ns.set[k] = struct{}{}
		}
	}

	return ns
}
