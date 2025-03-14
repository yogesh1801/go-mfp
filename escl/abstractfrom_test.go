// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// abstract.Scanner->eSCL conversions tests

package escl

import (
	"reflect"
	"slices"
	"testing"

	"github.com/alexpevzner/mfp/abstract"
	"github.com/alexpevzner/mfp/util/generic"
	"github.com/alexpevzner/mfp/util/optional"
)

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
	}

	for _, test := range tests {
		out := fromAbstractOptionalRange(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("input: %d\n"+
				"expected: %#v\n"+
				"present:  %#v",
				test.in, test.out, out)
		}
	}
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
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("input: %d\n"+
				"expected: %#v\n"+
				"present:  %#v",
				test.in, test.out, out)
		}
	}
}

// TestFromAbstractIntents tests fromAbstractIntents
func TestFromAbstractIntents(t *testing.T) {
	type testData struct {
		in  generic.Bitset[abstract.Intent]
		out []Intent
	}

	tests := []testData{
		{
			// Empty set
			in:  0,
			out: nil,
		},
		{
			// A couple of elements
			in: generic.MakeBitset(
				abstract.IntentDocument,
				abstract.IntentTextAndGraphic,
			),
			out: []Intent{Document, TextAndGraphic},
		},
		{
			// Full set
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
			// Set with some unknown element
			in: generic.MakeBitset(
				abstract.IntentDocument,
				abstract.IntentTextAndGraphic,
				30, // Unknown
			),
			out: []Intent{Document, TextAndGraphic},
		},
		{
			// Only unknown elements
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

		if !reflect.DeepEqual(out, expected) {
			t.Errorf("input: %d\n"+
				"expected: %#v\n"+
				"present:  %#v",
				test.in, test.out, out)
		}
	}
}

// TestFromAbstractBinaryRenderings tests fromAbstractBinaryRenderings
func TestFromAbstractBinaryRenderings(t *testing.T) {
	type testData struct {
		in  generic.Bitset[abstract.BinaryRendering]
		out []BinaryRendering
	}

	tests := []testData{
		{
			// Empty set
			in:  0,
			out: nil,
		},
		{
			// Single element
			in: generic.MakeBitset(
				abstract.BinaryRenderingHalftone,
			),
			out: []BinaryRendering{Halftone},
		},
		{
			// Full set
			in: generic.MakeBitset(
				abstract.BinaryRenderingHalftone,
				abstract.BinaryRenderingThreshold,
			),
			out: []BinaryRendering{Halftone, Threshold},
		},
		{
			// Set with some unknown element
			in: generic.MakeBitset(
				abstract.BinaryRenderingHalftone,
				30, // Unknown
			),
			out: []BinaryRendering{Halftone},
		},
		{
			// Only unknown elements
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

		if !reflect.DeepEqual(out, expected) {
			t.Errorf("input: %d\n"+
				"expected: %#v\n"+
				"present:  %#v",
				test.in, test.out, out)
		}
	}
}

// TestFromAbstractCCDChannels tests fromAbstractCCDChannels
func TestFromAbstractCCDChannels(t *testing.T) {
	type testData struct {
		in  generic.Bitset[abstract.CCDChannel]
		out []CCDChannel
	}

	tests := []testData{
		{
			// Empty set
			in:  0,
			out: nil,
		},
		{
			// A couple of elements
			in: generic.MakeBitset(
				abstract.CCDChannelRed,
				abstract.CCDChannelBlue,
			),
			out: []CCDChannel{Red, Blue},
		},
		{
			// Full set
			in: generic.MakeBitset(
				abstract.CCDChannelRed,
				abstract.CCDChannelGreen,
				abstract.CCDChannelBlue,
				abstract.CCDChannelNTSC,
				abstract.CCDChannelGrayCcd,
				abstract.CCDChannelGrayCcdEmulated,
			),
			out: []CCDChannel{Red, Green, Blue,
				NTSC, GrayCcd, GrayCcdEmulated},
		},
		{
			// Set with some unknown element
			in: generic.MakeBitset(
				abstract.CCDChannelRed,
				30, // Unknown
			),
			out: []CCDChannel{Red},
		},
		{
			// Only unknown elements
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

		if !reflect.DeepEqual(out, expected) {
			t.Errorf("input: %d\n"+
				"expected: %#v\n"+
				"present:  %#v",
				test.in, test.out, out)
		}
	}
}

// TestFromAbstractColorModes tests fromAbstractColorModes
func TestFromAbstractColorModes(t *testing.T) {
	type testData struct {
		modes  generic.Bitset[abstract.ColorMode]
		depths generic.Bitset[abstract.Depth]
		out    []ColorMode
	}

	tests := []testData{
		{
			// Empty set
			modes:  0,
			depths: 0,
			out:    nil,
		},
		{
			// All modes, 8 bit
			modes: generic.MakeBitset(
				abstract.ColorModeBinary,
				abstract.ColorModeMono,
				abstract.ColorModeColor,
			),
			depths: generic.MakeBitset(
				abstract.Depth8,
			),
			out: []ColorMode{BlackAndWhite1, Grayscale8, RGB24},
		},
		{
			// All modes, 8+16 bit
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
			// Unknown mode
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
			// Unknown depth
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
			// Only unknown modes
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

		if !reflect.DeepEqual(out, expected) {
			t.Errorf("input: modes: %s, depths: %s\n"+
				"expected: %#v\n"+
				"present:  %#v",
				test.modes, test.depths, test.out, out)
		}
	}
}

// TestFromAbstractInputSourceCaps tests fromAbstractInputSourceCaps
// function
func TestFromAbstractInputSourceCaps(t *testing.T) {
	type testData struct {
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
			// Bare minimum structure
			ver:     MakeVersion(2, 0),
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
				MaxXOffset:       optional.New(0),
				MaxYOffset:       optional.New(0),
				MaxScanRegions:   optional.New(1),
				SupportedIntents: []Intent{Document},
			},
		},
	}

	for _, test := range tests {
		out := fromAbstractInputSourceCaps(
			test.ver, test.formats, test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("input:        %#v\n"+
				"escl version: %s\n"+
				"formats:      %#v\n"+
				"expected:     %#v\n"+
				"present:      %#v",
				test.in, test.formats, test.ver, test.out, out)
		}
	}
}
