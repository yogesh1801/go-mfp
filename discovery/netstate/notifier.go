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
	snapLast snapshot      // Last network state known to the client
	queue    eventqueue    // Queue of not yet delivered events
	errSeq   int64         // Sequence number of next error
	lockChan chan struct{} // Channel-based lock
}

// NewNotifier creates a new Notifier.
//
// This is safe to use Notifier with multiple goroutines.
func NewNotifier() *Notifier {
	return &Notifier{
		lockChan: make(chan struct{}, 1),
	}
}

// Get waits until and returns a next [Event].
// The only case when it returns an error is the context cancellation.
func (not *Notifier) Get(ctx context.Context) (Event, error) {
	// Acquire the lock
	if err := not.lock(ctx); err != nil {
		return nil, err
	}

	defer not.unlock()

	// Check for queued events
	if evnt := not.queue.pull(); evnt != nil {
		return evnt, nil
	}

	// Wait for an event or context cancellation
	mon := getMonitor()
	for {
		snapNext, waitchan := mon.get()

		evnt, errSeq := mon.getError(not.errSeq)
		if evnt != nil {
			mon.errSeq = errSeq
			return evnt, nil
		}

		events := not.snapLast.sync(snapNext)
		if len(events) != 0 {
			not.snapLast = snapNext
			not.queue.push(events[1:]...)
			return events[0], nil
		}

		select {
		case <-waitchan:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

// lock acquires the lock.
//
// If lock is busy, it waits until lock is available or Context
// is canceled.
func (not *Notifier) lock(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case not.lockChan <- struct{}{}:
		return nil
	}
}

// unlock releases the lock, previously acquired by Notifier.lock.
func (not *Notifier) unlock() {
	<-not.lockChan
}
