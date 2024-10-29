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
	ctx        context.Context      // Cancelable context
	cancel     context.CancelFunc   // Its cancel function
	q          *querier             // Parent querier
	table      map[wsd.AnyURI]*host // Table of hosts
	lock       sync.Mutex           // hosts.table lock
	inputQueue chan wsd.Msg         // Messages from UDP
	done       sync.WaitGroup       // Wait for hosts.Close
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
		table:      make(map[wsd.AnyURI]*host),
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
			switch body := msg.Body.(type) {
			case wsd.AnnouncesBody:
				ht.handleAnnounces(body)
			case wsd.Bye:
				ht.handleBye(body)
			}
		}
	}
}

// handleBye handles received [wsd.Bye] message.
func (ht *hosts) handleBye(body wsd.Bye) {
}

// handleAnnounce is the common handler for WSD announce messages
// (i.e., Hello, ProbeMatch and ResolveMatch).
func (ht *hosts) handleAnnounces(body wsd.AnnouncesBody) {
	action := body.Action()
	anns := body.Announces()

	l := log.Begin(ht.ctx)
	for _, ann := range anns {
		l.Debug("%s: Types %s:", action, ann.Types)
		if len(ann.XAddrs) != 0 {
			l.Debug("Xaddrs:")
			for _, xaddr := range ann.XAddrs {
				l.Debug("  %s", xaddr)
			}
		}
	}
	l.Commit()
}

// getHost returns a host by addr. If host is not known yet,
// it can be created on demand.
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
