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
//
// There are 5 types of events:
//
//   - [EventAddInterface] and [EventDelInterface] generated when
//     network interface is added or deleted.
//   - [EventAddAddress] and [EventDelAddress] generated when IP
//     address is added or deleted.
//   - [EventError] is generated when some error occurs. Errors
//     are not fatal and EventError intended only for logging.
//     This is not guaranteed that all errors will be reported this
//     way, but this mechanism attempts to be as informative as possible.
//
// When address is added, events will come in the following order:
//  1. [EventAddInterface]
//  2. [EventAddAddress], that used previously added interface.
//
// When address is deleted, events will come in reverse order.
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
	_ Event = EventError{}
)

// EventAddInterface is fired when new interface added to the system.
//
// See [Event] description for details.
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
//
// See [Event] description for details.
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
// See [Event] description for details.
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
// See [Event] description for details.
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
// See [Event] description for details.
type EventError struct {
	Err error // The error
}

// String returns string representation of [EventError], for logging.
func (e EventError) String() string {
	return fmt.Sprintf("error: %s", e.Err)
}

// event implements an Event interface
func (EventError) event() {}
