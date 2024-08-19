// MFP - Miulti-Function Printers and scanners toolkit
// DNS-SD service discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery Backend

package dnssd

import (
	"context"

	"github.com/alexpevzner/go-avahi"
	"github.com/alexpevzner/mfp/discovery"
)

// backend is the [discovery.Backend] for DNS-SD discovery
type backend struct {
	clnt *avahi.Client
	chn  chan any
}

// NewBackend creates a new [discovery.Backend] for DNS-SD discovery.
func NewBackend(ctx context.Context) (discovery.Backend, error) {
	clnt, err := avahi.NewClient(avahi.ClientLoopbackWorkarounds)
	if err != nil {
		return nil, err
	}

	back := &backend{
		clnt: clnt,
		chn:  make(chan any),
	}

	return back, nil
}

// Chan returns an event channel.
func (back *backend) Chan() <-chan any {
	return back.chn
}

// Close closes the backend
func (back *backend) Close() {
	close(back.chn)
	back.clnt.Close()
}
