// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Functions for management collections of IP addresses

package discovery

import (
	"net/netip"
	"net/url"
	"sort"
	"strings"
)

// addrsFromEndpoints extracts addresses from endpoints
func addrsFromEndpoints(endpoints []string) (addrs []netip.Addr) {
	for _, endpoint := range endpoints {
		u, err := url.Parse(endpoint)
		if err != nil {
			continue
		}

		switch u.Scheme {
		case "ipp", "ipps", "lpd", "socket", "http", "https":
		// Only these schemes use IP addresses
		default:
			continue
		}

		// Extract host, drop port
		host := u.Host

		colonAt := strings.LastIndex(host, ":")
		bracketAt := strings.LastIndex(host, "]")
		if colonAt >= 0 && colonAt > bracketAt {
			host = host[:colonAt]
		}

		// Drop square brackets around IP6 address
		if strings.HasPrefix(host, "[") &&
			strings.HasSuffix(host, "]") {
			host = host[1 : len(host)-1]
		}

		addr, err := netip.ParseAddr(host)
		if err != nil {
			continue
		}

		addr = addr.Unmap() // Just in case

		// Save the address
		addrs = append(addrs, addr)
	}

	// Sort extracted addresses
	sort.Slice(addrs, func(i, j int) bool {
		return addrs[i].Less(addrs[j])
	})

	return
}

// addrsContain reports if collection of addresses contains
// the specified address.
//
// addresses assumed to be sorted.
func addrsContain(addrs []netip.Addr, addr netip.Addr) bool {
	_, found := sort.Find(len(addrs), func(i int) int {
		return addr.Compare(addrs[i])
	})
	return found
}

// addrsAdd adds address into the collection of addresses.
//
// It assumes sorted addresses on input and returns updated and
// properly sorted collection on output. It works by modifying
// the input collection in-place, so be aware.
//
// Additionally, it returns a bool flag that indicates that
// address was actually added.
func addrsAdd(addrs []netip.Addr, addr netip.Addr) ([]netip.Addr, bool) {
	i, found := sort.Find(len(addrs), func(i int) int {
		return addr.Compare(addrs[i])
	})

	if found {
		return addrs, false
	}

	addrs = append(addrs, netip.Addr{})
	copy(addrs[i+1:], addrs[i:])
	addrs[i] = addr

	return addrs, true
}

// addrsDel deletes address from the collection of addresses.
//
// It assumes sorted addresses on input and returns updated and
// properly sorted collection on output. It works by modifying
// the input collection in-place, so be aware.
//
// Additionally, it returns a bool flag that indicates that
// address was actually found in the collection and deleted.
func addrsDel(addrs []netip.Addr, addr netip.Addr) ([]netip.Addr, bool) {
	i, found := sort.Find(len(addrs), func(i int) int {
		return addr.Compare(addrs[i])
	})

	if !found {
		return addrs, false
	}

	copy(addrs[i:], addrs[i+1:])
	addrs = addrs[:len(addrs)-1]

	return addrs, true
}

// addrsMerge merges two collections of addresses.
//
// Input collections must be sorted and returned collection
// is sorted as well.
func addrsMerge(addrs1, addrs2 []netip.Addr) []netip.Addr {
	out := make([]netip.Addr, 0, len(addrs1)+len(addrs2))

	for len(addrs1) > 0 && len(addrs2) > 0 {
		cmp := addrs1[0].Compare(addrs2[0])
		switch {
		case cmp < 0:
			out = append(out, addrs1[0])
			addrs1 = addrs1[1:]
		case cmp > 0:
			out = append(out, addrs2[0])
			addrs2 = addrs2[1:]
		default:
			out = append(out, addrs1[0])
			addrs1 = addrs1[1:]
			addrs2 = addrs2[1:]
		}
	}

	out = append(out, addrs1...)
	out = append(out, addrs2...)

	return out
}

// addrsOverlap reports if two collections of addresses overlap,
// i.e., contains some addresses in common.
//
// Input collections must be sorted.
func addrsOverlap(addrs1, addrs2 []netip.Addr) bool {
	for len(addrs1) > 0 && len(addrs2) > 0 {
		cmp := addrs1[0].Compare(addrs2[0])
		switch {
		case cmp < 0:
			addrs1 = addrs1[1:]
		case cmp > 0:
			addrs2 = addrs2[1:]
		default:
			return true
		}
	}

	return false
}
