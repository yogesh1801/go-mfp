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
	"sync/atomic"

	"github.com/alexpevzner/mfp/discovery/netstate"
	"github.com/alexpevzner/mfp/log"
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/wsd"
)

// querier is responsible for transmission of WSDD queries
type querier struct {
	ctx   context.Context                // Logging context
	addrs map[netstate.Addr]*querierAddr // Per-local address contexts
}

// querierAddr is the per-local-address structure
type querierAddr struct {
	parent     *querier       // Parent querier
	addr       netstate.Addr  // Local address
	probeMsg   []byte         // Probe message
	probeSched *sched         // Probe scheduler
	dest       netip.AddrPort // Destination address
	conn       *uconn         // Connection for sending UDP multicasts
	closing    atomic.Bool    // Close in progress
	doneProber chan struct{}  // Closed when procProber is done
	doneReader chan struct{}  // Closed when procReader is done
}

// newQuerier creates a new querier
func newQuerier(ctx context.Context) *querier {
	q := &querier{
		ctx:   ctx,
		addrs: make(map[netstate.Addr]*querierAddr),
	}

	return q
}

// AddAddr adds local address
func (q *querier) AddAddr(netstate.Addr) {
}

// DelAddr deletes local address
func (q *querier) DelAddr(netstate.Addr) {
}

// newQuerierAddr returns a new querierAddr
func (q *querier) newQuerierAddr(addr netstate.Addr) *querierAddr {
	qa := &querierAddr{
		parent:     q,
		probeSched: newSched(false),
		doneProber: make(chan struct{}),
	}

	if addr.Is4() {
		qa.dest = wsddMulticastIP4
	} else {
		qa.dest = wsddMulticastIP6
	}

	go qa.procProber()

	return qa
}

// Close closes querierAddr
func (qa *querierAddr) Close() {
	// Mark querierAddr as being closing
	qa.closing.Store(true)

	// Kill querierAddr.procProber
	qa.probeSched.Close()
	<-qa.doneProber

	// Kil qa.procReader, if it is started
	if qa.conn != nil {
		qa.conn.Close()
		<-qa.doneReader
	}
}

// procProber creates UDP connection on demand and sends probes.
// It runs on its own goroutine and paced by the sa.probeSched scheduler.
func (qa *querierAddr) procProber() {
	defer close(qa.doneProber)

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
					qa.doneReader = make(chan struct{})
					go qa.procReader()
				}
			}

			// Update qa.probeMsg
			qa.updateProbeMsg()

		case schedSend:
			if qa.conn != nil {
				qa.conn.WriteToUDPAddrPort(qa.probeMsg, qa.dest)
			}
		}
	}
}

// procReader runs on its own goroutine receives messages from the qa.conn.
func (qa *querierAddr) procReader() {
	defer close(qa.doneReader)

	for {
		var buf [65536]byte
		n, from, err := qa.conn.RecvFrom(buf[:])

		if qa.closing.Load() {
			return
		}

		if err != nil {
			log.Error(qa.parent.ctx, "UDP recv: %s", err)
			return
		}

		log.Debug(qa.parent.ctx, "%d bytes received from %s",
			n, from)

		data := buf[:n]
		msg, err := wsd.DecodeMsg(data)
		if err != nil {
			log.Warning(qa.parent.ctx, "%s", err)
			continue
		}

		log.Debug(qa.parent.ctx, "%s message received",
			msg.Header.Action)
	}
}

// updateProbeMsg updates qa.probeMsg
func (qa *querierAddr) updateProbeMsg() {
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
