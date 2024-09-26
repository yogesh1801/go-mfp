// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Authentication modes

package discovery

import "strings"

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

// String converts MediaKind to string.
func (media MediaKind) String() string {
	s := []string{}

	if media&MediaOther != 0 {
		s = append(s, "other")
	}
	if media&MediaDisk != 0 {
		s = append(s, "disc")
	}
	if media&MediaDocument != 0 {
		s = append(s, "document")
	}
	if media&MediaEnvelope != 0 {
		s = append(s, "envelope")
	}
	if media&MediaLabel != 0 {
		s = append(s, "label")
	}
	if media&MediaLargeFormat != 0 {
		s = append(s, "large-format")
	}
	if media&MediaPhoto != 0 {
		s = append(s, "photo")
	}
	if media&MediaPostcard != 0 {
		s = append(s, "postcard")
	}
	if media&MediaReceipt != 0 {
		s = append(s, "receipt")
	}
	if media&MediaRoll != 0 {
		s = append(s, "roll")
	}

	return strings.Join(s, ",")
}
