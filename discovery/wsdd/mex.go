// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Metadate Exchange (MEX)

package wsdd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/alexpevzner/mfp/discovery/dnssd"
	"github.com/alexpevzner/mfp/discovery/netstate"
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/wsd"
)

// mexGetter retrieves WSD metadata by XAddr URL.
type mexGetter struct {
	back      *backend                    // Parent backend
	ctx       context.Context             // Cancelable context
	cancel    context.CancelFunc          // Its cancel function
	http      http.Client                 // HTTP client
	cache     map[mexCacheID]*mexCacheEnt // Cached metadata
	lock      sync.Mutex                  // Access lock
	closewait sync.WaitGroup              // for mexGetter.Close
}

// newMexgetter creates a new mexGetter
func newMexgetter(back *backend) *mexGetter {
	ctx, cancel := context.WithCancel(back.ctx)

	mg := &mexGetter{
		back:   back,
		ctx:    ctx,
		cancel: cancel,
		http: http.Client{
			Timeout: 5 * time.Second,
		},
		cache: make(map[mexCacheID]*mexCacheEnt),
	}

	return mg
}

// Close closes the mexGetter. It cancels all pending metadata
// requests.
func (mg *mexGetter) Close() {
	mg.cancel()
	mg.closewait.Wait()
}

// Get retrives the wsd.Metadata from the device.
//
// Parameters are:
//   - from is the address of local interface where request is initiated
//   - target is the WSD target service address
//   - ver is the MetadataVersion. It comes together with the
//     XAddr URLs as a part of the wsd.Announce structure.
//
// The MetadataVersion affects metadata caching. The never metadata
// overrides the cached version.
func (mg *mexGetter) Get(from netstate.Addr, target wsd.AnyURI,
	xaddr *url.URL, ver uint64) []wsd.Metadata {

	// Obtain cache entry or create a new one
	mg.lock.Lock()

	id := mexCacheID{from, target, xaddr.String()}
	ent := mg.cache[id]

	if ent == nil {
		ent = mg.newEnt(id, xaddr, ver)
		mg.closewait.Add(1)
		go func() {
			mg.fetch(from, target, ent)
			mg.closewait.Done()
		}()
	}

	mg.lock.Unlock()

	// Wait until metadata fetching is done
	ent.wait()

	// Return whatever we have
	return ent.metadata
}

// newEnt creates a new cache entry.
func (mg *mexGetter) newEnt(id mexCacheID, xaddr *url.URL,
	ver uint64) *mexCacheEnt {

	ent := &mexCacheEnt{
		xaddr:    xaddr,
		ver:      ver,
		waitchan: make(chan struct{}),
	}

	return ent
}

// fetch fetches the metadata
func (mg *mexGetter) fetch(from netstate.Addr, target wsd.AnyURI,
	ent *mexCacheEnt) {

	var xaddrs []*url.URL

	// Resolve XAddr. It may return multiple literal URLs.
	xaddrs, err := mg.resolve(from, ent.xaddr)
	if err != nil {
		ent.err = err
		return
	}

	// Fetch metadata
	var wait sync.WaitGroup
	var lock sync.Mutex
	var metadata []wsd.Metadata

	wait.Add(len(xaddrs))

	for _, xaddr := range xaddrs {
		go func(xaddr2 *url.URL) {
			meta, err := mg.fetchHTTP(from, target, xaddr2)
			lock.Lock()
			if err != nil {
				metadata = append(metadata, meta)
			}
			lock.Unlock()

			wait.Done()
		}(xaddr)
	}

	wait.Wait()

	// Update cache entry
	ent.metadata = append(ent.metadata, metadata...)
}

// fetchHTTP performs HTTP query for the WSD metadata
func (mg *mexGetter) fetchHTTP(from netstate.Addr, target wsd.AnyURI,
	xaddr *url.URL) (meta wsd.Metadata, err error) {

	// Create a request
	msgid := wsd.AnyURI(uuid.Must(uuid.Random()).URN())
	msg := wsd.Msg{
		Header: wsd.Header{
			Action:    wsd.ActGet,
			MessageID: msgid,
			To:        target,
		},
		Body: wsd.Get{},
	}

	data := msg.Encode()

	rq := &http.Request{
		Method: "POST",
		URL:    xaddr,
		Body:   io.NopCloser(bytes.NewReader(data)),
		Close:  true,
	}

	rq = rq.WithContext(mg.ctx)

	// Perform HTTP query
	rsp, err := mg.http.Do(rq)
	if err != nil {
		return
	}

	defer rsp.Body.Close()

	if rsp.StatusCode/100 != 2 {
		err = fmt.Errorf("Unexpected HTTP status: %s", rsp.Status)
		return
	}

	// Decode response
	data, err = io.ReadAll(io.LimitReader(rsp.Body,
		int64(wsddMetadataGetMaxResponse+1)))

	if err != nil {
		return
	}

	if len(data) > wsddMetadataGetMaxResponse {
		err = fmt.Errorf("HTTP response too large")
		return
	}

	msg, err = wsd.DecodeMsg(data)
	if err != nil {
		return
	}

	meta, ok := msg.Body.(wsd.Metadata)
	if !ok {
		err = fmt.Errorf("Unexpected WSD response: %s",
			msg.Header.Action)
	}

	return
}

// resolve returns one or more literal URLs by replacing
// hostname with the resolved IP addresses.
func (mg *mexGetter) resolve(from netstate.Addr,
	u *url.URL) ([]*url.URL, error) {

	// Literal URL doesn't need to be resolved
	if urlIsLiteral(u) {
		return []*url.URL{u}, nil
	}

	// Create a resolver
	res, err := dnssd.NewResolver()
	if err != nil {
		return nil, err
	}

	defer res.Close()

	// Resolve hostname
	ifidx := from.Interface().Index()
	addrs, err := res.LookupHost(mg.ctx, ifidx, u.Hostname())
	if err != nil {
		return nil, err
	}

	// Create literal URLs
	var urls []*url.URL
	for _, addr := range addrs {
		urls = append(urls, urlWithHostname(u, addr.String()))
	}

	return urls, nil
}

// mexCacheID is the MEX cache entry ID.
type mexCacheID struct {
	from   netstate.Addr // Local interface address
	target wsd.AnyURI    // Target service WSD "address"
	xaddr  string        // Requested URL string
}

// mexCacheEnt is the MEX cache entry.
type mexCacheEnt struct {
	xaddr    *url.URL       // Requested XAddr
	ver      uint64         // Requested metadata version
	metadata []wsd.Metadata // Cached metadata
	err      error          // Metadata query error
	waitchan chan struct{}  // Closed when fetch is complete
}

// done marks cache entry as done downloading and awakes all
// goroutines blocked at mexCacheEnt.wait
func (ent *mexCacheEnt) done() {
	close(ent.waitchan)
}

// wait waits until metadata retrieval completion.
func (ent *mexCacheEnt) wait() {
	<-ent.waitchan
}
