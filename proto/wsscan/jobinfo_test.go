// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Tests for jobList

package wsscan

import (
	"testing"
)

// TestJobListGet tests the get method.
func TestJobListGet(t *testing.T) {
	var jl jobList

	// get on empty list returns nil
	if jl.get(1) != nil {
		t.Error("get on empty list: expected nil")
	}

	jl.put(jobInfo{jobID: 1, jobToken: "a", state: JobStateProcessing})
	jl.put(jobInfo{jobID: 2, jobToken: "b", state: JobStateCompleted})

	// get existing job
	j := jl.get(1)
	if j == nil {
		t.Fatal("get(1): expected job, got nil")
	}
	if j.jobToken != "a" {
		t.Errorf("get(1).jobToken: expected %q, got %q", "a", j.jobToken)
	}

	j = jl.get(2)
	if j == nil {
		t.Fatal("get(2): expected job, got nil")
	}
	if j.jobToken != "b" {
		t.Errorf("get(2).jobToken: expected %q, got %q", "b", j.jobToken)
	}

	// get non-existing job returns nil
	if jl.get(99) != nil {
		t.Error("get(99): expected nil for non-existing job")
	}
}

// TestJobListPutInsert tests that put inserts new jobs and all are findable.
func TestJobListPutInsert(t *testing.T) {
	var jl jobList

	jl.put(jobInfo{jobID: 1, state: JobStateProcessing})
	jl.put(jobInfo{jobID: 2, state: JobStateCompleted})
	jl.put(jobInfo{jobID: 3, state: JobStateCanceled})

	if len(jl) != 3 {
		t.Fatalf("len: expected 3, got %d", len(jl))
	}

	for _, id := range []int{1, 2, 3} {
		if jl.get(id) == nil {
			t.Errorf("get(%d): expected job, got nil", id)
		}
	}
}

// TestJobListPutUpdate tests that put updates an existing job in place.
func TestJobListPutUpdate(t *testing.T) {
	var jl jobList

	jl.put(jobInfo{jobID: 1, state: JobStateProcessing})
	jl.put(jobInfo{jobID: 2, state: JobStateProcessing})

	// Update job 1 state
	jl.put(jobInfo{jobID: 1, state: JobStateCompleted, jobToken: "updated"})

	if len(jl) != 2 {
		t.Errorf("len: expected 2, got %d", len(jl))
	}

	j := jl.get(1)
	if j == nil {
		t.Fatal("get(1): expected job, got nil")
	}
	if j.state != JobStateCompleted {
		t.Errorf("state: expected Completed, got %s", j.state)
	}
	if j.jobToken != "updated" {
		t.Errorf("jobToken: expected %q, got %q", "updated", j.jobToken)
	}
}

// TestJobListGetReturnsPointer tests that get returns a pointer into
// the list so the caller can update the job in place.
func TestJobListGetReturnsPointer(t *testing.T) {
	var jl jobList

	jl.put(jobInfo{jobID: 1, state: JobStateProcessing})

	j := jl.get(1)
	if j == nil {
		t.Fatal("get(1): expected job, got nil")
	}

	// Mutate via pointer
	j.state = JobStateCompleted

	// Change must be visible in the list
	if jl.get(1).state != JobStateCompleted {
		t.Error("mutation via pointer not reflected in list")
	}
}

// TestJobListHistoryLimit tests that the list is trimmed to
// abstractServerJobHistorySize.
func TestJobListHistoryLimit(t *testing.T) {
	var jl jobList

	// Insert more jobs than the limit
	for i := 1; i <= abstractServerJobHistorySize+5; i++ {
		jl.put(jobInfo{jobID: i})
	}

	if len(jl) != abstractServerJobHistorySize {
		t.Errorf("len: expected %d, got %d",
			abstractServerJobHistorySize, len(jl))
	}

	// The 5 oldest jobs (IDs 1-5) must have been dropped
	for i := 1; i <= 5; i++ {
		if jl.get(i) != nil {
			t.Errorf("get(%d): expected nil (dropped), got job", i)
		}
	}

	// The 10 most recent jobs (IDs 6-15) must still be present
	for i := 6; i <= abstractServerJobHistorySize+5; i++ {
		if jl.get(i) == nil {
			t.Errorf("get(%d): expected job, got nil", i)
		}
	}
}
