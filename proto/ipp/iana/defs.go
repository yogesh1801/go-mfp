// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common definitions

package iana

import (
	"math"

	"github.com/OpenPrinting/goipp"
)

const (
	// MIN is the minimum bound for attribute value range
	MIN = math.MinInt32

	// MAX is the maximum bound for attribute value range
	MAX = math.MaxInt32
)

// Attribute is the attribute definition.
type Attribute struct {
	SetOf    bool                 // 1SetOf attribute
	Min, Max int32                // Allowed range of values
	Tags     []goipp.Tag          // Allowed value tags
	Members  map[string]Attribute // Collection members
}
