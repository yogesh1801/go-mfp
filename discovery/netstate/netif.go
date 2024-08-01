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
)

// NetIf represents a network interface.
//
// Unline [net.Interface], NetIf is comparable (it supports == and
// can be used as a map index) and immutable.
type NetIf struct {
	index int
	name  string
}

// MakeNetIf makes [NetIf] from index and name.
func MakeNetIf(index int, name string) NetIf {
	return NetIf{index, name}
}

// NetIfFromInterface makes [NetIf] from [net.Interface]
func NetIfFromInterface(ifi net.Interface) NetIf {
	return NetIf{
		index: ifi.Index,
		name:  ifi.Name,
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

// String returns string representation of the interface,
// for debugging purposes.
func (nif NetIf) String() string {
	return fmt.Sprintf("%s(#%d)", nif.name, nif.index)
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
