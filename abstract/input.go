// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input source

package abstract

// Input specifies the desired input source.
type Input int

// Known intents
const (
	InputUnset  Input = iota // Not set
	InputPlaten              // Scan from platen
	InputADF                 // Automatic Document Feeder
	InputCamera              // Scan from camera
)
