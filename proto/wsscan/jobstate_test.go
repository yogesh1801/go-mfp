// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for scan job state

package wsscan

import "testing"

var testJobState = testEnum[JobState]{
	decodeStr: DecodeJobState,
	decodeXML: decodeJobState,
	dataset: []testEnumData[JobState]{
		{JobStateAborted, "Aborted"},
		{JobStateCanceled, "Canceled"},
		{JobStateCompleted, "Completed"},
		{JobStateCreating, "Creating"},
		{JobStateHeld, "Held"},
		{JobStatePending, "Pending"},
		{JobStateProcessing, "Processing"},
		{JobStateStarted, "Started"},
		{JobStateTerminating, "Terminating"},
	},
}

// TestJobState tests [JobState] common methods and functions.
func TestJobState(t *testing.T) {
	testJobState.run(t)
}
