// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan request tests

package abstract

import (
	"slices"
	"testing"

	"github.com/alexpevzner/mfp/internal/testutils"
	"github.com/alexpevzner/mfp/util/generic"
	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/uuid"
)

// testUUID contains a parsed UUID
var testUUID = uuid.Must(
	uuid.Parse("418b75ab-1bd7-4d01-8178-75a84450a11c"),
)

// testIntents contains initialized set of Intents
var testIntents = generic.MakeBitset(
	IntentDocument,
	IntentTextAndGraphic,
	IntentPhoto,
	IntentPreview,
)

// testColorModes contains initialized ColorMode set
var testColorModes = generic.MakeBitset(
	ColorModeBinary,
	ColorModeMono,
	ColorModeColor,
)

// testDepth contains initialized abstract.Depth set
var testDepth = generic.MakeBitset(ColorDepth8)

// testBinaryRenderings contains initialized BinaryRendering set
var testBinaryRenderings = generic.MakeBitset(
	BinaryRenderingHalftone,
	BinaryRenderingThreshold,
)

// testCCDChannels contains initialized CCDChannel set
var testCCDChannels = generic.MakeBitset(
	CCDChannelRed,
	CCDChannelGreen,
	CCDChannelBlue,
	CCDChannelNTSC,
)

// testResolutions contains initialized []Resolution slice
var testResolutions = []Resolution{
	{XResolution: 200, YResolution: 100},
	{XResolution: 200, YResolution: 200},
	{XResolution: 200, YResolution: 400},
	{XResolution: 300, YResolution: 300},
	{XResolution: 400, YResolution: 400},
	{XResolution: 600, YResolution: 600},
}

// testSettingsProfiles contains initialized []SettingsProfile slice
var testSettingsProfiles = []SettingsProfile{
	{
		ColorModes:       testColorModes,
		Depths:           testDepth,
		BinaryRenderings: testBinaryRenderings,
		CCDChannels:      testCCDChannels,
		Resolutions:      testResolutions,
	},
}

// testSettingsProfilesHiRes contains initialized []SettingsProfile slice
// with extra resolutions
var testSettingsProfilesHiRes = []SettingsProfile{
	{
		ColorModes:       testColorModes,
		Depths:           testDepth,
		BinaryRenderings: testBinaryRenderings,
		CCDChannels:      testCCDChannels,
		Resolutions: append(testResolutions,
			Resolution{1200, 1200}),
	},
	{
		ColorModes: generic.MakeBitset(ColorModeMono),
		Depths:     generic.MakeBitset(ColorDepth8),
		Resolutions: append(testResolutions,
			Resolution{2400, 2400}),
	},
}

// testPlatenInputCapabilities contains InputCapabilities for
// the Platen source
var testPlatenInputCapabilities = &InputCapabilities{
	MinWidth:  DimensionFromDots(300, 118),
	MinHeight: DimensionFromDots(300, 118),
	MaxWidth:  DimensionFromDots(300, 2551),
	MaxHeight: DimensionFromDots(300, 3508),
	Intents:   testIntents,
	Profiles:  testSettingsProfilesHiRes,
}

// testADFenInputCapabilities contains InputCapabilities for
// the ADFen source
var testADFenInputCapabilities = &InputCapabilities{
	MinWidth:  DimensionFromDots(300, 591),
	MinHeight: DimensionFromDots(300, 591),
	MaxWidth:  DimensionFromDots(300, 2551),
	MaxHeight: DimensionFromDots(300, 4205),
	Intents:   testIntents,
	Profiles:  testSettingsProfiles,
}

// testScannerCapabilities contains initialized ScannerCapabilities
// structure
var testScannerCapabilities = &ScannerCapabilities{
	UUID:              testUUID,
	MakeAndModel:      "Abstract Scanner",
	SerialNumber:      "AS-12345",
	Manufacturer:      "Abstract Corp.",
	CompressionRange:  Range{Min: 1, Max: 5, Normal: 1},
	ADFCapacity:       75,
	BrightnessRange:   Range{Min: -100, Max: 100, Normal: 0},
	ContrastRange:     Range{Min: -127, Max: 127, Normal: 0},
	GammaRange:        Range{Min: 1, Max: 40, Normal: 20},
	HighlightRange:    Range{Min: 0, Max: 100, Normal: 60},
	NoiseRemovalRange: Range{Min: 0, Max: 10, Normal: 2},
	ShadowRange:       Range{Min: 0, Max: 100, Normal: 10},
	SharpenRange:      Range{Min: 0, Max: 100, Normal: 15},
	ThresholdRange:    Range{Min: 0, Max: 100, Normal: 50},
	Platen:            testPlatenInputCapabilities,
	ADFSimplex:        testADFenInputCapabilities,
	ADFDuplex:         testADFenInputCapabilities,
}

// Variations of the initialized ScannerCapabilities structure:
//   - testScannerCapabilitiesNoPlaten     - no platen source
//   - testScannerCapabilitiesNoADF        - no ADF
//   - testScannerCapabilitiesNoADFSimplex - no ADFSimplex
//   - testScannerCapabilitiesNoADFDuplex  - no ADFDuplex
//   - testScannerCapabilitiesNoInput      - no inputs at all
//   - testScannerCapabilitiesNoColor      - no ColorModeColor support
//   - testScannerCapabilitiesNoHalftone   - no BinaryRenderingHalftone
var testScannerCapabilitiesNoPlaten *ScannerCapabilities
var testScannerCapabilitiesNoADF *ScannerCapabilities
var testScannerCapabilitiesNoADFSimplex *ScannerCapabilities
var testScannerCapabilitiesNoADFDuplex *ScannerCapabilities
var testScannerCapabilitiesNoInput *ScannerCapabilities
var testScannerCapabilitiesNoColor *ScannerCapabilities
var testScannerCapabilitiesNoHalftone *ScannerCapabilities

func init() {
	testScannerCapabilitiesNoPlaten = testScannerCapabilities.Clone()
	testScannerCapabilitiesNoPlaten.Platen = nil

	testScannerCapabilitiesNoADF = testScannerCapabilities.Clone()
	testScannerCapabilitiesNoADF.ADFSimplex = nil
	testScannerCapabilitiesNoADF.ADFDuplex = nil

	testScannerCapabilitiesNoADFSimplex = testScannerCapabilities.Clone()
	testScannerCapabilitiesNoADFSimplex.ADFSimplex = nil

	testScannerCapabilitiesNoADFDuplex = testScannerCapabilities.Clone()
	testScannerCapabilitiesNoADFDuplex.ADFDuplex = nil

	testScannerCapabilitiesNoInput = testScannerCapabilities.Clone()
	testScannerCapabilitiesNoInput.Platen = nil
	testScannerCapabilitiesNoInput.ADFSimplex = nil
	testScannerCapabilitiesNoInput.ADFDuplex = nil

	testScannerCapabilitiesNoColor = testScannerCapabilities.Clone()
	for _, inpcaps := range []**InputCapabilities{
		&testScannerCapabilitiesNoColor.Platen,
		&testScannerCapabilitiesNoColor.ADFSimplex,
		&testScannerCapabilitiesNoColor.ADFDuplex,
	} {
		*inpcaps = (*inpcaps).Clone()
		(*inpcaps).Profiles = slices.Clone((*inpcaps).Profiles)
		for i := range (*inpcaps).Profiles {
			prof := &(*inpcaps).Profiles[i]
			prof.ColorModes.Del(ColorModeColor)
		}
	}

	testScannerCapabilitiesNoHalftone = testScannerCapabilities.Clone()
	for _, inpcaps := range []**InputCapabilities{
		&testScannerCapabilitiesNoHalftone.Platen,
		&testScannerCapabilitiesNoHalftone.ADFSimplex,
		&testScannerCapabilitiesNoHalftone.ADFDuplex,
	} {
		*inpcaps = (*inpcaps).Clone()
		(*inpcaps).Profiles = slices.Clone((*inpcaps).Profiles)
		for i := range (*inpcaps).Profiles {
			prof := &(*inpcaps).Profiles[i]
			prof.BinaryRenderings.Del(BinaryRenderingHalftone)
		}
	}

}

// TestScannerRequestValidate tests ScannerRequest.Validate function.
func TestScannerRequestValidate(t *testing.T) {
	type testData struct {
		comment  string
		scancaps *ScannerCapabilities
		req      *ScannerRequest
		err      error
	}

	tests := []testData{
		// Zero request tests
		{
			comment:  "all-default request",
			scancaps: testScannerCapabilities,
			req:      &ScannerRequest{},
		},

		{
			comment:  "all-default request, no input supported",
			scancaps: testScannerCapabilitiesNoInput,
			req:      &ScannerRequest{},
			err: ErrParam{
				ErrUnsupportedParam, "Input", InputUnset,
			},
		},

		// InputPlaten tests
		{
			comment:  "InputPlaten",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input: InputPlaten,
			},
		},

		{
			comment:  "InputPlaten, unsupported",
			scancaps: testScannerCapabilitiesNoPlaten,
			req: &ScannerRequest{
				Input: InputPlaten,
			},
			err: ErrParam{
				ErrUnsupportedParam, "Input", InputPlaten,
			},
		},

		// InputADF/ADFModeUnset tests
		{
			comment:  "InputADF/ADFModeUnset",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input: InputADF,
			},
		},

		{
			comment:  "InputADF/ADFModeUnset, NoADFSimplex",
			scancaps: testScannerCapabilitiesNoADFSimplex,
			req: &ScannerRequest{
				Input: InputADF,
			},
		},

		{
			comment:  "InputADF/ADFModeUnset, NoADFDuplex",
			scancaps: testScannerCapabilitiesNoADFDuplex,
			req: &ScannerRequest{
				Input: InputADF,
			},
		},

		{
			comment:  "InputADF/ADFModeUnset, NoADF",
			scancaps: testScannerCapabilitiesNoADF,
			req: &ScannerRequest{
				Input: InputADF,
			},
			err: ErrParam{
				ErrUnsupportedParam, "Input", InputADF,
			},
		},

		// InputADF/ADFModeSimplex tests
		{
			comment:  "InputADF/ADFModeSimplex",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input:   InputADF,
				ADFMode: ADFModeSimplex,
			},
		},

		{
			comment:  "InputADF/ADFModeSimplex, NoADF",
			scancaps: testScannerCapabilitiesNoADF,
			req: &ScannerRequest{
				Input:   InputADF,
				ADFMode: ADFModeSimplex,
			},
			err: ErrParam{
				ErrUnsupportedParam, "Input", InputADF,
			},
		},

		{
			comment:  "InputADF/ADFModeSimplex, NoADFSimplex",
			scancaps: testScannerCapabilitiesNoADFSimplex,
			req: &ScannerRequest{
				Input:   InputADF,
				ADFMode: ADFModeSimplex,
			},
			err: ErrParam{
				ErrUnsupportedParam, "ADFMode", ADFModeSimplex,
			},
		},

		// InputADF/ADFModeDuplex tests
		{
			comment:  "InputADF/ADFModeDuplex",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input:   InputADF,
				ADFMode: ADFModeDuplex,
			},
		},

		{
			comment:  "InputADF/ADFModeDuplex, NoADF",
			scancaps: testScannerCapabilitiesNoADF,
			req: &ScannerRequest{
				Input:   InputADF,
				ADFMode: ADFModeDuplex,
			},
			err: ErrParam{
				ErrUnsupportedParam, "Input", InputADF,
			},
		},

		{
			comment:  "InputADF/ADFModeDuplex, NoADFDuplex",
			scancaps: testScannerCapabilitiesNoADFDuplex,
			req: &ScannerRequest{
				Input:   InputADF,
				ADFMode: ADFModeDuplex,
			},
			err: ErrParam{
				ErrUnsupportedParam, "ADFMode", ADFModeDuplex,
			},
		},

		// InputADF/invalid mode
		{
			comment:  "InputADF/ADFModeDuplex",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input:   InputADF,
				ADFMode: adfModeMax,
			},
			err: ErrParam{
				ErrInvalidParam, "ADFMode", adfModeMax,
			},
		},

		// invalid input
		{
			comment:  "invalid input",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input: inputMax,
			},
			err: ErrParam{
				ErrInvalidParam, "Input", inputMax,
			},
		},

		// ColorMode tests
		{
			comment:  "ColorModeBinary",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode: ColorModeBinary,
			},
		},

		{
			comment:  "ColorModeMono",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode: ColorModeMono,
			},
		},

		{
			comment:  "ColorModeColor",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode: ColorModeColor,
			},
		},

		{
			comment:  "invalid ColorMode",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode: colorModeMax,
			},
			err: ErrParam{
				ErrInvalidParam, "ColorMode", colorModeMax,
			},
		},

		{
			comment:  "ColorModeColor, unsupported",
			scancaps: testScannerCapabilitiesNoColor,
			req: &ScannerRequest{
				ColorMode: ColorModeColor,
			},
			err: ErrParam{
				ErrUnsupportedParam, "ColorMode",
				ColorModeColor,
			},
		},

		// BinaryRendering tests
		{
			comment:  "BinaryRenderingUnset",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:       ColorModeBinary,
				BinaryRendering: BinaryRenderingUnset,
			},
		},

		{
			comment:  "BinaryRenderingHalftone",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:       ColorModeBinary,
				BinaryRendering: BinaryRenderingHalftone,
			},
		},

		{
			comment:  "ColorModeUnset+BinaryRenderingHalftone, unsupported",
			scancaps: testScannerCapabilitiesNoHalftone,
			req: &ScannerRequest{
				ColorMode:       ColorModeUnset,
				BinaryRendering: BinaryRenderingHalftone,
			},
			err: ErrParam{
				ErrUnsupportedParam, "BinaryRendering",
				BinaryRenderingHalftone,
			},
		},

		{
			comment:  "ColorModeBinary+BinaryRenderingHalftone, unsupported",
			scancaps: testScannerCapabilitiesNoHalftone,
			req: &ScannerRequest{
				ColorMode:       ColorModeBinary,
				BinaryRendering: BinaryRenderingHalftone,
			},
			err: ErrParam{
				ErrUnsupportedParam, "BinaryRendering",
				BinaryRenderingHalftone,
			},
		},

		{
			comment:  "ColorModeColor+BinaryRenderingHalftone, ignored",
			scancaps: testScannerCapabilitiesNoHalftone,
			req: &ScannerRequest{
				ColorMode:       ColorModeColor,
				BinaryRendering: BinaryRenderingHalftone,
			},
		},

		{
			comment:  "BinaryRenderingThreshold",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:       ColorModeBinary,
				BinaryRendering: BinaryRenderingThreshold,
			},
		},

		{
			comment:  "BinaryRenderingThreshold, out of range",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:       ColorModeBinary,
				BinaryRendering: BinaryRenderingThreshold,
				Threshold:       optional.New(200),
			},
			err: ErrParam{
				ErrUnsupportedParam, "Threshold",
				200,
			},
		},

		{
			comment:  "invalid BinaryRenderingUnset",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:       ColorModeBinary,
				BinaryRendering: binaryRenderingMax,
			},
			err: ErrParam{
				ErrInvalidParam, "BinaryRendering",
				binaryRenderingMax,
			},
		},

		// ColorDepth tests
		{
			comment:  "ColorDepthUnset",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorDepth: ColorDepthUnset,
			},
		},

		{
			comment:  "ColorDepth8",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:  ColorModeColor,
				ColorDepth: ColorDepth8,
			},
		},

		{
			comment:  "ColorModeUnset+ColorDepth16, unsupported",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:  ColorModeUnset,
				ColorDepth: ColorDepth16,
			},
			err: ErrParam{
				ErrUnsupportedParam, "ColorDepth",
				ColorDepth16,
			},
		},

		{
			comment:  "ColorModeColor+ColorDepth16, unsupported",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:  ColorModeColor,
				ColorDepth: ColorDepth16,
			},
			err: ErrParam{
				ErrUnsupportedParam, "ColorDepth",
				ColorDepth16,
			},
		},

		{
			comment:  "ColorModeBinary+ColorDepth16, ignored",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:  ColorModeBinary,
				ColorDepth: ColorDepth16,
			},
		},

		{
			comment:  "invalid ColorDepth",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:  ColorModeColor,
				ColorDepth: colorDepthMax,
			},
			err: ErrParam{
				ErrInvalidParam, "ColorDepth",
				colorDepthMax,
			},
		},

		// CCDChannel tests
		{
			comment:  "CCDChannelNTSC",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				CCDChannel: CCDChannelNTSC,
			},
		},

		{
			comment:  "ColorModeUnset+CCDChannelGrayCcd, unsupported",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:  ColorModeUnset,
				CCDChannel: CCDChannelGrayCcd,
			},
			err: ErrParam{
				ErrUnsupportedParam, "CCDChannel",
				CCDChannelGrayCcd,
			},
		},

		{
			comment:  "ColorModeBinary+CCDChannelGrayCcd, unsupported",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:  ColorModeBinary,
				CCDChannel: CCDChannelGrayCcd,
			},
			err: ErrParam{
				ErrUnsupportedParam, "CCDChannel",
				CCDChannelGrayCcd,
			},
		},

		{
			comment:  "ColorModeColor+CCDChannelGrayCcd, ignored",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				ColorMode:  ColorModeColor,
				CCDChannel: CCDChannelGrayCcd,
			},
		},

		{
			comment:  "invalid CCDChannel",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				CCDChannel: ccdChannelMax,
			},
			err: ErrParam{
				ErrInvalidParam, "CCDChannel",
				ccdChannelMax,
			},
		},

		// Intent tests
		{
			comment:  "IntentDocument",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Intent: IntentDocument,
			},
		},

		{
			comment:  "IntentBusinessCard, unsupported",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Intent: IntentBusinessCard,
			},
			err: ErrParam{
				ErrUnsupportedParam, "Intent",
				IntentBusinessCard,
			},
		},

		{
			comment:  "invalid Intent",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Intent: intentMax,
			},
			err: ErrParam{
				ErrInvalidParam, "Intent",
				intentMax,
			},
		},

		// Region tests
		{
			comment:  "Region: good",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Region: Region{
					Width:  A4Width,
					Height: A4Height,
				},
			},
		},

		{
			comment:  "invalid Region",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Region: Region{Width: -5},
			},
			err: ErrParam{
				ErrInvalidParam, "Region",
				Region{Width: -5},
			},
		},

		{
			comment:  "Region: too big",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Region: Region{
					Width:  A3Width,
					Height: A3Height,
				},
			},
			err: ErrParam{
				ErrUnsupportedParam, "Region",
				Region{Width: A3Width, Height: A3Height},
			},
		},

		// Resolution test
		{
			comment:  "Resolution: good",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Resolution: Resolution{300, 300},
			},
		},

		{
			comment:  "Resolution: invalid",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Resolution: Resolution{-300, -300},
			},
			err: ErrParam{
				ErrInvalidParam, "Resolution",
				Resolution{-300, -300},
			},
		},

		{
			comment:  "Resolution: unsupported",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Resolution: Resolution{4800, 4800},
			},
			err: ErrParam{
				ErrUnsupportedParam, "Resolution",
				Resolution{4800, 4800},
			},
		},

		{
			comment:  "Resolution: 1200x1200, platen: OK",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input:      InputPlaten,
				Resolution: Resolution{1200, 1200},
			},
		},

		{
			comment:  "Resolution: 1200x1200, ADF: unsupported",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input:      InputADF,
				Resolution: Resolution{1200, 1200},
			},
			err: ErrParam{
				ErrUnsupportedParam, "Resolution",
				Resolution{1200, 1200},
			},
		},

		{
			comment:  "Resolution: 1200x1200, unset: OK",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input:      InputUnset,
				Resolution: Resolution{1200, 1200},
			},
		},

		{
			comment:  "Resolution: 2400x2400, color=unset: OK",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input:      InputUnset,
				Resolution: Resolution{2400, 2400},
			},
		},

		{
			comment:  "Resolution: 2400x2400, color=mono: OK",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input:      InputUnset,
				ColorMode:  ColorModeMono,
				Resolution: Resolution{2400, 2400},
			},
		},

		{
			comment:  "Resolution: 2400x2400, color=color: not OK",
			scancaps: testScannerCapabilities,
			req: &ScannerRequest{
				Input:      InputUnset,
				ColorMode:  ColorModeColor,
				Resolution: Resolution{2400, 2400},
			},
			err: ErrParam{
				ErrUnsupportedParam, "Resolution",
				Resolution{2400, 2400},
			},
		},
	}

	for _, test := range tests {
		err := test.req.Validate(test.scancaps)
		diff := testutils.Diff(test.err, err)
		if diff != "" {
			t.Errorf("failed: %q:\n%s", test.comment, diff)
		}
	}
}
