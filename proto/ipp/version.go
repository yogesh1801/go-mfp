// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP version constants

package ipp

import "github.com/OpenPrinting/goipp"

const (
	// DefaultVersion is the default version of the IPP protocol.
	DefaultVersion = goipp.DefaultVersion

	// MinVersion is the minimal supported version of the IPP protocol.
	MinVersion goipp.Version = 0x0100

	// MaxVersion is the maximal supported version of the IPP protocol.
	MaxVersion goipp.Version = 0x0202
)
