// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner input source

package discovery

// ScanSource defines input sources, supported by scanner
type ScanSource int

// ScanSource bits:
const (
	ScanOther  = 1 << iota // Other input
	ScanPlaten             // Platen source
	ScanADF                // Automatic Document Feeder
)
