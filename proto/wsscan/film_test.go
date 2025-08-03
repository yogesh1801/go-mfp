// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol - unit tests
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions

package wsscan

import (
	"testing"
)

// Test Film XML encoding/decoding
func TestFilmXML(t *testing.T) {
	// Create a test Film structure with all fields populated
	film := Film{
		FilmColor:             RGB24,
		FilmMaximumSize:       Dimension{Width: 100, Height: 200},
		FilmMinimumSize:       Dimension{Width: 10, Height: 20},
		FilmOpticalResolution: Dimension{Width: 300, Height: 300},
		FilmResolutions: Resolutions{
			Widths:  []int{300, 600},
			Heights: []int{300, 600},
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

	if decoded.FilmMaximumSize != film.FilmMaximumSize {
		t.Error("FilmMaximumSize mismatch")
	}

	if decoded.FilmMinimumSize != film.FilmMinimumSize {
		t.Error("FilmMinimumSize mismatch")
	}

	if decoded.FilmOpticalResolution != film.FilmOpticalResolution {
		t.Error("FilmOpticalResolution mismatch")
	}

	// Compare resolutions
	if len(decoded.FilmResolutions.Widths) != len(film.FilmResolutions.Widths) ||
		len(decoded.FilmResolutions.Heights) != len(film.FilmResolutions.Heights) {
		t.Error("FilmResolutions length mismatch")
	} else {
		for i := range decoded.FilmResolutions.Widths {
			if decoded.FilmResolutions.Widths[i] != film.FilmResolutions.Widths[i] {
				t.Errorf("FilmResolutions width[%d] mismatch", i)
			}
		}
		for i := range decoded.FilmResolutions.Heights {
			if decoded.FilmResolutions.Heights[i] != film.FilmResolutions.Heights[i] {
				t.Errorf("FilmResolutions height[%d] mismatch", i)
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
		FilmColor:             RGB24,
		FilmMaximumSize:       Dimension{Width: 100, Height: 200},
		FilmMinimumSize:       Dimension{Width: 10, Height: 20},
		FilmOpticalResolution: Dimension{Width: 300, Height: 300},
		FilmResolutions: Resolutions{
			Widths:  []int{300},
			Heights: []int{300},
		},
		FilmScanModesSupported: []FilmScanMode{},
	}

	encoded = film.toXML("Film")
	_, err = decodeFilm(encoded)
	if err == nil {
		t.Error("Expected error when decoding Film with empty FilmScanModesSupported")
	}
}
