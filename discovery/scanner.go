// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner information

package discovery

import (
	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/util/generic"
)

// ScannerParameters represents the discoverable information about the printer.
type ScannerParameters struct {
	// Scanner capabilities
	Duplex  Option                             // Duplex mode supported
	Sources ScanSource                         // Supported sources
	Colors  generic.Bitset[abstract.ColorMode] // Supported color modes
	PDL     []string                           // Supported MIME types
}
