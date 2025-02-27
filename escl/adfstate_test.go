// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF state test

package escl

import "testing"

var testADFState = testEnum[ADFState]{
	decodeStr: DecodeADFState,
	decodeXML: decodeADFState,
	dataset: []testEnumData[ADFState]{
		{ScannerAdfProcessing, "ScannerAdfProcessing"},
		{ScannerAdfEmpty, "ScannerAdfEmpty"},
		{ScannerAdfJam, "ScannerAdfJam"},
		{ScannerAdfLoaded, "ScannerAdfLoaded"},
		{ScannerAdfMispick, "ScannerAdfMispick"},
		{ScannerAdfHatchOpen, "ScannerAdfHatchOpen"},
		{ScannerAdfDuplexPageTooShort, "ScannerAdfDuplexPageTooShort"},
		{ScannerAdfDuplexPageTooLong, "ScannerAdfDuplexPageTooLong"},
		{ScannerAdfMultipickDetected, "ScannerAdfMultipickDetected"},
		{ScannerAdfInputTrayFailed, "ScannerAdfInputTrayFailed"},
		{ScannerAdfInputTrayOverloaded, "ScannerAdfInputTrayOverloaded"},
	},
}

// TestADFState tests [ADFState] common methods and functions.
func TestADFState(t *testing.T) {
	testADFState.run(t)
}
