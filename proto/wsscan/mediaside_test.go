// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// MediaSide tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
)

func TestMediaSideRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		ms   MediaSide
	}{
		{
			name: "Empty MediaSide",
			ms:   MediaSide{},
		},
		{
			name: "MediaSide with ColorProcessing only",
			ms: MediaSide{
				ColorProcessing: optional.New(ColorProcessing(
					ValWithOptions[ColorEntry]{
						Text: RGB24,
					},
				)),
			},
		},
		{
			name: "MediaSide with Resolution only",
			ms: MediaSide{
				Resolution: optional.New(Resolution{
					Height: ValWithOptions[int]{Text: 300},
					Width:  ValWithOptions[int]{Text: 300},
				}),
			},
		},
		{
			name: "MediaSide with ScanRegion only",
			ms: MediaSide{
				ScanRegion: optional.New(ScanRegion{
					ScanRegionHeight: ValWithOptions[int]{Text: 1000},
					ScanRegionWidth:  ValWithOptions[int]{Text: 800},
				}),
			},
		},
		{
			name: "MediaSide with all elements",
			ms: MediaSide{
				ColorProcessing: optional.New(ColorProcessing(
					ValWithOptions[ColorEntry]{
						Text:      Grayscale8,
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode to XML
			xml := tt.ms.toXML(NsWSCN + ":MediaFront")

			// Decode back
			decoded, err := decodeMediaSide(xml)
			if err != nil {
				t.Fatalf("decodeMediaSide() error = %v", err)
			}

			// Compare using reflect.DeepEqual
			if !reflect.DeepEqual(tt.ms, decoded) {
				t.Errorf("Round trip failed.\nOriginal: %+v\nDecoded:  %+v",
					tt.ms, decoded)
			}
		})
	}
}
