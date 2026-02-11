// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// FilmScanModeElement: specifies the exposure type of the film to be scanned

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// FilmScanModeElement specifies the exposure type of the film to be scanned.
// The text value can be extended and subset.
type FilmScanModeElement ValWithOptions[string]

// decodeFilmScanModeElement decodes a FilmScanModeElement from an XML element.
func decodeFilmScanModeElement(root xmldoc.Element) (
	FilmScanModeElement, error) {
	var base ValWithOptions[string]
	decoded, err := base.decodeValWithOptions(root, stringValueDecoder)
	if err != nil {
		return FilmScanModeElement{}, err
	}
	return FilmScanModeElement(decoded), nil
}

// toXML converts a FilmScanModeElement to an XML element.
func (fsm FilmScanModeElement) toXML(name string) xmldoc.Element {
	return ValWithOptions[string](fsm).toXML(name, stringValueEncoder)
}
