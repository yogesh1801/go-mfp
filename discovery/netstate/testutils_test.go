// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common utilities for tests

package netstate

import (
	"fmt"
	"net"
	"slices"
	"sort"
	"sync"
	"sync/atomic"
)

// testNetIfMaker makes network interfaces for testing
type testNetIfMaker struct {
	index int32
}

// newTestNetIfMaker makes a new testNetIfMaker
func testNewNetIfMaker() *testNetIfMaker {
	return &testNetIfMaker{}
}

// testNetIfMaker makes a new interface
func (netifmaker *testNetIfMaker) new() NetIf {
	idx := atomic.AddInt32(&netifmaker.index, 1)
	nif := NetIf{
		index: int(idx),
		name:  fmt.Sprintf("net%d", idx),
	}
	return nif
}

// testAddr creates a new address for testing
// It returns address as *Addr
func testMakeAddr(nif NetIf, cidr string) *Addr {
	return &Addr{
		IPNet:     *testMakeIPNet(cidr),
		Interface: nif,
		Primary:   true,
	}
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
func testMonitorUpdateAddrs(addrs []*Addr) {
	addrs = slices.Clone(addrs)
	sort.Slice(addrs, func(i, j int) bool {
		return addrs[i].Less(addrs[j])
	})

	testMonitorInstanse.lock.Lock()
	defer testMonitorInstanse.lock.Unlock()

	testMonitorInstanse.snapshot = snapshot{addrs}
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
