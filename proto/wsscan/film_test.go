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
		FilmColor: RGB24,
		FilmMaximumSize: Dimensions{
			Width:  100,
			Height: 200,
		},
		FilmMinimumSize: Dimensions{
			Width:  10,
			Height: 20,
		},
		FilmOpticalResolution: Dimensions{
			Width:  300,
			Height: 300,
		},
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

	if decoded.FilmMinimumSize.Width != 10 {
		t.Errorf("FilmMinimumSize width mismatch: got %d, want %d",
			decoded.FilmMinimumSize.Width, 10)
	}
	if decoded.FilmMinimumSize.Height != 20 {
		t.Errorf("FilmMinimumSize height mismatch: got %d, want %d",
			decoded.FilmMinimumSize.Height, 20)
	}

	if decoded.FilmOpticalResolution.Width != 300 {
		t.Errorf("FilmOpticalResolution width mismatch: got %d, want %d",
			decoded.FilmOpticalResolution.Width, 300)
	}
	if decoded.FilmOpticalResolution.Height != 300 {
		t.Errorf("FilmOpticalResolution height mismatch: got %d, want %d",
			decoded.FilmOpticalResolution.Height, 300)
	}

	// Compare resolutions
	if len(decoded.FilmResolutions.Widths) != len(film.FilmResolutions.Widths) ||
		len(decoded.FilmResolutions.Heights) != len(film.FilmResolutions.Heights) {
		t.Error("FilmResolutions length mismatch")
	} else {
		expectedWidths := []int{300, 600}
		for i, w := range decoded.FilmResolutions.Widths {
			if w != expectedWidths[i] {
				t.Errorf("FilmResolutions width[%d] mismatch: got %d, want %d",
					i, w, expectedWidths[i])
			}
		}

		expectedHeights := []int{300, 600}
		for i, h := range decoded.FilmResolutions.Heights {
			if h != expectedHeights[i] {
				t.Errorf("FilmResolutions height[%d] mismatch: got %d, want %d",
					i, h, expectedHeights[i])
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
			Width:  100,
			Height: 200,
		},
		FilmMinimumSize: Dimensions{
			Width:  10,
			Height: 20,
		},
		FilmOpticalResolution: Dimensions{
			Width:  300,
			Height: 300,
		},
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
