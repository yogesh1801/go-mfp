// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Authentication modes

package discovery

// MediaKind bits lists the categories of printing supported
// by the printer.
type MediaKind int

// MediaKind values:
const (
	MediaOther       MediaKind = 1 << iota // Other kind
	MediaDisk                              // Prints on CD/DVD
	MediaDocument                          // Standard document printing
	MediaEnvelope                          // Prints on envelopes
	MediaLabel                             // Prints on cut labels
	MediaLargeFormat                       // Large format (>A3)
	MediaPhoto                             // Photo printer
	MediaPostcard                          // Prints on postcards
	MediaReceipt                           // Continuous rolls of receipts
	MediaRoll                              // Rolls of docs/photos
)
