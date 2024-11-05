// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Cancellable timer

package wsdd

import "time"

// timer is the cancellable timer. It can be used to pause program
// execution for the specified time, with the ability to cancel the
// pause and with the indication, if pause was expired or canceled.
type timer struct {
	c chan struct{}
}

// newTimer creates a new timer
func newTimer() timer {
	return timer{
		c: make(chan struct{}),
	}
}

// Sleep delays the calling goroutine execution until time expires
// or timer is canceled.
//
// It returns true if pause has expired or false if timer was canceled.
func (tmr timer) Sleep(d time.Duration) bool {
	// Already canceled?
	select {
	case <-tmr.c:
		return false
	default:
	}

	// Perform cancellable pause
	t := time.NewTimer(d)
	select {
	case <-t.C:
		return true
	case <-tmr.c:
		t.Stop()
		return false
	}
}

// Cancel cancels the timer. All pending and future timer.Sleep calls
// will return immediately with the false value.
//
// Once canceled, timer remains in the canceled state forever.
func (tmr timer) Cancel() {
	close(tmr.c)
}
