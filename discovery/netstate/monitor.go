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
	monitorInstance *monitor

	// monitorInitOnce ensures monitorInstance will be
	// initialized only once
	monitorInitOnce sync.Once
)

// getMonitor returns a network event monitor.
// Monitor is a singleton, shared between all Notifiers.
// If monitor is not exist yet, it will be created on demand.
func getMonitor() *monitor {
	monitorInitOnce.Do(func() {
		monitorInstance = newMonitor()
	})
	return monitorInstance
}
