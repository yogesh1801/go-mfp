// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The virtual scanner

package abstract

import "context"

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

	err := req.Validate(vscan.ScanCaps)
	if err != nil {
		return nil, err
	}

	images := [][]byte{vscan.PlatenImage}
	if req.Input == InputADF {
		images = vscan.ADFImages
	}

	doc := NewVirtualDocument(vscan.Resolution, images...)
	return doc, nil
}

// Close closes the scanner connection.
func (vscan *VirtualScanner) Close() error {
	return nil
}
