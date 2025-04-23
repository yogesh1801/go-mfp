// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// abstract.Scanner->eSCL conversions tests

package escl

import (
	"fmt"
	"slices"
	"testing"

	"github.com/alexpevzner/mfp/abstract"
	"github.com/alexpevzner/mfp/internal/testutils"
	"github.com/alexpevzner/mfp/util/generic"
	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/uuid"
)

// testAbstractResolutions contains initialized []abstract.Resolution slice
var testAbstractResolutions = []abstract.Resolution{
	{XResolution: 75, YResolution: 75},
	{XResolution: 150, YResolution: 150},
	{XResolution: 300, YResolution: 300},
	{XResolution: 600, YResolution: 600},
}

// testAbstractColorModes contains initialized abstract.ColorMode set
var testAbstractColorModes = generic.MakeBitset(
	abstract.ColorModeBinary,
	abstract.ColorModeMono,
	abstract.ColorModeColor,
)

// testAbstractDepth contains initialized abstract.Depth set
var testAbstractDepth = generic.MakeBitset(abstract.Depth8)

// testAbstractBinaryRenderings contains initialized
// abstract.BinaryRendering set
var testAbstractBinaryRenderings = generic.MakeBitset(
	abstract.BinaryRenderingHalftone,
	abstract.BinaryRenderingThreshold,
)

// testAbstractCCDChannels contains initialized abstract.CCDChannel set
var testAbstractCCDChannels = generic.MakeBitset(
	abstract.CCDChannelRed,
	abstract.CCDChannelGreen,
	abstract.CCDChannelBlue,
	abstract.CCDChannelNTSC,
	abstract.CCDChannelGrayCcd,
	abstract.CCDChannelGrayCcdEmulated,
)

// testAbstractSettingsProfiles contains initialized []abstract.SettingsProfile
// slice
var testAbstractSettingsProfiles = []abstract.SettingsProfile{
	{
		ColorModes:       testAbstractColorModes,
		Depths:           testAbstractDepth,
		BinaryRenderings: testAbstractBinaryRenderings,
		CCDChannels:      testAbstractCCDChannels,
		Resolutions:      testAbstractResolutions,
	},
}

// testAbstractUUID contains a parsed UUID
var testAbstractUUID = uuid.Must(
	uuid.Parse("418b75ab-1bd7-4d01-8178-75a84450a11c"),
)

// testAbstractInputCapabilities contains initialized
// abstract.InputCapabilities structure
var testAbstractInputCapabilities = &abstract.InputCapabilities{
	MinWidth:              15 * abstract.Millimeter,
	MinHeight:             20 * abstract.Millimeter,
	MaxWidth:              abstract.A4Width,
	MaxHeight:             abstract.A4Height,
	MaxXOffset:            2 * abstract.Inch,
	MaxYOffset:            3 * abstract.Inch,
	RiskyLeftMargins:      3 * abstract.Millimeter,
	RiskyRightMargins:     4 * abstract.Millimeter,
	RiskyTopMargins:       5 * abstract.Millimeter,
	RiskyBottomMargins:    6 * abstract.Millimeter,
	MaxOpticalXResolution: 1200,
	MaxOpticalYResolution: 600,
	Intents:               generic.MakeBitset(abstract.IntentDocument),
	Profiles:              testAbstractSettingsProfiles,
}

// testAbstractScannerCapabilities contains initialized
// abstract.ScannerCapabilities structure
var testAbstractScannerCapabilities = &abstract.ScannerCapabilities{
	UUID:              testAbstractUUID,
	MakeAndModel:      "Abstract Scanner",
	SerialNumber:      "AS-12345",
	Manufacturer:      "Abstract Corp.",
	AdminURI:          "http://192.168.0.1/admin",
	IconURI:           "http://192.168.0.1/icon.png",
	DocumentFormats:   []string{"image/jpeg", "application/pdf"},
	CompressionRange:  abstract.Range{Min: 2, Max: 10, Normal: 5},
	ADFCapacity:       35,
	BrightnessRange:   abstract.Range{Min: -100, Max: 100, Normal: 0},
	ContrastRange:     abstract.Range{Min: 0, Max: 100, Normal: 75},
	GammaRange:        abstract.Range{Min: 1, Max: 40, Normal: 20},
	HighlightRange:    abstract.Range{Min: 0, Max: 100, Normal: 60},
	NoiseRemovalRange: abstract.Range{Min: 0, Max: 10, Normal: 2},
	ShadowRange:       abstract.Range{Min: 0, Max: 100, Normal: 10},
	SharpenRange:      abstract.Range{Min: 0, Max: 100, Normal: 15},
	ThresholdRange:    abstract.Range{Min: 0, Max: 100, Normal: 50},
	Platen:            testAbstractInputCapabilities,
}

// TestFromAbstractOptionalInt tests fromAbstractOptionalInt
func TestFromAbstractOptionalInt(t *testing.T) {
	type testData struct {
		in  int
		out optional.Val[int]
	}

	tests := []testData{
		{0, nil},
		{1, optional.New(1)},
	}

	for _, test := range tests {
		out := fromAbstractOptionalInt(test.in)
		comment := fmt.Sprintf("input: %+v", test.in)

		testutils.CheckConvertionTest(t,
			"fromAbstractOptionalRange",
			comment, test.out, out)
	}
}

// TestFromAbstracOptionaltRange tests fromAbstractOptionalRange
func TestFromAbstractRange(t *testing.T) {
	type testData struct {
		in  abstract.Range
		out optional.Val[Range]
	}

	tests := []testData{
		{
			in:  abstract.Range{},
			out: nil,
		},
		{
			in: abstract.Range{Min: -100, Max: 100, Normal: 0},
			out: optional.New(Range{
				Min: -100, Max: 100, Normal: 0, Step: nil,
			}),
		},
		{
			in: abstract.Range{Min: -100, Max: 100, Step: 2},
			out: optional.New(Range{
				Min: -100, Max: 100, Normal: 0,
				Step: optional.New(2),
			}),
		},
	}

	for _, test := range tests {
		out := fromAbstractOptionalRange(test.in)
		comment := fmt.Sprintf("input: %+v", test.in)

		testutils.CheckConvertionTest(t,
			"fromAbstractOptionalRange",
			comment, test.out, out)
	}
}

// TestFromAbstractIntents tests fromAbstractIntents
func TestFromAbstractIntents(t *testing.T) {
	type testData struct {
		comment string
		in      generic.Bitset[abstract.Intent]
		out     []Intent
	}

	tests := []testData{
		{
			comment: "Empty set",
			in:      0,
			out:     nil,
		},
		{
			comment: "A couple of elements",
			in: generic.MakeBitset(
				abstract.IntentDocument,
				abstract.IntentTextAndGraphic,
			),
			out: []Intent{Document, TextAndGraphic},
		},
		{
			comment: "Full set",
			in: generic.MakeBitset(
				abstract.IntentDocument,
				abstract.IntentTextAndGraphic,
				abstract.IntentPhoto,
				abstract.IntentPreview,
				abstract.IntentObject,
				abstract.IntentBusinessCard,
			),
			out: []Intent{
				Document,
				TextAndGraphic,
				Photo,
				Preview,
				Object,
				BusinessCard,
			},
		},
		{
			comment: "Set with some unknown element",
			in: generic.MakeBitset(
				abstract.IntentDocument,
				abstract.IntentTextAndGraphic,
				30, // Unknown
			),
			out: []Intent{Document, TextAndGraphic},
		},
		{
			comment: "Only unknown elements",
			in: generic.MakeBitset(
				abstract.Intent(30), // Unknown
			),
			out: nil,
		},
	}

	for _, test := range tests {
		out := fromAbstractIntents(test.in)
		expected := slices.Clone(test.out)

		slices.Sort(out)
		slices.Sort(expected)

		testutils.CheckConvertionTest(t,
			"fromAbstractIntents",
			test.comment, expected, out)
	}
}

// TestFromAbstractBinaryRenderings tests fromAbstractBinaryRenderings
func TestFromAbstractBinaryRenderings(t *testing.T) {
	type testData struct {
		comment string
		in      generic.Bitset[abstract.BinaryRendering]
		out     []BinaryRendering
	}

	tests := []testData{
		{
			comment: "Empty set",
			in:      0,
			out:     nil,
		},
		{
			comment: "Single element",
			in: generic.MakeBitset(
				abstract.BinaryRenderingHalftone,
			),
			out: []BinaryRendering{Halftone},
		},
		{
			comment: "Full set",
			in:      testAbstractBinaryRenderings,
			out:     []BinaryRendering{Halftone, Threshold},
		},
		{
			comment: "Set with some unknown element",
			in: generic.MakeBitset(
				abstract.BinaryRenderingHalftone,
				30, // Unknown
			),
			out: []BinaryRendering{Halftone},
		},
		{
			comment: "Only unknown elements",
			in: generic.MakeBitset(
				abstract.BinaryRendering(30), // Unknown
			),
			out: nil,
		},
	}

	for _, test := range tests {
		out := fromAbstractBinaryRenderings(test.in)
		expected := slices.Clone(test.out)

		slices.Sort(out)
		slices.Sort(expected)

		testutils.CheckConvertionTest(t,
			"fromAbstractBinaryRenderings",
			test.comment, expected, out)
	}
}

// TestFromAbstractCCDChannels tests fromAbstractCCDChannels
func TestFromAbstractCCDChannels(t *testing.T) {
	type testData struct {
		comment string
		in      generic.Bitset[abstract.CCDChannel]
		out     []CCDChannel
	}

	tests := []testData{
		{
			comment: "Empty set",
			in:      0,
			out:     nil,
		},
		{
			comment: "A couple of elements",
			in: generic.MakeBitset(
				abstract.CCDChannelRed,
				abstract.CCDChannelBlue,
			),
			out: []CCDChannel{Red, Blue},
		},
		{
			comment: "Full set",
			in:      testAbstractCCDChannels,
			out: []CCDChannel{Red, Green, Blue,
				NTSC, GrayCcd, GrayCcdEmulated},
		},
		{
			comment: "Set with some unknown element",
			in: generic.MakeBitset(
				abstract.CCDChannelRed,
				30, // Unknown
			),
			out: []CCDChannel{Red},
		},
		{
			comment: "Only unknown elements",
			in: generic.MakeBitset(
				abstract.CCDChannel(30), // Unknown
			),
			out: nil,
		},
	}

	for _, test := range tests {
		out := fromAbstractCCDChannels(test.in)
		expected := slices.Clone(test.out)

		slices.Sort(out)
		slices.Sort(expected)

		testutils.CheckConvertionTest(t,
			"fromAbstractCCDChannels",
			test.comment, expected, out)
	}
}

// TestFromAbstractColorModes tests fromAbstractColorModes
func TestFromAbstractColorModes(t *testing.T) {
	type testData struct {
		comment string
		modes   generic.Bitset[abstract.ColorMode]
		depths  generic.Bitset[abstract.Depth]
		out     []ColorMode
	}

	tests := []testData{
		{
			comment: "Empty set",
			modes:   0,
			depths:  0,
			out:     nil,
		},
		{
			comment: "All modes, 8 bit",
			modes:   testAbstractColorModes,
			depths:  testAbstractDepth,
			out:     []ColorMode{BlackAndWhite1, Grayscale8, RGB24},
		},
		{
			comment: "All modes, 8+16 bit",
			modes: generic.MakeBitset(
				abstract.ColorModeBinary,
				abstract.ColorModeMono,
				abstract.ColorModeColor,
			),
			depths: generic.MakeBitset(
				abstract.Depth8,
				abstract.Depth16,
			),
			out: []ColorMode{BlackAndWhite1, Grayscale8, RGB24,
				Grayscale16, RGB48},
		},
		{
			comment: "Unknown mode",
			modes: generic.MakeBitset(
				abstract.ColorModeBinary,
				abstract.ColorModeMono,
				abstract.ColorModeColor,
				30,
			),
			depths: generic.MakeBitset(
				abstract.Depth8,
			),
			out: []ColorMode{BlackAndWhite1, Grayscale8, RGB24},
		},
		{
			comment: "Unknown depth",
			modes: generic.MakeBitset(
				abstract.ColorModeBinary,
				abstract.ColorModeMono,
				abstract.ColorModeColor,
			),
			depths: generic.MakeBitset(
				abstract.Depth8,
				30, // Unknown
			),
			out: []ColorMode{BlackAndWhite1, Grayscale8, RGB24},
		},
		{
			comment: "Only unknown modes",
			modes: generic.MakeBitset(
				abstract.ColorMode(30),
			),
			depths: generic.MakeBitset(
				abstract.Depth8,
			),
			out: nil,
		},
	}

	for _, test := range tests {
		out := fromAbstractColorModes(test.modes, test.depths)
		expected := slices.Clone(test.out)

		slices.Sort(out)
		slices.Sort(expected)

		testutils.CheckConvertionTest(t,
			"fromAbstractResolutions",
			test.comment, expected, out)
	}
}

// TestFromAbstractResolutions tests fromAbstractResolutions
// function
func TestFromAbstractResolutions(t *testing.T) {
	type testData struct {
		comment     string
		absres      []abstract.Resolution
		absresrange abstract.ResolutionRange
		out         SupportedResolutions
	}

	tests := []testData{
		{
			comment: "Only discrete resolutions",
			absres:  testAbstractResolutions,
			out: SupportedResolutions{
				DiscreteResolutions: []DiscreteResolution{
					{75, 75},
					{150, 150},
					{300, 300},
					{600, 600},
				},
			},
		},

		{
			comment: "Range of resolutions",
			absresrange: abstract.ResolutionRange{
				XMin: 200, XMax: 1200, XStep: 100, XNormal: 400,
				YMin: 100, YMax: 600, YStep: 50, YNormal: 300,
			},

			out: SupportedResolutions{
				ResolutionRange: optional.New(ResolutionRange{
					XResolutionRange: Range{
						Min:    200,
						Max:    1200,
						Normal: 400,
						Step:   optional.New(100),
					},
					YResolutionRange: Range{
						Min:    100,
						Max:    600,
						Normal: 300,
						Step:   optional.New(50),
					},
				}),
			},
		},

		{
			comment: "Range of resolutions with missed step",
			absresrange: abstract.ResolutionRange{
				XMin: 200, XMax: 1200, XNormal: 400,
				YMin: 100, YMax: 600, YNormal: 300,
			},

			out: SupportedResolutions{
				ResolutionRange: optional.New(ResolutionRange{
					XResolutionRange: Range{
						Min:    200,
						Max:    1200,
						Normal: 400,
						Step:   optional.New(1),
					},
					YResolutionRange: Range{
						Min:    100,
						Max:    600,
						Normal: 300,
						Step:   optional.New(1),
					},
				}),
			},
		},

		{
			comment: "Mix of discrete and range resolutions",
			absres:  testAbstractResolutions,
			absresrange: abstract.ResolutionRange{
				XMin: 200, XMax: 1200, XStep: 100, XNormal: 400,
				YMin: 100, YMax: 600, YStep: 50, YNormal: 300,
			},
			out: SupportedResolutions{
				DiscreteResolutions: []DiscreteResolution{
					{75, 75},
					{150, 150},
					{300, 300},
					{600, 600},
				},
				ResolutionRange: optional.New(ResolutionRange{
					XResolutionRange: Range{
						Min:    200,
						Max:    1200,
						Normal: 400,
						Step:   optional.New(100),
					},
					YResolutionRange: Range{
						Min:    100,
						Max:    600,
						Normal: 300,
						Step:   optional.New(50),
					},
				}),
			},
		},
	}

	for _, test := range tests {
		out := fromAbstractResolutions(test.absres, test.absresrange)

		testutils.CheckConvertionTest(t,
			"fromAbstractResolutions",
			test.comment, test.out, out)
	}
}

// TestFromAbstractSettingsProfiles tests fromAbstractSettingsProfiles
// function
func TestFromAbstractSettingsProfiles(t *testing.T) {
	type testData struct {
		comment string // Not used
		ver     Version
		formats []string
		in      []abstract.SettingsProfile
		out     []SettingProfile
	}

	formats := []string{"image/jpeg", "application/pdf"}

	tests := []testData{
		{
			// eSCL 2.0
			ver:     MakeVersion(2, 0),
			formats: formats,
			in:      testAbstractSettingsProfiles,
			out: []SettingProfile{
				{
					ColorModes: fromAbstractColorModes(
						testAbstractColorModes,
						testAbstractDepth),
					ColorSpaces:     []ColorSpace{SRGB},
					DocumentFormats: formats,
					CCDChannels: fromAbstractCCDChannels(
						testAbstractCCDChannels),
					BinaryRenderings: fromAbstractBinaryRenderings(
						testAbstractBinaryRenderings),
					SupportedResolutions: fromAbstractResolutions(
						testAbstractResolutions,
						abstract.ResolutionRange{}),
				},
			},
		},

		{
			// eSCL 2.1+
			ver:     MakeVersion(2, 1),
			formats: formats,
			in:      testAbstractSettingsProfiles,
			out: []SettingProfile{
				{
					ColorModes: fromAbstractColorModes(
						testAbstractColorModes,
						testAbstractDepth),
					ColorSpaces:        []ColorSpace{SRGB},
					DocumentFormats:    formats,
					DocumentFormatsExt: formats,
					CCDChannels: fromAbstractCCDChannels(
						testAbstractCCDChannels),
					BinaryRenderings: fromAbstractBinaryRenderings(
						testAbstractBinaryRenderings),
					SupportedResolutions: fromAbstractResolutions(
						testAbstractResolutions,
						abstract.ResolutionRange{}),
				},
			},
		},
	}

	for _, test := range tests {
		out := fromAbstractSettingsProfiles(
			test.ver, test.formats, test.in)

		comment := fmt.Sprintf("eSCL %s", test.ver)

		testutils.CheckConvertionTest(t,
			"fromAbstractScannerCapabilities",
			comment, test.out, out)
	}
}

// TestFromAbstractInputSourceCaps tests fromAbstractInputSourceCaps
// function
func TestFromAbstractInputSourceCaps(t *testing.T) {
	type testData struct {
		comment string
		ver     Version
		formats []string
		in      *abstract.InputCapabilities
		out     InputSourceCaps
	}

	formats := []string{"image/jpeg", "application/pdf"}
	intents := generic.MakeBitset(
		abstract.IntentDocument,
	)

	tests := []testData{
		{
			comment: "Bare minimum structure",
			ver:     DefaultVersion,
			formats: formats,
			in: &abstract.InputCapabilities{
				MinWidth:  3 * abstract.Millimeter,
				MinHeight: 5 * abstract.Millimeter,
				MaxWidth:  abstract.A4Width,
				MaxHeight: abstract.A4Height,
				Intents:   intents,
			},
			out: InputSourceCaps{
				MinWidth:         (3 * abstract.Millimeter).Dots(300),
				MinHeight:        (5 * abstract.Millimeter).Dots(300),
				MaxWidth:         abstract.A4Width.Dots(300),
				MaxHeight:        abstract.A4Height.Dots(300),
				MaxXOffset:       nil,
				MaxYOffset:       nil,
				MaxScanRegions:   optional.New(1),
				SupportedIntents: []Intent{Document},
			},
		},

		{
			comment: " Full-data test",
			ver:     DefaultVersion,
			formats: formats,
			in:      testAbstractInputCapabilities,
			out: InputSourceCaps{
				MinWidth:  (15 * abstract.Millimeter).Dots(300),
				MinHeight: (20 * abstract.Millimeter).Dots(300),
				MaxWidth:  abstract.A4Width.Dots(300),
				MaxHeight: abstract.A4Height.Dots(300),
				MaxXOffset: optional.New(
					(2 * abstract.Inch).Dots(300)),
				MaxYOffset: optional.New(
					(3 * abstract.Inch).Dots(300)),
				MaxOpticalXResolution: optional.New(1200),
				MaxOpticalYResolution: optional.New(600),
				RiskyLeftMargins: optional.New(
					(3 * abstract.Millimeter).Dots(300)),
				RiskyRightMargins: optional.New(
					(4 * abstract.Millimeter).Dots(300)),
				RiskyTopMargins: optional.New(
					(5 * abstract.Millimeter).Dots(300)),
				RiskyBottomMargins: optional.New(
					(6 * abstract.Millimeter).Dots(300)),
				MaxScanRegions:   optional.New(1),
				SupportedIntents: []Intent{Document},
				SettingProfiles: fromAbstractSettingsProfiles(
					DefaultVersion,
					formats,
					testAbstractSettingsProfiles,
				),
			},
		},
	}

	for _, test := range tests {
		out := fromAbstractInputSourceCaps(
			test.ver, test.formats, test.in)

		comment := fmt.Sprintf("%s (eSCL %s)", test.comment, test.ver)

		testutils.CheckConvertionTest(t,
			"fromAbstractScannerCapabilities",
			comment, test.out, out)
	}
}

// TestFromAbstractScannerCapabilities tests fromAbstractScannerCapabilities
// function
func TestFromAbstractScannerCapabilities(t *testing.T) {
	type testData struct {
		comment string
		in      *abstract.ScannerCapabilities
		out     ScannerCapabilities
	}

	formats := []string{"image/jpeg", "application/pdf"}

	capacity := 25

	abscaps := fromAbstractInputSourceCaps(
		DefaultVersion, formats, testAbstractInputCapabilities)

	platen := Platen{
		optional.New(abscaps),
	}

	adfSimplex := ADF{
		ADFSimplexInputCaps: optional.New(abscaps),
		FeederCapacity:      optional.New(capacity),
	}

	adfDuplex := ADF{
		ADFSimplexInputCaps: optional.New(abscaps),
		ADFDuplexInputCaps:  optional.New(abscaps),
		FeederCapacity:      optional.New(capacity),
		ADFOptions:          []ADFOption{Duplex},
	}

	tests := []testData{
		{
			comment: "Bare minimum",
			in: &abstract.ScannerCapabilities{
				UUID: testAbstractUUID,
			},
			out: ScannerCapabilities{
				Version: DefaultVersion,
				UUID:    optional.New(testAbstractUUID),
			},
		},

		{
			comment: "Bare minimim with Platen source",
			in: &abstract.ScannerCapabilities{
				UUID:            testAbstractUUID,
				DocumentFormats: formats,
				Platen:          testAbstractInputCapabilities,
			},
			out: ScannerCapabilities{
				Version: DefaultVersion,
				UUID:    optional.New(testAbstractUUID),
				Platen:  optional.New(platen),
			},
		},

		{
			comment: "Bare minimim with Platen + ADF Simplex",
			in: &abstract.ScannerCapabilities{
				UUID:            testAbstractUUID,
				DocumentFormats: formats,
				ADFCapacity:     capacity,
				Platen:          testAbstractInputCapabilities,
				ADFSimplex:      testAbstractInputCapabilities,
			},
			out: ScannerCapabilities{
				Version: DefaultVersion,
				UUID:    optional.New(testAbstractUUID),
				Platen:  optional.New(platen),
				ADF:     optional.New(adfSimplex),
			},
		},

		{
			comment: "Bare minimim with Platen + ADF Duplex",
			in: &abstract.ScannerCapabilities{
				UUID:            testAbstractUUID,
				DocumentFormats: formats,
				ADFCapacity:     capacity,
				Platen:          testAbstractInputCapabilities,
				ADFSimplex:      testAbstractInputCapabilities,
				ADFDuplex:       testAbstractInputCapabilities,
			},
			out: ScannerCapabilities{
				Version: DefaultVersion,
				UUID:    optional.New(testAbstractUUID),
				Platen:  optional.New(platen),
				ADF:     optional.New(adfDuplex),
			},
		},

		{
			comment: "Full-data test",
			in:      testAbstractScannerCapabilities,
			out: ScannerCapabilities{
				Version:      DefaultVersion,
				UUID:         optional.New(testAbstractUUID),
				MakeAndModel: optional.New("Abstract Scanner"),
				SerialNumber: optional.New("AS-12345"),
				Manufacturer: optional.New("Abstract Corp."),
				AdminURI:     optional.New("http://192.168.0.1/admin"),
				IconURI:      optional.New("http://192.168.0.1/icon.png"),
				Platen:       optional.New(platen),
				BrightnessSupport: optional.New(
					Range{Min: -100, Max: 100, Normal: 0}),
				CompressionFactorSupport: optional.New(
					Range{Min: 2, Max: 10, Normal: 5}),
				ContrastSupport: optional.New(
					Range{Min: 0, Max: 100, Normal: 75}),
				GammaSupport: optional.New(
					Range{Min: 1, Max: 40, Normal: 20}),
				HighlightSupport: optional.New(
					Range{Min: 0, Max: 100, Normal: 60}),
				NoiseRemovalSupport: optional.New(
					Range{Min: 0, Max: 10, Normal: 2}),
				ShadowSupport: optional.New(
					Range{Min: 0, Max: 100, Normal: 10}),
				SharpenSupport: optional.New(
					Range{Min: 0, Max: 100, Normal: 15}),
				ThresholdSupport: optional.New(
					Range{Min: 0, Max: 100, Normal: 50}),
			},
		},
	}

	for _, test := range tests {
		out := fromAbstractScannerCapabilities(
			DefaultVersion, test.in)

		testutils.CheckConvertionTest(t,
			"fromAbstractScannerCapabilities",
			test.comment, test.out, out)
	}
}

// TestFromAbstractScanSettings tests fromAbstractScanSettings
// function
func TestFromAbstractScanSettings(t *testing.T) {
	type testData struct {
		comment string
		ver     Version
		in      *abstract.ScannerRequest
		out     ScanSettings
	}

	tests := []testData{
		// Empty request; all default
		{
			comment: "Empty request",
			ver:     DefaultVersion,
			in:      &abstract.ScannerRequest{},
			out: ScanSettings{
				Version: DefaultVersion,
			},
		},

		// Intents support
		{
			comment: "Intent=Document",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				Intent: abstract.IntentDocument,
			},
			out: ScanSettings{
				Version: DefaultVersion,
				Intent:  optional.New(Document),
			},
		},

		{
			comment: "Intent=Invaid",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				Intent: abstract.Intent(-1),
			},
			out: ScanSettings{
				Version: DefaultVersion,
				Intent:  nil, // Must default to nil
			},
		},

		// Inputs
		{
			comment: "InputPlaten",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				Input: abstract.InputPlaten,
			},
			out: ScanSettings{
				Version:     DefaultVersion,
				InputSource: optional.New(InputPlaten),
			},
		},

		{
			comment: "InputFeeder+Simplex(by default)",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				Input: abstract.InputADF,
			},
			out: ScanSettings{
				Version:     DefaultVersion,
				InputSource: optional.New(InputFeeder),
			},
		},

		{
			comment: "InputFeeder+Simplex",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				Input:   abstract.InputADF,
				ADFMode: abstract.ADFModeSimplex,
			},
			out: ScanSettings{
				Version:     DefaultVersion,
				InputSource: optional.New(InputFeeder),
			},
		},

		{
			comment: "InputFeeder+Duplex",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				Input:   abstract.InputADF,
				ADFMode: abstract.ADFModeDuplex,
			},
			out: ScanSettings{
				Version:     DefaultVersion,
				InputSource: optional.New(InputFeeder),
				Duplex:      optional.New(true),
			},
		},

		// Color modes support
		{
			comment: "BlackAndWhite1",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode: abstract.ColorModeBinary,
			},
			out: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(BlackAndWhite1),
			},
		},

		{
			comment: "BlackAndWhite1+Halftone",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode:       abstract.ColorModeBinary,
				BinaryRendering: abstract.BinaryRenderingHalftone,
			},
			out: ScanSettings{
				Version:         DefaultVersion,
				ColorMode:       optional.New(BlackAndWhite1),
				BinaryRendering: optional.New(Halftone),
			},
		},

		{
			comment: "BlackAndWhite1+Threshold(default)",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode:       abstract.ColorModeBinary,
				BinaryRendering: abstract.BinaryRenderingThreshold,
			},
			out: ScanSettings{
				Version:         DefaultVersion,
				ColorMode:       optional.New(BlackAndWhite1),
				BinaryRendering: optional.New(Threshold),
			},
		},

		{
			comment: "BlackAndWhite1+Threshold=50",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode:       abstract.ColorModeBinary,
				BinaryRendering: abstract.BinaryRenderingThreshold,
				Threshold:       optional.New(50),
			},
			out: ScanSettings{
				Version:         DefaultVersion,
				ColorMode:       optional.New(BlackAndWhite1),
				BinaryRendering: optional.New(Threshold),
				Threshold:       optional.New(50),
			},
		},

		{
			comment: "Grayscale8",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode: abstract.ColorModeMono,
				Depth:     abstract.Depth8,
			},
			out: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(Grayscale8),
			},
		},

		{
			comment: "Grayscale8 (default depth)",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode: abstract.ColorModeMono,
			},
			out: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(Grayscale8),
			},
		},

		{
			comment: "Grayscale16",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode: abstract.ColorModeMono,
				Depth:     abstract.Depth16,
			},
			out: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(Grayscale16),
			},
		},

		{
			comment: "RGB24",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode: abstract.ColorModeColor,
				Depth:     abstract.Depth8,
			},
			out: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(RGB24),
			},
		},

		{
			comment: "RGB24 (default depth)",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode: abstract.ColorModeColor,
			},
			out: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(RGB24),
			},
		},

		{
			comment: "RGB24",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				ColorMode: abstract.ColorModeColor,
				Depth:     abstract.Depth16,
			},
			out: ScanSettings{
				Version:   DefaultVersion,
				ColorMode: optional.New(RGB48),
			},
		},

		// DocumentFormat/DocumentFormatExt (depends on version)
		{
			comment: "DocumentFormat",
			ver:     MakeVersion(2, 0),
			in: &abstract.ScannerRequest{
				DocumentFormat: "image/jpeg",
			},
			out: ScanSettings{
				Version:           MakeVersion(2, 0),
				DocumentFormat:    optional.New("image/jpeg"),
				DocumentFormatExt: nil, // eSCL < 2.1
			},
		},

		{
			comment: "DocumentFormat",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				DocumentFormat: "image/jpeg",
			},
			out: ScanSettings{
				Version:           DefaultVersion,
				DocumentFormat:    optional.New("image/jpeg"),
				DocumentFormatExt: optional.New("image/jpeg"),
			},
		},

		// Region
		{
			comment: "Region",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				Region: abstract.Region{
					XOffset: abstract.Inch,
					YOffset: abstract.Inch * 2,
					Width:   abstract.Inch * 10,
					Height:  abstract.Inch * 11,
				},
			},
			out: ScanSettings{
				Version: DefaultVersion,
				ScanRegions: []ScanRegion{
					{
						XOffset:            300,
						YOffset:            600,
						Width:              3000,
						Height:             3300,
						ContentRegionUnits: ThreeHundredthsOfInches,
					},
				},
			},
		},

		// Resolution
		{
			comment: "Resolution",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				Resolution: abstract.Resolution{
					XResolution: 600,
					YResolution: 300,
				},
			},
			out: ScanSettings{
				Version:     DefaultVersion,
				XResolution: optional.New(600),
				YResolution: optional.New(300),
			},
		},

		// CCDChannel
		{
			comment: "CCDChannel",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				CCDChannel: abstract.CCDChannelNTSC,
			},
			out: ScanSettings{
				Version:    DefaultVersion,
				CCDChannel: optional.New(NTSC),
			},
		},

		// Image processing parameters
		{
			comment: "CCDChannel",
			ver:     DefaultVersion,
			in: &abstract.ScannerRequest{
				Brightness:   optional.New(100),
				Contrast:     optional.New(80),
				Gamma:        optional.New(20),
				Highlight:    optional.New(85),
				NoiseRemoval: optional.New(30),
				Shadow:       optional.New(15),
				Sharpen:      optional.New(10),
				Compression:  optional.New(5),
			},
			out: ScanSettings{
				Version:           DefaultVersion,
				Brightness:        optional.New(100),
				Contrast:          optional.New(80),
				Gamma:             optional.New(20),
				Highlight:         optional.New(85),
				NoiseRemoval:      optional.New(30),
				Shadow:            optional.New(15),
				Sharpen:           optional.New(10),
				CompressionFactor: optional.New(5),
			},
		},
	}

	for _, test := range tests {
		out := fromAbstractScanSettings(
			test.ver, test.in)

		comment := fmt.Sprintf("%s (eSCL %s)", test.comment, test.ver)

		testutils.CheckConvertionTest(t,
			"fromAbstractScannerCapabilities",
			comment, test.out, out)
	}
}
