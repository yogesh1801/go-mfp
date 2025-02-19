// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color mode

package escl

import "testing"

var testADFOption = testEnum[ADFOption]{
	decodeStr: DecodeADFOption,
	decodeXML: decodeADFOption,
	dataset: []testEnumData[ADFOption]{
		{DetectPaperLoaded, "DetectPaperLoaded"},
		{SelectSinglePage, "SelectSinglePage"},
		{Duplex, "Duplex"},
	},
}

// TestADFOption tests [ADFOption] common methods and functions.
func TestADFOption(t *testing.T) {
	testADFOption.run(t)
}
