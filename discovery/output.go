// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Output generation

package discovery

import "time"

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
	// TODO
	return nil
}
