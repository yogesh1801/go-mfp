// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol - unit tests
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions

package wsscan

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
)

// Test Film XML encoding/decoding
func TestFilmXML(t *testing.T) {
	// Create a test Film structure with all fields populated
	film := Film{
		FilmColor: RGB24,
		FilmMaximumSize: Dimensions{
			Width:  TextWithOverrideAndDefault{Text: "100"},
			Height: TextWithOverrideAndDefault{Text: "200"},
		},
		FilmMinimumSize: Dimensions{
			Width:  TextWithOverrideAndDefault{Text: "10"},
			Height: TextWithOverrideAndDefault{Text: "20"},
		},
		FilmOpticalResolution: Dimensions{
			Width:  TextWithOverrideAndDefault{Text: "300"},
			Height: TextWithOverrideAndDefault{Text: "300"},
		},
		FilmResolutions: Resolutions{
			Widths: []TextWithOverrideAndDefault{
				{
					Text:        "300",
					Override:    optional.New(BooleanElement("true")),
					UsedDefault: optional.New(BooleanElement("false")),
				},
				{
					Text:        "600",
					Override:    optional.New(BooleanElement("false")),
					UsedDefault: optional.New(BooleanElement("true")),
				},
			},
			Heights: []TextWithOverrideAndDefault{
				{
					Text:        "300",
					Override:    optional.New(BooleanElement("true")),
					UsedDefault: optional.New(BooleanElement("false")),
				},
				{
					Text:        "600",
					Override:    optional.New(BooleanElement("false")),
					UsedDefault: optional.New(BooleanElement("true")),
				},
			},
		},
		FilmScanModesSupported: []FilmScanMode{
			ColorSlideFilm,
			ColorNegativeFilm,
			BlackandWhiteNegativeFilm,
		},
	}

	// Convert to XML
	encoded := film.toXML("Film")

	// Verify structure
	if encoded.Name != "Film" {
		t.Errorf("Expected root element name 'Film', got %q", encoded.Name)
	}

	// Decode back
	decoded, err := decodeFilm(encoded)
	if err != nil {
		t.Errorf("Failed to decode Film: %v", err)
	}

	// Compare original and decoded
	if decoded.FilmColor != film.FilmColor {
		t.Error("FilmColor mismatch")
	}

	if decoded.FilmMinimumSize.Width.Text != "10" {
		t.Error("FilmMinimumSize width mismatch")
	}
	if decoded.FilmMinimumSize.Height.Text != "20" {
		t.Error("FilmMinimumSize height mismatch")
	}

	if decoded.FilmOpticalResolution.Width.Text != "300" {
		t.Error("FilmOpticalResolution width mismatch")
	}
	if decoded.FilmOpticalResolution.Height.Text != "300" {
		t.Error("FilmOpticalResolution height mismatch")
	}

	// Compare resolutions
	if len(decoded.FilmResolutions.Widths) != len(film.FilmResolutions.Widths) ||
		len(decoded.FilmResolutions.Heights) != len(film.FilmResolutions.Heights) {
		t.Error("FilmResolutions length mismatch")
	} else {
		expectedWidths := []string{"300", "600"}
		for i, w := range decoded.FilmResolutions.Widths {
			if w.Text != expectedWidths[i] {
				t.Errorf("FilmResolutions width[%d] mismatch: got %s, want %s",
					i, w.Text, expectedWidths[i])
			}
		}

		// Check Override and UsedDefault attributes
		for i, width := range decoded.FilmResolutions.Widths {
			if optional.Get(width.Override) != optional.Get(film.FilmResolutions.Widths[i].Override) {
				t.Errorf("FilmResolutions width[%d] Override mismatch", i)
			}
			if optional.Get(width.UsedDefault) != optional.Get(film.FilmResolutions.Widths[i].UsedDefault) {
				t.Errorf("FilmResolutions width[%d] UsedDefault mismatch", i)
			}
		}

		expectedHeights := []string{"300", "600"}
		for i, h := range decoded.FilmResolutions.Heights {
			if h.Text != expectedHeights[i] {
				t.Errorf("FilmResolutions height[%d] mismatch: got %s, want %s",
					i, h.Text, expectedHeights[i])
			}
		}

		// Check Override and UsedDefault attributes
		for i, height := range decoded.FilmResolutions.Heights {
			if optional.Get(height.Override) != optional.Get(film.FilmResolutions.Heights[i].Override) {
				t.Errorf("FilmResolutions height[%d] Override mismatch", i)
			}
			if optional.Get(height.UsedDefault) != optional.Get(film.FilmResolutions.Heights[i].UsedDefault) {
				t.Errorf("FilmResolutions height[%d] UsedDefault mismatch", i)
			}
		}
	}

	if len(decoded.FilmScanModesSupported) != len(film.FilmScanModesSupported) {
		t.Errorf("FilmScanModesSupported length mismatch: got %d, want %d",
			len(decoded.FilmScanModesSupported), len(film.FilmScanModesSupported))
	} else {
		for i := range decoded.FilmScanModesSupported {
			if decoded.FilmScanModesSupported[i] != film.FilmScanModesSupported[i] {
				t.Errorf("FilmScanModesSupported[%d] mismatch: got %v, want %v",
					i, decoded.FilmScanModesSupported[i], film.FilmScanModesSupported[i])
			}
		}
	}
}

// Test Film with missing required fields
func TestFilmMissingFields(t *testing.T) {
	// Create a minimal Film structure with missing required fields
	film := Film{
		FilmScanModesSupported: []FilmScanMode{ColorSlideFilm},
	}

	// Convert to XML and try to decode
	encoded := film.toXML("Film")
	_, err := decodeFilm(encoded)
	if err == nil {
		t.Error("Expected error when decoding Film with missing required fields")
	}

	// Test with empty FilmScanModesSupported
	film = Film{
		FilmColor: RGB24,
		FilmMaximumSize: Dimensions{
			Width:  TextWithOverrideAndDefault{Text: "100"},
			Height: TextWithOverrideAndDefault{Text: "200"},
		},
		FilmMinimumSize: Dimensions{
			Width:  TextWithOverrideAndDefault{Text: "10"},
			Height: TextWithOverrideAndDefault{Text: "20"},
		},
		FilmOpticalResolution: Dimensions{
			Width:  TextWithOverrideAndDefault{Text: "300"},
			Height: TextWithOverrideAndDefault{Text: "300"},
		},
		FilmResolutions: Resolutions{
			Widths:  []TextWithOverrideAndDefault{{Text: "300"}},
			Heights: []TextWithOverrideAndDefault{{Text: "300"}},
		},
		FilmScanModesSupported: []FilmScanMode{},
	}

	encoded = film.toXML("Film")
	_, err = decodeFilm(encoded)
	if err == nil {
		t.Error("Expected error when decoding Film with empty FilmScanModesSupported")
	}
}
