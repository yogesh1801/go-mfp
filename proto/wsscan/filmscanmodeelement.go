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
// The optional attributes MustHonor, Override, and UsedDefault are Boolean values.
// The text value can be extended and subset.
type FilmScanModeElement struct {
	TextWithBoolAttrs[string]
}

// decodeFilmScanModeElement decodes a FilmScanModeElement from an XML element.
func decodeFilmScanModeElement(root xmldoc.Element) (FilmScanModeElement, error) {
	var fsm FilmScanModeElement
	decoded, err := fsm.TextWithBoolAttrs.decodeTextWithBoolAttrs(root, func(s string) (string, error) {
		return s, nil
	})
	if err != nil {
		return fsm, err
	}
	fsm.TextWithBoolAttrs = decoded
	return fsm, nil
}

// toXML converts a FilmScanModeElement to an XML element.
func (fsm FilmScanModeElement) toXML(name string) xmldoc.Element {
	return fsm.TextWithBoolAttrs.toXML(name, func(s string) string {
		return s
	})
}
