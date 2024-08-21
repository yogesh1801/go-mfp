// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Lookup options

package dnssd

// LookupFlags define DNS-SD lookup mode.
type LookupFlags int

// LookupFlags bits.
const (
	// Use classical (unicast) DNS
	LookupClassical LookupFlags = 1 << iota

	// Use Multicast DNS (mDNS)
	LookupMulticast

	// Use both methods. This is default, if none bits are set
	LookupBoths = LookupClassical | LookupMulticast
)

const (
	// This is the default lookup domain for DNS-SD.
	LookupDomain = "local"
)
