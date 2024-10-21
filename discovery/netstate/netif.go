// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network interface

package netstate

import (
	"fmt"
	"net"
	"strings"
)

// NetIf represents a network interface.
//
// Unline [net.Interface], NetIf is comparable (it supports == and
// can be used as a map index) and immutable.
type NetIf struct {
	index int
	name  string
	flags NetIfFlags
}

// NetIfFlags contains network interface flags
type NetIfFlags int

// NetIfFlags bits:
const (
	NetIfBroadcast NetIfFlags = 1 << iota
	NetIfLoopback
	NetIfMulticast
)

// All returns true if all flags, specified by mask, are set.
func (flags NetIfFlags) All(mask NetIfFlags) bool {
	return flags&mask == mask
}

// Any returns true if any of the flags, specified by mask, are set.
func (flags NetIfFlags) Any(mask NetIfFlags) bool {
	return flags&mask != 0
}

// String returns string representation of the NetIfFlags, for debugging.
func (flags NetIfFlags) String() string {
	s := make([]string, 0, 3)

	if flags&NetIfBroadcast != 0 {
		s = append(s, "broadcast")
	}
	if flags&NetIfLoopback != 0 {
		s = append(s, "loopback")
	}
	if flags&NetIfMulticast != 0 {
		s = append(s, "multicast")
	}

	return strings.Join(s, ",")
}

// MakeNetIf makes [NetIf] from index and name.
func MakeNetIf(index int, name string, flags NetIfFlags) NetIf {
	return NetIf{index, name, flags}
}

// NetIfFromInterface makes [NetIf] from [net.Interface]
func NetIfFromInterface(ifi net.Interface) NetIf {
	var flags NetIfFlags
	if ifi.Flags&net.FlagBroadcast != 0 {
		flags |= NetIfBroadcast
	}
	if ifi.Flags&net.FlagLoopback != 0 {
		flags |= NetIfLoopback
	}
	if ifi.Flags&net.FlagMulticast != 0 {
		flags |= NetIfMulticast
	}

	return NetIf{
		index: ifi.Index,
		name:  ifi.Name,
		flags: flags,
	}
}

// Index returns interface index
func (nif NetIf) Index() int {
	return nif.index
}

// Name returns interface name
func (nif NetIf) Name() string {
	return nif.name
}

// Flags returns interface flags
func (nif NetIf) Flags() NetIfFlags {
	return nif.flags
}

// String returns string representation of the interface,
// for debugging purposes.
func (nif NetIf) String() string {
	return fmt.Sprintf("%s(%s)", nif.name, nif.flags.String())
}

// Less reports whether nif sorts before nif2.
// If compares interfaces first by index, then by name.
func (nif NetIf) Less(nif2 NetIf) bool {
	return nif.index < nif2.index ||
		nif.index == nif2.index && nif.name < nif2.name
}

// AsInterface returns [net.Interface] that corresponds to [NetIf].
//
// This function may fail, for example, because actual interface has
// disappear from the network stack.
func (nif NetIf) AsInterface() (*net.Interface, error) {
	ifi, err := net.InterfaceByName(nif.name)
	if err != nil {
		ifi, err = net.InterfaceByIndex(nif.index)
	}
	return ifi, err
}
