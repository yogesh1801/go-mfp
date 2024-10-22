// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of local UDP ports

package wsdd

import (
	"net/netip"
	"sync"

	"github.com/alexpevzner/mfp/internal/generic"
)

// ports represents a set of own UDP ports.
// It is used to filter off own UDP messages.
type ports struct {
	set  generic.Set[netip.AddrPort] // Local addr:ports
	lock sync.Mutex                  // ports.set lock
}

// newPorts returns a new set of UDP ports
func newPorts() *ports {
	return &ports{
		set: generic.NewSet[netip.AddrPort](),
	}
}

// Add adds address to the set
func (p *ports) Add(addr netip.AddrPort) {
	p.lock.Lock()
	p.set.Add(addr)
	p.lock.Unlock()
}

// Del deletes address to the set
func (p *ports) Del(addr netip.AddrPort) {
	p.lock.Lock()
	p.set.Del(addr)
	p.lock.Unlock()
}

// Contains reports if address in the set.
func (p *ports) Contains(addr netip.AddrPort) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.set.Contains(addr)
}

// Clear clears the set
func (p *ports) Clear() {
	p.lock.Lock()
	p.set.Clear()
	p.lock.Unlock()
}
