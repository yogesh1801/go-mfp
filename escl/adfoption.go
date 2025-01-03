// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF options

package escl

// ADFOption specifies options, supported by the Automatic Document Feeder.
type ADFOption int

// Known color modes:
const (
	UnknownADFOption  ADFOption = iota - 1 // Unknown color mode
	DetectPaperLoaded                      // Can detect if paper is loaded
	SelectSinglePage                       // Can scan part of loaded pages
	Duplex                                 // Duplex support
)

// String returns a string representation of the [ADFOption]
func (opt ADFOption) String() string {
	switch opt {
	case DetectPaperLoaded:
		return "DetectPaperLoaded"
	case SelectSinglePage:
		return "SelectSinglePage"
	case Duplex:
		return "Duplex"
	}

	return "Unknown"
}

// DecodeADFOption decodes [ADFOption] out of its XML string representation.
func DecodeADFOption(s string) ADFOption {
	switch s {
	case "DetectPaperLoaded":
		return DetectPaperLoaded
	case "SelectSinglePage":
		return SelectSinglePage
	case "Duplex":
		return Duplex
	}

	return UnknownADFOption
}
