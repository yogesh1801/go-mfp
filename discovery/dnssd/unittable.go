// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Table of discovered units

package dnssd

import "github.com/alexpevzner/mfp/discovery"

// unitTable maintains all discovered units
type unitTable struct {
	units  map[discovery.UnitID]*unit // Table of units
	events []discovery.Event          // Generated events
}

// newUnitTable creates a new unitTable
func newUnitTable() *unitTable {
	return &unitTable{
		units: make(map[discovery.UnitID]*unit),
	}
}

// Get returns unit by id
func (untab *unitTable) Get(id discovery.UnitID) *unit {
	return untab.units[id]
}

// Put inserts unit into the table
func (untab *unitTable) Put(un *unit) {
	un.untab = untab
	untab.units[un.id] = un
	un.SendInitEvents()
}

// PushEvent pushes event into the event queue
func (untab *unitTable) PushEvent(e discovery.Event) {
	untab.events = append(untab.events, e)
}
