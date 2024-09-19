// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery cache

package discovery

import (
	"errors"
	"fmt"
	"time"
)

// cache represents the discovery cache
type cache struct {
	ctime   time.Time            // Cache creation time
	entries map[UnitID]*cacheEnt // Cache entries
}

// cacheUnit is the cache entry for print/scan/faxout units.
type cacheUnit struct {
	id        UnitID   // Unit identity
	meta      Metadata // Unit metadata
	params    any      // PrinterParameters or ScannerParameters
	endpoints []string // Unit endpoints
}

// cacheEnt is the cache entry for print/scan/faxout units.
type cacheEnt struct {
	cacheUnit
	hasMeta          bool      // Metadata is received
	hasParams        bool      // Parameters are received
	stagingEndpoints []string  // Newly discovered endpoints, on quarantine
	stagingDoneAt    time.Time // End of staging time
}

// newCache creates the new discovery cache
func newCache() *cache {
	return &cache{
		ctime:   time.Now(),
		entries: make(map[UnitID]*cacheEnt),
	}
}

// AddUnit adds new unit. Called when EventAddUnit is received.
func (c *cache) AddUnit(id UnitID) error {
	if c.entries[id] != nil {
		return errors.New("unit already added")
	}

	c.entries[id] = &cacheEnt{cacheUnit: cacheUnit{id: id}}
	return nil
}

// DelUnit deletes the unit. Called when EventDelUnit is received.
func (c *cache) DelUnit(id UnitID) error {
	if c.entries[id] == nil {
		return errors.New("unknown UnitID")
	}

	delete(c.entries, id)
	return nil
}

// SetMetadata saves unit Metadata.
// Called when EventMetadata is received
func (c *cache) SetMetadata(id UnitID, meta Metadata) error {
	ent := c.entries[id]
	if ent == nil {
		return errors.New("unknown UnitID")
	}

	ent.meta = meta
	ent.hasMeta = true

	return nil
}

// SetPrinterParameters saves printer parameters.
// Called when EventPrinterParameters is received
func (c *cache) SetPrinterParameters(id UnitID, p PrinterParameters) error {
	return c.setParameters(id, ServicePrinter, p)
}

// SetScannerParameters saves scanner parameters.
// Called when EventScannerParameters is received
func (c *cache) SetScannerParameters(id UnitID, p ScannerParameters) error {
	return c.setParameters(id, ServiceScanner, p)
}

// AddEndpoint adds unit endpoint.
func (c *cache) AddEndpoint(id UnitID, endpoint string) error {
	ent := c.entries[id]
	if ent == nil {
		return errors.New("unknown UnitID")
	}

	if endpointsContain(ent.endpoints, endpoint) ||
		endpointsContain(ent.stagingEndpoints, endpoint) {
		return errors.New("endpoint already added")
	}

	ent.stagingBegin()
	ent.stagingEndpoints, _ = endpointsAdd(ent.stagingEndpoints, endpoint)

	return nil
}

// AddEndpoint deletes unit endpoint.
func (c *cache) DelEndpoint(id UnitID, endpoint string) error {
	ent := c.entries[id]
	if ent == nil {
		return errors.New("unknown UnitID")
	}

	switch {
	case endpointsContain(ent.endpoints, endpoint):
		ent.endpoints, _ = endpointsDel(ent.endpoints, endpoint)

	case endpointsContain(ent.stagingEndpoints, endpoint):
		ent.stagingEndpoints, _ = endpointsDel(ent.stagingEndpoints,
			endpoint)
	default:
		return errors.New("unknown endpoint")
	}

	return nil
}

// SetFaxoutParameters saves faxout parameters.
// Called when EventFaxoutParameters is received.
//
// Note, Faxout parameters are represented by the PrinterParameters type.
func (c *cache) SetFaxoutParameters(id UnitID, p PrinterParameters) error {
	return c.setParameters(id, ServiceFaxout, p)
}

// setParameters saves unit parameters.
func (c *cache) setParameters(id UnitID, svcMustBe ServiceType, p any) error {
	ent := c.entries[id]
	if ent == nil {
		return errors.New("unknown UnitID")
	}

	if ent.id.SvcType != svcMustBe {
		return fmt.Errorf("unit is %s, must be %s",
			ent.id.SvcType, svcMustBe)
	}

	ent.params = p
	ent.hasParams = true

	return nil
}

// stagingBegin starts staging interval for discovered endpoints.
//
// The purpose of this interval is to merge together endpoints
// discovery events, which may come not in the same time (for example,
// if device has IP4 and IP6 addresses, these addresses may come
// spread out over time).
//
// So when new endpoint is discovered, we don't publish it immediately,
// but instead save into the staging area and start the timer. The subsequent
// endpoints are added to the staging area, but already running timer is
// not restarted. When the timer expires, all endpoints collected at
// the staging area is published.
func (ent *cacheEnt) stagingBegin() {
	if !ent.stagingCheck() {
		ent.stagingDoneAt = time.Now().Add(StabilizationTime)
	}
}

// stagingEnd publishes all endpoints collected in the staging area.
func (ent *cacheEnt) stagingEnd() {
	ent.endpoints = endpointsMerge(ent.endpoints, ent.stagingEndpoints)
	ent.stagingEndpoints = ent.stagingEndpoints[:]
}

// stagingCheck checks if staging interval is expired. If so,
// it merges publishes the staged endpoints and clears staging area.
//
// It returns 'true' if the previously started staging interval still
// in progress.
func (ent *cacheEnt) stagingCheck() bool {
	switch {
	case len(ent.stagingEndpoints) == 0:
		return false
	case ent.stagingDoneAt.After(time.Now()):
		ent.stagingEnd()
		return false
	}

	return true
}
