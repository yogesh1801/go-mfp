// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Job state reason

package escl

import (
	"fmt"

	"github.com/alexpevzner/mfp/xmldoc"
)

// JobStateReason accompanies the [JobState] and gives additional
// information why job has reached the particular state.
//
// The Mopria eSCL specification doesn't provide detailed specification
// of this type and simply refers to pwd:JobStateReasons.
//
// So here we represent it as a string and list all known values, defined
// by the PWG. See JobStateReasonsWKV at the PWG site:
//
//	https://www.pwg.org/schemas/smpjt3d10/PwgWellKnownValues_xsd.html
//
// Most likely, not all these values are applicable for scanners.
type JobStateReason string

// Known values for JobStateReason
const (
	// Standard PWG values
	UnknownJobStateReason            JobStateReason = ""
	AbortedBySystem                  JobStateReason = "AbortedBySystem"
	AccountAuthorizationFailed       JobStateReason = "AccountAuthorizationFailed"
	AccountClosed                    JobStateReason = "AccountClosed"
	AccountInfoNeeded                JobStateReason = "AccountInfoNeeded"
	AccountLimitReached              JobStateReason = "AccountLimitReached"
	CompressionError                 JobStateReason = "CompressionError"
	ConflictingAttributes            JobStateReason = "ConflictingAttributes"
	ConnectedToDestination           JobStateReason = "ConnectedToDestination"
	ConnectingToDestination          JobStateReason = "ConnectingToDestination"
	DestinationURIFailed             JobStateReason = "DestinationUriFailed"
	DigitalSignatureDidNotVerify     JobStateReason = "DigitalSignatureDidNotVerify"
	DigitalSignatureTypeNotSupported JobStateReason = "DigitalSignatureTypeNotSupported"
	DocumentAccessError              JobStateReason = "DocumentAccessError"
	DocumentFormatError              JobStateReason = "DocumentFormatError"
	DocumentPasswordError            JobStateReason = "DocumentPasswordError"
	DocumentPermissionError          JobStateReason = "DocumentPermissionError"
	DocumentSecurityError            JobStateReason = "DocumentSecurityError"
	DocumentUnprintableError         JobStateReason = "DocumentUnprintableError"
	ErrorsDetected                   JobStateReason = "ErrorsDetected"
	JobCanceledAtDevice              JobStateReason = "JobCanceledAtDevice"
	JobCanceledByOperator            JobStateReason = "JobCanceledByOperator"
	JobCanceledByUser                JobStateReason = "JobCanceledByUser"
	JobCompletedSuccessfully         JobStateReason = "JobCompletedSuccessfully"
	JobCompletedWithErrors           JobStateReason = "JobCompletedWithErrors"
	JobCompletedWithWarnings         JobStateReason = "JobCompletedWithWarnings"
	JobDataInsufficient              JobStateReason = "JobDataInsufficient"
	JobDelayOutputUntilSpecified     JobStateReason = "JobDelayOutputUntilSpecified"
	JobDigitalSignatureWait          JobStateReason = "JobDigitalSignatureWait"
	JobFetchable                     JobStateReason = "JobFetchable"
	JobHeldForReview                 JobStateReason = "JobHeldForReview"
	JobHoldUntilSpecified            JobStateReason = "JobHoldUntilSpecified"
	JobIncoming                      JobStateReason = "JobIncoming"
	JobInterpreting                  JobStateReason = "JobInterpreting"
	JobOutgoing                      JobStateReason = "JobOutgoing"
	JobPasswordWait                  JobStateReason = "JobPasswordWait"
	JobPrintedSuccessfully           JobStateReason = "JobPrintedSuccessfully"
	JobPrintedWithErrors             JobStateReason = "JobPrintedWithErrors"
	JobPrintedWithWarnings           JobStateReason = "JobPrintedWithWarnings"
	JobPrinting                      JobStateReason = "JobPrinting"
	JobQueued                        JobStateReason = "JobQueued"
	JobQueuedForMarker               JobStateReason = "JobQueuedForMarker"
	JobReleaseWait                   JobStateReason = "JobReleaseWait"
	JobRestartable                   JobStateReason = "JobRestartable"
	JobResuming                      JobStateReason = "JobResuming"
	JobSavedSuccessfully             JobStateReason = "JobSavedSuccessfully"
	JobSavedWithErrors               JobStateReason = "JobSavedWithErrors"
	JobSavedWithWarnings             JobStateReason = "JobSavedWithWarnings"
	JobSaving                        JobStateReason = "JobSaving"
	JobSpooling                      JobStateReason = "JobSpooling"
	JobStreaming                     JobStateReason = "JobStreaming"
	JobSuspended                     JobStateReason = "JobSuspended"
	JobSuspendedByOperator           JobStateReason = "JobSuspendedByOperator"
	JobSuspendedBySystem             JobStateReason = "JobSuspendedBySystem"
	JobSuspendedByUser               JobStateReason = "JobSuspendedByUser"
	JobSuspending                    JobStateReason = "JobSuspending"
	JobTransferring                  JobStateReason = "JobTransferring"
	JobTransforming                  JobStateReason = "JobTransforming"
	None                             JobStateReason = "None"
	PrinterStopped                   JobStateReason = "PrinterStopped"
	PrinterStoppedPartly             JobStateReason = "PrinterStoppedPartly"
	ProcessingToStopPoint            JobStateReason = "ProcessingToStopPoint"
	QueuedInDevice                   JobStateReason = "QueuedInDevice"
	ResourcesAreNotReady             JobStateReason = "ResourcesAreNotReady"
	ResourcesAreNotSupported         JobStateReason = "ResourcesAreNotSupported"
	ServiceOffLine                   JobStateReason = "ServiceOffLine"
	SubmissionInterrupted            JobStateReason = "SubmissionInterrupted"
	UnsupportedAttributesOrValues    JobStateReason = "UnsupportedAttributesOrValues"
	UnsupportedCompression           JobStateReason = "UnsupportedCompression"
	UnsupportedDocumentFormat        JobStateReason = "UnsupportedDocumentFormat"
	WaitingForUserAction             JobStateReason = "WaitingForUserAction"
	WarningsDetected                 JobStateReason = "WarningsDetected"

	// Additional values, mentioned in the eSCL specification
	JobScanning                JobStateReason = "JobScanning"
	JobHeldByService           JobStateReason = "JobHeldByService"
	JobScanningAndTransferring JobStateReason = "JobScanningAndTransferring"
)

// decodeJobStateReason decodes [JobStateReason] from the XML tree.
func decodeJobStateReason(root xmldoc.Element) (reason JobStateReason, err error) {
	var v string
	v, err = decodeNMTOKEN(root)

	if err != nil {
		err = fmt.Errorf("invalid JobStateReason: %q",
			root.Text)
		err = xmldoc.XMLErrWrap(root, err)
		return
	}

	reason = JobStateReason(v)
	return
}

// toXML generates XML tree for the [JobStateReason].
func (reason JobStateReason) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: reason.String(),
	}
}

// String returns a string representation of the [JobStateReasons]
//
// Although JobStateReason already defines as a string, we provide
// this method for consistency with other similar types.
func (reason JobStateReason) String() string {
	return string(reason)
}
