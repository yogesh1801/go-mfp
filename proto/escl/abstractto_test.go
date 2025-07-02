// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// eSCL->abstract.Scanner conversions tests

package escl

import (
	"bytes"
	"testing"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/uuid"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
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

// TestScannerCapabilitiesToAbstract tests ScannerCapabilities.ToAbstract
func TestScannerCapabilitiesToAbstract(t *testing.T) {
	rd := bytes.NewReader(testutils.Kyocera.
		ECOSYS.M2040dn.ESCL.ScannerCapabilities)
	xml, err := xmldoc.Decode(NsMap, rd)
	assert.NoError(err)

	scancaps, err := DecodeScannerCapabilities(xml)
	assert.NoError(err)

	abscaps := scancaps.ToAbstract()
	expected := abstract.ScannerCapabilities{
		UUID:              uuid.MustParse("4509a320-00a0-008f-00b6-002507510eca"),
		MakeAndModel:      "Kyocera ECOSYS M2040dn",
		SerialNumber:      "VCF9192281",
		Manufacturer:      "",
		AdminURI:          "https://KM7B6A91.local/airprint",
		IconURI:           "https://KM7B6A91.local/printer-icon/machine_128.png",
		DocumentFormats:   []string{"image/jpeg", "application/pdf"},
		CompressionRange:  abstract.Range{Min: 1, Max: 5, Normal: 1, Step: 1},
		ADFCapacity:       75,
		BrightnessRange:   abstract.Range{},
		ContrastRange:     abstract.Range{},
		GammaRange:        abstract.Range{},
		HighlightRange:    abstract.Range{},
		NoiseRemovalRange: abstract.Range{},
		ShadowRange:       abstract.Range{},
		SharpenRange:      abstract.Range{Min: -3, Max: 3, Normal: 0, Step: 1},
		ThresholdRange:    abstract.Range{},
		Platen: &abstract.InputCapabilities{
			MinWidth:  999,
			MaxWidth:  21598,
			MinHeight: 999,
			MaxHeight: 29701,
			Intents: generic.MakeBitset(
				abstract.IntentDocument,
				abstract.IntentTextAndGraphic,
				abstract.IntentPhoto,
				abstract.IntentPreview,
			),
			Profiles: []abstract.SettingsProfile{
				{
					ColorModes: generic.MakeBitset(
						abstract.ColorModeBinary,
						abstract.ColorModeMono,
						abstract.ColorModeColor,
					),
					Depths: generic.MakeBitset(
						abstract.ColorDepth8,
					),
					Resolutions: []abstract.Resolution{
						{XResolution: 200, YResolution: 100},
						{XResolution: 200, YResolution: 200},
						{XResolution: 200, YResolution: 400},
						{XResolution: 300, YResolution: 300},
						{XResolution: 400, YResolution: 400},
						{XResolution: 600, YResolution: 600},
					},
				},
			},
		},
		ADFSimplex: &abstract.InputCapabilities{
			MinWidth:  5004,
			MaxWidth:  21598,
			MinHeight: 5004,
			MaxHeight: 35602,
			Intents: generic.MakeBitset(
				abstract.IntentDocument,
				abstract.IntentTextAndGraphic,
				abstract.IntentPhoto,
				abstract.IntentPreview,
			),
			Profiles: []abstract.SettingsProfile{
				{
					ColorModes: generic.MakeBitset(
						abstract.ColorModeBinary,
						abstract.ColorModeMono,
						abstract.ColorModeColor,
					),
					Depths: generic.MakeBitset(
						abstract.ColorDepth8,
					),
					Resolutions: []abstract.Resolution{
						{XResolution: 200, YResolution: 100},
						{XResolution: 200, YResolution: 200},
						{XResolution: 200, YResolution: 400},
						{XResolution: 300, YResolution: 300},
						{XResolution: 400, YResolution: 400},
						{XResolution: 600, YResolution: 600},
					},
				},
			},
		},
		ADFDuplex: &abstract.InputCapabilities{
			MinWidth:  5004,
			MaxWidth:  21598,
			MinHeight: 5004,
			MaxHeight: 35602,
			Intents: generic.MakeBitset(
				abstract.IntentDocument,
				abstract.IntentTextAndGraphic,
				abstract.IntentPhoto,
				abstract.IntentPreview,
			),
			Profiles: []abstract.SettingsProfile{
				{
					ColorModes: generic.MakeBitset(
						abstract.ColorModeBinary,
						abstract.ColorModeMono,
						abstract.ColorModeColor,
					),
					Depths: generic.MakeBitset(
						abstract.ColorDepth8,
					),
					Resolutions: []abstract.Resolution{
						{XResolution: 200, YResolution: 100},
						{XResolution: 200, YResolution: 200},
						{XResolution: 200, YResolution: 400},
						{XResolution: 300, YResolution: 300},
						{XResolution: 400, YResolution: 400},
						{XResolution: 600, YResolution: 600},
					},
				},
			},
		},
	}

	diff := testutils.Diff(abscaps, expected)
	if diff != "" {
		t.Errorf("ScannerCapabilities.toAbstract:\n%s", diff)
	}
}
