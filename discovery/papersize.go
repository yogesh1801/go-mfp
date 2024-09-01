// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Paper size

package discovery

// PaperSize roughly defines the maximum paper size supported by printer
type PaperSize int

// PaperSize values:
const (
	PaperUnknown PaperSize = iota // Paper size is not known
	PaperA4Minus                  // Smaller that A4
	PaperA4                       // A4 (US Legal)
	PaperA3                       // A3 (IS Tabloid)
	PaperA2                       // A2 (ISO-C)
	PaperA2Plus                   // Large that A2
)
