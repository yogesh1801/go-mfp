// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner information

package discovery

// ScanUnit represents a scan unit
type ScanUnit struct {
	ID        UnitID            // Unit identity
	Meta      Metadata          // Unit metadata
	Params    ScannerParameters // Scanner parameters
	Endpoints []string          // URLs of printer endpoints
}

// ScannerParameters represents the discoverable information about the printer.
type ScannerParameters struct {
	// Scanner description
	AdminURL string // Scanner configuration page
	IconURL  string // Icon URL ("" if not available)
	Location string // E.g., "2nd Floor Computer Lab"

	// Scanner capabilities
	Duplex  bool        // Duplex mode supported
	Sources InputSource // Supported sources
	Colors  ColorMode   // Supported color modes
	PDL     []string    // Supported MIME types
}
