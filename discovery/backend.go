// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery backend

package discovery

// Backend scans/monitors its search [Realm] and reports discovered
// devices using channel of events.
type Backend interface {
	// Chan returns an event channel. Backend uses this channel
	// to reports events of the following types:
	//   [*EventAddPrinter]
	//   [*EventDelPrinter]
	//   [*EventAddScanner]
	//   [*EventDelScanner]
	//   [*EventAddEndpoint]
	//   [*EventDelEndpoint]
	Chan() <-chan Event

	// Close closes the backend and releases resources it holds.
	// It also closes backent's event channel, effectively unblocking
	// pending readers.
	Close()
}
