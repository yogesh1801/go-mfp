// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package discovery

// Option is the 3-state Boolean value, which can be true, false
// or unknown. It is used to specify device characteristics like
// duplex or color support or similar.
type Option int

// Option values:
const (
	OptUnknown Option = iota
	OptFalse
	OptTrue
)
