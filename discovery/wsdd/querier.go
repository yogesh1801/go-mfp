// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSDD Querier

package wsdd

import (
	"context"
	"net/netip"
	"sync"

	"github.com/alexpevzner/mfp/discovery/netstate"
	"github.com/alexpevzner/mfp/log"
	"github.com/alexpevzner/mfp/wsd"
)

// querier is responsible for transmission of WSDD queries
type querier struct {
	ctx    context.Context    // Logging context
	netmon *netstate.Notifier // Network state monitor
	mconn4 *mconn             // For IP4 multicasts reception
	mconn6 *mconn             // For IP6 multicasts reception
	links  *links             // Per-local address links
	hosts  *hosts             // Hosts table

	// querier.procNetmon closing synchronization
	ctxNetmon    context.Context    // Cancelable context for procNetmon
	cancelNetmon context.CancelFunc // Its cancellation function
	doneNetmon   sync.WaitGroup     // Wait for procNetmon termination

	// querier.procMconn closing synchronization
	doneMconn sync.WaitGroup // Wait for procMconn termination
}

// newQuerier creates a new querier
func newQuerier(ctx context.Context) (*querier, error) {
	// Create multicast sockets
	mconn4, err := newMconn(wsddMulticastIP4)
	if err != nil {
		return nil, err
	}

	mconn6, err := newMconn(wsddMulticastIP6)
	if err != nil {
		mconn4.Close()
		return nil, err
	}

	// Create querier structure
	q := &querier{
		ctx:    ctx,
		netmon: netstate.NewNotifier(),
		mconn4: mconn4,
		mconn6: mconn6,
	}

	q.hosts = newHosts(ctx, q)
	q.links = newLinks(ctx, q)

	return q, nil
}

// Start starts querier operations.
func (q *querier) Start() {
	// Start q.procNetmon
	q.ctxNetmon, q.cancelNetmon = context.WithCancel(q.ctx)
	q.doneNetmon.Add(1)
	go q.procNetmon()

	// Start q.procMconn, one per connection
	q.doneMconn.Add(2)
	go q.procMconn(q.mconn4)
	go q.procMconn(q.mconn6)
}

// Close closes the querier
func (q *querier) Close() {
	// Stop procNetmon
	q.cancelNetmon()
	q.doneNetmon.Wait()

	// Stop multicasts reception
	q.mconn4.Close()
	q.mconn6.Close()
	q.doneMconn.Wait()

	// Close all links
	q.links.Close()

	// Close hosts
	q.hosts.Close()
}

// Input handles received UDP messages.
func (q *querier) Input(data []byte, from, to netip.AddrPort, ifidx int) {
	// Silently drop looped packets
	if q.links.IsLocalPort(from) {
		return
	}

	// Decode the message
	log.Debug(q.ctx, "%d bytes received from %s%%%d",
		len(data), from, ifidx)

	msg, err := wsd.DecodeMsg(data)
	if err != nil {
		log.Warning(q.ctx, "%s", err)
		return
	}

	// Dispatch the message
	log.Debug(q.ctx, "%s message received", msg.Header.Action)

	switch msg.Header.Action {
	case wsd.ActHello, wsd.ActBye, wsd.ActProbeMatches,
		wsd.ActResolveMatches:
		q.hosts.InputFromUDP(msg)
	}
}

// netmonproc processes netstate.Notifier events.
// It runs on its own goroutine.
func (q *querier) procNetmon() {
	defer q.doneNetmon.Done()

	for {
		evnt, err := q.netmon.Get(q.ctxNetmon)
		if err != nil {
			return
		}

		log.Debug(q.ctx, "%s", evnt)

		switch evnt := evnt.(type) {
		case netstate.EventAddPrimaryAddress:
			q.links.Add(evnt.Addr)
		case netstate.EventDelPrimaryAddress:
			q.links.Del(evnt.Addr)
		}
	}
}

// procMconn receives UDP multicast messages from the multicast conection.
func (q *querier) procMconn(mc *mconn) {
	defer q.doneMconn.Done()

	for {
		var buf [65536]byte
		n, from, cmsg, err := mc.RecvFrom(buf[:])

		if mc.IsClosed() {
			return
		}

		if err != nil {
			log.Error(q.ctx, "UDP recv: %s", err)
			continue
		}

		q.Input(buf[:n], from, mc.LocalAddrPort(), cmsg.IfIndex)
	}
}
