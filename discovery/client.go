// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery client

package discovery

import (
	"context"
	"fmt"
	"sync"

	"github.com/alexpevzner/mfp/log"
)

// Client implements a client side of devices discovery.
type Client struct {
	ctx      context.Context
	cancel   context.CancelFunc
	queue    *Eventqueue
	backends map[Backend]struct{}
	lock     sync.Mutex
	done     sync.WaitGroup
}

// NewClient creates a new discovery [Client].
func NewClient(ctx context.Context) *Client {
	// Set log prefix
	ctx = log.WithPrefix(ctx, "discovery")

	// Create cancelable context
	ctx, cancel := context.WithCancel(ctx)

	// Create client structure
	clnt := &Client{
		ctx:      ctx,
		cancel:   cancel,
		queue:    NewEventqueue(),
		backends: make(map[Backend]struct{}),
	}

	// Start work thread
	clnt.done.Add(1)
	go clnt.proc()

	return clnt
}

// Close closes all attached backends and then closes the Client
// and releases all resources it holds.
func (clnt *Client) Close() {
	clnt.cancel()
	clnt.done.Wait()
}

// AddBackend adds a discovery [Backend] to the [Client].
func (clnt *Client) AddBackend(bk Backend) {
	clnt.lock.Lock()
	defer clnt.lock.Unlock()

	if _, found := clnt.backends[bk]; found {
		err := fmt.Errorf("backend %s already added", bk.Name())
		panic(err)
	}

	log.Debug(clnt.ctx, "%s: backend added", bk.Name())
	clnt.backends[bk] = struct{}{}
	bk.Start(clnt.queue)
}

// proc runs the discovery event loop on its separate goroutine.
func (clnt *Client) proc() {
	defer clnt.done.Done()

	for {
		evnt, err := clnt.queue.pull(clnt.ctx)
		if err != nil {
			return
		}

		log.Debug(clnt.ctx, "%s", evnt.Name())
	}
}
