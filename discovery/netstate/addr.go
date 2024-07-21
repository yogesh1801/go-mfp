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

// Addr represents a single IP address with mask, assigned to the network
// interface.
//
// Interface may have multiple addresses which may belong to the same
// or different IP networks. Belonging addresses to IP networks divides
// addresses into groups. One and only one address of each group will
// be marked as Primary.
//
// In another words, if all interface addresses will belong to the different
// IP networks, all of them will be marked as Primary. If some of the
// interface addresses belong to the same IP network, only one of these
// addresses will be chosen as Primary.
//
// Addresses considered belonging to the same IP network, if ranges, taking
// address Mask into account, overlap. [Addr.Overlaps] can be used to
// test any two addresses for overlapping. Strictly speaking, ranges
// covered by two overlapping addresses either equal, if masks are the
// same, or nest, if mask of the "inner" address.
//
// For overlapping addresses, [Addr.Narrower] reports whether of addresses
// are narrower.
type Addr struct {
	net.IPNet               // IP address with mask
	Interface net.Interface // Interface that owns the address
	Primary   bool          // It's a primary address
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
		return addr.Narrower(addr2)
	}
}

// Overlaps reports whether two addresses overlap.
//
// Overlapped addressed are addresses for which all following is true:
//   - they belong to the same network interface
//   - they belong to the same address family, both either IP4 or IP6
//   - their address range, taking Mask into account, overlap
func (addr *Addr) Overlaps(addr2 *Addr) bool {
	var answer bool
	if addr.SameInterface(addr2) {
		answer = addr.Contains(addr2.IP) || addr2.Contains(addr.IP)
	}
	return answer
}

// Narrower reports whether addr is narrower that addr2.
//
// It means the following:
//   - addr and addr2 overlap (see [Addr.Overlap] for definition).
//   - mask of addr is narrower (contains more leading ones and less
//     trailing zeroes) that mask of addr2
func (addr *Addr) Narrower(addr2 *Addr) bool {
	var answer bool
	if addr.Overlaps(addr2) {
		ones, _ := addr.Mask.Size()
		ones2, _ := addr2.Mask.Size()
		return ones > ones2
	}
	return answer
}

// Is4 tells is [Addr] is IP4 address.
func (addr *Addr) Is4() bool {
	return addr.IP.To4() != nil
}
