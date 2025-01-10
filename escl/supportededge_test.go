// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image edges, for edge auto detection test

package escl

import "testing"

// TestSupportedEdgeString tests SupportedEdge.String
func TestSupportedEdgeString(t *testing.T) {
	type testData struct {
		intent SupportedEdge
		s      string
	}

	tests := []testData{
		{TopEdge, "TopEdge"},
		{LeftEdge, "LeftEdge"},
		{BottomEdge, "BottomEdge"},
		{RightEdge, "RightEdge"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.intent.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.intent), test.s, s)
		}
	}
}

// TestDecodeSupportedEdge tests DecodeSupportedEdge
func TestDecodeSupportedEdge(t *testing.T) {
	type testData struct {
		intent SupportedEdge
		s      string
	}

	tests := []testData{
		{TopEdge, "TopEdge"},
		{LeftEdge, "LeftEdge"},
		{BottomEdge, "BottomEdge"},
		{RightEdge, "RightEdge"},
		{UnknownSupportedEdge, "XXX"},
	}

	for _, test := range tests {
		intent := DecodeSupportedEdge(test.s)
		if intent != test.intent {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.intent, intent)
		}
	}
}
