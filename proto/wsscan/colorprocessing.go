// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ColorProcessing: specifies the color-processing mode of the input source

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ColorProcessing specifies the color-processing mode of the
// input source on the scanner.
// The text value is a ColorEntry (e.g., BlackAndWhite1, Grayscale8, RGB24 etc.)
type ColorProcessing ValWithOptions[ColorEntry]

// colorProcessingDecoder is the decoder function for use with ValWithOptions
func colorProcessingDecoder(s string) (ColorEntry, error) {
	val := DecodeColorEntry(s)
	if val == UnknownColorEntry {
		return val, fmt.Errorf("unknown ColorProcessing value: %s", s)
	}
	return val, nil
}

// colorProcessingEncoder is the encoder function for use with ValWithOptions
func colorProcessingEncoder(ce ColorEntry) string {
	return ce.String()
}

// decodeColorProcessing decodes ColorProcessing from the XML tree
func decodeColorProcessing(root xmldoc.Element) (ColorProcessing, error) {
	var base ValWithOptions[ColorEntry]
	decoded, err := base.decodeValWithOptions(root, colorProcessingDecoder)
	if err != nil {
		return ColorProcessing{}, err
	}
	return ColorProcessing(decoded), nil
}

// toXML generates XML tree for the ColorProcessing
func (cp ColorProcessing) toXML(name string) xmldoc.Element {
	return ValWithOptions[ColorEntry](cp).toXML(name, colorProcessingEncoder)
}
