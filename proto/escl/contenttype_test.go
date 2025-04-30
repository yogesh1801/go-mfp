// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan "content type" test

package escl

import "testing"

var testContentType = testEnum[ContentType]{
	decodeStr: DecodeContentType,
	decodeXML: decodeContentType,
	dataset: []testEnumData[ContentType]{
		{ContentTypePhoto, "Photo"},
		{ContentTypeText, "Text"},
		{ContentTypeTextAndPhoto, "TextAndPhoto"},
		{ContentTypeLineArt, "LineArt"},
		{ContentTypeMagazine, "Magazine"},
		{ContentTypeHalftone, "Halftone"},
		{ContentTypeAuto, "Auto"},
	},
}

// TestContentType tests [ContentType] common methods and functions.
func TestContentType(t *testing.T) {
	testContentType.run(t)
}
