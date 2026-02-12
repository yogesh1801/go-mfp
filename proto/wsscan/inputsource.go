// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// InputSource: specifies the source from which the document should be scanned

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// InputSource represents the source from which a document should be scanned.
type InputSource ValWithOptions[InputSourceValue]

// InputSourceValue represents the possible values for InputSource
type InputSourceValue int

// Known InputSource values
const (
	UnknownInputSource InputSourceValue = iota
	InputSourceADF                      // Document feeding device,
	// scanning only the front side of each page
	InputSourceADFDuplex // Document feeding device,
	// scanning both sides of each page
	InputSourceFilm   // Film scanning option
	InputSourcePlaten // Scanner platen
)

// String returns the string representation of InputSourceValue
func (isv InputSourceValue) String() string {
	switch isv {
	case InputSourceADF:
		return "ADF"
	case InputSourceADFDuplex:
		return "ADFDuplex"
	case InputSourceFilm:
		return "Film"
	case InputSourcePlaten:
		return "Platen"
	}
	return "Unknown"
}

// DecodeInputSourceValue decodes InputSourceValue
// from its XML string representation
func DecodeInputSourceValue(s string) InputSourceValue {
	switch s {
	case "ADF":
		return InputSourceADF
	case "ADFDuplex":
		return InputSourceADFDuplex
	case "Film":
		return InputSourceFilm
	case "Platen":
		return InputSourcePlaten
	}
	return UnknownInputSource
}

// inputSourceValueDecoder is the decoder function for use with ValWithOptions
func inputSourceValueDecoder(s string) (InputSourceValue, error) {
	val := DecodeInputSourceValue(s)
	if val == UnknownInputSource {
		return val, fmt.Errorf("unknown InputSource value: %s", s)
	}
	return val, nil
}

// inputSourceValueEncoder is the encoder function for use with ValWithOptions
func inputSourceValueEncoder(isv InputSourceValue) string {
	return isv.String()
}

// decodeInputSource decodes InputSource from the XML tree
func decodeInputSource(root xmldoc.Element) (InputSource, error) {
	var base ValWithOptions[InputSourceValue]
	decoded, err := base.decodeValWithOptions(root,
		inputSourceValueDecoder)
	if err != nil {
		return InputSource{}, err
	}
	return InputSource(decoded), nil
}

// toXML generates XML tree for the InputSource
func (is InputSource) toXML(name string) xmldoc.Element {
	return ValWithOptions[InputSourceValue](is).toXML(name,
		inputSourceValueEncoder)
}
