// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Tests of network state snapshot

package netstate

import (
	"fmt"
	"slices"
	"testing"
)

// TestSnapshotEqual tests snapshot.equal
func TestSnapshotEqual(t *testing.T) {
	type testData struct {
		addrset1, addrset2 []Addr // Two sets of addresses
		eq                 bool   // Expected snapshot.equal answer
	}

	if0 := NetIf{0, "if0"}

	tests := []testData{
		{
			addrset1: nil,
			addrset2: nil,
			eq:       true,
		},

		{
			addrset1: []Addr{
				testMakeAddr(if0, "127.0.0.1/24"),
				testMakeAddr(if0, "192.168.0.1/24"),
			},
			addrset2: []Addr{
				testMakeAddr(if0, "192.168.0.1/24"),
				testMakeAddr(if0, "127.0.0.1/24"),
			},
			eq: true,
		},

		{
			addrset1: []Addr{
				testMakeAddr(if0, "127.0.0.1/24"),
			},
			addrset2: []Addr{
				testMakeAddr(if0, "192.168.0.1/24"),
				testMakeAddr(if0, "127.0.0.1/24"),
			},
			eq: false,
		},

		{
			addrset1: []Addr{
				testMakeAddr(if0, "127.0.0.1/24"),
				testMakeAddr(if0, "192.168.0.1/24"),
			},
			addrset2: []Addr{
				testMakeAddr(if0, "127.0.0.1/24"),
			},
			eq: false,
		},
	}

	for _, test := range tests {
		snap1 := newSnapshotFromAddrs(slices.Clone(test.addrset1))
		snap2 := newSnapshotFromAddrs(slices.Clone(test.addrset2))
		eq := snap1.Equal(snap2)
		if eq != test.eq {
			t.Errorf("snapshot.equal: expected %v, present %v\n"+
				"addrset1: %s\n"+
				"addrset2: %s",
				test.eq, eq,
				test.addrset1, test.addrset2)
		}
	}
}

// TestSnapshotMake tests creation of new snapshot from provided addresses.
func TestSnapshotMake(t *testing.T) {
	type testData struct {
		addrs   []Addr  // Input addresses
		netifs  []NetIf // Expected interfaces
		primary []Addr  // Expected primary addresses
	}

	if0 := NetIf{0, "if0"}
	if1 := NetIf{1, "if1"}

	tests := []testData{
		// Empty set
		{
			addrs: []Addr{},
		},

		// A single address
		{
			addrs: []Addr{
				testMakeAddr(if0, "127.0.0.1/24"),
			},
			netifs: []NetIf{if0},
			primary: []Addr{
				testMakeAddr(if0, "127.0.0.1/24"),
			},
		},

		// Two addresses from different interfaces
		{
			addrs: []Addr{
				testMakeAddr(if0, "192.168.0.1/24"),
				testMakeAddr(if1, "192.168.1.1/24"),
			},
			netifs: []NetIf{if0, if1},
			primary: []Addr{
				testMakeAddr(if0, "192.168.0.1/24"),
				testMakeAddr(if1, "192.168.1.1/24"),
			},
		},

		// One interface connected to 3 IP networks
		{
			addrs: []Addr{
				testMakeAddr(if0, "192.168.0.1/24"),
				testMakeAddr(if0, "192.168.1.1/24"),
				testMakeAddr(if0, "192.168.2.1/24"),
			},
			netifs: []NetIf{if0},
			primary: []Addr{
				testMakeAddr(if0, "192.168.0.1/24"),
				testMakeAddr(if0, "192.168.1.1/24"),
				testMakeAddr(if0, "192.168.2.1/24"),
			},
		},

		// Overlapping addresses, first is narrower
		{
			addrs: []Addr{
				testMakeAddr(if0, "192.168.0.1/32"),
				testMakeAddr(if0, "192.168.0.55/24"),
				testMakeAddr(if0, "192.168.1.1/32"),
				testMakeAddr(if0, "192.168.1.55/24"),
				testMakeAddr(if0, "192.168.2.1/32"),
				testMakeAddr(if0, "192.168.2.55/24"),
			},
			netifs: []NetIf{if0},
			primary: []Addr{
				testMakeAddr(if0, "192.168.0.55/24"),
				testMakeAddr(if0, "192.168.1.55/24"),
				testMakeAddr(if0, "192.168.2.55/24"),
			},
		},

		// Overlapping addresses, second is narrower
		{
			addrs: []Addr{
				testMakeAddr(if0, "192.168.0.1/24"),
				testMakeAddr(if0, "192.168.0.55/32"),
				testMakeAddr(if0, "192.168.1.1/24"),
				testMakeAddr(if0, "192.168.1.55/32"),
				testMakeAddr(if0, "192.168.2.1/24"),
				testMakeAddr(if0, "192.168.2.55/32"),
			},
			netifs: []NetIf{if0},
			primary: []Addr{
				testMakeAddr(if0, "192.168.0.1/24"),
				testMakeAddr(if0, "192.168.1.1/24"),
				testMakeAddr(if0, "192.168.2.1/24"),
			},
		},
	}

	for _, test := range tests {
		snap := newSnapshotFromAddrs(test.addrs)

		// Obtain output slices
		outAddrs := snap.Addrs()
		outNetifs := snap.Interfaces()
		outPrimary := snap.PrimaryAddrs()

		// Compare input vs output
		badAddrs := !testAddrsEqual(test.addrs, outAddrs)
		badNetifs := !testNetifsEqual(test.netifs, outNetifs)
		badPrimary := !testAddrsEqual(test.primary, outPrimary)

		if badAddrs || badNetifs || badPrimary {
			t.Errorf("for input: %s", test.addrs)
			if badAddrs {
				t.Errorf("unexpected addresses: %s", outAddrs)
			}

			if badNetifs {
				t.Errorf("interfaces:\n"+
					"expected: %s\n"+
					"present:  %s",
					test.netifs, outNetifs)
			}

			if badPrimary {
				t.Errorf("primary addresses:\n"+
					"expected: %s\n"+
					"present:  %s",
					test.primary, outPrimary)
			}
		}
	}
}

// TestSnapshotSync tests snapshot.Sync
func TestSnapshotSync(t *testing.T) {
	type testData struct {
		addrs []Addr // Addresses at the next step
	}

	lo := NetIf{0, "lo"}
	if1 := NetIf{1, "if1"}

	tests := []testData{
		{
			[]Addr{
				testMakeAddr(lo, "127.0.0.1/24"),
			},
		},
		{
			[]Addr{
				testMakeAddr(lo, "127.0.0.1/24"),
				testMakeAddr(if1, "192.168.0.1/24"),
			},
		},
		{
			[]Addr{
				testMakeAddr(lo, "127.0.0.1/24"),
				testMakeAddr(if1, "192.168.0.1/24"),
				testMakeAddr(if1, "192.168.0.2/32"),
			},
		},
	}

	var prev snapshot
	for _, test := range tests {
		next := newSnapshotFromAddrs(test.addrs)
		events := prev.Sync(next)
		//t.Errorf("%s", events)
		err := testVerifyEvents(prev, events)
		if err != nil {
			t.Errorf("%s", err)
			return
		}

		prev = next
	}
}

// testVerifyEvents verifies "received" events against expected
func testVerifyEvents(snap snapshot, events []Event) error {

	interfaces := make(map[NetIf]struct{})
	for _, ifnet := range snap.Interfaces() {
		interfaces[ifnet] = struct{}{}
	}

	addrs := make(map[Addr]struct{})
	for _, addr := range snap.Addrs() {
		addrs[addr] = struct{}{}
	}

	primary := make(map[Addr]struct{})
	for _, addr := range snap.PrimaryAddrs() {
		primary[addr] = struct{}{}
	}

	hasInterface := func(nif NetIf) bool {
		_, found := interfaces[nif]
		return found
	}

	hasAddr := func(addr Addr) bool {
		_, found := addrs[addr]
		return found
	}

	hasPrimary := func(addr Addr) bool {
		_, found := primary[addr]
		return found
	}

	for _, evnt := range events {
		makeErr := func(msg string) error {
			return fmt.Errorf("%s: %s", evnt, msg)
		}

		switch evnt := evnt.(type) {
		case EventAddInterface:
			nif := evnt.Interface
			if hasInterface(nif) {
				return makeErr("interface already in set")
			}
			interfaces[nif] = struct{}{}

		case EventDelInterface:
			nif := evnt.Interface
			if !hasInterface(nif) {
				return makeErr("interface not in set")
			}
			delete(interfaces, nif)

		case EventAddAddress:
			addr := evnt.Addr
			nif := addr.Interface()

			if !hasInterface(nif) {
				return makeErr("interface not in set")
			}

			if hasAddr(addr) {
				return makeErr("address already in set")
			}

			addrs[addr] = struct{}{}

		case EventDelAddress:
			addr := evnt.Addr
			nif := addr.Interface()

			if !hasInterface(nif) {
				return makeErr("interface not in set")
			}

			if !hasAddr(addr) {
				return makeErr("address not in set")
			}

			delete(addrs, addr)

		case EventAddPrimaryAddress:
			addr := evnt.Addr
			nif := addr.Interface()

			if !hasInterface(nif) {
				return makeErr("interface not in set")
			}

			if !hasAddr(addr) {
				return makeErr("address not in set")
			}

			if hasPrimary(addr) {
				return makeErr("address already primary")
			}

			primary[addr] = struct{}{}

		case EventDelPrimaryAddress:
			addr := evnt.Addr
			nif := addr.Interface()

			if !hasInterface(nif) {
				return makeErr("interface not in set")
			}

			if !hasAddr(addr) {
				return makeErr("address not in set")
			}

			if !hasPrimary(addr) {
				return makeErr("address is not primary")
			}

			delete(primary, addr)
		}
	}

	return nil
}
