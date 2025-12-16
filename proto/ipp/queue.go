// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Job Queue

package ipp

import (
	"math"
	"strings"
	"sync"

	"github.com/OpenPrinting/go-mfp/util/uuid"
)

// queue manages queue of jobs
type queue struct {
	lock   sync.Mutex      // Access lock
	nextid int32           // Next JOB ID
	jobs   []*job          // Queued jobs
	byID   map[int]*job    // Job by JobID
	byURI  map[string]*job // Job by JobURI
}

// newQueue creates new queue.
func newQueue() *queue {
	q := &queue{
		nextid: 1,
		byID:   make(map[int]*job),
		byURI:  make(map[string]*job),
	}
	return q
}

// Push pushes new job into the queue.
// It assigns the job.JobId
func (q *queue) Push(j *job) {
	q.lock.Lock()
	defer q.lock.Unlock()

	j.JobID = q.allocJobID()

	q.jobs = append(q.jobs, j)
	q.byID[j.JobID] = j
	q.byURI[j.JobURI] = j
}

// allocJobID allocates the next JobID.
// It must be called under q.lock.
func (q *queue) allocJobID() int {
	for {
		id := int(q.nextid)
		if q.nextid == math.MaxInt32 {
			q.nextid = 1
		} else {
			q.nextid++
		}

		if q.byID[id] == nil {
			return id
		}
	}
}

// job represents a single job in the queue
type job struct {
	JobStatus          // Job status attributes
	JobCreateOperation // Job create-time operation attributes
	JobAttributes      // Job creation attributes
}

// newJob creates a new job.
func newJob(ops *JobCreateOperation, attrs *JobAttributes) *job {
	uu := uuid.Must(uuid.Random())
	uri := strings.Join([]string{ops.PrinterURI, "jobs", uu.String()}, "/")

	j := &job{
		JobStatus: JobStatus{
			JobURI:          uri,
			JobState:        EnJobStatePendingHeld,
			JobStateReasons: []KwJobStateReasons{KwJobStateReasonsJobIncoming},
		},
		JobCreateOperation: *ops,
		JobAttributes:      *attrs,
	}

	return j
}
