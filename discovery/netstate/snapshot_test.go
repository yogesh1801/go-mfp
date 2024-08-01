// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Tests of network state snapshot

package netstate

import (
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
