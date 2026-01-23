// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The Printer object

package discovery

import (
	"strings"

	"github.com/OpenPrinting/go-mfp/util/optional"
)

// PrinterParameters represents the discoverable information about the printer.
//
// It is defined in the [IPP.Everywhere] and [Apple Bounjour Printing]
// terms, but usable with other discovery protocols.
//
// [IPP.Everywhere]: https://ftp.pwg.org/pub/pwg/candidates/cs-ippeve11-20200515-5100.14.pdf
// [Apple Bounjour Printing]: https://developer.apple.com/bonjour/printing-specification/bonjourprinting-1.2.1.pdf
type PrinterParameters struct {
	// Printer description
	Auth  AuthMode  // Required authentication type
	Paper PaperSize // Max paper size
	Media MediaKind // Kind of output media

	// Printer capabilities
	Bind    optional.Val[bool] // Printer can bind output
	Collate optional.Val[bool] // Printer can collate copies
	Color   optional.Val[bool] // Printer can print in color
	Copies  optional.Val[bool] // Printer can make copies in hardware
	Duplex  optional.Val[bool] // Printer supports duplex printing
	Punch   optional.Val[bool] // Printer can punch output
	Sort    optional.Val[bool] // Printer can sort output
	Staple  optional.Val[bool] // Printer can staple output

	// Operational parameters
	PSProduct string   // PS Product name (helps PPD location)
	PDL       []string // Supported MIME types
	Queue     string   // Queue name
	Priority  int      // Queue priority, 0(highest)...99(lowest)
}

// fixup fixes PrinterParameters, received from backend
func (p *PrinterParameters) fixup() {
	if p.Auth == 0 {
		p.Auth = AuthNone
	}
	if p.Media == 0 {
		p.Media = MediaOther
	}
}

// Flags formats printer flags (Bind, Color, Collate etc) into a
// single string ("bind,color,collate,...").
func (p PrinterParameters) Flags() string {
	s := []string{}

	if optional.Get(p.Bind) {
		s = append(s, "bind")
	}
	if optional.Get(p.Collate) {
		s = append(s, "collate")
	}
	if optional.Get(p.Color) {
		s = append(s, "color")
	}
	if optional.Get(p.Copies) {
		s = append(s, "copies")
	}
	if optional.Get(p.Duplex) {
		s = append(s, "duplex")
	}
	if optional.Get(p.Punch) {
		s = append(s, "punch")
	}
	if optional.Get(p.Sort) {
		s = append(s, "sort")
	}
	if optional.Get(p.Staple) {
		s = append(s, "staple")
	}

	return strings.Join(s, ",")
}
