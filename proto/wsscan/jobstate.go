// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scan job state

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// JobState defines the current state of a scan job.
type JobState int

// known job states:
const (
	UnknownJobState     JobState = iota
	JobStateAborted              // The system aborted the job
	JobStateCanceled             // The job was canceled by a client or by means outside the scope of the WSD Scan Service
	JobStateCompleted            // The job is finished processing and all of the image data has been sent to the client
	JobStateCreating             // The job is being initialized
	JobStateHeld                 // The job is waiting to be processed but is unavailable for scheduling
	JobStatePending              // The job has been initialized and is waiting to be processed
	JobStateProcessing           // The job data is being digitized, transformed, or transferred
	JobStateStarted              // The scan device has started processing the job (transient state)
	JobStateTerminating          // The job was canceled or aborted and is terminating
)

// decodeJobState decodes [JobState] from the XML tree.
func decodeJobState(root xmldoc.Element) (js JobState, err error) {
	return decodeEnum(root, DecodeJobState)
}

// toXML generates XML tree for the [JobState].
func (js JobState) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: js.String(),
	}
}

// String returns a string representation of the [JobState]
func (js JobState) String() string {
	switch js {
	case JobStateAborted:
		return "Aborted"
	case JobStateCanceled:
		return "Canceled"
	case JobStateCompleted:
		return "Completed"
	case JobStateCreating:
		return "Creating"
	case JobStateHeld:
		return "Held"
	case JobStatePending:
		return "Pending"
	case JobStateProcessing:
		return "Processing"
	case JobStateStarted:
		return "Started"
	case JobStateTerminating:
		return "Terminating"
	}

	return "Unknown"
}

// DecodeJobState decodes [JobState] out of its XML string representation.
func DecodeJobState(s string) JobState {
	switch s {
	case "Aborted":
		return JobStateAborted
	case "Canceled":
		return JobStateCanceled
	case "Completed":
		return JobStateCompleted
	case "Creating":
		return JobStateCreating
	case "Held":
		return JobStateHeld
	case "Pending":
		return JobStatePending
	case "Processing":
		return JobStateProcessing
	case "Started":
		return JobStateStarted
	case "Terminating":
		return JobStateTerminating
	}

	return UnknownJobState
}
