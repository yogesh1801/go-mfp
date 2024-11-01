// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Metadate Exchange (MEX)

package wsdd

import (
	"errors"
	"net/url"

	"github.com/alexpevzner/mfp/wsd"
)

// mexGetter retrieves WSD metadata by XAddr URL.
type mexGetter struct {
	back  *backend                    // Parent backend
	cache map[mexCacheID]*mexCacheEnt // Cached metadata
}

// mexCacheID is the MEX cache entry ID.
type mexCacheID struct {
	dev   wsd.AnyURI // Device address
	xaddr string     // Requested URL string
}

// mexCacheEnt is the MEX cache entry.
type mexCacheEnt struct {
	xaddr         *url.URL     // Requested XAddr
	xarrdResolved *url.URL     // XAddr with resolved hostname
	ver           uint64       // Requested metadata version
	metadata      wsd.Metadata // Cached metadata
}

// newMexgetter creates a new mexGetter
func newMexgetter(back *backend) *mexGetter {
	mg := &mexGetter{
		back:  back,
		cache: make(map[mexCacheID]*mexCacheEnt),
	}

	return mg
}

// Close closes the mexGetter. It cancels all pending metadata
// requests.
func (mg *mexGetter) Close() {
}

// Get retrives the wsd.Metadata from the device.
//
// dev is the WSD device address
//
// ver is the MetadataVersion. It comes together with the
// XAddr URLs as a part of the wsd.Announce structure.
//
// The MetadataVersion affects metadata caching. The never metadata
// overrides the cached version.
func (mg *mexGetter) Get(dev wsd.AnyURI, ver uint64,
	xaddr wsd.AnyURI) (*wsd.Metadata, error) {
	return nil, errors.New("not implemented")
}
