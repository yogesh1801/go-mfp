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
//
// Note, cache API is not reentrant and requires external locking.
type cache struct {
	readyAt time.Time            // Time when cache is warmed up and ready
	entries map[UnitID]*cacheEnt // Cache entries
	out     output               // Cached output
}

// cacheEnt is the cache entry for print/scan/faxout units.
type cacheEnt struct {
	unit
	hasMeta          bool      // Metadata is received
	hasParams        bool      // Parameters are received
	stagingEndpoints []string  // Newly discovered endpoints, on quarantine
	stagingDoneAt    time.Time // End of staging time. Zero if no staging.
}

// newCache creates the new discovery cache
func newCache() *cache {
	return &cache{
		readyAt: time.Now().Add(WarmUpTime),
		entries: make(map[UnitID]*cacheEnt),
	}
}

// ReadyAt returns time when cache is ready to be exported, according to
// the cache state and export Mode
func (c *cache) ReadyAt(m Mode) time.Time {
	// Handle simple cases
	switch m {
	case ModeNormal:
		return c.readyAt
	case ModeSnapshot:
		return time.Time{}
	}

	// If cache is ready, no need to wait
	if c.out.Cached() != nil {
		return time.Time{}
	}

	// for ModeWaitIncomplete we need to do some more work
	ready := c.readyAt

	for _, ent := range c.entries {
		ent.stagingCheck()
		if ent.stagingInProgress() {
			ready = timeLatest(ready, ent.stagingDoneAt)
		}
	}

	return ready
}

// Export exports the cached data.
func (c *cache) Export() []Device {
	// If cached output available, just return it now
	if devices := c.out.Cached(); devices != nil {
		return devices
	}

	// Re-generate the output
	units := make([]unit, 0, len(c.entries))
	ttl := time.Now().Add(365 * 24 * time.Hour) // Far in a future

	for _, ent := range c.entries {
		if ent.ready() {
			units = append(units, ent.unit)
			if ent.stagingInProgress() {
				ttl = timeEarliest(ttl, ent.stagingDoneAt)
			}
		}
	}

	return c.out.Generate(ttl, units)
}

// Snapshot exports the cached data in the ModeSnapshot mode.
func (c *cache) Snapshot() []Device {
	units := make([]unit, 0, len(c.entries))
	ttl := time.Now().Add(365 * 24 * time.Hour) // Far in a future

	for _, ent := range c.entries {
		unit, ok := ent.snapshot()
		if ok {
			units = append(units, unit)
		}
	}

	var out output
	return out.Generate(ttl, units)
}

// AddUnit adds new unit. Called when EventAddUnit is received.
func (c *cache) AddUnit(id UnitID) error {
	if c.entries[id] != nil {
		return errors.New("unit already added")
	}

	c.entries[id] = &cacheEnt{unit: unit{id: id}}
	c.out.Invalidate()

	return nil
}

// DelUnit deletes the unit. Called when EventDelUnit is received.
func (c *cache) DelUnit(id UnitID) error {
	if c.entries[id] == nil {
		return errors.New("unknown UnitID")
	}

	delete(c.entries, id)
	c.out.Invalidate()

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

	c.out.Invalidate()

	return nil
}

// SetPrinterParameters saves printer parameters.
// Called when EventPrinterParameters is received
func (c *cache) SetPrinterParameters(id UnitID, p PrinterParameters) error {
	p.fixup()
	return c.setParameters(id, ServicePrinter, p)
}

// SetScannerParameters saves scanner parameters.
// Called when EventScannerParameters is received
func (c *cache) SetScannerParameters(id UnitID, p ScannerParameters) error {
	return c.setParameters(id, ServiceScanner, p)
}

// SetFaxoutParameters saves faxout parameters.
// Called when EventFaxoutParameters is received.
//
// Note, Faxout parameters are represented by the PrinterParameters type.
func (c *cache) SetFaxoutParameters(id UnitID, p PrinterParameters) error {
	p.fixup()
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

	c.out.Invalidate()

	return nil
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

	c.out.Invalidate()

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

	c.out.Invalidate()

	return nil
}

// ready checks if cache entry is ready to be exported.
func (ent *cacheEnt) ready() bool {
	if ent.hasMeta && ent.hasParams {
		ent.stagingCheck() // Merge staged endpoints, if any
		return len(ent.endpoints) > 0
	}
	return false
}

// snapshot makes a snapshot from the cache entry.
func (ent *cacheEnt) snapshot() (un unit, ok bool) {
	if ent.hasMeta && ent.hasParams {
		un = ent.unit
		un.endpoints = endpointsMerge(un.endpoints,
			ent.stagingEndpoints)
		if len(un.endpoints) > 0 {
			return un, true
		}
	}

	return unit{}, false
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
	if !ent.stagingInProgress() {
		ent.stagingDoneAt = time.Now().Add(StabilizationTime)
	}
}

// stagingEnd publishes all endpoints collected in the staging area.
func (ent *cacheEnt) stagingEnd() {
	ent.endpoints = endpointsMerge(ent.endpoints, ent.stagingEndpoints)
	ent.stagingEndpoints = ent.stagingEndpoints[:0]
	ent.stagingDoneAt = time.Time{}
}

// stagingInProgress tells if staging interval is in progress.
func (ent *cacheEnt) stagingInProgress() bool {
	return !ent.stagingDoneAt.IsZero()
}

// stagingCheck checks if staging interval is expired. If so,
// it merges publishes the staged endpoints and clears staging area.
//
// It returns 'true' if the previously started staging interval still
// in progress.
func (ent *cacheEnt) stagingCheck() {
	if ent.stagingInProgress() && !ent.stagingDoneAt.After(time.Now()) {
		ent.stagingEnd()
	}
}
