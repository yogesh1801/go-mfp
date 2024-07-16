// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of network interfaces

package netstate

import (
	"net"
)

// ifnetset represents a set of network interfaces
type ifnetset struct {
	ift map[ifnetsetKey]net.Interface
}

// contains reports if interface is in the set
func (set ifnetset) contains(ifi net.Interface) bool {
	if set.ift == nil {
		return false
	}

	key := makeIfnetsetKey(ifi)
	_, found := set.ift[key]

	return found
}

// equal reports if two sets are equal
func (set ifnetset) equal(set2 ifnetset) bool {
	// If both maps are empty, they are equal
	if len(set.ift) == 0 && len(set2.ift) == 0 {
		return true
	}

	// If maps lengths don't match, maps are not equal
	if len(set.ift) != len(set2.ift) {
		return false
	}

	// The long way, check item by item.
	for key := range set.ift {
		if _, found := set2.ift[key]; !found {
			return false
		}
	}

	for key := range set2.ift {
		if _, found := set.ift[key]; !found {
			return false
		}
	}

	return true
}

// missed checks for some interfaces, present in 'set' and missed in
// 'set2'.
//
// If at least one interface found, it returns interface and 'true'.
// Otherwise, it returns net.Interface{}, false
//
// If there are more that one missed interfaces, any of them may
// be returned.
func (set ifnetset) missed(set2 ifnetset) (net.Interface, bool) {
	for _, ifi := range set.ift {
		if !set2.contains(ifi) {
			return ifi, true
		}
	}

	return net.Interface{}, false
}

// add adds interface to the set.
// It returns true, if interface was not in the set before.
func (set ifnetset) add(ifi net.Interface) bool {
	key := makeIfnetsetKey(ifi)

	found := false
	if set.ift == nil {
		set.ift = make(map[ifnetsetKey]net.Interface)
	} else {
		_, found = set.ift[key]
	}

	if !found {
		set.ift[key] = ifi
	}

	return !found
}

// del deletes interface from the set.
// It returns true, if interface was in the set before.
func (set ifnetset) del(ifi net.Interface) bool {
	if set.ift == nil {
		return false
	}

	key := makeIfnetsetKey(ifi)
	_, found := set.ift[key]

	if found {
		delete(set.ift, key)
	}

	return found
}

// ifnetsetKey represents a search key for network interface
type ifnetsetKey struct {
	name  string // net.Interface.Name
	index int    // net.interface.Index
}

// makeIfnetsetKey returns ifnetsetKey for the net.Interface
func makeIfnetsetKey(ifi net.Interface) ifnetsetKey {
	return ifnetsetKey{
		name:  ifi.Name,
		index: ifi.Index,
	}
}
