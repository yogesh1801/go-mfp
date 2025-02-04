// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan intent test

package escl

import "testing"

var testIntent = testEnum[Intent]{
	decodeStr: DecodeIntent,
	decodeXML: decodeIntent,
	ns:        NsScan,
	dataset: []testEnumData[Intent]{
		{Document, "Document"},
		{TextAndGraphic, "TextAndGraphic"},
		{Photo, "Photo"},
		{Preview, "Preview"},
		{Object, "Object"},
		{BusinessCard, "BusinessCard"},
	},
}

// TestIntent tests [Intent] common methods and functions.
func TestIntent(t *testing.T) {
	testIntent.run(t)
}
