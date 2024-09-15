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

	// ModeSnapshot returns content of discovery cache immediately
	// and doesn't wait for cache warm-up.
	ModeSnapshot
)
