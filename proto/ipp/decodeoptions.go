// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Decode options

package ipp

// DecodeOptions represent options used when [Object] is being
// decoded from the [goipp.Attributes].
type DecodeOptions struct {
	// KeepTrying, if set, instructs decoder do not stop on
	// value decoding errors, but just skip problematic value
	// and continue.
	KeepTrying bool
}
