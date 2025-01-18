// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input source test

package escl

import "testing"

// TestInputSourceString tests InputSource.String
func TestInputSourceString(t *testing.T) {
	type testData struct {
		input InputSource
		s     string
	}

	tests := []testData{
		{InputPlaten, "Platen"},
		{InputFeeder, "Feeder"},
		{InputCamera, "Camera"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.input.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.input), test.s, s)
		}
	}
}

// TestDecodeInputSource tests DecodeInputSource
func TestDecodeInputSource(t *testing.T) {
	type testData struct {
		input InputSource
		s     string
	}

	tests := []testData{
		{InputPlaten, "Platen"},
		{InputFeeder, "Feeder"},
		{InputCamera, "Camera"},
		{UnknownInputSource, "XXX"},
	}

	for _, test := range tests {
		input := DecodeInputSource(test.s)
		if input != test.input {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.input, input)
		}
	}
}
