// MFP - Miulti-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Film scan mode values

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// FilmScanMode defines the type of film being scanned, such as
// color slide film or black and white negative film.
type FilmScanMode int

// Known film scan modes:
const (
	UnknownFilmScanMode       FilmScanMode = iota
	NotApplicable                          // Default scan input source is not film
	ColorSlideFilm                         // Normal color space film images
	ColorNegativeFilm                      // Negative color space film images
	BlackandWhiteNegativeFilm              // Black and white negative film images
)

// decodeFilmScanMode decodes [FilmScanMode] from the XML tree.
func decodeFilmScanMode(root xmldoc.Element) (fsm FilmScanMode, err error) {
	return decodeEnum(root, DecodeFilmScanMode)
}

// toXML generates XML tree for the [FilmScanMode].
func (fsm FilmScanMode) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: fsm.String(),
	}
}

// String returns a string representation of the [FilmScanMode]
func (fsm FilmScanMode) String() string {
	switch fsm {
	case NotApplicable:
		return "NotApplicable"
	case ColorSlideFilm:
		return "ColorSlideFilm"
	case ColorNegativeFilm:
		return "ColorNegativeFilm"
	case BlackandWhiteNegativeFilm:
		return "BlackandWhiteNegativeFilm"
	}

	return "Unknown"
}

// DecodeFilmScanMode decodes [FilmScanMode] out of its XML string representation.
func DecodeFilmScanMode(s string) FilmScanMode {
	switch s {
	case "NotApplicable":
		return NotApplicable
	case "ColorSlideFilm":
		return ColorSlideFilm
	case "ColorNegativeFilm":
		return ColorNegativeFilm
	case "BlackandWhiteNegativeFilm":
		return BlackandWhiteNegativeFilm
	}

	return UnknownFilmScanMode
}
