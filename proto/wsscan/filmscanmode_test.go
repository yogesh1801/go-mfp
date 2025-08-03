// MFP - Miulti-Function Printers and scanners toolkit
// WS-Scan core protocol - unit tests
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions

package wsscan

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Test FilmScanMode.String()
func TestFilmScanModeString(t *testing.T) {
	tests := []struct {
		fsm      FilmScanMode
		expected string
	}{
		{UnknownFilmScanMode, "Unknown"},
		{NotApplicable, "NotApplicable"},
		{ColorSlideFilm, "ColorSlideFilm"},
		{ColorNegativeFilm, "ColorNegativeFilm"},
		{BlackandWhiteNegativeFilm, "BlackandWhiteNegativeFilm"},
	}

	for _, test := range tests {
		if s := test.fsm.String(); s != test.expected {
			t.Errorf("%v.String(): expected %q, got %q",
				test.fsm, test.expected, s)
		}
	}
}

// Test DecodeFilmScanMode()
func TestDecodeFilmScanMode(t *testing.T) {
	tests := []struct {
		str      string
		expected FilmScanMode
	}{
		{"NotApplicable", NotApplicable},
		{"ColorSlideFilm", ColorSlideFilm},
		{"ColorNegativeFilm", ColorNegativeFilm},
		{"BlackandWhiteNegativeFilm", BlackandWhiteNegativeFilm},
		{"Undefined", UnknownFilmScanMode},
		{"", UnknownFilmScanMode},
	}

	for _, test := range tests {
		if fsm := DecodeFilmScanMode(test.str); fsm != test.expected {
			t.Errorf("DecodeFilmScanMode(%q): expected %v, got %v",
				test.str, test.expected, fsm)
		}
	}
}

// Test FilmScanMode XML encoding/decoding
func TestFilmScanModeXML(t *testing.T) {
	tests := []struct {
		fsm      FilmScanMode
		expected string
	}{
		{NotApplicable, "NotApplicable"},
		{ColorSlideFilm, "ColorSlideFilm"},
		{ColorNegativeFilm, "ColorNegativeFilm"},
		{BlackandWhiteNegativeFilm, "BlackandWhiteNegativeFilm"},
	}

	const elementName = "FilmScanModeValue"

	for _, test := range tests {
		// Test encoding
		encoded := test.fsm.toXML(elementName)
		if encoded.Text != test.expected {
			t.Errorf("%v.toXML(): expected text %q, got %q",
				test.fsm, test.expected, encoded.Text)
		}

		// Test decoding
		decoded, err := decodeFilmScanMode(xmldoc.Element{
			Name: elementName,
			Text: test.expected,
		})

		if err != nil {
			t.Errorf("decodeFilmScanMode(%q): %s", test.expected, err)
		} else if decoded != test.fsm {
			t.Errorf("decodeFilmScanMode(%q): expected %v, got %v",
				test.expected, test.fsm, decoded)
		}
	}
}
