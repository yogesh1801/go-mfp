// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package wsdd

import (
	"context"
	"sync"

	"github.com/alexpevzner/mfp/discovery"
	"github.com/alexpevzner/mfp/discovery/netstate"
	"github.com/alexpevzner/mfp/log"
)

// backend is the [discovery.Backend] for WSD device discovery.
type backend struct {
	ctx    context.Context       // For logging and backend.Close
	cancel context.CancelFunc    // Context's cancel function
	queue  *discovery.Eventqueue // Event queue
	netmon *netstate.Notifier    // Network state monitor
	mconn4 *mconn                // IP4 multicasts reception connection
	mconn6 *mconn                // IP6 multicasts reception connection
	done   sync.WaitGroup        // For backend.Close synchronization
}

// NewBackend creates a new [discovery.Backend] for WSD device discovery.
func NewBackend(ctx context.Context) (discovery.Backend, error) {
	// Set log prefix
	ctx = log.WithPrefix(ctx, "wsdd")

	// Create multicast sockets
	mconn4, err := newMconn(wsddMulticastIP4)
	if err != nil {
		return nil, err
	}

	mconn6, err := newMconn(wsddMulticastIP6)
	if err != nil {
		mconn4.Close()
		return nil, err
	}

	// Create cancelable context
	ctx, cancel := context.WithCancel(ctx)

	// Create backend structure
	back := &backend{
		ctx:    ctx,
		cancel: cancel,
		netmon: netstate.NewNotifier(),
		mconn4: mconn4,
		mconn6: mconn6,
	}
	return back, nil
}

// Name returns backend name.
func (back *backend) Name() string {
	return "wsdd"
}

// Start starts Backend operations.
func (back *backend) Start(queue *discovery.Eventqueue) {
	back.queue = queue

	back.done.Add(1)
	go back.netmonproc()

	log.Debug(back.ctx, "backend started")
}

// Close closes the backend
func (back *backend) Close() {
	back.cancel()
	back.done.Wait()
}

// netmonproc processes netstate.Notifier events.
// It runs on its own goroutine.
func (back *backend) netmonproc() {
	defer back.done.Done()

	for {
		evnt, err := back.netmon.Get(back.ctx)
		if err != nil {
			return
		}

		log.Debug(back.ctx, "%s", evnt)
	}
}
