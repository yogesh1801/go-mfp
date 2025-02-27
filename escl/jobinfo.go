// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package escl

import (
	"time"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/uuid"
)

// JobInfo reports the state of a particular scan job.
//
// eSCL Technical Specification, 9.1.
type JobInfo struct {
	JobUUID          optional.Val[uuid.UUID]     // Unique, persistent
	JobURI           string                      // Unique Job URI, that identifies the job
	Age              optional.Val[time.Duration] // Time since last update
	ImagesCompleted  optional.Val[int]           // Images completed so far
	ImagesToTransfer optional.Val[int]           // Images to transfer
	JobState         JobState                    // Job state
	JobStateReasons  JobStateReason              // Reason of the job state
}
