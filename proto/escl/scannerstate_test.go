// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Job state

package escl

import "testing"

var testScannerState = testEnum[ScannerState]{
	decodeStr: DecodeScannerState,
	decodeXML: decodeScannerState,
	dataset: []testEnumData[ScannerState]{
		{ScannerIdle, "Idle"},
		{ScannerProcessing, "Processing"},
		{ScannerTesting, "Testing"},
		{ScannerStopped, "Stopped"},
		{ScannerDown, "Down"},
	},
}

// TestScannerState tests [ScannerState] common methods and functions.
func TestScannerState(t *testing.T) {
	testScannerState.run(t)
}
