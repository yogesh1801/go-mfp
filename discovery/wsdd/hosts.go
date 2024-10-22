// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Table of discovered hosts

package wsdd

import (
	"sync"

	"github.com/alexpevzner/mfp/wsd"
)

// hosts contains a table of discovered hosts.
//
// Hosts, in the WSD terms, more or less corresponds to devices on
// a network. Each device contains some "hosted" services (print
// service, scan service etc).
//
// Please note, WSD uses 3-level addressing architecture:
//  1. Each device has a stable device "address", persistent
//     across reboots. Formally this is "anyURI", but usually,
//     uri:uuid:...
//  2. Each "hosted service" (i.e., printer, scanner, ...) has
//     one or more "transport addresses" (HTTP URLs). The transport
//     addresses are suitable to query device metadata. Printer
//     and scanner XAddrs are often the same, but this is not
//     necessarily.
//  3. Querying device metadata via XAddrs return one or more
//     service endpoints, suitable for printing or scanning,
//     depending on a service type.
type hosts struct {
	table map[wsd.AnyURI]*host // Table of hosts
	lock  sync.Mutex           // hosts.table lock
}

// newHosts creates a new table of hosts
func newHosts() *hosts {
	ht := &hosts{
		table: make(map[wsd.AnyURI]*host),
	}

	return ht
}

// Close closes the host table and cancels all ongoing discovery activity,
// like fetching host's metadata
func (ht *hosts) Close() {
}

// host represents a discovered host (a device). Each host may
// be a home of oner or more hosted services.
type host struct {
	Address     wsd.AnyURI // Host "address" (stable identifier)
	XAddrsScan  []string   // Scanner transport addresses
	XAddrsPrint []string   // Printer transport addresses
}
