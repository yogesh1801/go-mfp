// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovered unit

package dnssd

import (
	"net/netip"
	"net/url"
	"strconv"

	"github.com/alexpevzner/mfp/discovery"
)

// unit represents a discovered print or scan unit.
// It accepts RR updates and generates events.
type unit struct {
	queue   *discovery.Eventqueue // Event queue
	id      discovery.UnitID      // Unit ID
	svcType string                // Service type
	addrs   set[netip.Addr]       // IP addresses of the unit
	port    uint16                // IP port
	txtPrn  txtPrinter            // Parsed TXT for print unit
	txtScn  txtScanner            // Parsed TXT for scan unit
}

// newPrinterUnit creates a new printer unit
func newPrinterUnit(queue *discovery.Eventqueue,
	id discovery.UnitID, txt txtPrinter, port uint16) *unit {

	un := &unit{
		queue:   queue,
		id:      id,
		svcType: txt.svcType,
		addrs:   newSet[netip.Addr](),
		port:    port,
		txtPrn:  txt,
	}

	un.queue.Push(&discovery.EventAddUnit{ID: un.id})
	un.queue.Push(&discovery.EventPrinterParameters{
		ID:      un.id,
		Printer: *txt.params,
	})

	return un
}

// newScannerUnit creates a new scanner unit
func newScannerUnit(queue *discovery.Eventqueue,
	id discovery.UnitID, txt txtScanner, port uint16) *unit {
	un := &unit{
		queue:   queue,
		id:      id,
		svcType: txt.svcType,
		addrs:   newSet[netip.Addr](),
		port:    port,
		txtScn:  txt,
	}

	un.queue.Push(&discovery.EventAddUnit{ID: un.id})
	un.queue.Push(&discovery.EventScannerParameters{
		ID:      un.id,
		Scanner: *txt.params,
	})

	return un
}

// Delete deletes the unit
func (un *unit) Delete() {
	un.queue.Push(&discovery.EventDelUnit{ID: un.id})
}

// IsPrinter reports if unit is the print unit
func (un *unit) IsPrinter() bool {
	return un.txtPrn.params != nil
}

// setTxtPrinter changes TXT record for printer unit
func (un *unit) SetTxtPrinter(txt txtPrinter) {
	un.txtPrn = txt
	un.queue.Push(&discovery.EventPrinterParameters{
		ID:      un.id,
		Printer: *txt.params,
	})
}

// setTxtScanner changes TXT record for scanner unit
func (un *unit) SetTxtScanner(txt txtScanner) {
	un.txtScn = txt
	un.queue.Push(&discovery.EventScannerParameters{
		ID:      un.id,
		Scanner: *txt.params,
	})
}

// addAddr adds unit's IP address
func (un *unit) AddAddr(addr netip.Addr) {
	if !un.addrs.Contains(addr) {
		un.addrs.Add(addr)
		if un.port != 0 {
			evnt := &discovery.EventAddEndpoint{
				ID:       un.id,
				Endpoint: un.endpoint(addr),
			}
			un.queue.Push(evnt)
		}
	}
}

// delAddr deletes unit's IP address
func (un *unit) DelAddr(addr netip.Addr) {
	if un.addrs.Contains(addr) {
		un.addrs.Del(addr)
		if un.port != 0 {
			evnt := &discovery.EventDelEndpoint{
				ID:       un.id,
				Endpoint: un.endpoint(addr),
			}
			un.queue.Push(evnt)
		}
	}
}

// endpoint creates an endpoint URL
func (un *unit) endpoint(addr netip.Addr) string {
	host := addr.WithZone("").String()
	if addr.Is6() {
		// We need to put address into square brackets and
		// append zone for the link-local addresses
		host = "[" + host
		if zone := addr.Zone(); zone != "" {
			host += "%" + zone
		}
		host += "]"
	}

	port := int(un.port)
	var url url.URL

	switch un.svcType {
	case svcTypeAppSocket:
		// socket://host[:port]
		// default port: 9100
		url.Scheme = "socket"
		url.Host = host
		if port != 9100 {
			url.Host += ":" + strconv.Itoa(port)
		}

	case svcTypeIPP, svcTypeIPPS:
		// ipp://host[:port]/queue or ipps://host[:port]/queue
		// default port: 631
		url.Scheme = "ipp"
		if un.svcType == svcTypeIPPS {
			url.Scheme = "ipps"
		}

		url.Host = host
		if port != 631 {
			url.Host += ":" + strconv.Itoa(port)
		}

		url.Path = un.txtPrn.params.Queue

	case svcTypeLPD:
		// lpd://host[:port]/queue
		// default port: 515
		url.Scheme = "lpd"
		url.Host = host
		if port != 515 {
			url.Host += ":" + strconv.Itoa(port)
		}
		url.Path = un.txtPrn.params.Queue

	case svcTypeESCL, svcTypeESCLS:
		// http://host[:port]/path or https://host[:port]/path
		// default port: 80
		url.Host = host
		if un.svcType == svcTypeESCL {
			url.Scheme = "http"
			if port != 80 {
				url.Host += ":" + strconv.Itoa(port)
			}
		} else {
			url.Scheme = "https"
			if port != 443 {
				url.Host += ":" + strconv.Itoa(port)
			}
		}
		url.Path = un.txtScn.uriPath
	}

	println("url:", url.String())
	return url.String()
}
