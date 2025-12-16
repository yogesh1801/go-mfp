// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP enums

package ipp

import "reflect"

// EnJobState represents "job-state" values.
//
// See RFC8011, 5.3.7.
type EnJobState int

const (
	// EnJobStatePending means the job is candidate to start processing,
	// but not yet processing.
	EnJobStatePending EnJobState = 3

	// EnJobStatePendingHeld means the job is not candidate for processing
	// until hold is removed.
	EnJobStatePendingHeld EnJobState = 4

	// EnJobStateProcessing means the job is being processed.
	EnJobStateProcessing EnJobState = 5

	// EnJobStateProcessingStopped means the job processing has been stopped
	// for some reason.
	EnJobStateProcessingStopped EnJobState = 6

	// EnJobStateCanceled means the job has been canceled by the
	// Cancel-Job operation.
	EnJobStateCanceled EnJobState = 7

	// EnJobStateAborted means the job has been aborted by the system.
	EnJobStateAborted EnJobState = 8

	// EnJobStateCompleted means the job has been completed.
	EnJobStateCompleted EnJobState = 9
)

// kwRegisteredTypes lists all registered keyword types for IPP codec.
var enRegisteredTypes = map[reflect.Type]struct{}{
	reflect.TypeOf(EnJobState(0)):    struct{}{},
	reflect.TypeOf(EnPrinterType(0)): struct{}{},
}
