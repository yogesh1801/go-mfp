// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// mDNS resolver

package dnssd

import (
	"context"
	"errors"
	"net/netip"
	"sync/atomic"
	"time"

	"github.com/alexpevzner/go-avahi"
)

// ResolverWaitTime is the [Resolver] timeout in case resolving
// doesn't finish earlier.
const ResolverWaitTime = 2 * time.Second

// Resolver performs DNS lookups, using mDNS
type Resolver struct {
	avahiClnt *avahi.Client // The avahi.Client
	closed    atomic.Bool   // Resolver is closed
}

// NewResolver returns a new resolver.
// Resolver owns some resources and must be closed when not needed anymore.
func NewResolver() (*Resolver, error) {
	clnt, err := avahi.NewClient(avahi.ClientLoopbackWorkarounds)
	if err != nil {
		return nil, err
	}

	return &Resolver{avahiClnt: clnt}, nil
}

// Close closes the resolver.
func (res *Resolver) Close() {
	res.closed.Store(true)
	res.avahiClnt.Close()
}

// LookupHost looks up the given host using the mDNS resolver.
func (res *Resolver) LookupHost(ctx context.Context,
	ifidx int, host string) (addrs []netip.Addr, err error) {

	// Create resolvers, separate for IP4 and IP6
	avahiResolverIP4, err := avahi.NewHostNameResolver(
		res.avahiClnt,
		avahi.IfIndex(ifidx),
		avahi.ProtocolIP4,
		host,
		avahi.ProtocolIP4,
		0)

	if err != nil {
		return nil, err
	}

	defer avahiResolverIP4.Close()

	avahiResolverIP6, err := avahi.NewHostNameResolver(
		res.avahiClnt,
		avahi.IfIndex(ifidx),
		avahi.ProtocolIP6,
		host,
		avahi.ProtocolIP6,
		0)

	if err != nil {
		return nil, err
	}

	defer avahiResolverIP6.Close()

	// Setup timer
	tm := time.NewTimer(2 * time.Second)
	defer tm.Stop()

	// Poll for events
	doneIP4 := false
	doneIP6 := false

	for !(doneIP4 && doneIP6) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()

		case <-tm.C:
			doneIP4, doneIP6 = true, true

		case <-res.avahiClnt.Chan():
			// Just drain the channel

		case evnt := <-avahiResolverIP4.Chan():
			// Note, we ignore avahi.ResolverFailure event,
			// because it may mean just that resolving didn't
			// finish in time.
			if evnt.Event == avahi.ResolverFound {
				doneIP4 = true
				addrs = append(addrs, evnt.Addr)
			}

		case evnt := <-avahiResolverIP6.Chan():
			// Note, we ignore avahi.ResolverFailure event,
			// because it may mean just that resolving didn't
			// finish in time.
			if evnt.Event == avahi.ResolverFound {
				doneIP6 = true
				addrs = append(addrs, evnt.Addr)
			}
		}

		if res.closed.Load() {
			// To avoid infinite looping in a case resolver
			// is closed, as we ignore avahi.ResolverFailure
			// events.
			return nil, errors.New("Resolver is closed")
		}
	}

	return
}
