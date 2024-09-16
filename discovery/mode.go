// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package discovery

// Mode represents discovery mode.
//
// It mostly affects discovery system behavior when searching
// for network devices using some kind of multicast-based device
// discovery protocol (i.e., DNS-SD, WSD and so on).
type Mode int

// Mode values
const (
	// Normally, discovery system warms up cache after initialization
	// or refresh command, and then keeps cache up to date.
	//
	// In this mode, if cache is not yet warmed up, discovery system
	// will wait until it happens. Otherwise, it will return cached
	// data immediately.
	ModeNormal = iota

	// Due to the nature of most discovery protocols, when a new device
	// joins the network, the information describing the device arrives
	// in parts, and these parts do not necessarily arrive at the same
	// time.
	//
	// For example, the IPv6 address of a device may be discovered
	// significantly later than its IPv4 address.
	//
	// As a result, newly discovered (or changed) devices may remain in
	// an "incomplete" state for some time.
	//
	// In the ModeNormal mode, the discovery system will not wait for
	// these incomplete devices to stabilize and will simply return
	// the previous stable state in the output. In the ModeWaitIncomplete
	// mode, if incomplete devices exist in the cache, the discovery
	// system will wait for a period, allowing them the opportunity
	// to stabilize.
	ModeWaitIncomplete

	// ModeSnapshot returns content of discovery cache immediately
	// and doesn't wait for cache warm-up.
	ModeSnapshot
)
