// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The virtual scanner

package abstract

import (
	"context"

	"github.com/OpenPrinting/go-mfp/log"
)

// VirtualScanner implements the [Scanner] interface for the virtual
// (simulated) scanner.
type VirtualScanner struct {
	ScanCaps    *ScannerCapabilities // Scanner capabilities
	Resolution  Resolution           // Images resolution
	PlatenImage []byte               // Image "loaded" into Platen
	ADFImages   [][]byte             // Images "loaded" into ADF
}

// Capabilities returns the [ScannerCapabilities].
// Caller should not modify the returned structure.
func (vscan *VirtualScanner) Capabilities() *ScannerCapabilities {
	return vscan.ScanCaps
}

// Scan supplies the scan request.
func (vscan *VirtualScanner) Scan(ctx context.Context, req ScannerRequest) (
	Document, error) {

	log.Begin(ctx).
		Debug("VSCAN: scan requested:").
		Object(log.LevelDebug, 4, &req).
		Commit()

	err := req.Validate(vscan.ScanCaps)
	if err != nil {
		log.Debug(ctx, "VSCAN: %s", err)
		return nil, err
	}

	images := [][]byte{vscan.PlatenImage}
	if req.Input == InputADF {
		images = vscan.ADFImages
	}

	doc := NewVirtualDocument(vscan.Resolution, images...)

	filter := NewFilter(doc)
	filter.SetResolution(req.Resolution)
	filter.SetRegion(req.Region)

	return filter, nil
}

// Close closes the scanner connection.
func (vscan *VirtualScanner) Close() error {
	return nil
}
