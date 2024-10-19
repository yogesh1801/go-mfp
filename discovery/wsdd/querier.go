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
	ctx       context.Context             // Logging context
	netmon    *netstate.Notifier          // Network state monitor
	mconn4    *mconn                      // For IP4 multicasts reception
	mconn6    *mconn                      // For IP6 multicasts reception
	links     map[netip.Addr]*querierLink // Per-local address contexts
	linksLock sync.Mutex                  // querier.links lock

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
	dest       netip.AddrPort // Destination (multicast) address
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
		links:  make(map[netip.Addr]*querierLink),
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
	for addr, ql := range q.links {
		ql.Close()
		delete(q.links, addr)
	}
}

// input handles received UDP messages.
func (q *querier) input(data []byte, from, to netip.AddrPort, ifidx int) {
	if q.hasLocalAddr(from.Addr()) {
		//log.Debug(q.ctx, "skipped message from self (%s%%%d)",
		//	from, ifidx)
		return
	}

	log.Debug(q.ctx, "%d bytes received from %s%%%d",
		len(data), from, ifidx)

	msg, err := wsd.DecodeMsg(data)
	if err != nil {
		log.Warning(q.ctx, "%s", err)
		return
	}

	log.Debug(q.ctx, "%s message received", msg.Header.Action)
}

// addLocalAddr adds local address
func (q *querier) addLocalAddr(addr netstate.Addr) {
	// Ignore non-multicast links
	flags := addr.Interface().Flags()
	if !flags.All(netstate.NetIfMulticast) {
		return
	}

	// Add link
	ql := q.newQuerierLink(addr)

	q.linksLock.Lock()
	q.links[addr.Addr()] = ql
	q.linksLock.Unlock()
}

// delLocalAddr deletes local address
func (q *querier) delLocalAddr(addr netstate.Addr) {
	// Ignore non-multicast links
	flags := addr.Interface().Flags()
	if !flags.All(netstate.NetIfMulticast) {
		return
	}

	// Del link
	q.linksLock.Lock()
	ql := q.links[addr.Addr()]
	delete(q.links, addr.Addr())
	q.linksLock.Unlock()

	ql.Close()
}

// hasLocalAddr reports if address is known as local
func (q *querier) hasLocalAddr(addr netip.Addr) bool {
	q.linksLock.Lock()
	_, found := q.links[addr]
	q.linksLock.Unlock()
	return found
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
			q.addLocalAddr(evnt.Addr)
		case netstate.EventDelPrimaryAddress:
			q.delLocalAddr(evnt.Addr)
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

		q.input(buf[:n], from, mc.LocalAddrPort(), cmsg.IfIndex)
	}
}

// newQuerierLink returns a new querierLink
func (q *querier) newQuerierLink(addr netstate.Addr) *querierLink {
	ql := &querierLink{
		addr:       addr,
		parent:     q,
		probeSched: newSched(false),
	}

	if addr.Is4() {
		ql.dest = wsddMulticastIP4
	} else {
		ql.dest = wsddMulticastIP6
	}

	ql.doneProber.Add(1)
	go ql.procProber()

	return ql
}

// Close closes querierLink
func (ql *querierLink) Close() {
	// Kill querierLink.procProber
	ql.probeSched.Close()
	ql.doneProber.Wait()

	// Kil ql.procReader, if it is started
	if ql.conn != nil {
		ql.conn.Close()
		ql.doneReader.Wait()
	}
}

// procProber creates UDP connection on demand and sends probes.
// It runs on its own goroutine and paced by the sa.probeSched scheduler.
func (ql *querierLink) procProber() {
	defer ql.doneProber.Done()

	for {
		evnt := <-ql.probeSched.Chan()
		switch evnt {
		case schedClosed:
			return

		case schedNewMessage:
			// Open connection on demand
			if ql.conn == nil {
				var err error
				ql.conn, err = newUconn(ql.addr, 0)
				if err != nil {
					log.Debug(ql.parent.ctx, "%s", err)
				}

				if ql.conn != nil {
					ql.doneReader.Add(1)
					go ql.procReader()
				}
			}

			// Update ql.probeMsg
			ql.updateProbeMsg()

		case schedSend:
			if ql.conn != nil {
				ql.conn.WriteToUDPAddrPort(ql.probeMsg, ql.dest)
				log.Debug(ql.parent.ctx,
					"%s message sent to %s%%%d",
					wsd.ActProbe, ql.dest,
					ql.addr.Interface().Index())
			}
		}
	}
}

// procReader runs on its own goroutine and receives messages from the ql.conn.
func (ql *querierLink) procReader() {
	defer ql.doneReader.Done()

	ifidx := ql.addr.Interface().Index()
	to := ql.conn.LocalAddrPort()

	for {
		var buf [65536]byte
		n, from, err := ql.conn.RecvFrom(buf[:])

		if ql.conn.IsClosed() {
			return
		}

		if err != nil {
			log.Error(ql.parent.ctx, "UDP recv: %s", err)
			continue
		}

		ql.parent.input(buf[:n], from, to, ifidx)
	}
}

// updateProbeMsg updates ql.probeMsg
func (ql *querierLink) updateProbeMsg() {
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
	ql.probeMsg = msg.Encode()
}
