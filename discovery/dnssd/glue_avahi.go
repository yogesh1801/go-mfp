// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Avahi glue

package dnssd

import (
	"fmt"
	"net"
	"net/netip"
	"time"

	"github.com/alexpevzner/go-avahi"
)

// Parameters:
const (
	// Avahi client may fail to start if avahi-daemon is not
	// running or for for the similar reasons.
	//
	// If failed, Avahi client will automatically restart
	// with the following interval between attempts.
	avahiClientRestartInterval = 1 * time.Second
)

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

// avahiLookupFlagsTable contains mapping of LookupFlags to avahi.LookupFlags
var avahiLookupFlagsTable = [...]avahi.LookupFlags{
	0:               avahi.LookupUseMulticast,
	LookupClassical: avahi.LookupUseWideArea,
	LookupMulticast: avahi.LookupUseMulticast,
	LookupBoths:     avahi.LookupUseMulticast,
}

// avahiLookupFlags maps LookupFlags to avahi.LookupFlags
func avahiLookupFlags(flags LookupFlags) avahi.LookupFlags {
	return avahiLookupFlagsTable[flags&LookupBoths]
}
