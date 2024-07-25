// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network state monitor test

package netstate

import "testing"

// TestGetMonitor performs testing of getMonitor() function
func TestGetMonitor(t *testing.T) {
	var newMonitorCallCount int

	newMonitorHook := func() monitor {
		newMonitorCallCount++
		return newTestMonitor()
	}

	testGetMonitorReset()
	saveNewMonitor := hookNewMonitor
	hookNewMonitor = newMonitorHook
	defer func() { hookNewMonitor = saveNewMonitor }()

	for i := 0; i < 10; i++ {
		mon := getMonitor()
		if newMonitorCallCount != 1 {
			t.Errorf("newTestMonitor call count: expected 1, present %d)",
				newMonitorCallCount)
		}

		if mon != testMonitorInstanse {
			t.Errorf("getMonitor() returned invalid result")
		}
	}
}
