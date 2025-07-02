// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL->abstract.Scanner conversions tests

package escl

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/optional"
)

// TestScanSettingsToAbstract tests ScanSettings.toAbstract method
func TestScanSettingsToAbstract(t *testing.T) {
	type testData struct {
		comment string
		ss      ScanSettings
		out     abstract.ScannerRequest
	}

	tests := []testData{
		// Empty request
		{
			comment: "Bare minimum",
			ss: ScanSettings{
				Version: DefaultVersion,
			},
			out: abstract.ScannerRequest{},
		},

		// Input + ADFMode
		{
			comment: "InputPlaten",
			ss: ScanSettings{
				Version:     DefaultVersion,
				InputSource: optional.New(InputPlaten),
			},
			out: abstract.ScannerRequest{
				Input: abstract.InputPlaten,
			},
		},

		{
			comment: "InputFeeder,simplex",
			ss: ScanSettings{
				Version:     DefaultVersion,
				InputSource: optional.New(InputFeeder),
			},
			out: abstract.ScannerRequest{
				Input:   abstract.InputADF,
				ADFMode: abstract.ADFModeSimplex,
			},
		},

		{
			comment: "InputFeeder,duplex",
			ss: ScanSettings{
				Version:     DefaultVersion,
				InputSource: optional.New(InputFeeder),
				Duplex:      optional.New(true),
			},
			out: abstract.ScannerRequest{
				Input:   abstract.InputADF,
				ADFMode: abstract.ADFModeDuplex,
			},
		},

		// ColorMode, ColorDepth, BinaryRendering and Threshold
		{
			comment: "BlackAndWhite1",
			ss: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(BlackAndWhite1),
				Threshold: optional.New(50), // Ignored
			},
			out: abstract.ScannerRequest{
				ColorMode: abstract.ColorModeBinary,
			},
		},

		{
			comment: "BlackAndWhite1+Halftone",
			ss: ScanSettings{
				Version:         DefaultVersion,
				ColorMode:       optional.New(BlackAndWhite1),
				BinaryRendering: optional.New(Halftone),
				Threshold:       optional.New(50), // Ignored
			},
			out: abstract.ScannerRequest{
				ColorMode:       abstract.ColorModeBinary,
				BinaryRendering: abstract.BinaryRenderingHalftone,
			},
		},

		{
			comment: "BlackAndWhite1+Threshold",
			ss: ScanSettings{
				Version:         DefaultVersion,
				ColorMode:       optional.New(BlackAndWhite1),
				BinaryRendering: optional.New(Threshold),
			},
			out: abstract.ScannerRequest{
				ColorMode:       abstract.ColorModeBinary,
				BinaryRendering: abstract.BinaryRenderingThreshold,
			},
		},

		{
			comment: "BlackAndWhite1+Threshold+Threshold=50",
			ss: ScanSettings{
				Version:         DefaultVersion,
				ColorMode:       optional.New(BlackAndWhite1),
				BinaryRendering: optional.New(Threshold),
				Threshold:       optional.New(50),
			},
			out: abstract.ScannerRequest{
				ColorMode:       abstract.ColorModeBinary,
				BinaryRendering: abstract.BinaryRenderingThreshold,
				Threshold:       optional.New(50),
			},
		},

		{
			comment: "Grayscale8",
			ss: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(Grayscale8),
			},
			out: abstract.ScannerRequest{
				ColorMode:  abstract.ColorModeMono,
				ColorDepth: abstract.ColorDepth8,
			},
		},

		{
			comment: "Grayscale16",
			ss: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(Grayscale16),
			},
			out: abstract.ScannerRequest{
				ColorMode:  abstract.ColorModeMono,
				ColorDepth: abstract.ColorDepth16,
			},
		},

		{
			comment: "RGB24",
			ss: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(RGB24),
			},
			out: abstract.ScannerRequest{
				ColorMode:  abstract.ColorModeColor,
				ColorDepth: abstract.ColorDepth8,
			},
		},

		{
			comment: "RGB48",
			ss: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(RGB48),
			},
			out: abstract.ScannerRequest{
				ColorMode:  abstract.ColorModeColor,
				ColorDepth: abstract.ColorDepth16,
			},
		},

		// CCDChannel
		{
			comment: "CCDChannel: Red",
			ss: ScanSettings{
				Version:    DefaultVersion,
				CCDChannel: optional.New(Red),
			},
			out: abstract.ScannerRequest{
				CCDChannel: abstract.CCDChannelRed,
			},
		},

		{
			comment: "CCDChannel: Green",
			ss: ScanSettings{
				Version:    DefaultVersion,
				CCDChannel: optional.New(Green),
			},
			out: abstract.ScannerRequest{
				CCDChannel: abstract.CCDChannelGreen,
			},
		},

		{
			comment: "CCDChannel: Blue",
			ss: ScanSettings{
				Version:    DefaultVersion,
				CCDChannel: optional.New(Blue),
			},
			out: abstract.ScannerRequest{
				CCDChannel: abstract.CCDChannelBlue,
			},
		},

		{
			comment: "CCDChannel: NTSC",
			ss: ScanSettings{
				Version:    DefaultVersion,
				CCDChannel: optional.New(NTSC),
			},
			out: abstract.ScannerRequest{
				CCDChannel: abstract.CCDChannelNTSC,
			},
		},

		{
			comment: "CCDChannel: GrayCcd",
			ss: ScanSettings{
				Version:    DefaultVersion,
				CCDChannel: optional.New(GrayCcd),
			},
			out: abstract.ScannerRequest{
				CCDChannel: abstract.CCDChannelGrayCcd,
			},
		},

		{
			comment: "CCDChannel: GrayCcdEmulated",
			ss: ScanSettings{
				Version:    DefaultVersion,
				CCDChannel: optional.New(GrayCcdEmulated),
			},
			out: abstract.ScannerRequest{
				CCDChannel: abstract.CCDChannelGrayCcdEmulated,
			},
		},

		{
			comment: "CCDChannel: Unknown",
			ss: ScanSettings{
				Version:    DefaultVersion,
				CCDChannel: optional.New(CCDChannel(-1)),
			},
			out: abstract.ScannerRequest{
				CCDChannel: abstract.CCDChannelUnset,
			},
		},

		// DocumentFormat/DocumentFormatExt
		{
			comment: "DocumentFormat",
			ss: ScanSettings{
				Version: DefaultVersion,
				DocumentFormat: optional.New(
					"image/jpeg"),
			},
			out: abstract.ScannerRequest{
				DocumentFormat: "image/jpeg",
			},
		},

		{
			comment: "DocumentFormatExt",
			ss: ScanSettings{
				Version: DefaultVersion,
				DocumentFormatExt: optional.New(
					"application/pdf"),
			},
			out: abstract.ScannerRequest{
				DocumentFormat: "application/pdf",
			},
		},

		{
			comment: "DocumentFormat+DocumentFormatExt",
			ss: ScanSettings{
				Version: DefaultVersion,
				DocumentFormat: optional.New(
					"image/jpeg"),
				DocumentFormatExt: optional.New(
					"application/pdf"),
			},
			out: abstract.ScannerRequest{
				DocumentFormat: "application/pdf",
			},
		},

		// ScanRegions
		{
			comment: "ScanRegions",
			ss: ScanSettings{
				Version: DefaultVersion,
				ScanRegions: []ScanRegion{
					{
						XOffset:            150,
						YOffset:            300,
						Width:              5 * 300,
						Height:             7 * 300,
						ContentRegionUnits: ThreeHundredthsOfInches,
					},
				},
			},
			out: abstract.ScannerRequest{
				Region: abstract.Region{
					XOffset: abstract.Inch / 2,
					YOffset: abstract.Inch,
					Width:   5 * abstract.Inch,
					Height:  7 * abstract.Inch,
				},
			},
		},

		// Resolution
		{
			comment: "Resolution",
			ss: ScanSettings{
				Version:     DefaultVersion,
				XResolution: optional.New(300),
				YResolution: optional.New(600),
			},
			out: abstract.ScannerRequest{
				Resolution: abstract.Resolution{
					XResolution: 300,
					YResolution: 600,
				},
			},
		},

		// Intent
		{
			comment: "Intent: Document",
			ss: ScanSettings{
				Version: DefaultVersion,
				Intent:  optional.New(Document),
			},
			out: abstract.ScannerRequest{
				Intent: abstract.IntentDocument,
			},
		},

		{
			comment: "Intent: TextAndGraphic",
			ss: ScanSettings{
				Version: DefaultVersion,
				Intent:  optional.New(TextAndGraphic),
			},
			out: abstract.ScannerRequest{
				Intent: abstract.IntentTextAndGraphic,
			},
		},

		{
			comment: "Intent: Photo",
			ss: ScanSettings{
				Version: DefaultVersion,
				Intent:  optional.New(Photo),
			},
			out: abstract.ScannerRequest{
				Intent: abstract.IntentPhoto,
			},
		},

		{
			comment: "Intent: Preview",
			ss: ScanSettings{
				Version: DefaultVersion,
				Intent:  optional.New(Preview),
			},
			out: abstract.ScannerRequest{
				Intent: abstract.IntentPreview,
			},
		},

		{
			comment: "Intent: Object",
			ss: ScanSettings{
				Version: DefaultVersion,
				Intent:  optional.New(Object),
			},
			out: abstract.ScannerRequest{
				Intent: abstract.IntentObject,
			},
		},

		{
			comment: "Intent: BusinessCard",
			ss: ScanSettings{
				Version: DefaultVersion,
				Intent:  optional.New(BusinessCard),
			},
			out: abstract.ScannerRequest{
				Intent: abstract.IntentBusinessCard,
			},
		},

		{
			comment: "Intent: Unknown",
			ss: ScanSettings{
				Version: DefaultVersion,
				Intent:  optional.New(Intent(-1)),
			},
			out: abstract.ScannerRequest{
				Intent: abstract.IntentUnset,
			},
		},
	}

	for _, test := range tests {
		out := test.ss.ToAbstract()

		testutils.CheckConvertionTest(t,
			"ScanSettings.toAbstract",
			test.comment, test.out, out)
	}
}
