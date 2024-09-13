// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery backend

package discovery

// Backend scans/monitors its search [Realm] and reports discovered
// devices by sending series of [Event] into the provided [Eventqueue].
type Backend interface {
	// Name returns backend name.
	Name() string

	// Start starts Backend operations.
	Start(*Eventqueue)

	// Close closes the Backend and releases resources it holds.
	Close()
}
