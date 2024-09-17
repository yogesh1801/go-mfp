// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network Address Family

package discovery

import "fmt"

// AddressFamily represents network address family
type AddressFamily int

// AddressFamily constants:
const (
	AddressFamilyNA  AddressFamily = iota // Not applicable
	AddressFamilyIP4                      // IPv4
	AddressFamilyIP6                      // IPp6
)

// String returns AddressFamily name, for debugging
func (af AddressFamily) String() string {
	switch af {
	case AddressFamilyNA:
		return "n/a"
	case AddressFamilyIP4:
		return "ip4"
	case AddressFamilyIP6:
		return "ip6"
	}

	return fmt.Sprintf("unknown (%d)", int(af))
}
