// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Table of discovered units

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

// units contains a table of discovered units.
//
// Please note, WSD uses 3-level addressing architecture:
//  1. Each device has a stable device "address", persistent
//     across reboots. Formally this is "xsd:anyURI", but usually,
//     uri:uuid:...
//  2. Each "hosted service" (i.e., printer, scanner, ...) has
//     one or more "transport addresses" (HTTP URLs). The transport
//     addresses are suitable to query device metadata. Printer
//     and scanner XAddrs are often the same, but this is not
//     necessarily.
//  3. Querying device metadata via XAddrs return one or more
//     service endpoints, suitable for printing or scanning,
//     depending on a service type.
//
// Every unit represents a discovered unit (the "hosted service"
// in the WSD terms), localized to the particular network interface.
// If the same device is visible over multiple network interfaces,
// the discovery system when merge them together.
type units struct {
	back   *backend                   // Parent backend
	ctx    context.Context            // Cancelable context
	cancel context.CancelFunc         // Its cancel function
	q      *querier                   // Parent querier
	table  map[discovery.UnitID]*unit // Discovered units
	lock   sync.Mutex                 // units.table lock
	done   sync.WaitGroup             // Wait for hosts.Close
}

// newUnits creates a new table of units
func newUnits(back *backend, q *querier) *units {
	// Create cancelable context
	ctx, cancel := context.WithCancel(back.ctx)

	// Create units structure
	ut := &units{
		back:   back,
		ctx:    ctx,
		cancel: cancel,
		q:      q,
		table:  make(map[discovery.UnitID]*unit),
	}

	return ut
}

// Close closes the unit table and cancels all ongoing discovery activity,
// like fetching unit's metadata
func (ut *units) Close() {
	ut.cancel()
	ut.done.Wait()
}

// InputFromUSB handles WSD message, received from UDP
func (ut *units) InputFromUDP(msg wsd.Msg) {
	ut.lock.Lock()
	defer ut.lock.Unlock()

	switch msg.Body.(type) {
	case wsd.AnnouncesBody:
		ut.handleAnnounces(msg)
	case wsd.Bye:
		ut.handleBye(msg)
	}
}

// handleBye handles received [wsd.Bye] message.
func (ut *units) handleBye(msg wsd.Msg) {
}

// handleAnnounce is the common handler for WSD announce messages
// (i.e., Hello, ProbeMatch and ResolveMatch).
func (ut *units) handleAnnounces(msg wsd.Msg) {
	body := msg.Body.(wsd.AnnouncesBody)
	action := body.Action()
	anns := body.Announces()

	// Parse and dispatch XAddrs. Log the event.
	for _, ann := range anns {
		addr := ann.EndpointReference.Address
		ver := ann.MetadataVersion

		log.Debug(ut.ctx, "%s received:", action)
		log.Debug(ut.ctx, "  Address         %s:", addr)
		log.Debug(ut.ctx, "  Types           %s:", ann.Types)
		log.Debug(ut.ctx, "  MetadataVersion %d:", ver)

		printUnitID := ut.makeUnitID(msg.IfIdx,
			discovery.ServicePrinter, addr)
		scanUnitID := ut.makeUnitID(msg.IfIdx,
			discovery.ServiceScanner, addr)

		if len(ann.XAddrs) != 0 {
			log.Debug(ut.ctx, "  Xaddrs:")

			// Parse and collect XAddrs
			var xaddrs []*url.URL
			for _, s := range ann.XAddrs {
				if u := urlParse(s); u != nil {
					log.Debug(ut.ctx, "    %s", s)
					xaddrs = append(xaddrs, u)
				} else {
					log.Warning(ut.ctx, "    %s (invalid)", s)
				}
			}

			// Dispatch XAddrs
			if len(xaddrs) != 0 {
				if ann.Types&wsd.TypePrinter != 0 {
					un := ut.getUnit(printUnitID, true)
					un.handleXaddrs(xaddrs, ver)
				}

				if ann.Types&wsd.TypeScanner != 0 {
					un := ut.getUnit(scanUnitID, true)
					un.handleXaddrs(xaddrs, ver)
				}
			}
		}
	}
}

// makeUnitID creates a discovery.UnitID for the discovered
// service
func (ut *units) makeUnitID(ifidx int, svctype discovery.ServiceType,
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
func (ut *units) getUnit(id discovery.UnitID, create bool) *unit {
	un := ut.table[id]
	if un == nil && create {
		un = newUnit(id)
	}
	return un
}

// unit represents a discovered unit (a device)
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
