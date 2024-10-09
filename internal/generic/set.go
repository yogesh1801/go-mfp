// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Generic sets

package generic

// Set is the generic set of any comparable objects.
type Set[T comparable] struct {
	members map[T]struct{} // Members of the set
}

// NewSet creates a new Set
func NewSet[T comparable]() Set[T] {
	return Set[T]{
		members: make(map[T]struct{}),
	}
}

// Empty reports if set is empty
func (s Set[T]) Empty() bool {
	return len(s.members) == 0
}

// Contains reports if member already in the set
func (s Set[T]) Contains(member T) bool {
	_, found := s.members[member]
	return found
}

// Add adds member to the set
func (s Set[T]) Add(member T) {
	s.members[member] = struct{}{}
}

// Del deletes member from the set
func (s Set[T]) Del(member T) {
	delete(s.members, member)
}

// ForEach applies function to the each member of the set
func (s Set[T]) ForEach(f func(T)) {
	for member := range s.members {
		f(member)
	}
}
