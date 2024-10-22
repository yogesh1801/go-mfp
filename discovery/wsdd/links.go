// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// UDP data links

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

// links contains a table of the per-local address links
type links struct {
	ctx   context.Context      // Logging context
	q     *querier             // Parent querier
	table map[netip.Addr]*link // Per-local address links
	lock  sync.Mutex           // links.table lock
	ports *ports               // Set of Local ports
}

// newLinks creates a new links table
func newLinks(ctx context.Context, q *querier) *links {
	return &links{
		ctx:   ctx,
		q:     q,
		table: make(map[netip.Addr]*link),
		ports: newPorts(),
	}
}

// Close closes links table and all links it owns.
func (lt *links) Close() {
	lt.ports.Clear()

	lt.lock.Lock()
	defer lt.lock.Unlock()

	for addr, l := range lt.table {
		l.Close()
		delete(lt.table, addr)
	}
}

// Add adds a local address
func (lt *links) Add(addr netstate.Addr) {
	// Ignore non-multicast links
	flags := addr.Interface().Flags()
	if !flags.All(netstate.NetIfMulticast) {
		return
	}

	// Add link
	l := newLink(lt, addr)

	lt.lock.Lock()
	lt.table[addr.Addr()] = l
	lt.lock.Unlock()
}

// Del deletes local address
func (lt *links) Del(addr netstate.Addr) {
	// Ignore non-multicast links
	flags := addr.Interface().Flags()
	if !flags.All(netstate.NetIfMulticast) {
		return
	}

	// Del link
	lt.lock.Lock()
	l := lt.table[addr.Addr()]
	delete(lt.table, addr.Addr())
	lt.lock.Unlock()

	l.Close()
}

// link is a per-local address link. It implements sending
// of the UDP multicast packets from that address and reception
// of responses.
//
// It also automatically sends Probe requests.
type link struct {
	parent     *links         // Parent links table
	addr       netstate.Addr  // Local address
	probeMsg   []byte         // Probe message
	probeSched *sched         // Probe scheduler
	dest       netip.AddrPort // Destination (multicast) address
	conn       *uconn         // Connection for sending UDP multicasts
	doneProber sync.WaitGroup // Wait for procProber termination
	doneReader sync.WaitGroup // Wait for procReader termination
}

// newLink creates a new link
func newLink(lt *links, addr netstate.Addr) *link {
	l := &link{
		addr:       addr,
		parent:     lt,
		probeSched: newSched(false),
	}

	if addr.Is4() {
		l.dest = wsddMulticastIP4
	} else {
		l.dest = wsddMulticastIP6
	}

	l.doneProber.Add(1)
	go l.procProber()

	return l
}

// close closes the link
func (l *link) Close() {
	// Kill link.procProber
	l.probeSched.Close()
	l.doneProber.Wait()

	// Kill l.procReader, if it is started
	if l.conn != nil {
		l.parent.ports.Del(l.conn.LocalAddrPort())
		l.conn.Close()
		l.doneReader.Wait()
	}
}

// procProber creates UDP connection on demand and sends probes.
// It runs on its own goroutine and paced by the sa.probeSched scheduler.
func (l *link) procProber() {
	defer l.doneProber.Done()

	for {
		evnt := <-l.probeSched.Chan()
		switch evnt {
		case schedClosed:
			return

		case schedNewMessage:
			// Open connection on demand
			if l.conn == nil {
				var err error
				l.conn, err = newUconn(l.addr, 0)
				if err != nil {
					log.Debug(l.parent.ctx, "%s", err)
				}

				if l.conn != nil {
					l.parent.ports.Add(
						l.conn.LocalAddrPort())
					l.doneReader.Add(1)
					go l.procReader()
				}
			}

			// Update l.probeMsg
			l.updateProbeMsg()

		case schedSend:
			if l.conn != nil {
				l.conn.WriteToUDPAddrPort(l.probeMsg, l.dest)
				log.Debug(l.parent.ctx,
					"%s message sent to %s%%%s",
					wsd.ActProbe, l.dest,
					l.addr.Interface().Name())
			}
		}
	}
}

// procReader runs on its own goroutine and receives messages from the l.conn.
func (l *link) procReader() {
	defer l.doneReader.Done()

	ifidx := l.addr.Interface().Index()
	to := l.conn.LocalAddrPort()

	for {
		// Receive next packet
		var buf [65536]byte
		n, from, err := l.conn.RecvFrom(buf[:])

		if l.conn.IsClosed() {
			return
		}

		if err != nil {
			log.Error(l.parent.ctx, "UDP recv: %s", err)
			continue
		}

		// Silently drop looped packets
		if l.parent.ports.Contains(from) {
			continue
		}

		// Pass it to the querier's Input routine
		l.parent.q.Input(buf[:n], from, to, ifidx)
	}
}

// updateProbeMsg updates l.probeMsg
func (l *link) updateProbeMsg() {
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
	l.probeMsg = msg.Encode()
}
