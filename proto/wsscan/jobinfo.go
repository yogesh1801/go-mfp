// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Job info: per-job state tracked by the AbstractServer

package wsscan

import "time"

// abstractServerJobHistorySize is the maximum number of jobs
// the AbstractServer keeps in its history.
const abstractServerJobHistorySize = 10

// jobInfo holds the server-side state for a single scan job.
type jobInfo struct {
	jobID          int
	jobToken       string
	state          JobState
	scanTicket     ScanTicket
	scansCompleted int
	createdTime    time.Time
	completedTime  time.Time
}

// jobList is a bounded list of [jobInfo] entries.
// Its maximum size is [abstractServerJobHistorySize].
type jobList []jobInfo

// get returns a pointer to the [jobInfo] with the given jobID,
// or nil if no such job exists.
func (jl jobList) get(jobID int) *jobInfo {
	for i := range jl {
		if jl[i].jobID == jobID {
			return &jl[i]
		}
	}
	return nil
}

// put inserts or replaces a [jobInfo] in the list.
//
// If a job with the same jobID already exists it is updated in place.
// Otherwise the job is appended and the oldest entry is dropped if
// the list exceeds [abstractServerJobHistorySize].
func (jl *jobList) put(job jobInfo) {
	// Update in place if job already exists.
	for i := range *jl {
		if (*jl)[i].jobID == job.jobID {
			(*jl)[i] = job
			return
		}
	}

	// Append new job, dropping the oldest entry if over the limit.
	*jl = append(*jl, job)
	if len(*jl) > abstractServerJobHistorySize {
		*jl = (*jl)[1:]
	}
}
