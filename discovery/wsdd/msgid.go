// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// MessageID generation

package wsdd

import (
	"sync"
	"time"

	"github.com/alexpevzner/mfp/discovery/netstate"
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/wsd"
)

const (
	// msgIDTimeToLive defines how long generated msgIDs are
	// kept in the msgID table.
	msgIDTimeToLive = 5 * time.Second

	// msgID expected table capacity.
	msgIDTableCapacity = 128

	// msgIDGCPeriod defines, how often garbage collection
	// is performed.
	msgIDGCPeriod = msgIDTimeToLive
)

// msgIDGen generates and validates WSD MessageID
//
// Each generated MessageID has a small portion of the attached
// data, the [netstate.Addr], local IP address with interface.
//
// MessageID has a limited time to live. Expired IDs removed
// from the table.
type msgIDGen struct {
	table     map[wsd.AnyURI]msgIDEnt
	lock      sync.Mutex
	closechan chan struct{}
	done      sync.WaitGroup
}

// newMsgIDGen creates a new MessageID generator.
func newMsgIDGen() *msgIDGen {
	gen := &msgIDGen{
		table:     make(map[wsd.AnyURI]msgIDEnt, msgIDTableCapacity),
		closechan: make(chan struct{}),
	}

	gen.done.Add(1)
	go gen.gc()

	return gen
}

// Close closes the MessageID generator.
func (gen *msgIDGen) Close() {
	close(gen.closechan)
	gen.done.Wait()
}

// gc runs on its own goroutine and performs garbage collection
func (gen *msgIDGen) gc() {
	defer gen.done.Done()

	for {
		// Wait for some time
		t := time.NewTimer(msgIDGCPeriod)
		select {
		case <-t.C:
		case <-gen.closechan:
			t.Stop()
			return
		}

		// Perform garbage collection
		now := time.Now()

		gen.lock.Lock()
		for urn, ent := range gen.table {
			if ent.expired(now) {
				delete(gen.table, urn)
			}
		}
		gen.lock.Unlock()
	}
}

// New generates a new MessageID.
//
// MessageID is the UUID URI of the following form:
//
//	urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (gen *msgIDGen) New(addr netstate.Addr) wsd.AnyURI {
	urn := wsd.AnyURI(uuid.Must(uuid.Random()).URN())
	ent := msgIDEnt{
		addr:    addr,
		expires: time.Now().Add(msgIDTimeToLive),
	}

	gen.lock.Lock()
	defer gen.lock.Unlock()

	gen.table[urn] = ent

	return urn
}

// Get validates MessageID URN, and if it is OK and not expired,
// returns local address, associated with it.
func (gen *msgIDGen) Get(urn wsd.AnyURI) (addr netstate.Addr, ok bool) {
	// Lock the table
	gen.lock.Lock()
	defer gen.lock.Unlock()

	// Get msgIDEnt; check for expiration
	ent, ok := gen.table[urn]
	if ok && ent.expired(time.Now()) {
		delete(gen.table, urn)
		ok = false
	}

	// Report the result
	if !ok {
		return netstate.Addr{}, false
	}

	return ent.addr, ok
}

// msgIDEnt is the per-MessageID saved data
type msgIDEnt struct {
	addr    netstate.Addr // Local address
	expires time.Time     // Expiration time
}

// expired reports if msgIDEnt is expired
func (ent msgIDEnt) expired(now time.Time) bool {
	return !now.Before(ent.expires)
}
