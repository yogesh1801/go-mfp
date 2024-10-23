// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Table of discovered hosts

package wsdd

import (
	"context"
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
	ctx   context.Context      // Logging context
	q     *querier             // Parent querier
	table map[wsd.AnyURI]*host // Table of hosts
	lock  sync.Mutex           // hosts.table lock
}

// newHosts creates a new table of hosts
func newHosts(ctx context.Context, q *querier) *hosts {
	ht := &hosts{
		ctx:   ctx,
		q:     q,
		table: make(map[wsd.AnyURI]*host),
	}

	return ht
}

// Close closes the host table and cancels all ongoing discovery activity,
// like fetching host's metadata
func (ht *hosts) Close() {
}

// HandleProbeMatches handles received [wsd.Hello] message.
func (ht *hosts) HandleHello(body wsd.Hello) {
	ht.handleAnnounce(body.EndpointReference.Address,
		body.Types, body.XAddrs, body.MetadataVersion)
}

// HandleBye handles received [wsd.Bye] message.
func (ht *hosts) HandleBye(body wsd.Bye) {
}

// HandleProbeMatches handles received [wsd.ProbeMatches] message.
func (ht *hosts) HandleProbeMatches(body wsd.ProbeMatches) {
	for _, match := range body.ProbeMatch {
		ht.handleAnnounce(match.EndpointReference.Address,
			match.Types, match.XAddrs, match.MetadataVersion)
	}
}

// HandleResolveMatches handles received [wsd.ResolveMatches] message.
func (ht *hosts) HandleResolveMatches(body wsd.ResolveMatches) {
	for _, match := range body.ResolveMatch {
		ht.handleAnnounce(match.EndpointReference.Address,
			match.Types, match.XAddrs, match.MetadataVersion)
	}
}

// handleAnnounce is the common handler for WSD announce messages
// (i.e., Hello, ProbeMatch and ResolveMatch)
func (ht *hosts) handleAnnounce(addr wsd.AnyURI,
	types wsd.Types, xaddrs wsd.XAddrs, ver uint64) {
}

// getHost
func (ht *hosts) getHost(addr wsd.AnyURI, create bool) *host {
	h := ht.table[addr]
	if h == nil && create {
		h = &host{Address: addr}
	}
	return h
}

// host represents a discovered host (a device). Each host may
// be a home of oner or more hosted services.
type host struct {
	Address     wsd.AnyURI // Host "address" (stable identifier)
	XAddrsScan  []string   // Scanner transport addresses
	XAddrsPrint []string   // Printer transport addresses
}
