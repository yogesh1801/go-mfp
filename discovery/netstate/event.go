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
)

// Event is the common interface of all events.
//
// There are 7 types of events:
//
//   - [EventAddInterface] and [EventDelInterface] generated when
//     network interface is added or deleted.
//   - [EventAddAddress] and [EventDelAddress] generated when IP
//     address is added or deleted.
//   - [EventAddPrimaryAddress] and [EventDelPrimaryAddress] generated
//     when primary address is added or deleted or when existent address
//     changes its status. See [Addr] description for definition of
//     the primary address.
//   - [EventError] is generated when some error occurs. Errors
//     are not fatal and EventError intended only for logging.
//     This is not guaranteed that all errors will be reported this
//     way, but this mechanism attempts to be as informative as possible.
//
// When address is added, events will come in the following order:
//  1. [EventAddInterface]
//  2. [EventAddAddress], that used previously added interface.
//  3. [EventAddPrimaryAddress], that used previously added address.
//
// When address is deleted, events will come in reverse order.
//
// Please notice, that primary addresses will be reported twice,
// using EventAddAddress/EventDelAddress events and using
// EventAddPrimaryAddress/EventDelPrimaryAddress events.
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
	_ Event = EventAddPrimaryAddress{}
	_ Event = EventDelPrimaryAddress{}
	_ Event = EventError{}
)

// EventAddInterface is fired when new interface added to the system.
//
// See [Event] description for details.
type EventAddInterface struct {
	Interface NetIf // Added interface
}

// String returns string representation of [EventAddInterface], for logging.
func (e EventAddInterface) String() string {
	return fmt.Sprintf("add-interface: %s(#%d)",
		e.Interface.Name(), e.Interface.Index())
}

// event implements an [Event] interface
func (EventAddInterface) event() {}

// EventDelInterface is fired when network interface is removed from the system.
//
// See [Event] description for details.
type EventDelInterface struct {
	Interface NetIf // Deleted interface
}

// String returns string representation of [EventDelInterface], for logging.
func (e EventDelInterface) String() string {
	return fmt.Sprintf("del-interface: %s(#%d)",
		e.Interface.Name(), e.Interface.Index())
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
	return fmt.Sprintf("add-address: %s%%%s(#%d)",
		e.Addr.String(),
		e.Addr.Interface().Name(), e.Addr.Interface().Index())
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
	return fmt.Sprintf("del-address: %s%%%s(#%d)",
		e.Addr.String(),
		e.Addr.Interface().Name(), e.Addr.Interface().Index())
}

// event implements an Event interface
func (EventDelAddress) event() {}

// EventAddPrimaryAddress is fired when primary IP address added to some
// interface or when existing address changes its status to Primary.
//
// See [Event] description for details.
type EventAddPrimaryAddress struct {
	Addr Addr // Added address
}

// String returns string representation of [EventAddPrimaryAddress],
// for logging.
func (e EventAddPrimaryAddress) String() string {
	return fmt.Sprintf("add-primary: %s%%%s(#%d)",
		e.Addr.String(),
		e.Addr.Interface().Name(), e.Addr.Interface().Index())
}

// event implements an Event interface
func (EventAddPrimaryAddress) event() {}

// EventDelPrimaryAddress is fired when Primary IP address is removed
// or looses its status of the Primary address.
//
// See [Event] description for details.
type EventDelPrimaryAddress struct {
	Addr Addr // Deleted address
}

// String returns string representation of [EventDelPrimaryAddress],
// for logging.
func (e EventDelPrimaryAddress) String() string {
	return fmt.Sprintf("del-primary: %s%%%s(#%d)",
		e.Addr.String(),
		e.Addr.Interface().Name(), e.Addr.Interface().Index())
}

// event implements an Event interface
func (EventDelPrimaryAddress) event() {}

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
