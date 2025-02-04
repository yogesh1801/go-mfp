// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input source

package escl

import "github.com/alexpevzner/mfp/xmldoc"

// InputSource specifies the desired input source.
type InputSource int

// Known intents
const (
	UnknownInputSource InputSource = iota // Unknown input
	InputPlaten                           // Scan from platen
	InputFeeder                           // Scan from feeder
	InputCamera                           // Scan from camera
)

// decodeInputSource decodes [InputSource] from the XML tree.
func decodeInputSource(root xmldoc.Element) (input InputSource, err error) {
	return decodeEnum(root, DecodeInputSource, NsScan)
}

// toXML generates XML tree for the [InputSource].
func (input InputSource) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: NsScan + ":" + input.String(),
	}
}

// String returns a string representation of the [InputSource]
func (input InputSource) String() string {
	switch input {
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
