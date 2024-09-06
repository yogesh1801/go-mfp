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

	"github.com/alexpevzner/mfp/discovery"
)

// unit represents a discovered print or scan unit.
// It accepts RR updates and generates events.
type unit struct {
	untab  *unitTable              // Back link to unitsTable
	id     discovery.UnitID        // Unit ID
	addrs  map[netip.Addr]struct{} // IP addresses
	port   uint16                  // IP port
	txtPrn txtPrinter              // Parsed TXT for print unit
	txtScn txtScanner              // Parsed TXT for scan unit
}

// newPrinterUnit creates a new printer unit
func newPrinterUnit(id discovery.UnitID, txt txtPrinter, port uint16) *unit {
	un := &unit{
		id:     id,
		addrs:  make(map[netip.Addr]struct{}),
		port:   port,
		txtPrn: txt,
	}

	return un
}

// newScannerUnit creates a new scanner unit
func newScannerUnit(id discovery.UnitID, txt txtScanner, port uint16) *unit {
	un := &unit{
		id:     id,
		addrs:  make(map[netip.Addr]struct{}),
		port:   port,
		txtScn: txt,
	}

	return un
}

// setTxtPrinter changes TXT record for printer unit
func (un *unit) setTxtPrinter(txt txtPrinter) {
	un.txtPrn = txt
}

// setTxtScanner changes TXT record for scanner unit
func (un *unit) setTxtScanner(txt txtScanner) {
	un.txtScn = txt
}

// addAddr adds unit's IP address
func (un *unit) addAddr(addr netip.Addr) {
}

// delAddr deletes unit's IP address
func (un *unit) delAddr(addr netip.Addr) {
}

// pushEvent pushes event into the event queue
func (un *unit) pushEvent(e discovery.Event) {
	un.untab.PushEvent(e)
}
