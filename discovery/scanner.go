// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner information

package discovery

// ScannerInfo represents the discoverable information about the printer.
type ScannerInfo struct {
	// Scanner identity
	ID DeviceID // Device identity

	// Scanner description
	AdminURL string // Scanner configuration page
	IconURL  string // Icon URL ("" if not available)
	Location string // E.g., "2nd Floor Computer Lab"

	// Scanner capabilities
	Duplex bool        // Duplex mode supported
	Source InputSource // Supported sources
	Color  ColorMode   // Supported color modes
	PDL    []string    // Supported MIME types
}
