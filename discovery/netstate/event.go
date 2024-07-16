// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Network notifier events

package netstate

import (
	"fmt"
	"net"
)

// Event is the common interface of all events.
type Event interface {
	// String returns string representation of the Event,
	// for logging.
	String() string

	// Dummy unexported method to disallow definition of Event
	// types outside of the package
	event()
}

var (
	_ Event = EventAddInterface{}
	_ Event = EventDelInterface{}
	_ Event = EventAddAddress{}
	_ Event = EventDelAddress{}
)

// EventAddInterface is fired when new interface added to the system.
type EventAddInterface struct {
	Interface net.Interface // Added interface
}

// String returns string representation of [EventAddInterface], for logging.
func (e EventAddInterface) String() string {
	return fmt.Sprintf("add-interface: interface=%q, index=%d",
		e.Interface.Name, e.Interface.Index)
}

// event implements an [Event] interface
func (EventAddInterface) event() {}

// EventDelInterface is fired when network interface is removed from the system.
type EventDelInterface struct {
	Interface net.Interface // Deleted interface
}

// String returns string representation of [EventDelInterface], for logging.
func (e EventDelInterface) String() string {
	return fmt.Sprintf("del-interface: interface=%q, index=%d",
		e.Interface.Name, e.Interface.Index)
}

// event implements an Event interface
func (EventDelInterface) event() {}

// EventAddAddress is fired when new IP address added to some interface.
//
// This is guaranteed that [EventAddInterface] will be delivered before
// corresponding EventAddAddress events.
type EventAddAddress struct {
	Addr Addr // Added address
}

// String returns string representation of EventAddInterface, for logging.
func (e EventAddAddress) String() string {
	return fmt.Sprintf("add-address: interface=%q, index=%d, addr=%s",
		e.Addr.Interface.Name, e.Addr.Interface.Index, e.Addr.String())
}

// event implements an Event interface
func (EventAddAddress) event() {}

// EventDelAddress is fired when IP address is removed.
//
// This is guaranteed that all appropriate [EventDeldAddress] events
// will be delivered before [EventDelInterface] of the appropriate
// interface.
type EventDelAddress struct {
	Addr Addr // Deleted address
}

// String returns string representation of [EventDelAddress], for logging.
func (e EventDelAddress) String() string {
	return fmt.Sprintf("del-address: interface=%q, index=%d, addr=%s",
		e.Addr.Interface.Name, e.Addr.Interface.Index, e.Addr.String())
}

// event implements an Event interface
func (EventDelAddress) event() {}

// EventError is fired when some error occurs.
//
// Errors are not fatal and delivered barely for logging.
// If series of errors occurs, this is not guaranteed that
// all will be delivered.
type EventError struct {
	Err error // The error
}

// String returns string representation of [EventError], for logging.
func (e EventError) String() string {
	return fmt.Sprintf("error: %s", e.Err)
}

// event implements an Event interface
func (EventError) event() {}
