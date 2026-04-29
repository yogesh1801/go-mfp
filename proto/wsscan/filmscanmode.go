// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Film scan mode values

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// FilmScanMode defines the type of film being scanned, such as
// color slide film or black and white negative film.
//
// It is backed by string so that it can represent both the standard values
// defined in the WS-Scan specification and vendor-specific extensions that
// may appear in the wild.
type FilmScanMode string

// Known film scan modes. Additional vendor-defined values may exist; those
// are represented by arbitrary [FilmScanMode] values constructed from their
// string form.
const (
	UnknownFilmScanMode       FilmScanMode = ""
	NotApplicable             FilmScanMode = "NotApplicable"
	ColorSlideFilm            FilmScanMode = "ColorSlideFilm"
	ColorNegativeFilm         FilmScanMode = "ColorNegativeFilm"
	BlackandWhiteNegativeFilm FilmScanMode = "BlackandWhiteNegativeFilm"
)

// decodeFilmScanMode decodes [FilmScanMode] from the XML tree.
// Empty values are rejected; any other value is preserved as-is, allowing
// vendor extensions to round-trip.
func decodeFilmScanMode(root xmldoc.Element) (fsm FilmScanMode, err error) {
	if root.Text == "" {
		err = fmt.Errorf("invalid FilmScanMode: empty")
		err = xmldoc.XMLErrWrap(root, err)
		return UnknownFilmScanMode, err
	}
	return FilmScanMode(root.Text), nil
}

// toXML generates XML tree for the [FilmScanMode].
func (fsm FilmScanMode) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: string(fsm),
	}
}

// String returns a string representation of the [FilmScanMode].
func (fsm FilmScanMode) String() string {
	return string(fsm)
}

// DecodeFilmScanMode decodes [FilmScanMode] from its XML string representation.
// Vendor-defined values are preserved verbatim; only the empty string is mapped
// to [UnknownFilmScanMode].
func DecodeFilmScanMode(s string) FilmScanMode {
	return FilmScanMode(s)
}

// filmScanModeDecoder is the decoder function for use with ValWithOptions.
// It rejects only the empty string; vendor extensions are preserved.
func filmScanModeDecoder(s string) (FilmScanMode, error) {
	if s == "" {
		return UnknownFilmScanMode,
			fmt.Errorf("invalid FilmScanMode: empty")
	}
	return FilmScanMode(s), nil
}

// filmScanModeEncoder is the encoder function for use with ValWithOptions.
func filmScanModeEncoder(fsm FilmScanMode) string {
	return string(fsm)
}
