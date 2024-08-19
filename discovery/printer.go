// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The Printer object

package discovery

// PrinterInfo represents the discoverable information about the printer.
//
// It is defined in the [IPP.Everywhere] and [Apple Bounjour Printing]
// terms, but usable with other discovery protocols.
//
// [IPP.Everywhere]: https://ftp.pwg.org/pub/pwg/candidates/cs-ippeve11-20200515-5100.14.pdf
// [Apple Bounjour Printing]: https://developer.apple.com/bonjour/printing-specification/bonjourprinting-1.2.1.pdf
type PrinterInfo struct {
	// Printer identity
	ID DeviceID // Device identity

	// Printer description
	AuthInfoRequired KwAuthInfo   // Required authentication type
	AdminURL         string       // Printer configuration page
	Location         string       // E.g., "2nd Floor Computer Lab"
	TLSVersion       KwTLSVersion // Highest supported TLS version

	// Printer capabilities
	CanBind    bool     // Printer can bind output
	CanCollate bool     // Printer can collate copies
	CanColor   bool     // Printer can print in color
	CanCopies  bool     // Printer can make copies in hardware
	CanDuplex  bool     // Printer supports duplex printing
	CanPunch   bool     // Printer can punch output
	CanSort    bool     // Printer can sort output
	CanStaple  bool     // Printer can staple output
	PDL        []string // Supported MIME types

	// Printer endpoints
	Endpoints []string // Collection of URLs
}
