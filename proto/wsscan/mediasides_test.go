// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// MediaSides tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestMediaSidesRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		ms   MediaSides
	}{
		{
			name: "MediaSides with MediaFront only (empty)",
			ms: MediaSides{
				MediaFront: MediaSide{},
			},
		},
		{
			name: "MediaSides with MediaFront only (empty) and MustHonor",
			ms: MediaSides{
				MediaFront: MediaSide{},
				MustHonor:  optional.New(BooleanElement("true")),
			},
		},
		{
			name: "MediaSides with MediaFront with ColorProcessing",
			ms: MediaSides{
				MediaFront: MediaSide{
					ColorProcessing: optional.New(ColorProcessing(
						ValWithOptions[ColorEntry]{
							Text: RGB24,
						},
					)),
				},
			},
		},
		{
			name: "MediaSides with MediaFront and MediaBack",
			ms: MediaSides{
				MediaFront: MediaSide{},
				MediaBack: optional.New(MediaSide{
					Resolution: optional.New(Resolution{
						Height: ValWithOptions[int]{Text: 300},
						Width:  ValWithOptions[int]{Text: 300},
					}),
				}),
			},
		},
		{
			name: "MediaSides with both sides populated",
			ms: MediaSides{
				MediaFront: MediaSide{
					ColorProcessing: optional.New(ColorProcessing(
						ValWithOptions[ColorEntry]{
							Text: RGB24,
						},
					)),
					Resolution: optional.New(Resolution{
						Height: ValWithOptions[int]{Text: 600},
						Width:  ValWithOptions[int]{Text: 600},
					}),
				},
				MediaBack: optional.New(MediaSide{
					ColorProcessing: optional.New(ColorProcessing(
						ValWithOptions[ColorEntry]{
							Text: Grayscale8,
						},
					)),
					ScanRegion: optional.New(ScanRegion{
						ScanRegionHeight: ValWithOptions[int]{Text: 1000},
						ScanRegionWidth:  ValWithOptions[int]{Text: 800},
					}),
				}),
			},
		},
		{
			name: "MediaSides with all attributes and elements",
			ms: MediaSides{
				MustHonor: optional.New(BooleanElement("false")),
				MediaFront: MediaSide{
					ColorProcessing: optional.New(ColorProcessing(
						ValWithOptions[ColorEntry]{
							Text:      RGB24,
							MustHonor: optional.New(BooleanElement("true")),
						},
					)),
					Resolution: optional.New(Resolution{
						Height:    ValWithOptions[int]{Text: 600},
						Width:     ValWithOptions[int]{Text: 600},
						MustHonor: optional.New(BooleanElement("false")),
					}),
					ScanRegion: optional.New(ScanRegion{
						ScanRegionHeight: ValWithOptions[int]{Text: 2000},
						ScanRegionWidth:  ValWithOptions[int]{Text: 1500},
						ScanRegionXOffset: optional.New(
							ValWithOptions[int]{Text: 100}),
						ScanRegionYOffset: optional.New(
							ValWithOptions[int]{Text: 50}),
					}),
				},
				MediaBack: optional.New(MediaSide{
					ColorProcessing: optional.New(ColorProcessing(
						ValWithOptions[ColorEntry]{
							Text: Grayscale8,
						},
					)),
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode to XML
			xml := tt.ms.toXML(NsWSCN + ":MediaSides")

			// Decode back
			decoded, err := decodeMediaSides(xml)
			if err != nil {
				t.Fatalf("decodeMediaSides() error = %v", err)
			}

			// Compare using reflect.DeepEqual
			if !reflect.DeepEqual(tt.ms, decoded) {
				t.Errorf("Round trip failed.\nOriginal: %+v\nDecoded:  %+v",
					tt.ms, decoded)
			}
		})
	}
}

func TestMediaSidesInvalidMustHonor(t *testing.T) {
	tests := []struct {
		name      string
		mustHonor string
	}{
		{"invalid value", "invalid"},
		{"numeric 2", "2"},
		{"empty string", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elm := xmldoc.Element{
				Name: NsWSCN + ":MediaSides",
				Attrs: []xmldoc.Attr{
					{Name: NsWSCN + ":MustHonor", Value: tt.mustHonor},
				},
			}
			_, err := decodeMediaSides(elm)
			if err == nil {
				t.Error("Expected error for invalid MustHonor, got nil")
			}
		})
	}
}

func TestMediaSidesMissingMediaFront(t *testing.T) {
	// MediaFront is required, so decoding should fail without it
	elm := xmldoc.Element{
		Name: NsWSCN + ":MediaSides",
	}
	_, err := decodeMediaSides(elm)
	if err == nil {
		t.Error("Expected error for missing MediaFront, got nil")
	}
}
