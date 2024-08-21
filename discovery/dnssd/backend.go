// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery Backend

package dnssd

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/alexpevzner/go-avahi"
	"github.com/alexpevzner/mfp/discovery"
	"github.com/alexpevzner/mfp/log"
)

// Parameters:
const (
	// Avahi client may fail to start if avahi-daemon is not
	// running for for the similar reasons.
	//
	// If failed, Avahi client will automatically restart
	// with the following interval between attempts.
	avahiClientRestartInterval = 1 * time.Second
)

// backend is the [discovery.Backend] for DNS-SD discovery
type backend struct {
	ctx         context.Context
	clnt        *avahi.Client
	poller      *avahi.Poller
	chn         chan any
	lookupFlags LookupFlags
	domain      string
	cancel      context.CancelFunc
	done        sync.WaitGroup

	resolvers map[avahiSvcInstanceKey]*avahi.ServiceResolver
}

// NewBackend creates a new [discovery.Backend] for DNS-SD discovery.
func NewBackend(ctx context.Context,
	domain string, flags LookupFlags) (discovery.Backend, error) {

	// Set log prefix
	ctx = log.WithPrefix(ctx, "dnssd")

	// Create Avahi client.
	clnt, err := avahi.NewClient(avahi.ClientLoopbackWorkarounds)
	if err != nil {
		log.Error(ctx, "%s", err)
		return nil, err
	}

	// Create cancelable context
	ctx, cancel := context.WithCancel(ctx)

	back := &backend{
		ctx:         ctx,
		clnt:        clnt,
		poller:      avahi.NewPoller(),
		chn:         make(chan any),
		lookupFlags: flags & LookupBoths,
		domain:      domain,
		cancel:      cancel,
	}

	back.makeMaps()

	// Start event loop goroutine
	back.done.Add(1)
	go back.proc()

	log.Debug(ctx, "backend started")

	return back, nil
}

// makeMaps makes back.resolvers etc maps
func (back *backend) makeMaps() {
	back.resolvers = make(
		map[avahiSvcInstanceKey]*avahi.ServiceResolver,
	)
}

// makeMaps purges back.resolvers etc maps
func (back *backend) purgeMaps() {
	// Until we move to Go 1.21 with its clear() builtin,
	// just re-create all maps.
	back.makeMaps()
}

// Chan returns an event channel.
func (back *backend) Chan() <-chan any {
	return back.chn
}

// Close closes the backend
func (back *backend) Close() {
	back.cancel()
	back.done.Wait()

	close(back.chn)
	back.clnt.Close()
}

// proc runs the backend event loop on its separate goroutine.
func (back *backend) proc() {
	defer back.done.Done()

	for {
		// Start service browsers.
		err := back.startServiceBrowsers()

		// Handle events
		for err == nil {
			var evnt any
			evnt, err = back.poller.Poll(back.ctx)

			switch evnt := evnt.(type) {
			case *avahi.ClientEvent:
				err = back.onClientEvent(evnt)

			case *avahi.ServiceBrowserEvent:
				err = back.onServiceBrowserEvent(evnt)
			}
		}

		// Try to restart Avahi Client until success or
		// back.ctx expiration.
		for err != nil {
			err = back.ctx.Err()
			if err != nil {
				return
			}

			err = back.clientRestart()
		}
	}
}

// clientRestart pauses for avahiClientRestartInterval and then
// restarts the Avahi client.
func (back *backend) clientRestart() error {
	// It effectively closes all Browsers and Resolvers
	back.clnt.Close()
	back.purgeMaps()

	// Pause for avahiClientRestartInterval.
	// Return immediately, if back.ctx has expired.
	timer := time.NewTimer(avahiClientRestartInterval)
	defer timer.Stop()

	select {
	case <-back.ctx.Done():
		return back.ctx.Err()
	case <-timer.C:
	}

	// Try to restart the Client
	clnt, err := avahi.NewClient(avahi.ClientLoopbackWorkarounds)
	if err != nil {
		return err
	}

	back.clnt = clnt
	return nil
}

// startServiceBrowsers starts service browsers for all service
// types mentioned in the svcTypes.
//
// The newly created browsers are added to the back.poller
func (back *backend) startServiceBrowsers() error {
	for _, svctype := range svcTypes {
		browser, err := avahi.NewServiceBrowser(
			back.clnt,
			avahi.IfIndexUnspec,
			avahi.ProtocolUnspec,
			svctype,
			back.domain,
			avahiLookupFlags[back.lookupFlags],
		)

		if err != nil {
			log.Error(back.ctx, "%s")
			return err
		}

		log.Debug(back.ctx, "service browse: %q", svctype)
		back.poller.AddServiceBrowser(browser)
	}

	return nil
}

// onClientEvent handles avahi.ClientEvent.
func (back *backend) onClientEvent(evnt *avahi.ClientEvent) error {
	switch evnt.State {
	case avahi.ClientStateFailure:
		return evnt.Err
	}

	return nil
}

// onServiceBrowserEvent handles avahi.ServiceBrowserEvent
func (back *backend) onServiceBrowserEvent(
	evnt *avahi.ServiceBrowserEvent) error {

	switch evnt.Event {
	case avahi.BrowserNew:
		key := avahiSvcInstanceKeyFromBrowserEvent(evnt)
		_, duplicate := back.resolvers[key]

		if !duplicate {
			log.Debug(back.ctx, "service found %s", key)
		} else {
			log.Debug(back.ctx, "service found %s (duplicate)", key)
			return nil
		}

		resolver, err := avahi.NewServiceResolver(
			back.clnt,
			evnt.IfIdx,
			evnt.Proto,
			evnt.InstanceName,
			evnt.SvcType,
			evnt.Domain,
			avahi.ProtocolUnspec,
			avahiLookupFlags[back.lookupFlags],
		)

		if err != nil {
			return err
		}

		back.poller.AddServiceResolver(resolver)

	case avahi.BrowserRemove:
		key := avahiSvcInstanceKeyFromBrowserEvent(evnt)
		delete(back.resolvers, key)

	case avahi.BrowserFailure:
		return evnt.Err
	}

	return nil
}

// avahiLookupFlags maps LookupFlags to avahi.LookupFlags
var avahiLookupFlags = [...]avahi.LookupFlags{
	0:               avahi.LookupUseMulticast,
	LookupClassical: avahi.LookupUseWideArea,
	LookupMulticast: avahi.LookupUseMulticast,
	LookupBoths:     avahi.LookupUseMulticast,
}

// avahiSvcInstanceKey identifies a particular instance of
// service of the particular type.
type avahiSvcInstanceKey struct {
	IfIdx        avahi.IfIndex  // Network interface index
	Proto        avahi.Protocol // Network protocol
	InstanceName string         // Service instance name
	SvcType      string         // Service type
	Domain       string         // Service domain
}

// String returns string representation of the avahiSvcInstanceKey,
// for debugging.
func (key avahiSvcInstanceKey) String() string {
	nm := key.InstanceName + "." + key.SvcType + "." + key.Domain

	var ifname string

	if ifi, err := net.InterfaceByIndex(int(key.IfIdx)); err == nil {
		ifname = ifi.Name
	} else {
		ifname = fmt.Sprintf("%d", key.IfIdx)
	}

	return fmt.Sprintf("%q (%s,%s)", nm, key.Proto, ifname)
}

// avahiSvcInstanceKeyFromBrowserEvent makes avahiSvcInstanceKey
// from the avahi.ServiceBrowserEvent.
func avahiSvcInstanceKeyFromBrowserEvent(
	evnt *avahi.ServiceBrowserEvent) avahiSvcInstanceKey {

	return avahiSvcInstanceKey{
		IfIdx:        evnt.IfIdx,
		Proto:        evnt.Proto,
		InstanceName: evnt.InstanceName,
		SvcType:      evnt.SvcType,
		Domain:       evnt.Domain,
	}
}
