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

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// snapshot represents a snapshot of current network state.
type snapshot struct {
	addrs []snapshotAddr // Known addresses
}

// snapshotAddr represents network interface address equipped with
// additional information for internal use
type snapshotAddr struct {
	Addr         // Wrapped address
	primary bool // It's a primary address
}

// less orders [Addr] for sorting.
func (saddr snapshotAddr) less(saddr2 snapshotAddr) bool {
	return saddr.Addr.Less(saddr2.Addr)
}

// lessUnmasked compares addresses line snapshotAddr.less, but
// difference in address masks is ignored.
func (saddr snapshotAddr) lessUnmasked(saddr2 snapshotAddr) bool {
	return saddr.Unmasked().Less(saddr2.Unmasked())
}

// narrower reports whether addr is narrower that addr2.
func (saddr snapshotAddr) narrower(saddr2 snapshotAddr) bool {
	return saddr.Addr.Narrower(saddr2.Addr)
}

// SameInterface reports if two addresses belong to the same
// network interface.
func (saddr snapshotAddr) sameInterface(saddr2 snapshotAddr) bool {
	return saddr.Addr.SameInterface(saddr2.Addr)
}

// newNetstate creates a snapshot of a current network state
func newSnapshot() (snapshot, error) {
	// Get interfaces
	ift, err := hookNetInterfaces()
	if err != nil {
		return snapshot{}, err
	}

	// Get addresses
	addrs := make([]Addr, 0, len(ift))
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
		for _, ifa := range ifat {
			// Interface addresses must be of the type *net.IPNet,
			// but be prepared if they are not, just in case
			if ipnet, ok := ifa.(*net.IPNet); ok {
				addr := AddrFromIPNet(*ipnet, nif)
				addrs = append(addrs, addr)
			}
		}
	}

	// Make a snapshot
	return newSnapshotFromAddrs(addrs), nil
}

// newNetstate creates a snapshot of a current network state
// from provided addresses.
func newSnapshotFromAddrs(addrs []Addr) snapshot {
	// Copy and sort addresses
	saddrs := make([]snapshotAddr, len(addrs))
	for i := range addrs {
		saddrs[i] = snapshotAddr{addrs[i], true}
	}

	sort.Slice(saddrs, func(i, j int) bool {
		return saddrs[i].less(saddrs[j])
	})

	// Markup primary addresses
	for beg := 0; beg < len(saddrs); {
		end := beg + 1
		for end < len(saddrs) &&
			saddrs[beg].sameInterface(saddrs[end]) {
			end++
		}

		// Now saddrs[beg:end] belongs to the same interface.
		// Markup primary addresses within it.
		for i := beg; i < end; i++ {
			for j := beg; j < end; j++ {
				if i != j {
					if saddrs[i].primary &&
						saddrs[i].narrower(saddrs[j]) {
						saddrs[i].primary = false
					}
				}
			}
		}

		beg = end
	}

	return snapshot{saddrs}
}

// Equal tells if two snapshots are equal
func (snap snapshot) Equal(snap2 snapshot) bool {
	return generic.EqualSlices(snap.addrs, snap2.addrs)
}

// Interfaces returns slice of network interfaces in the snapshot.
func (snap snapshot) Interfaces() (netifs []NetIf) {
	ifset := make(map[NetIf]struct{})
	for _, addr := range snap.addrs {
		nif := addr.Interface()
		if _, found := ifset[nif]; !found {
			ifset[nif] = struct{}{}
			netifs = append(netifs, nif)
		}
	}

	return
}

// Addrs returns slice of network addresses in the snapshot.
func (snap snapshot) Addrs() (addrs []Addr) {
	addrs = make([]Addr, len(snap.addrs))
	for i := range snap.addrs {
		addrs[i] = snap.addrs[i].Addr
	}

	return
}

// PrimaryAddrs returns slice of primary network addresses in the snapshot.
func (snap snapshot) PrimaryAddrs() (addrs []Addr) {
	addrs = make([]Addr, 0, len(snap.addrs))

	for i := range snap.addrs {
		if snap.addrs[i].primary {
			addrs = append(addrs, snap.addrs[i].Addr)
		}
	}

	return
}

// sync generates a series of events in order to bring 'snap'
// into the same state as snap2.
func (snap snapshot) Sync(snap2 snapshot) (events []Event) {
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
		nextDone := len(next) == 0
		prevDone := len(prev) == 0
		noneDone := !(nextDone || prevDone)

		switch {
		case nextDone, noneDone && prev[0].lessUnmasked(next[0]):

			addr := prev[0]
			prev = prev[1:]

			nif := addr.Addr.Interface()
			cnt := interfaces[nif]
			interfaces[nif] = cnt - 1

			if addr.primary {
				events = append(events,
					EventDelPrimaryAddress{addr.Addr})
			}

			events = append(events, EventDelAddress{addr.Addr})

			if cnt == 1 {
				// If interface is not longer in use, report
				// its removal.
				events = append(events, EventDelInterface{nif})
			}

		case prevDone, noneDone && next[0].lessUnmasked(prev[0]):

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

			if addr.primary {
				events = append(events,
					EventAddPrimaryAddress{addr.Addr})
			}

		default:
			// Note:
			//   - Neither prev or next are empty.
			//   - Neither prev[0].Less(next[0]) or visa versa
			//
			// It means, prev[0] and next[0] are equal by their
			// IP address.
			//
			// So here we just report possible changes in
			// address mask or Primary address state.
			aprev, anext := prev[0], next[0]
			prev, next = prev[1:], next[1:]

			if aprev == anext {
				// The fast path: nothing changed
				break
			}

			if aprev.primary {
				events = append(events,
					EventDelPrimaryAddress{aprev.Addr})
			}

			if aprev.Addr != anext.Addr {
				events = append(events,
					EventDelAddress{aprev.Addr})
				events = append(events,
					EventAddAddress{anext.Addr})
			}

			if anext.primary {
				events = append(events,
					EventAddPrimaryAddress{anext.Addr})
			}
		}
	}

	return
}
