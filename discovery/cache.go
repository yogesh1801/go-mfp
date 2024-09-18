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
	ctime time.Time             // Cache creation time
	units map[UnitID]*cacheUnit // Cached units
}

// cacheUnit is the cache entry for print/scan/faxout units.
type cacheUnit struct {
	id               UnitID    // Unit identity
	meta             Metadata  // Unit metadata
	params           any       // PrinterParameters or ScannerParameters
	hasMeta          bool      // Metadata is received
	hasParams        bool      // Parameters are received
	endpoints        []string  // Unit endpoints
	stagingEndpoints []string  // Newly discovered endpoints, on quarantine
	stagingDoneAt    time.Time // End of staging time
}

// newCache creates the new discovery cache
func newCache() *cache {
	return &cache{
		ctime: time.Now(),
		units: make(map[UnitID]*cacheUnit),
	}
}

// AddUnit adds new unit. Called when EventAddUnit is received.
func (c *cache) AddUnit(id UnitID) error {
	if c.units[id] != nil {
		return errors.New("unit already added")
	}

	c.units[id] = &cacheUnit{id: id}
	return nil
}

// DelUnit deletes the unit. Called when EventDelUnit is received.
func (c *cache) DelUnit(id UnitID) error {
	if c.units[id] == nil {
		return errors.New("unknown UnitID")
	}

	delete(c.units, id)
	return nil
}

// SetMetadata saves unit Metadata.
// Called when EventMetadata is received
func (c *cache) SetMetadata(id UnitID, meta Metadata) error {
	un := c.units[id]
	if un == nil {
		return errors.New("unknown UnitID")
	}

	un.meta = meta
	un.hasMeta = true

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
	un := c.units[id]
	if un == nil {
		return errors.New("unknown UnitID")
	}

	if endpointsContain(un.endpoints, endpoint) ||
		endpointsContain(un.stagingEndpoints, endpoint) {
		return errors.New("endpoint already added")
	}

	un.stagingBegin()
	un.stagingEndpoints, _ = endpointsAdd(un.stagingEndpoints, endpoint)

	return nil
}

// AddEndpoint deletes unit endpoint.
func (c *cache) DelEndpoint(id UnitID, endpoint string) error {
	un := c.units[id]
	if un == nil {
		return errors.New("unknown UnitID")
	}

	switch {
	case endpointsContain(un.endpoints, endpoint):
		un.endpoints, _ = endpointsDel(un.endpoints, endpoint)

	case endpointsContain(un.stagingEndpoints, endpoint):
		un.stagingEndpoints, _ = endpointsDel(un.stagingEndpoints,
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
	un := c.units[id]
	if un == nil {
		return errors.New("unknown UnitID")
	}

	if un.id.SvcType != svcMustBe {
		return fmt.Errorf("unit is %s, must be %s",
			un.id.SvcType, svcMustBe)
	}

	un.params = p
	un.hasParams = true

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
func (un *cacheUnit) stagingBegin() {
	if !un.stagingCheck() {
		un.stagingDoneAt = time.Now().Add(StabilizationTime)
	}
}

// stagingEnd publishes all endpoints collected in the staging area.
func (un *cacheUnit) stagingEnd() {
	un.endpoints = endpointsMerge(un.endpoints, un.stagingEndpoints)
	un.stagingEndpoints = un.stagingEndpoints[:]
}

// stagingCheck checks if staging interval is expired. If so,
// it merges publishes the staged endpoints and clears staging area.
//
// It returns 'true' if the previously started staging interval still
// in progress.
func (un *cacheUnit) stagingCheck() bool {
	switch {
	case len(un.stagingEndpoints) == 0:
		return false
	case un.stagingDoneAt.After(time.Now()):
		un.stagingEnd()
		return false
	}

	return true
}
