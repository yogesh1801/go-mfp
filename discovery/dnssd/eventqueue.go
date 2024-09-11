// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Event queue

package dnssd

import "github.com/alexpevzner/mfp/discovery"

// eventqueue is the queue of discovery.Event
type eventqueue struct {
	events []discovery.Event // Generated events
}

// newEventqueue creates a new event queue
func newEventqueue() *eventqueue {
	return &eventqueue{
		events: make([]discovery.Event, 0, 16),
	}
}

// Push adds event to the queue
func (q *eventqueue) Push(evnt discovery.Event) {
	q.events = append(q.events, evnt)
}
