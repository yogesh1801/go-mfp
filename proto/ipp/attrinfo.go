// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package ipp

import "github.com/OpenPrinting/goipp"

// AttrInfo contains information about some IPP attribute.
type AttrInfo struct {
	Name string    // Attribute name
	Tag  goipp.Tag // Value tag
}
