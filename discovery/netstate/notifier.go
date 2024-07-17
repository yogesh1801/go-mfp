// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Event notifier

package netstate

import "context"

// Notifier provides network state change events notifications.
type Notifier struct {
	laststate netstate // Last network state known to the client
	errSeq    int64    // Sequence number of next error
}

// NewNotifier creates a new Notifier.
func NewNotifier() *Notifier {
	return &Notifier{}
}

// Get waits until and returns a next [Event].
// The only case when it returns an error is the context cancellation.
func (not *Notifier) Get(ctx context.Context) (Event, error) {
	// Quick check for a pending Context error
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	// Wait for an event or context cancellation
	mon := gewMonitor()
	for {
		laststate, waitchan := mon.get()

		evnt, errSeq := mon.getError(not.errSeq)
		if evnt != nil {
			mon.errSeq = errSeq
			return evnt, nil
		}

		evnt = not.laststate.sync(laststate)
		if evnt != nil {
			return evnt, nil
		}

		select {
		case <-waitchan:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
