// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner input source

package discovery

import "strings"

// ScanSource defines input sources, supported by scanner
type ScanSource int

// ScanSource bits:
const (
	ScanOther  = 1 << iota // Other input
	ScanPlaten             // Platen source
	ScanADF                // Automatic Document Feeder
)

// String formats ScanSource as string, for printing and logging
func (ss ScanSource) String() string {
	s := []string{}

	if ss&ScanOther != 0 {
		s = append(s, "other")
	}
	if ss&ScanPlaten != 0 {
		s = append(s, "platen")
	}
	if ss&ScanADF != 0 {
		s = append(s, "ADF")
	}

	return strings.Join(s, ",")
}
