// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Job state

package escl

import "testing"

var testJobState = testEnum[JobState]{
	decodeStr: DecodeJobState,
	decodeXML: decodeJobState,
	dataset: []testEnumData[JobState]{
		{JobCanceled, "Canceled"},
		{JobAborted, "Aborted"},
		{JobCompleted, "Completed"},
		{JobPending, "Pending"},
		{JobProcessing, "Processing"},
		{JobPendingHeld, "PendingHeld"},
	},
}

// TestJobState tests [JobState] common methods and functions.
func TestJobState(t *testing.T) {
	testJobState.run(t)
}
