// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package discovery

// SearchScope specified a search scope.
type SearchScope int

// SearchScope values:
const (
	SearchDNSSD       SearchScope = iota // DNS-SD
	SearchWSDiscovery                    // Microsoft WS-Discovery
)
