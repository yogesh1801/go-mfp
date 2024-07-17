// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of network interfaces

package netstate

import "net"

// setOfInterfaces represents set of network interfaces with usage counters
type setOfInterfaces struct {
	ift map[setOfInterfacesKey]int
}

// setOfInterfacesKey identifies network interface for the setOfInterfaces.
type setOfInterfacesKey struct {
	name  string // Interface name
	index int    // Interface index<F2>
}

// newSetOfInterfaces returns a new setOfInterfaces
func newSetOfInterfaces() *setOfInterfaces {
	return &setOfInterfaces{
		ift: make(map[setOfInterfacesKey]int),
	}
}

// key returns setOfInterfacesKey for the net.Interface
func (*setOfInterfaces) key(ifi net.Interface) setOfInterfacesKey {
	return setOfInterfacesKey{ifi.Name, ifi.Index}
}

// contains returns usage counter for the network interface in the set.
func (set *setOfInterfaces) contains(ifi net.Interface) int {
	key := set.key(ifi)
	return set.ift[key]
}

// add adds network interface to the set and returns the previous
// value of the usage counter.
//
// Interface may be added multiple times; each time its usage counter
// is incremented.
func (set *setOfInterfaces) add(ifi net.Interface) int {
	key := set.key(ifi)
	prev := set.ift[key]
	set.ift[key] = prev + 1
	return prev
}

// del deletes network interface from the set and returns the previous
// value of the usage counter.
func (set *setOfInterfaces) del(ifi net.Interface) int {
	key := set.key(ifi)
	prev := set.ift[key]
	set.ift[key] = prev - 1
	return prev
}

// addAddrs adds all interfaces mentioned in the slice of addresses
// into the set. If interface used multiple times, it will be added
// multiple times.
func (set *setOfInterfaces) addAddrs(addrs []*Addr) {
	for _, addr := range addrs {
		set.add(addr.Interface)
	}
}
