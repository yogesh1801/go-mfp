// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The Printer object

package discovery

import "strings"

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
	Bind    bool // Printer can bind output
	Collate bool // Printer can collate copies
	Color   bool // Printer can print in color
	Copies  bool // Printer can make copies in hardware
	Duplex  bool // Printer supports duplex printing
	Punch   bool // Printer can punch output
	Sort    bool // Printer can sort output
	Staple  bool // Printer can staple output

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

	if p.Bind {
		s = append(s, "bind")
	}
	if p.Collate {
		s = append(s, "collate")
	}
	if p.Color {
		s = append(s, "color")
	}
	if p.Copies {
		s = append(s, "copies")
	}
	if p.Duplex {
		s = append(s, "duplex")
	}
	if p.Punch {
		s = append(s, "punch")
	}
	if p.Sort {
		s = append(s, "sort")
	}
	if p.Staple {
		s = append(s, "staple")
	}

	return strings.Join(s, ",")
}
