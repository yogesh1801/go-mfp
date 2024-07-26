// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network state snapshot

package netstate

import (
	"net"
	"sort"
)

// snapshot represents a snapshot of current network state.
type snapshot struct {
	addrs []Addr // Known addresses
}

// newNetstate creates a snapshot of a current network state
func newSnapshot() (snapshot, error) {
	// Get interfaces
	ift, err := hookNetInterfaces()
	if err != nil {
		return snapshot{}, err
	}

	// Get addresses
	addrs := []Addr{}
	for _, ifi := range ift {
		// Get addresses of the interface
		ifat, err := hookNetInterfacesAddrs(&ifi)
		if err != nil {
			// Interface might disappear just from our hands,
			// so just skip it in a case of error.
			continue
		}

		// Convert obtained addresses into []*Addr
		nif := NetIfFromInterface(ifi)
		ifaddrs := make([]Addr, 0, len(ifat))
		for _, ifa := range ifat {
			// Interface addresses must be of the type *net.IPNet,
			// but be prepared if they are not, just in case
			if ipnet, ok := ifa.(*net.IPNet); ok {
				addr := AddrFromIPNet(*ipnet, nif)
				ifaddrs = append(addrs, addr)
			}
		}

		// Markup Primary addresses
		for _, a1 := range ifaddrs {
			for _, a2 := range ifaddrs {
				if a1 != a2 && a1.Primary && a1.Narrower(a2) {
					a1.Primary = false
				}
			}
		}

		// Append addresses to the main slice
		addrs = append(addrs, ifaddrs...)
	}

	// Return a snapshot
	return newSnapshotFromAddrs(addrs), nil
}

// newNetstate creates a snapshot of a current network state
// from provided addresses.
//
// It takes ownership on slice of addresses and may modify it.
func newSnapshotFromAddrs(addrs []Addr) snapshot {
	sort.Slice(addrs, func(i, j int) bool {
		return addrs[i].Less(addrs[j])
	})

	ns := snapshot{
		addrs: addrs,
	}

	return ns
}

// equal tells if two snapshots are equal
func (snap snapshot) equal(snap2 snapshot) bool {
	prev := snap.addrs
	next := snap2.addrs

	// Skip common addresses
	for len(prev) > 0 && len(next) > 0 && prev[0].Equal(next[0]) {
		prev = prev[1:]
		next = next[1:]
	}

	return len(prev) == 0 && len(next) == 0
}

// sync generates a series of events in order to bring 'snap'
// into the same state as snap2.
func (snap snapshot) sync(snap2 snapshot) (events []Event) {
	// Initialize things
	prev := snap.addrs
	next := snap2.addrs

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

			addr := prev[0]
			prev = prev[1:]

			nif := addr.Interface()
			cnt := interfaces.del(nif)

			if addr.Primary {
				events = append(events,
					EventDelPrimaryAddress{addr})
			}

			events = append(events, EventDelAddress{addr})

			if cnt == 1 {
				// If interface is not longer in use, report
				// its removal.
				events = append(events, EventDelInterface{nif})
			}

		case len(prev) == 0,
			len(prev) > 0 && len(next) > 0 && next[0].Less(prev[0]):

			addr := next[0]
			next = next[1:]

			nif := addr.Interface()
			cnt := interfaces.add(nif)

			if cnt == 0 {
				// If interface was not in use before, report
				// its arrival.
				events = append(events, EventAddInterface{nif})
			}

			events = append(events, EventAddAddress{addr})

			if addr.Primary {
				events = append(events,
					EventAddPrimaryAddress{addr})
			}

		default:
			// Note:
			//   - Neither prev or next are empty.
			//   - Neither prev[0].Less(next[0]) or visa versa
			//
			// It means, prev[0] and next[0] are equal.
			//
			// So just report change of Primary address
			// state if it occurs.
			aprev, anext := prev[0], next[0]
			prev, next = prev[1:], next[1:]

			switch {
			case aprev.Primary && !anext.Primary:
				events = append(events,
					EventDelPrimaryAddress{aprev})
			case !aprev.Primary && anext.Primary:
				events = append(events,
					EventAddPrimaryAddress{anext})
			}
		}
	}

	return
}
