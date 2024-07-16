// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network state snapshot and event generation

package netstate

import (
	"net"
	"sort"
)

// netstate keeps current network state
type netstate struct {
	err        error    // Last error to be reported
	addrs      []*Addr  // Known addresses
	interfaces ifnetset // Known interfaces
}

// newNetstate creates a snapshot of a current network state
func newNetstate() netstate {
	// Get interfaces
	ift, err := net.Interfaces()
	if err != nil {
		return netstate{err: err}
	}

	// Get addresses
	addrs := []*Addr{}
	for _, ifi := range ift {
		ifat, err := ifi.Addrs()
		if err != nil {
			// Interface might disappear just from our hands,
			// so just skip it in a case of an error.
			continue
		}

		for _, ifa := range ifat {
			// Interface addresses must be of the type *net.IPNet,
			// but be prepared if they are not, just in case
			if ipnet, ok := ifa.(*net.IPNet); ok {
				addr := &Addr{*ipnet, ifi}
				addrs = append(addrs, addr)
			}
		}
	}

	sort.Slice(addrs, func(i, j int) bool {
		return addrs[i].Less(addrs[j])
	})

	// And build table of interfaces. As net.Interface is
	// not comparable, we can't use Go map here. So we
	// represent interfaces as a slice of fake Addresses,
	// with a single address per interface.
	//
	// Notice, we don't use output of net.Interfaces, because
	// our list of interfaces is filtered. It doesn't contain
	// interfaces without addresses. But we use size of that
	// output as a size hint for our own table.
	//
	// As table of interfaces is build from a sorted table of
	// addresses, it appears to be naturally sorted.
	interfaces := ifnetset{}
	for _, addr := range addrs {
		interfaces.add(addr.Interface)
	}

	// Return a netstate
	return netstate{
		addrs:      addrs,
		interfaces: interfaces,
	}
}

// setError sets an error. It returns true if something changed
func (ns *netstate) setError(err error) bool {
	ret := ns.err != err
	ns.err = err
	return ret
}

// equal tells if two netstates are equal
func (ns netstate) equal(ns2 netstate) bool {
	if ns.err != ns2.err {
		return false
	}

	if !ns.interfaces.equal(ns2.interfaces) {
		return false
	}

	evnt, _ := ns.addrsCmp(ns.addrs, ns2.addrs)
	return evnt == nil
}

// sync generates a series of events in order to bring 'ns'
// into the same state as nsNext.
//
// Each call to netstate.sync generates a single event and updates ns state,
// as if this event was applied to the state.
//
// If two network states are equal, it returns nil.
func (ns *netstate) sync(nsNext netstate) Event {
	if ns.err == nil && nsNext.err != nil {
		ns.err = nsNext.err
		return EventError{ns.err}
	}

	// Check for pending EventDelInterface events
	if ifi, ok := ns.interfaces.missed(nsNext.interfaces); ok {
		ns.interfaces.del(ifi)
		return EventDelInterface{ifi}
	}

	// Check for changes in addresses
	evnt, newlist := ns.addrsCmp(ns.addrs, nsNext.addrs)
	if add, ok := evnt.(EventAddAddress); ok {
		// Make sure EventAddInterface raised before
		// EventAddAddress.
		if ns.interfaces.add(add.Addr.Interface) {
			return EventAddInterface{add.Addr.Interface}
		}
	}

	ns.addrs = newlist
	return evnt
}

// addrsCmp find the first difference between two sorted
// slices of addresses.
//
// It returns EventAddAddress, if the first mismatching
// address present in 'prev' and missed in 'next', EventDelAddress,
// if the first mismatching address present in 'next' and missed
// in 'prev' or nil, if lists are equal.
//
// Additionaly, as a second return value, it returns an updated
// version of 'prev', as if event was applied to it.
func (*netstate) addrsCmp(prev, next []*Addr) (Event, []*Addr) {
	// Skip common addresses
	i := 0
	for i < len(prev) && i < len(next) && prev[i].Equal(next[i]) {
		i++
	}

	// Generate an Event
	switch {
	case i == len(prev) && i == len(next):
		// Lists are equal
		return nil, prev

	case i < len(prev) && i == len(next):
		// Prev contains at least one extra event
		evnt := EventDelAddress{*prev[i]}
		prev = append(prev[0:i], prev[i+1:]...)
		return evnt, prev

	case len(prev) == 0 && len(next) < 0:
		// Next contains at least one extra event
		evnt := EventAddAddress{*next[i]}
		prev = append(prev[0:i], next[i])
		return evnt, prev
	}

	// prev[i] and next[i] are different.
	//
	// We may generate either EventDelAddress{*prev[i]) or
	// EventAddAddress{*next[i]}, both are correct.
	//
	// Prefer EventDelAddress{...} without some special reason.
	evnt := EventDelAddress{*prev[i]}
	prev = append(prev[0:i], prev[i+1:]...)
	return evnt, prev
}
