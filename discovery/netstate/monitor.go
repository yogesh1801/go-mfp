// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network state monitor -- the common part

package netstate

import "sync"

var (
	// monitorInstance shared between all users.
	monitorInstance monitor

	// monitorInitOnce ensures monitorInstance will be
	// initialized only once
	monitorInitOnce sync.Once
)

// monitor keeps track on a current network state and provides
// notifications when something changes.
//
// monitor defined as an interface, so it can be hooked for testing
type monitor interface {
	// Get returns last known network state and channel to wait for updates.
	//
	// The returned channel will be closed by monitor when state changes
	// or new error occurs.
	Get() (snapshot, <-chan struct{})

	// GetError returns the latest error, if its sequence number
	// is greater that supplied by the caller (i.e., caller has
	// not seen this error yet). The returned error is wrapped into
	// the EventError structure.
	//
	// If there is no new error, it returns nil.
	//
	// Additionally it returns a sequence number for the next call.
	// The first call should use zero sequence number.
	GetError(seq int64) (Event, int64)
}

// getMonitor returns a network event monitor.
// Monitor is a singleton, shared between all Notifiers.
// If monitor is not exist yet, it will be created on demand.
func getMonitor() monitor {
	monitorInitOnce.Do(func() {
		monitorInstance = newMonitor()
	})
	return monitorInstance
}
