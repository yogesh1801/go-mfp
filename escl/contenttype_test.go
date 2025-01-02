// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan "content type" test

package escl

import "testing"

// TestContentTypeString tests ContentType.String
func TestContentTypeString(t *testing.T) {
	type testData struct {
		ct ContentType
		s  string
	}

	tests := []testData{
		{ContentTypePhoto, "Photo"},
		{ContentTypeText, "Text"},
		{ContentTypeTextAndPhoto, "TextAndPhoto"},
		{ContentTypeLineArt, "LineArt"},
		{ContentTypeMagazine, "Magazine"},
		{ContentTypeHalftone, "Halftone"},
		{ContentTypeAuto, "Auto"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.ct.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.ct), test.s, s)
		}
	}
}

// TestDecodeContentType tests DecodeContentType
func TestDecodeContentType(t *testing.T) {
	type testData struct {
		ct ContentType
		s  string
	}

	tests := []testData{
		{ContentTypePhoto, "Photo"},
		{ContentTypeText, "Text"},
		{ContentTypeTextAndPhoto, "TextAndPhoto"},
		{ContentTypeLineArt, "LineArt"},
		{ContentTypeMagazine, "Magazine"},
		{ContentTypeHalftone, "Halftone"},
		{ContentTypeAuto, "Auto"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		ct := DecodeContentType(test.s)
		if ct != test.ct {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.ct, ct)
		}
	}
}
