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
func (netifmaker *testNetIfMaker) new() net.Interface {
	idx := atomic.AddInt32(&netifmaker.index, 1)
	ifi := net.Interface{
		Index: int(idx),
		Name:  fmt.Sprintf("net%d", idx),
	}
	return ifi
}

// testAddr creates a new address for testing
// It returns address as *Addr
func testMakeAddr(netif net.Interface, cidr string) *Addr {
	return &Addr{
		IPNet:     *testMakeIPNet(cidr),
		Interface: netif,
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
