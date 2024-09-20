// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Utility functions for time.Time

package discovery

import "time"

// timeLatest returns latest of two times
func timeLatest(t1, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t1
	}
	return t2
}

// timeEarliest returns earliest of two times
func timeEarliest(t1, t2 time.Time) time.Time {
	if t1.Before(t2) {
		return t1
	}
	return t2
}
