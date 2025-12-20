// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scan job state reason

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// JobStateReason defines one reason why a job is in its current state.
type JobStateReason int

// known job state reasons:
const (
	UnknownJobStateReason JobStateReason = iota
	InvalidScanTicket
	DocumentFormatError
	ImageTransferError
	JobCanceledAtDevice
	JobCompletedWithErrors
	JobCompletedWithWarnings
	JobScanning
	JobScanningAndTransferring
	JobTimedOut
	JobTransferring
	None
	ScannerStopped
)

// decodeJobStateReason decodes [JobStateReason] from the XML tree.
func decodeJobStateReason(root xmldoc.Element) (jsr JobStateReason, err error) {
	return decodeEnum(root, DecodeJobStateReason)
}

// toXML generates XML tree for the [JobStateReason].
func (jsr JobStateReason) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: jsr.String(),
	}
}

// String returns a string representation of the [JobStateReason]
func (jsr JobStateReason) String() string {
	switch jsr {
	case InvalidScanTicket:
		return "InvalidScanTicket"
	case DocumentFormatError:
		return "DocumentFormatError"
	case ImageTransferError:
		return "ImageTransferError"
	case JobCanceledAtDevice:
		return "JobCanceledAtDevice"
	case JobCompletedWithErrors:
		return "JobCompletedWithErrors"
	case JobCompletedWithWarnings:
		return "JobCompletedWithWarnings"
	case JobScanning:
		return "JobScanning"
	case JobScanningAndTransferring:
		return "JobScanningAndTransferring"
	case JobTimedOut:
		return "JobTimedOut"
	case JobTransferring:
		return "JobTransferring"
	case None:
		return "None"
	case ScannerStopped:
		return "ScannerStopped"
	}

	return "Unknown"
}

// DecodeJobStateReason decodes [JobStateReason] out of its XML string representation.
func DecodeJobStateReason(s string) JobStateReason {
	switch s {
	case "InvalidScanTicket":
		return InvalidScanTicket
	case "DocumentFormatError":
		return DocumentFormatError
	case "ImageTransferError":
		return ImageTransferError
	case "JobCanceledAtDevice":
		return JobCanceledAtDevice
	case "JobCompletedWithErrors":
		return JobCompletedWithErrors
	case "JobCompletedWithWarnings":
		return JobCompletedWithWarnings
	case "JobScanning":
		return JobScanning
	case "JobScanningAndTransferring":
		return JobScanningAndTransferring
	case "JobTimedOut":
		return JobTimedOut
	case "JobTransferring":
		return JobTransferring
	case "None":
		return None
	case "ScannerStopped":
		return ScannerStopped
	}

	return UnknownJobStateReason
}
