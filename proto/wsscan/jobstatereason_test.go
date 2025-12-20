// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for scan job state reason

package wsscan

import "testing"

var testJobStateReason = testEnum[JobStateReason]{
	decodeStr: DecodeJobStateReason,
	decodeXML: decodeJobStateReason,
	dataset: []testEnumData[JobStateReason]{
		{InvalidScanTicket, "InvalidScanTicket"},
		{DocumentFormatError, "DocumentFormatError"},
		{ImageTransferError, "ImageTransferError"},
		{JobCanceledAtDevice, "JobCanceledAtDevice"},
		{JobCompletedWithErrors, "JobCompletedWithErrors"},
		{JobCompletedWithWarnings, "JobCompletedWithWarnings"},
		{JobScanning, "JobScanning"},
		{JobScanningAndTransferring, "JobScanningAndTransferring"},
		{JobTimedOut, "JobTimedOut"},
		{JobTransferring, "JobTransferring"},
		{None, "None"},
		{ScannerStopped, "ScannerStopped"},
	},
}

// TestJobStateReason tests [JobStateReason] common methods and functions.
func TestJobStateReason(t *testing.T) {
	testJobStateReason.run(t)
}
