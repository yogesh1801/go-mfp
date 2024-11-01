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
	"net/url"
	"sync"

	"github.com/alexpevzner/mfp/discovery"
	"github.com/alexpevzner/mfp/internal/generic"
	"github.com/alexpevzner/mfp/internal/zone"
	"github.com/alexpevzner/mfp/log"
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
	ctx        context.Context            // Cancelable context
	cancel     context.CancelFunc         // Its cancel function
	q          *querier                   // Parent querier
	table      map[discovery.UnitID]*unit // Discovered units
	lock       sync.Mutex                 // hosts.table lock
	inputQueue chan wsd.Msg               // Messages from UDP
	done       sync.WaitGroup             // Wait for hosts.Close
}

// newHosts creates a new table of hosts
func newHosts(ctx context.Context, q *querier) *hosts {
	// Create cancelable context
	ctx, cancel := context.WithCancel(ctx)

	// Create hosts structure
	ht := &hosts{
		ctx:        ctx,
		cancel:     cancel,
		q:          q,
		table:      make(map[discovery.UnitID]*unit),
		inputQueue: make(chan wsd.Msg, wsddUDPInputQueueSize),
	}

	ht.done.Add(1)
	go ht.inputProc()

	return ht
}

// Close closes the host table and cancels all ongoing discovery activity,
// like fetching host's metadata
func (ht *hosts) Close() {
	ht.cancel()
	ht.done.Wait()
}

// InputFromUSB handles WSD message, received from UDP
func (ht *hosts) InputFromUDP(msg wsd.Msg) {
	select {
	// We don't worry too much in a (very unlike) case of the
	// queue overflow. Just drop the message. This is UDP,
	// after all.
	case ht.inputQueue <- msg:
	default:
	}
}

// inputProc runs on its own goroutine and handles all received
// messages
func (ht *hosts) inputProc() {
	defer ht.done.Done()

	for ht.ctx.Err() == nil {
		select {
		case <-ht.ctx.Done():
		case msg := <-ht.inputQueue:
			switch msg.Body.(type) {
			case wsd.AnnouncesBody:
				ht.handleAnnounces(msg)
			case wsd.Bye:
				ht.handleBye(msg)
			}
		}
	}
}

// handleBye handles received [wsd.Bye] message.
func (ht *hosts) handleBye(msg wsd.Msg) {
}

// handleAnnounce is the common handler for WSD announce messages
// (i.e., Hello, ProbeMatch and ResolveMatch).
func (ht *hosts) handleAnnounces(msg wsd.Msg) {
	body := msg.Body.(wsd.AnnouncesBody)
	action := body.Action()
	anns := body.Announces()

	// Parse and dispatch XAddrs. Log the event.
	l := log.Begin(ht.ctx)
	for _, ann := range anns {
		addr := ann.EndpointReference.Address
		ver := ann.MetadataVersion

		l.Debug("%s received:", action)
		l.Debug("  Address         %s:", addr)
		l.Debug("  Types           %s:", ann.Types)
		l.Debug("  MetadataVersion %d:", ver)

		printUnitID := ht.makeUnitID(msg.IfIdx,
			discovery.ServicePrinter, addr)
		scanUnitID := ht.makeUnitID(msg.IfIdx,
			discovery.ServiceScanner, addr)

		if len(ann.XAddrs) != 0 {
			l.Debug("  Xaddrs:")

			// Parse and collect XAddrs
			var xaddrs []*url.URL
			for _, s := range ann.XAddrs {
				if u := urlParse(s); u != nil {
					l.Debug("    %s", s)
					xaddrs = append(xaddrs, u)
				} else {
					l.Warning("    %s (invalid)", s)
				}
			}

			// Dispatch XAddrs
			if len(xaddrs) != 0 {
				if ann.Types&wsd.TypePrinter != 0 {
					un := ht.getUnit(printUnitID, true)
					un.handleXaddrs(xaddrs, ver)
				}

				if ann.Types&wsd.TypeScanner != 0 {
					un := ht.getUnit(scanUnitID, true)
					un.handleXaddrs(xaddrs, ver)
				}
			}
		}
	}
	l.Commit()
}

// makeUnitID creates a discovery.UnitID for the discovered
// service
func (ht *hosts) makeUnitID(ifidx int, svctype discovery.ServiceType,
	addr wsd.AnyURI) discovery.UnitID {
	return discovery.UnitID{
		UUID:     addr.UUID(),
		Realm:    discovery.RealmWSD,
		Zone:     zone.Name(ifidx),
		SvcType:  svctype,
		SvcProto: discovery.ServiceWSD,
	}
}

// getUnit returns an unit by id. If unit is not known yet,
// it can be created on demand.
func (ht *hosts) getUnit(id discovery.UnitID, create bool) *unit {
	un := ht.table[id]
	if un == nil && create {
		un = newUnit(id)
	}
	return un
}

// host represents a discovered unit (a device)
type unit struct {
	id         discovery.UnitID           // Unit ID
	xaddrsSeen *generic.LockedSet[string] // Known XAddrs
}

// newUnut creates an unit structure
func newUnit(id discovery.UnitID) *unit {
	return &unit{
		id:         id,
		xaddrsSeen: generic.NewLockedSet[string](),
	}
}

// handleXaddrs handles newly discovered XAddrs
func (un *unit) handleXaddrs(xaddrs []*url.URL, ver uint64) {
}
