// MFP - Miulti-Function Printers and scanners toolkit
// IEEE 1284 definitions
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package ieee1284

import (
	"context"
	"io"

	"github.com/OpenPrinting/go-mfp/log"
)

// Printer implements the IEEE-1284 printer
type Printer struct {
	ctx context.Context // Logging context
}

// NewPrinter creates a new printer.
func NewPrinter(ctx context.Context) *Printer {
	return &Printer{
		ctx: ctx,
	}
}

// Write consumes data sent to the Printer from the host.
// It implements the [io.Writer] interface.
func (p *Printer) Write(data []byte) (int, error) {
	log.Dump(p.ctx, log.LevelTrace, data)
	return len(data), nil
}

// Read returns data sent by the Printer to the host.
// It implements the [io.Reader] interface.
func (p *Printer) Read(buf []byte) (int, error) {
	return 0, io.EOF
}

// Close closes the printer.
// It implements the [io.Closer] interface.
func (p *Printer) Close() error {
	return nil
}
