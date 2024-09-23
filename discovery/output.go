// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Output generation

package discovery

import (
	"time"
)

// output generates and manages the final discovery output from
// the internal representation of the discovered information,
// gathered in the cache
type output struct {
	devices []Device  // Cached output data
	ttl     time.Time // Cache valid until this time
}

// Cached returns the cached output data (created by latest output.Generate)
// It may return nil, if this information is not available.
func (out *output) Cached() []Device {
	if out.devices != nil && !out.ttl.After(time.Now()) {
		return out.devices
	}
	return nil
}

// Invalidate drops the cached output
func (out *output) Invalidate() {
	out.devices = nil
}

// Generate generates the discovery output from the discovery
// information, gathered in the cache.
func (out *output) Generate(ttl time.Time, units []unit) []Device {
	// Merge variants
	units = out.genMergeUnitVariants(units)

	// Classify units by DeviceName+UUID+Realm
	devices := out.genMergeDevicesByNameUUID(units)

	// Generate final output, save and returns
	outdevs := make([]Device, len(devices))
	for i := range devices {
		outdevs[i] = devices[i].Export()
	}

	out.devices = outdevs
	out.ttl = ttl

	return outdevs
}

// genMergeUnitVariants merges together units with distinct UnitID.Variant,
// but otherwise similar.
//
// We need it because if some variant (say, ip6) will disappear
// from the network, remaining variants still make the unit
// visible.
//
// Some units may come several times but in distinct variants
// (say ip4/ip6). This function merges them together, effectively
// remove duplicates.
func (out *output) genMergeUnitVariants(units []unit) []unit {
	scratchpad := make(map[UnitID]unit)
	for _, un := range units {
		un.id.Variant = ""
		key := un.id

		if prev, found := scratchpad[key]; found {
			// Keep the first found unit; merge endpoints
			prev.Merge(un)
			scratchpad[key] = prev
		} else {
			// Add new unit
			scratchpad[key] = un
		}
	}

	units = make([]unit, 0, len(scratchpad))
	for _, un := range scratchpad {
		units = append(units, un)
	}

	return units
}

// genMergeUnitSameFunction merges together units that differ
// in UnitID.Zone, but otherwise similar.
//
// It may happen, for example, when units from different
// network protocols appeared to belong to the same device.
func (out *output) genMergeUnitCrossZones(units []unit) []unit {
	scratchpad := make(map[UnitID]unit)

	for _, un := range units {
		un.id.Zone = ""
		key := un.id

		if prev, found := scratchpad[key]; found {
			// Keep the first found unit; merge endpoints
			prev.Merge(un)
			scratchpad[key] = prev
		} else {
			// Add new unit
			scratchpad[key] = un
		}
	}

	units = make([]unit, 0, len(scratchpad))
	for _, un := range scratchpad {
		units = append(units, un)
	}

	return units
}

// genMergeDevicesByNameUUID merges units by DeviceName+UUID+Realm
// and returns result as a slice of device-s.
func (out *output) genMergeDevicesByNameUUID(units []unit) []device {
	// Classify units by DeviceName+UUID+Realm
	scratchpad := make(map[UnitID][]unit)
	for _, un := range units {
		key := UnitID{
			DeviceName: un.id.DeviceName,
			UUID:       un.id.UUID,
			Realm:      un.id.Realm,
		}
		scratchpad[key] = append(scratchpad[key], un)
	}

	// TODO -- merge by addr

	// Build slice of devices
	devices := make([]device, 0, len(scratchpad))
	for _, devunits := range scratchpad {
		dev := device{devunits}
		devices = append(devices, dev)
	}

	// Post-process the devices
	for i := range devices {
		dev := &devices[i]
		dev.units = out.genMergeUnitCrossZones(dev.units)
	}

	return devices
}
