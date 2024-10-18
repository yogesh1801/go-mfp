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
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/wsd"
)

// querier is responsible for transmission of WSDD queries
type querier struct {
	ctx    context.Context                // Logging context
	netmon *netstate.Notifier             // Network state monitor
	mconn4 *mconn                         // For IP4 multicasts reception
	mconn6 *mconn                         // For IP6 multicasts reception
	addrs  map[netstate.Addr]*querierLink // Per-local address contexts

	// querier.procNetmon closing synchronization
	ctxNetmon    context.Context    // Cancelable context for procNetmon
	cancelNetmon context.CancelFunc // Its cancellation function
	doneNetmon   sync.WaitGroup     // Wait for procNetmon termination

	// querier.procMconn closing synchronization
	doneMconn sync.WaitGroup // Wait for procMconn termination
}

// querierLink is the per-local-address structure
type querierLink struct {
	parent     *querier       // Parent querier
	addr       netstate.Addr  // Local address
	probeMsg   []byte         // Probe message
	probeSched *sched         // Probe scheduler
	dest       netip.AddrPort // Destination address
	conn       *uconn         // Connection for sending UDP multicasts
	doneProber sync.WaitGroup // Wait for procProber termination
	doneReader sync.WaitGroup // Wait for procReader termination
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
		addrs:  make(map[netstate.Addr]*querierLink),
	}

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

	// Close all querierLink-s
	for addr, qa := range q.addrs {
		qa.Close()
		delete(q.addrs, addr)
	}
}

// input handles received UDP messages.
func (q *querier) input(data []byte, from, to netip.AddrPort, ifidx int) {
	log.Debug(q.ctx, "%d bytes received from %s%%%d",
		len(data), from, ifidx)

	msg, err := wsd.DecodeMsg(data)
	if err != nil {
		log.Warning(q.ctx, "%s", err)
		return
	}

	log.Debug(q.ctx, "%s message received", msg.Header.Action)
}

// AddAddr adds local address
func (q *querier) AddAddr(netstate.Addr) {
}

// DelAddr deletes local address
func (q *querier) DelAddr(netstate.Addr) {
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

		q.input(buf[:n], from, mc.LocalAddrPort(), cmsg.IfIndex)
	}
}

// newQuerierLink returns a new querierLink
func (q *querier) newQuerierLink(addr netstate.Addr) *querierLink {
	qa := &querierLink{
		parent:     q,
		probeSched: newSched(false),
	}

	if addr.Is4() {
		qa.dest = wsddMulticastIP4
	} else {
		qa.dest = wsddMulticastIP6
	}

	qa.doneProber.Add(1)
	go qa.procProber()

	return qa
}

// Close closes querierLink
func (qa *querierLink) Close() {
	// Kill querierLink.procProber
	qa.probeSched.Close()
	qa.doneProber.Wait()

	// Kil qa.procReader, if it is started
	if qa.conn != nil {
		qa.conn.Close()
		qa.doneReader.Wait()
	}
}

// procProber creates UDP connection on demand and sends probes.
// It runs on its own goroutine and paced by the sa.probeSched scheduler.
func (qa *querierLink) procProber() {
	defer qa.doneProber.Done()

	for {
		evnt := <-qa.probeSched.Chan()
		switch evnt {
		case schedClosed:
			return

		case schedNewMessage:
			// Open connection on demand
			if qa.conn == nil {
				qa.conn, _ = newUconn(qa.addr, 0)
				if qa.conn != nil {
					qa.doneReader.Add(1)
					go qa.procReader()
				}
			}

			// Update qa.probeMsg
			qa.updateProbeMsg()

		case schedSend:
			if qa.conn != nil {
				qa.conn.WriteToUDPAddrPort(qa.probeMsg, qa.dest)
				log.Debug(qa.parent.ctx, "%s message sent",
					wsd.ActProbe)
			}
		}
	}
}

// procReader runs on its own goroutine and receives messages from the qa.conn.
func (qa *querierLink) procReader() {
	defer qa.doneReader.Done()

	ifidx := qa.addr.Interface().Index()
	to := qa.conn.LocalAddrPort()

	for {
		var buf [65536]byte
		n, from, err := qa.conn.RecvFrom(buf[:])

		if qa.conn.IsClosed() {
			return
		}

		if err != nil {
			log.Error(qa.parent.ctx, "UDP recv: %s", err)
			continue
		}

		qa.parent.input(buf[:n], from, to, ifidx)
	}
}

// updateProbeMsg updates qa.probeMsg
func (qa *querierLink) updateProbeMsg() {
	msgid := wsd.AnyURI(uuid.Must(uuid.Random()).URN())
	msg := wsd.Msg{
		Header: wsd.Header{
			Action:    wsd.ActProbe,
			MessageID: msgid,
			To:        wsd.ToDiscovery,
		},
		Body: wsd.Probe{
			Types: wsd.TypeDevice,
		},
	}
	qa.probeMsg = msg.Encode()
}
