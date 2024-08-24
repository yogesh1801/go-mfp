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
	"net/netip"
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
	services    map[avahiServiceKey]*avahiService
	hostnames   map[avahiHostnameKey]*avahiHostname
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
	back.services = make(map[avahiServiceKey]*avahiService)
	back.hostnames = make(map[avahiHostnameKey]*avahiHostname)
}

// makeMaps purges back.resolvers etc maps
func (back *backend) purgeMaps() {
	for _, service := range back.services {
		back.delService(service)
	}

	if len(back.hostnames) != 0 {
		panic("internal error")
	}
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
	back.purgeMaps()
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

			case *avahi.ServiceResolverEvent:
				err = back.onServiceResolverEvent(evnt)

			case *avahi.RecordBrowserEvent:
				switch evnt.RType {
				case avahi.DNSTypeTXT:
					err = back.onTxtBrowserEvent(evnt)
				case avahi.DNSTypeA, avahi.DNSTypeAAAA:
					err = back.onAddrBrowserEvent(evnt)
				}
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

		title := fmt.Sprintf("svc-browse: start %q", svctype)

		if err != nil {
			log.Error(back.ctx, "%s: %s", title, err)
			return err
		}

		log.Debug(back.ctx, "%s: OK", title)
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
		key := avahiServiceKeyFromServiceBrowserEvent(evnt)
		title := fmt.Sprintf("svc-browse: found %s", key)

		if !back.hasService(key) {
			log.Debug(back.ctx, "%s", title)
		} else {
			log.Debug(back.ctx, "%s (duplicate)", title)
			return nil
		}

		return back.addService(key)

	case avahi.BrowserRemove:
		key := avahiServiceKeyFromServiceBrowserEvent(evnt)
		title := fmt.Sprintf("svc-browse: removed %s", key)

		service := back.getService(key)
		if service != nil {
			log.Debug(back.ctx, "%s", title)
			back.delService(service)
		} else {
			log.Debug(back.ctx, "%s (not found)", title)
		}

	case avahi.BrowserFailure:
		key := avahiServiceKeyFromServiceBrowserEvent(evnt)
		title := fmt.Sprintf("svc-browse: failed  %s", key)

		log.Error(back.ctx, "%s: %s", title, evnt.Err)
		return evnt.Err
	}

	return nil
}

// onServiceResolverEvent handles avahi.ServiceResolverEvent
func (back *backend) onServiceResolverEvent(
	evnt *avahi.ServiceResolverEvent) error {
	switch evnt.Event {
	case avahi.ResolverFound:
		key := avahiServiceKeyFromResolverEvent(evnt)
		title := fmt.Sprintf("svc-resolve: found %s", key)

		service := back.getService(key)
		if service == nil {
			// It may be out of order avahi.ResolverFound
			// event, received while service already removed,
			// so just log and return.
			log.Debug(back.ctx, "%s (unknown service)", title)
			return nil
		}

		log.Begin(back.ctx).
			Debug("%s:", title).
			Debug("  host: %s", evnt.Hostname).
			Debug("  port: %d", evnt.Port).
			Commit()

		service.port = evnt.Port
		back.addServiceHostname(service, evnt.Hostname)

	case avahi.ResolverFailure:
		key := avahiServiceKeyFromResolverEvent(evnt)
		title := fmt.Sprintf("svc-resolve: failed  %s", key)

		log.Error(back.ctx, "%s: %s", title, evnt.Err)
		return evnt.Err
	}

	return nil
}

// onTxtBrowserEvent handles avahi.RecordBrowserEvent
// for per-service TXT record browser
func (back *backend) onTxtBrowserEvent(evnt *avahi.RecordBrowserEvent) error {
	switch evnt.Event {
	case avahi.BrowserNew:
		key := avahiServiceKeyFromRecordBrowserEvent(evnt)
		title := fmt.Sprintf("txt-browse: found %s", key)
		log.Debug(back.ctx, "%s", title)

	case avahi.BrowserFailure:
		key := avahiServiceKeyFromRecordBrowserEvent(evnt)
		title := fmt.Sprintf("txt-browse: failed %s", key)

		log.Error(back.ctx, "%s: %s", title, evnt.Err)
		return evnt.Err
	}

	return nil
}

// onAddrBrowserEvent handles avahi.RecordBrowserEvent
// for per-hostname A and AAAA record browsers
func (back *backend) onAddrBrowserEvent(
	evnt *avahi.RecordBrowserEvent) error {

	switch evnt.Event {
	case avahi.BrowserNew, avahi.BrowserRemove:
		key := avahiHostnameKeyFromRecordBrowserEvent(evnt)
		var title string

		if evnt.Event == avahi.BrowserNew {
			title = fmt.Sprintf("addr-browse: found %s", key)
		} else {
			title = fmt.Sprintf("addr-browse: removed %s", key)
		}

		// Find hostname resolver
		hostname := back.getHostname(key)
		if hostname == nil {
			log.Debug(back.ctx, "%s: unknown hostname", title)
			return nil
		}

		// Decode address
		var addr netip.Addr
		if key.Proto == avahi.ProtocolIP4 {
			addr = avahi.DNSDecodeA(evnt.RData)
		} else {
			addr = avahi.DNSDecodeAAAA(evnt.RData)
		}

		if addr == (netip.Addr{}) {
			log.Debug(back.ctx, "%s: invalid addr", title)
			return nil
		}

		// Add or delete the address
		hasAddr := hostname.hasAddr(addr)
		switch {
		case evnt.Event == avahi.BrowserNew && hasAddr:
			log.Debug(back.ctx, "%s: %s (duplicate)", title, addr)
			return nil
		case evnt.Event == avahi.BrowserRemove && !hasAddr:
			log.Debug(back.ctx, "%s: %s (unknown addr)", title,
				addr)
			return nil

		case evnt.Event == avahi.BrowserNew:
			hostname.addAddr(addr)
		case evnt.Event == avahi.BrowserRemove:
			hostname.delAddr(addr)
		}

		log.Debug(back.ctx, "%s: %s", title, addr)

	case avahi.BrowserFailure:
		title := fmt.Sprintf("addr-browse: failed %s", evnt.Name)

		log.Error(back.ctx, "%s: %s", title, evnt.Err)
		return evnt.Err
	}

	return nil
}

// hasService reports if avahiService already exist for the key
func (back *backend) hasService(key avahiServiceKey) bool {
	return back.getService(key) != nil
}

// getService returns existent avahiService
func (back *backend) getService(key avahiServiceKey) *avahiService {
	return back.services[key]
}

// addService creates a new avahiService and registers it
// in the back.poller and back.services.
func (back *backend) addService(key avahiServiceKey) error {
	// Create service resolver
	flags := avahiLookupFlags[back.lookupFlags]
	flags |= avahi.LookupNoAddress | avahi.LookupNoTXT

	svcResolver, err := avahi.NewServiceResolver(
		back.clnt,
		key.IfIdx,
		key.Proto,
		key.InstanceName,
		key.SvcType,
		key.Domain,
		avahi.ProtocolUnspec,
		flags,
	)

	title := fmt.Sprintf("svc-resolve: start %s", key)

	if err != nil {
		log.Error(back.ctx, "%s: %s", title, err)
		return err
	}

	log.Debug(back.ctx, "%s: OK", title)

	// Create TXT record browser
	txtBrowser, err := avahi.NewRecordBrowser(
		back.clnt,
		key.IfIdx,
		key.Proto,
		key.FQDN(),
		avahi.DNSClassIN,
		avahi.DNSTypeTXT,
		avahiLookupFlags[back.lookupFlags],
	)

	title = fmt.Sprintf("txt-browse: start %s", key)

	if err != nil {
		log.Error(back.ctx, "%s: %s", title, err)
		return err
	}

	log.Debug(back.ctx, "%s: OK", title)

	// Create avahiService
	service := &avahiService{
		key:         key,
		svcResolver: svcResolver,
		txtBrowser:  txtBrowser,
	}

	back.poller.AddServiceResolver(svcResolver)
	back.poller.AddRecordBrowser(txtBrowser)

	back.services[key] = service
	return nil
}

// delService deletes the avahiService
func (back *backend) delService(service *avahiService) {
	log.Debug(back.ctx, "svc-resolve: stop %s", service.key)
	service.svcResolver.Close()

	log.Debug(back.ctx, "txt-browse: stop %s", service.key)
	service.txtBrowser.Close()

	delete(back.services, service.key)
	back.delServiceHostname(service)
}

// addServiceHostname creates an associon between avahiService
// and avahiHostname.
//
// If avahiHostname doesn't yet exist, it will be created on demand.
//
// If old association already exist, it will be removed.
func (back *backend) addServiceHostname(service *avahiService,
	name string) error {

	// Do nothing if service already has a hostname and
	// name was not changed.
	if service.hostname != nil && service.hostname.key.Hostname == name {
		return nil
	}

	// Prepare avahiHostnameKey
	key := avahiHostnameKey{
		IfIdx:    service.key.IfIdx,
		Proto:    service.key.Proto,
		Hostname: name,
	}

	// Find or create avahiHostname
	hostname := back.hostnames[key]
	if hostname == nil {
		var err error
		hostname, err = back.addHostname(key)
		if err != nil {
			return err
		}
	}

	service.hostname = hostname
	hostname.services[service] = struct{}{}

	return nil
}

// delServiceHostname removes association between avahiService
// and avahiHostname (if this association exists).
func (back *backend) delServiceHostname(service *avahiService) {
	hostname := service.hostname
	if hostname != nil {
		delete(hostname.services, service)
		if len(hostname.services) == 0 {
			back.delHostname(hostname)
		}
	}
}

// getHostname returns existent avahiHostname
func (back *backend) getHostname(key avahiHostnameKey) *avahiHostname {
	return back.hostnames[key]
}

// addHostname adds a new avahiHostname for the key
func (back *backend) addHostname(key avahiHostnameKey) (*avahiHostname, error) {
	// Create address resolver
	rtype := avahi.DNSTypeA
	if key.Proto == avahi.ProtocolIP6 {
		rtype = avahi.DNSTypeAAAA
	}

	addrBrowser, err := avahi.NewRecordBrowser(
		back.clnt,
		key.IfIdx,
		key.Proto,
		key.Hostname,
		avahi.DNSClassIN,
		rtype,
		avahiLookupFlags[back.lookupFlags],
	)

	title := fmt.Sprintf("addr-browse: start %s", key)

	if err != nil {
		log.Error(back.ctx, "%s: %s", title, err)
		return nil, err
	}

	log.Debug(back.ctx, "%s: OK", title)

	// Create avahiHostname
	hostname := &avahiHostname{
		key:         key,
		addrBrowser: addrBrowser,
		addrs:       make(map[netip.Addr]struct{}),
		services:    make(map[*avahiService]struct{}),
	}

	back.poller.AddRecordBrowser(addrBrowser)
	back.hostnames[key] = hostname

	return hostname, nil
}

// delHostname deletes avahiHostname for the key
func (back *backend) delHostname(hostname *avahiHostname) {
	log.Debug(back.ctx, "addr-browse: stop %s", hostname.key)
	hostname.addrBrowser.Close()

	delete(back.hostnames, hostname.key)
}

// avahiLookupFlags maps LookupFlags to avahi.LookupFlags
var avahiLookupFlags = [...]avahi.LookupFlags{
	0:               avahi.LookupUseMulticast,
	LookupClassical: avahi.LookupUseWideArea,
	LookupMulticast: avahi.LookupUseMulticast,
	LookupBoths:     avahi.LookupUseMulticast,
}

// avahiService is the per-service-instance structure
// that manages resources associated with the service
type avahiService struct {
	key         avahiServiceKey        // Identity
	svcResolver *avahi.ServiceResolver // Service resolver
	txtBrowser  *avahi.RecordBrowser   // TXT record resolver
	port        uint16                 // IP port
	hostname    *avahiHostname         // Hostname resolver
}

// avahiServiceKey identifies a particular instance of
// service of the particular type.
type avahiServiceKey struct {
	IfIdx        avahi.IfIndex  // Network interface index
	Proto        avahi.Protocol // Network protocol
	InstanceName string         // Service instance name
	SvcType      string         // Service type
	Domain       string         // Service domain
}

// FQDN returns the full-qualified domain name for the
// service instance.
func (key avahiServiceKey) FQDN() string {
	return avahi.DomainServiceNameJoin(key.InstanceName,
		key.SvcType, key.Domain)
}

// String returns string representation of the avahiServiceKey,
// for debugging.
func (key avahiServiceKey) String() string {
	var ifname string

	if ifi, err := net.InterfaceByIndex(int(key.IfIdx)); err == nil {
		ifname = ifi.Name
	} else {
		ifname = fmt.Sprintf("%d", key.IfIdx)
	}

	return fmt.Sprintf("%q (%s,%s)", key.FQDN(), key.Proto, ifname)
}

// avahiServiceKeyFromServiceBrowserEvent makes avahiServiceKey
// from the avahi.ServiceBrowserEvent.
func avahiServiceKeyFromServiceBrowserEvent(
	evnt *avahi.ServiceBrowserEvent) avahiServiceKey {

	return avahiServiceKey{
		IfIdx:        evnt.IfIdx,
		Proto:        evnt.Proto,
		InstanceName: evnt.InstanceName,
		SvcType:      evnt.SvcType,
		Domain:       evnt.Domain,
	}
}

// avahiServiceKeyFromResolverEvent makes avahiServiceKey
// from the avahi.ServiceResolverEvent.
func avahiServiceKeyFromResolverEvent(
	evnt *avahi.ServiceResolverEvent) avahiServiceKey {

	return avahiServiceKey{
		IfIdx:        evnt.IfIdx,
		Proto:        evnt.Proto,
		InstanceName: evnt.InstanceName,
		SvcType:      evnt.SvcType,
		Domain:       evnt.Domain,
	}
}

// avahiServiceKeyFromRecordBrowserEvent makes avahiServiceKey
// from the avahi.RecordBrowserEvent
func avahiServiceKeyFromRecordBrowserEvent(
	evnt *avahi.RecordBrowserEvent) avahiServiceKey {

	instance, svctype, domain := avahi.DomainServiceNameSplit(evnt.Name)

	return avahiServiceKey{
		IfIdx:        evnt.IfIdx,
		Proto:        evnt.Proto,
		InstanceName: instance,
		SvcType:      svctype,
		Domain:       domain,
	}
}

// avahiHostname is the per-hostname structure that manages
// resources associated with the hostname
type avahiHostname struct {
	key         avahiHostnameKey           // Identity
	addrBrowser *avahi.RecordBrowser       // A/AAAA record resolver
	addrs       map[netip.Addr]struct{}    // Resolved addresses
	services    map[*avahiService]struct{} // Dependent services
}

// hasAddr reports if address already known
func (hostname *avahiHostname) hasAddr(addr netip.Addr) bool {
	_, found := hostname.addrs[addr]
	return found
}

// addAddr adds the address
func (hostname *avahiHostname) addAddr(addr netip.Addr) {
	hostname.addrs[addr] = struct{}{}
}

// delAddr deletes the address
func (hostname *avahiHostname) delAddr(addr netip.Addr) {
	delete(hostname.addrs, addr)
}

// avahiHostnameKey identifies a particular instance of
// the avahiHostname
type avahiHostnameKey struct {
	IfIdx    avahi.IfIndex  // Network interface index
	Proto    avahi.Protocol // Network protocol
	Hostname string         // Hostname string (FQDN)
}

// String returns string representation of the avahiHostnameKey,
// for debugging.
func (key avahiHostnameKey) String() string {
	var ifname string

	if ifi, err := net.InterfaceByIndex(int(key.IfIdx)); err == nil {
		ifname = ifi.Name
	} else {
		ifname = fmt.Sprintf("%d", key.IfIdx)
	}

	return fmt.Sprintf("%q (%s,%s)", key.Hostname, key.Proto, ifname)
}

// avahiHostnameKeyFromRecordBrowserEvent makes avahiHostnameKey
// from the avahi.RecordBrowserEvent
func avahiHostnameKeyFromRecordBrowserEvent(
	evnt *avahi.RecordBrowserEvent) avahiHostnameKey {

	return avahiHostnameKey{
		IfIdx:    evnt.IfIdx,
		Proto:    evnt.Proto,
		Hostname: evnt.Name,
	}
}
