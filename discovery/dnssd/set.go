// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Generic sets

package dnssd

// addrset manages set of addresses
type set[T comparable] struct {
	members map[T]struct{} // Members of the set
}

// newAddrset creates a new addrset
func newSet[T comparable]() set[T] {
	return set[T]{
		members: make(map[T]struct{}),
	}
}

// Empty reports if set is empty
func (s set[T]) Empty() bool {
	return len(s.members) == 0
}

// Contains reports if member already in the set
func (s set[T]) Contains(member T) bool {
	_, found := s.members[member]
	return found
}

// Add adds member to the set
func (s set[T]) Add(member T) {
	s.members[member] = struct{}{}
}

// Del deletes member from the set
func (s set[T]) Del(member T) {
	delete(s.members, member)
}
