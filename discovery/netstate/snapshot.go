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
	"slices"
	"sort"
)

// snapshot represents a snapshot of current network state.
type snapshot struct {
	addrs []snapshotAddr // Known addresses
}

// snapshotAddr represents network interface address equipped with
// additional information for internal use
type snapshotAddr struct {
	Addr         // Interface address
	Primary bool // It's a primary address
}

func (saddr snapshotAddr) Less(saddr2 snapshotAddr) bool {
	return saddr.Addr.Less(saddr2.Addr)
}

func (saddr snapshotAddr) Narrower(saddr2 snapshotAddr) bool {
	return saddr.Addr.Narrower(saddr2.Addr)
}

// newNetstate creates a snapshot of a current network state
func newSnapshot() (snapshot, error) {
	// Get interfaces
	ift, err := hookNetInterfaces()
	if err != nil {
		return snapshot{}, err
	}

	// Get addresses
	addrs := []snapshotAddr{}
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
		ifaddrs := make([]snapshotAddr, 0, len(ifat))
		for _, ifa := range ifat {
			// Interface addresses must be of the type *net.IPNet,
			// but be prepared if they are not, just in case
			if ipnet, ok := ifa.(*net.IPNet); ok {
				addr := AddrFromIPNet(*ipnet, nif)
				ifaddrs = append(addrs,
					snapshotAddr{addr, true})
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
	sort.Slice(addrs, func(i, j int) bool {
		return addrs[i].Less(addrs[j])
	})

	return snapshot{addrs}, nil
}

// newNetstate creates a snapshot of a current network state
// from provided addresses.
func newSnapshotFromAddrs(addrs []Addr) snapshot {
	saddrs := make([]snapshotAddr, len(addrs))
	for i := range addrs {
		saddrs[i].Addr = addrs[i]
	}

	sort.Slice(saddrs, func(i, j int) bool {
		return saddrs[i].Less(saddrs[j])
	})

	return snapshot{saddrs}
}

// equal tells if two snapshots are equal
func (snap snapshot) equal(snap2 snapshot) bool {
	return slices.Equal(snap.addrs, snap2.addrs)
}

// sync generates a series of events in order to bring 'snap'
// into the same state as snap2.
func (snap snapshot) sync(snap2 snapshot) (events []Event) {
	// Initialize things
	prev := snap.addrs
	next := snap2.addrs

	interfaces := make(map[NetIf]int)
	for _, addr := range prev {
		nif := addr.Interface()
		interfaces[nif] = interfaces[nif] + 1
	}

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
			cnt := interfaces[nif]
			interfaces[nif] = cnt - 1

			if addr.Primary {
				events = append(events,
					EventDelPrimaryAddress{addr.Addr})
			}

			events = append(events, EventDelAddress{addr.Addr})

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
			cnt := interfaces[nif]
			interfaces[nif] = cnt + 1

			if cnt == 0 {
				// If interface was not in use before, report
				// its arrival.
				events = append(events, EventAddInterface{nif})
			}

			events = append(events, EventAddAddress{addr.Addr})

			if addr.Primary {
				events = append(events,
					EventAddPrimaryAddress{addr.Addr})
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
					EventDelPrimaryAddress{aprev.Addr})
			case !aprev.Primary && anext.Primary:
				events = append(events,
					EventAddPrimaryAddress{anext.Addr})
			}
		}
	}

	return
}
