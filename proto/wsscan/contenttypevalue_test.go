// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ContentTypeValue

package wsscan

import "testing"

var testContentTypeValue = testEnum[ContentTypeValue]{
	decodeStr: DecodeContentTypeValue,
	decodeXML: decodeContentTypeValue,
	dataset: []testEnumData[ContentTypeValue]{
		{Auto, "Auto"},
		{Text, "Text"},
		{Photo, "Photo"},
		{Halftone, "Halftone"},
		{Mixed, "Mixed"},
	},
}

// TestContentTypeValue tests [ContentTypeValue] common methods and functions.
func TestContentTypeValue(t *testing.T) {
	testContentTypeValue.run(t)
}
