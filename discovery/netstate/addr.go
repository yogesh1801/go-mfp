// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network Interface addresses

package netstate

import (
	"bytes"
	"net"
)

// Addr represents a single IP address assigned to the network interface.
type Addr struct {
	net.IPNet               // IP address with mask
	Interface net.Interface // Interface that owns the address
}

// Equal reports if two addresses are equal.
func (addr *Addr) Equal(addr2 *Addr) bool {
	return addr.SameInterface(addr2) &&
		bytes.Equal(addr.IPNet.IP, addr2.IPNet.IP) &&
		bytes.Equal(addr.IPNet.Mask, addr2.IPNet.Mask)
}

// SameInterface reports if two addresses belong to the same
// network interface.
//
// Note, we consider two interfaces equal if they have equal
// [net.Interface.Index] and [net.Interface.Name]. Other parts
// of the [net.Interface] considered interface parameters, not
// interface identity.
func (addr *Addr) SameInterface(addr2 *Addr) bool {
	return addr.Interface.Index == addr2.Interface.Index &&
		addr.Interface.Name == addr2.Interface.Name
}

// Less orders [Addr] for sorting.
//
// The sorting order is following:
//
//   - if addresses belongs to different interfaces, they are
//     sorted by [net.Interface.Index], in acceding order
//   - if interface indices are the same, but name differ, addresses
//     are sorted by interface name, in acceding order
//   - otherwise, if addresses belong to the different address
//     families, they are sorted by address family, IPv4 first
//   - otherwise, if IP addresses not the same, they are sorted
//     by IP address, in lexicographical acceding order
//   - otherwise, if masks are different, addresses are sorted by
//     network mask, the narrowed first
//   - otherwise, addresses are equal
func (addr *Addr) Less(addr2 *Addr) bool {
	switch {
	case addr.Interface.Index != addr2.Interface.Index:
		// Sort by net.Interface.Index
		return addr.Interface.Index < addr2.Interface.Index
	case addr.Interface.Name != addr2.Interface.Name:
		// Sort by net.Interface.Name
		return addr.Interface.Name < addr2.Interface.Name
	case addr.Is4() != addr2.Is4():
		// Sort by address family, IP4 first
		return addr.Is4()
	case !addr.IP.Equal(addr2.IP):
		// Sort by IP address, lexicographically
		return bytes.Compare(addr.IP, addr2.IP) < 0
	default:
		// Sort by network mask, the narrowed first.
		// The more network mask contains leading ones,
		// the narrower it is.
		ones, _ := addr.Mask.Size()
		ones2, _ := addr.Mask.Size()
		return ones > ones2
	}
}

// Is4 tells is [Addr] is IP4 address.
func (addr *Addr) Is4() bool {
	return addr.IP.To4() != nil
}
