// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ScannerConfiguration

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func createValidPlaten() Platen {
	return Platen{
		PlatenColor: []ColorEntry{BlackAndWhite1, RGB24},
		PlatenMaximumSize: Dimensions{
			Width:  210,
			Height: 297,
		},
		PlatenMinimumSize: Dimensions{
			Width:  50,
			Height: 100,
		},
		PlatenOpticalResolution: Dimensions{
			Width:  600,
			Height: 600,
		},
		PlatenResolutions: Resolutions{
			Widths:  []int{300, 600},
			Heights: []int{300, 600},
		},
	}
}

func createValidFilm() Film {
	return Film{
		FilmColor:             BlackAndWhite1,
		FilmMaximumSize:       Dimensions{Width: 35, Height: 135},
		FilmMinimumSize:       Dimensions{Width: 24, Height: 36},
		FilmOpticalResolution: Dimensions{Width: 2400, Height: 2400},
		FilmResolutions: Resolutions{
			Widths:  []int{1200, 2400},
			Heights: []int{1200, 2400},
		},
		FilmScanModesSupported: []FilmScanMode{ColorSlideFilm},
	}
}

// TestScannerConfiguration_RoundTrip_AllChildren verifies that a
// ScannerConfiguration with all optional child elements (ADF, Film, Platen)
// present encodes to XML and decodes back to an identical value.
func TestScannerConfiguration_RoundTrip_AllChildren(t *testing.T) {
	orig := ScannerConfiguration{
		ADF: optional.New(ADF{
			ADFSupportsDuplex: BooleanElement("true")}),
		DeviceSettings: createValidDeviceSettings(),
		Film:           optional.New(createValidFilm()),
		Platen:         optional.New(createValidPlaten()),
	}
	elm := orig.toXML(NsWSCN + ":ScannerConfiguration")
	if elm.Name != NsWSCN+":ScannerConfiguration" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":ScannerConfiguration", elm.Name)
	}

	parsed, err := decodeScannerConfiguration(elm)
	if err != nil {
		t.Fatalf("decodeScannerConfiguration returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScannerConfiguration_RoundTrip_DeviceSettingsOnly verifies that a
// ScannerConfiguration with only the required DeviceSettings child (no ADF,
// Film, or Platen) encodes to XML and decodes back to an identical value.
func TestScannerConfiguration_RoundTrip_DeviceSettingsOnly(t *testing.T) {
	orig := ScannerConfiguration{
		DeviceSettings: createValidDeviceSettings(),
	}
	elm := orig.toXML(NsWSCN + ":ScannerConfiguration")

	parsed, err := decodeScannerConfiguration(elm)
	if err != nil {
		t.Fatalf("decodeScannerConfiguration returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScannerConfiguration_MissingDeviceSettings verifies that decoding a
// ScannerConfiguration element without the required DeviceSettings child
// returns an error.
func TestScannerConfiguration_MissingDeviceSettings(t *testing.T) {
	orig := ScannerConfiguration{
		ADF:    optional.New(ADF{ADFSupportsDuplex: BooleanElement("true")}),
		Platen: optional.New(createValidPlaten()),
	}
	elm := orig.toXML(NsWSCN + ":ScannerConfiguration")
	// Remove DeviceSettings child
	var filtered []xmldoc.Element
	for _, child := range elm.Children {
		if child.Name != NsWSCN+":DeviceSettings" {
			filtered = append(filtered, child)
		}
	}
	elm.Children = filtered

	_, err := decodeScannerConfiguration(elm)
	if err == nil {
		t.Error("expected error for missing DeviceSettings, got nil")
	}
}
