// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The Printer object

package discovery

// PrintUnit represents a print unit
type PrintUnit struct {
	ID        UnitID            // Unit identity
	Params    PrinterParameters // Printer parameters
	Endpoints []string          // URLs of printer endpoints
}

// PrinterParameters represents the discoverable information about the printer.
//
// It is defined in the [IPP.Everywhere] and [Apple Bounjour Printing]
// terms, but usable with other discovery protocols.
//
// [IPP.Everywhere]: https://ftp.pwg.org/pub/pwg/candidates/cs-ippeve11-20200515-5100.14.pdf
// [Apple Bounjour Printing]: https://developer.apple.com/bonjour/printing-specification/bonjourprinting-1.2.1.pdf
type PrinterParameters struct {
	// Printer description
	Auth     AuthMode  // Required authentication type
	AdminURL string    // Printer configuration page
	Location string    // E.g., "2nd Floor Computer Lab"
	Paper    PaperSize // Max paper size
	Media    MediaKind // Kind of output media

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
	PPD      string   // PPD file name, if any
	PDL      []string // Supported MIME types
	Queue    string   // Queue name
	Priority int      // Queue priority, 0(highest)...99(lowest)
}
