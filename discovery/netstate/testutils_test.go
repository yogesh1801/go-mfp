// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common utilities for tests

package netstate

import (
	"net"
	"slices"
	"sort"
	"sync"
)

// testAddr creates a new address for testing
// It returns address as *Addr
func testMakeAddr(nif NetIf, cidr string) Addr {
	ipnet := *testMakeIPNet(cidr)
	return AddrFromIPNet(ipnet, nif)
}

// testAddr creates a new address for testing
// It returns address as net.IPNet
func testMakeIPNet(cidr string) *net.IPNet {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(err)
	}

	return &net.IPNet{IP: ip, Mask: ipnet.Mask}
}

// testNetifsEqual tells if two slices of network interfaces are equal.
//
// It ignores difference in addresses interfaces and difference
// between empty and nil slices.
func testNetifsEqual(netifs1, netifs2 []NetIf) bool {
	// Handle empty slices
	if len(netifs1) == 0 && len(netifs2) == 0 {
		return true
	}

	// Clone and sort slices
	netifs1 = slices.Clone(netifs1)
	netifs2 = slices.Clone(netifs2)

	sort.Slice(netifs1, func(i, j int) bool {
		return netifs1[i].Less(netifs1[j])
	})

	sort.Slice(netifs2, func(i, j int) bool {
		return netifs2[i].Less(netifs2[j])
	})

	// Compare slices
	return slices.Equal(netifs1, netifs2)
}

// testAddrsEqual tells if two slices of addresses are equal.
//
// It ignores difference in addresses ordering and difference
// between empty and nil slices.
func testAddrsEqual(addrs1, addrs2 []Addr) bool {
	// Handle empty slices
	if len(addrs1) == 0 && len(addrs2) == 0 {
		return true
	}

	// Clone and sort slices
	addrs1 = slices.Clone(addrs1)
	addrs2 = slices.Clone(addrs2)

	sort.Slice(addrs1, func(i, j int) bool {
		return addrs1[i].Less(addrs1[j])
	})

	sort.Slice(addrs2, func(i, j int) bool {
		return addrs2[i].Less(addrs2[j])
	})

	// Compare slices
	return slices.Equal(addrs1, addrs2)
}

// testMonitor implements monitor interface, for testing
type testMonitor struct {
	lock     sync.Mutex
	snapshot snapshot
	waitchan chan struct{}
}

var testMonitorInstanse = &testMonitor{
	waitchan: make(chan struct{}),
}

// newTestMonitor creates a new testMonitor
func newTestMonitor() monitor {
	return testMonitorInstanse
}

// testMonitorUpdateAddrs updates network addresses, exposed by
// the testMonitorInstanse
func testMonitorUpdateAddrs(addrs []Addr) {
	addrs = slices.Clone(addrs)
	sort.Slice(addrs, func(i, j int) bool {
		return addrs[i].Less(addrs[j])
	})

	testMonitorInstanse.lock.Lock()
	defer testMonitorInstanse.lock.Unlock()

	testMonitorInstanse.snapshot = newSnapshotFromAddrs(addrs)
	close(testMonitorInstanse.waitchan)
}

// Get returns last known network state and channel to wait for updates.
//
// The returned channel will be closed by monitor when state changes.
func (mon *testMonitor) Get() (snapshot, <-chan struct{}) {
	testMonitorInstanse.lock.Lock()
	defer testMonitorInstanse.lock.Unlock()

	return mon.snapshot, mon.waitchan
}

// GetError returns the latest error, if its sequence number
func (mon *testMonitor) GetError(seq int64) (Event, int64) {
	return nil, 0
}
