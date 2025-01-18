// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input source

package escl

// InputSource specifies the desired input source.
type InputSource int

// Known intents
const (
	UnknownInputSource InputSource = iota // Unknown intent
	InputPlaten                           // Scan from platen
	InputFeeder                           // Scan from feeder
	InputCamera                           // Scan from camera
)

// String returns a string representation of the [InputSource]
func (intent InputSource) String() string {
	switch intent {
	case InputPlaten:
		return "Platen"
	case InputFeeder:
		return "Feeder"
	case InputCamera:
		return "Camera"
	}

	return "Unknown"
}

// DecodeInputSource decodes [InputSource] out of its XML string representation.
func DecodeInputSource(s string) InputSource {
	switch s {
	case "Platen":
		return InputPlaten
	case "Feeder":
		return InputFeeder
	case "Camera":
		return InputCamera
	}

	return UnknownInputSource
}
