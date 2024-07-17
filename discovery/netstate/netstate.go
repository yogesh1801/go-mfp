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
	addrs []*Addr // Known addresses
}

// newNetstate creates a snapshot of a current network state
func newNetstate() (netstate, error) {
	// Get interfaces
	ift, err := net.Interfaces()
	if err != nil {
		return netstate{}, err
	}

	// Get addresses
	addrs := []*Addr{}
	for _, ifi := range ift {
		ifat, err := ifi.Addrs()
		if err != nil {
			// Interface might disappear just from our hands,
			// so just skip it in a case of error.
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

	// Return a netstate
	ns := netstate{
		addrs: addrs,
	}

	return ns, nil
}

// equal tells if two netstates are equal
func (ns netstate) equal(ns2 netstate) bool {
	prev := ns.addrs
	next := ns2.addrs

	// Skip common addresses
	for len(prev) > 0 && len(next) > 0 && prev[0].Equal(next[0]) {
		prev = prev[1:]
		next = next[1:]
	}

	return len(prev) > 0 || len(next) > 0
}

// sync generates a series of events in order to bring 'ns'
// into the same state as nsNext.
func (ns netstate) sync(nsNext netstate) (events []Event) {
	// Initialize things
	prev := ns.addrs
	next := nsNext.addrs

	interfaces := newSetOfInterfaces()
	interfaces.addAddrs(prev)

	// Generate Events
	//
	// This is a variation of the classical algorithm which
	// "merges" two sorted sequences.
	//
	// The algorithm processes sequences item by item, taking
	// the "lowest" item of both above sequences are not empty,
	// or the next item of non-empty sequence, if other is empty,
	// until both sequences are processed.
	//
	// Note, as our goal is to generate events that will "transform"
	// "prev" to "next", taking item from the "prev" sequence generates
	// EventDelAddress event, and taking sequence from the "next"
	// sequence generates EventAddAddress event
	for len(prev) > 0 || len(next) > 0 {
		switch {
		case len(next) == 0,
			len(prev) > 0 && len(next) > 0 && prev[0].Less(next[0]):

			addr := *prev[0]
			prev = prev[1:]

			ifi := addr.Interface
			cnt := interfaces.del(ifi)

			events = append(events, EventDelAddress{addr})

			if cnt == 1 {
				// If interface is not longer in use, report
				// its removal.
				events = append(events, EventDelInterface{ifi})
			}

		case len(prev) == 0,
			len(prev) > 0 && len(next) > 0 && next[0].Less(prev[0]):

			addr := *next[0]
			next = next[1:]

			ifi := addr.Interface
			cnt := interfaces.add(ifi)

			if cnt == 0 {
				// If interface was not in use before, report
				// its arrival.
				events = append(events, EventAddInterface{ifi})
			}

			events = append(events, EventAddAddress{addr})

		default:
			// Note:
			//   - Neither prev or next are empty.
			//   - Neither prev[0].Less(next[0]) or visa versa
			//
			// It means, prev[0] and next[0] are equal,
			// so just skip both
			prev = prev[1:]
			prev = prev[1:]
		}
	}

	return
}
