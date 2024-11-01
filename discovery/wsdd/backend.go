// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSDD backend

package wsdd

import (
	"context"

	"github.com/alexpevzner/mfp/discovery"
	"github.com/alexpevzner/mfp/log"
)

// backend is the [discovery.Backend] for WSD device discovery.
type backend struct {
	ctx     context.Context       // For logging and backend.Close
	queue   *discovery.Eventqueue // Event queue
	querier *querier              // WSDD querier
}

// NewBackend creates a new [discovery.Backend] for WSD device discovery.
func NewBackend(ctx context.Context) (discovery.Backend, error) {
	// Set log prefix
	ctx = log.WithPrefix(ctx, "wsdd")

	// Create backend structure
	back := &backend{
		ctx: ctx,
	}

	// Create querier
	var err error
	back.querier, err = newQuerier(back)
	if err != nil {
		return nil, err
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
	back.querier.Start()

	log.Debug(back.ctx, "backend started")
}

// Close closes the backend
func (back *backend) Close() {
	back.querier.Close()
}

// Debug writes a LevelDebug message on behalf of the backend.
func (back *backend) debug(format string, args ...any) {
	log.Debug(back.ctx, format, args...)
}

// Warning writes a LevelWarning message on behalf of the backend.
func (back *backend) warning(format string, args ...any) {
	log.Warning(back.ctx, format, args...)
}

// Error writes a LevelError message on behalf of the backend.
func (back *backend) error(format string, args ...any) {
	log.Error(back.ctx, format, args...)
}
