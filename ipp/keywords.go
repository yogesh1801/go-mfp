// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP keywords

package ipp

import (
	"reflect"
	"strings"
)

// KwCompression represents standard keyword values for compression.
//
// See RFC8011, 5.4.32.
type KwCompression string

const (
	// KwCompressionNone is no compression
	KwCompressionNone KwCompression = "none"

	// KwCompressionDeflate is RFC1951 ZIP inflate/deflate
	KwCompressionDeflate KwCompression = "deflate"

	// KwCompressionGzip is RFC1952 GNU zip
	KwCompressionGzip KwCompression = "gzip"

	// KwCompressionCompress is RFC1977 UNIX compression
	KwCompressionCompress KwCompression = "compress"
)

// KwJobDelayOutputUntil represents standard keyword values for
// "job-delay-output-until" attribute.
//
// See PWG5100.7, 6.8.4.
type KwJobDelayOutputUntil string

const (
	// KwJobDelayOutputUntilDayTime means delay output until daylight,
	// typically 6am to 6pm.
	KwJobDelayOutputUntilDayTime KwJobDelayOutputUntil = "day-time"

	// KwJobDelayOutputUntilEvening means delay output until the evening,
	// typically from 6pm to 12am.
	KwJobDelayOutputUntilEvening KwJobDelayOutputUntil = "evening"

	// KwJobDelayOutputUntilIndefinite means delay output indefinitely;
	// the time period can be changed using the Set- Job-Attributes
	// operation
	KwJobDelayOutputUntilIndefinite KwJobDelayOutputUntil = "indefinite"

	// KwJobDelayOutputUntilNight means delay output until the night,
	// typically from 12am to 6am.
	KwJobDelayOutputUntilNight KwJobDelayOutputUntil = "night"

	// KwJobDelayOutputUntilNoDelayOutput means do not delay the output.
	KwJobDelayOutputUntilNoDelayOutput KwJobDelayOutputUntil = "no-delay-output"

	// KwJobDelayOutputUntilSecondShift means delay output until the
	// second work shift, typically from 4pm to 12am.
	KwJobDelayOutputUntilSecondShift KwJobDelayOutputUntil = "second-shift"

	// KwJobDelayOutputUntilThirdShift means delay output until the
	// third work shift, typically from 12am to 8am.
	KwJobDelayOutputUntilThirdShift KwJobDelayOutputUntil = "third-shift"

	// KwJobDelayOutputUntilWeekend means delay output until the weekend.
	KwJobDelayOutputUntilWeekend KwJobDelayOutputUntil = "weekend"
)

// KwJobHoldUntil represents standard keyword values for
// "job-hold-until" attribute.
//
// See PWG5100.7, 5.2.2.
type KwJobHoldUntil string

const (
	// KwJobHoldUntilNoHold means do not delay the output.
	KwJobHoldUntilNoHold KwJobHoldUntil = "no-hold"

	// KwJobHoldUntilIndefinite means delay output indefinitely;
	// the time period can be changed using the Set- Job-Attributes
	// operation
	KwJobHoldUntilIndefinite KwJobHoldUntil = "indefinite"

	// KwJobHoldUntilDayTime means delay output until daylight,
	// typically 6am to 6pm.
	KwJobHoldUntilDayTime KwJobHoldUntil = "day-time"

	// KwJobHoldUntilEvening means delay output until the evening,
	// typically from 6pm to 12am.
	KwJobHoldUntilEvening KwJobHoldUntil = "evening"

	// KwJobHoldUntilNight means delay output until the night,
	// typically from 12am to 6am.
	KwJobHoldUntilNight KwJobHoldUntil = "night"

	// KwJobHoldUntilWeekend means delay output until the weekend.
	KwJobHoldUntilWeekend KwJobHoldUntil = "weekend"

	// KwJobHoldUntilSecondShift means delay output until the
	// second work shift, typically from 4pm to 12am.
	KwJobHoldUntilSecondShift KwJobHoldUntil = "second-shift"

	// KwJobHoldUntilThirdShift means delay output until the
	// third work shift, typically from 12am to 8am.
	KwJobHoldUntilThirdShift KwJobHoldUntil = "third-shift"
)

// KwJobSheets represents standard keyword values for
// "job-sheets" attribute.
//
// See: RFC8011, 5.2.3., PWG5100.7, 8.1.
type KwJobSheets string

const (
	// ----- RFC8011 values -----

	// KwJobSheetsNone means no Job sheet is printed
	KwJobSheetsNone KwJobSheets = "none"

	// KwJobSheetsStandard means one or more site-specific
	// standard Job sheets are printed
	KwJobSheetsStandard KwJobSheets = "standard"

	// ----- PWG5100.7 values -----

	// KwJobSheetsJobStartSheet means that a ob Sheet is printed to
	// indicate the start of the Job.
	KwJobSheetsJobStartSheet KwJobSheets = "job-start-sheet "

	// KwJobSheetsJobEndSheet means that a ob Sheet is printed to
	// indicate the end of the Job.
	KwJobSheetsJobEndSheet KwJobSheets = "job-end-sheet"

	// KwJobSheetsJobBothSheets instructs Printer to print
	// Job Sheets o indicate the start and end of all the
	// output associated with the Job.
	KwJobSheetsJobBothSheets KwJobSheets = "job-both-sheets"

	// KwJobSheetsFirstPrintStreamPage instructs Printer to print the
	// first input Page in the Document Data as the Job Sheet.
	// The Printer's standard Job Sheet is suppressed.
	KwJobSheetsFirstPrintStreamPage KwJobSheets = "first-print-stream-page"
)

// KwJobSpooling represents standard keyword values for
// "job-spooling-supported" attribute.
//
// See PWG5100.7, 6.9.31.
type KwJobSpooling string

const (
	// KwJobSpoolingAutomatic means that it is up to the
	// Printer when to process the Document data.
	KwJobSpoolingAutomatic KwJobSpooling = "automatic"

	// KwJobSpoolingSpool means that the Document data is
	// processed after it has been spooled (stored).
	KwJobSpoolingSpool KwJobSpooling = "spool"

	// KwJobSpoolingStream means that the Document data is
	// processed as it is received.
	KwJobSpoolingStream KwJobSpooling = "stream"
)

// KwPdlOverride represents standard keyword values for
// "pdl-override-supported" attribute.
//
// See RFC8011, 5.4.28.
type KwPdlOverride string

const (
	// KwPdlOverrideAattempted indicates that Printer attempts to
	// make the IPP attribute values take precedence over embedded
	// instructions in the Document data.
	KwPdlOverrideAattempted KwPdlOverride = "attempted"

	// KwPdlOverrideNotAttempted indicates that the Printer makes no
	// attempt to make the IPP attribute values take precedence over
	// embedded instructions in the Document data.
	KwPdlOverrideNotAttempted KwPdlOverride = "not-attempted"
)

// KwPrinterStateReasons represents standard keyword values for
// "printer-state-reasons" attribute.
//
// See RFC8011, 5.4.12.
type KwPrinterStateReasons string

// Standard values for KwPrinterStateReasons
const (
	// The following are standard keyword values of "printer-state-reasons"
	// attribute:
	KwPrinterStateNone                           KwPrinterStateReasons = "none"
	KwPrinterStateOther                          KwPrinterStateReasons = "other"
	KwPrinterStateConnectingToDevice             KwPrinterStateReasons = "connecting-to-device"
	KwPrinterStateCoverOpen                      KwPrinterStateReasons = "cover-open"
	KwPrinterStateDeveloperEmpty                 KwPrinterStateReasons = "developer-empty"
	KwPrinterStateDeveloperLow                   KwPrinterStateReasons = "developer-low"
	KwPrinterStateDoorOpen                       KwPrinterStateReasons = "door-open"
	KwPrinterStateFuserOverTemp                  KwPrinterStateReasons = "fuser-over-temp"
	KwPrinterStateFuserUnderTemp                 KwPrinterStateReasons = "fuser-under-temp"
	KwPrinterStateInputTrayMissing               KwPrinterStateReasons = "input-tray-missing"
	KwPrinterStateInterlockOpen                  KwPrinterStateReasons = "interlock-open"
	KwPrinterStateInterpreterResourceUnavailable KwPrinterStateReasons = "interpreter-resource-unavailable"
	KwPrinterStateMarkerSupplyEmpty              KwPrinterStateReasons = "marker-supply-empty"
	KwPrinterStateMarkerSupplyLow                KwPrinterStateReasons = "marker-supply-low"
	KwPrinterStateMarkerWasteAlmostFull          KwPrinterStateReasons = "marker-waste-almost-full"
	KwPrinterStateMarkerWasteFull                KwPrinterStateReasons = "marker-waste-full"
	KwPrinterStateMediaEmpty                     KwPrinterStateReasons = "media-empty"
	KwPrinterStateMediaJam                       KwPrinterStateReasons = "media-jam"
	KwPrinterStateMediaLow                       KwPrinterStateReasons = "media-low"
	KwPrinterStateMediaNeeded                    KwPrinterStateReasons = "media-needed"
	KwPrinterStateMovingToPaused                 KwPrinterStateReasons = "moving-to-paused"
	KwPrinterStateOpcLifeOver                    KwPrinterStateReasons = "opc-life-over"
	KwPrinterStateOpcNearEol                     KwPrinterStateReasons = "opc-near-eol"
	KwPrinterStateOutputAreaAlmostFull           KwPrinterStateReasons = "output-area-almost-full"
	KwPrinterStateOutputAreaFull                 KwPrinterStateReasons = "output-area-full"
	KwPrinterStateOutputTrayMissing              KwPrinterStateReasons = "output-tray-missing"
	KwPrinterStatePaused                         KwPrinterStateReasons = "paused"
	KwPrinterStateShutdown                       KwPrinterStateReasons = "shutdown"
	KwPrinterStateSpoolAreaFull                  KwPrinterStateReasons = "spool-area-full"
	KwPrinterStateStoppedPartly                  KwPrinterStateReasons = "stopped-partly"
	KwPrinterStateStopping                       KwPrinterStateReasons = "stopping"
	KwPrinterStateTimedOut                       KwPrinterStateReasons = "timed-out"
	KwPrinterStateTonerEmpty                     KwPrinterStateReasons = "toner-empty"
	KwPrinterStateTonerLow                       KwPrinterStateReasons = "toner-low"

	// "printer-state-reasons" may also have one of the following
	// standard suffixes, indicating its level of severity:
	KwPrinterStateReport  KwPrinterStateReasons = "-report"
	KwPrinterStateWarning KwPrinterStateReasons = "-warning"
	KwPrinterStateError   KwPrinterStateReasons = "-error"
)

// Split splits KwPrinterStateReasons string into reason and severity suffix
// (which can be one of "-report", "-warning" or "-error".
//
// If string contains one of the known suffixes, reason itself and
// suffix will be returned separately. Otherwise, the first returned
// value will be unmodified string, and second will be "":
//
//     "media-low-warning"  ->  "media-low", "-warning"
//     "media-jam-error"    ->  "media-jam", "-error"
//     "shutdown"           ->  "shutdown",  ""
func (s KwPrinterStateReasons) Split() (
	reason, severity KwPrinterStateReasons) {

	idx := strings.LastIndexByte(string(s), '-')

	if idx >= 0 {
		prefix, suffix := s[:idx], s[idx:]
		switch suffix {
		case KwPrinterStateReport,
			KwPrinterStateWarning,
			KwPrinterStateError:
			return prefix, suffix
		}
	}

	return s, ""
}

// Reason splits KwPrinterStateReasons into reason and severity
// and returns its reason part.
func (s KwPrinterStateReasons) Reason() KwPrinterStateReasons {
	reason, _ := s.Split()
	return reason
}

// Severity splits KwPrinterStateReasons into reason and severity
// and returns its severity part.
func (s KwPrinterStateReasons) Severity() KwPrinterStateReasons {
	_, severity := s.Split()
	return severity
}

// KwURIAuthentication represents standard keyword values for
// "uri-authentication-supported" attribute.
//
// See RFC8011, 5.4.2.
type KwURIAuthentication string

const (
	// KwURIAuthenticationNone means that there is no
	// authentication mechanism associated with the URI.
	KwURIAuthenticationNone KwURIAuthentication = "none"

	// KwURIAuthenticationRequestingUserName means that Client
	// sends authenticated user by the "requesting-user-name"
	// operation attribute.
	KwURIAuthenticationRequestingUserName KwURIAuthentication = "requesting-user-name"

	// KwURIAuthenticationBasic means HTTP basic authentication.
	KwURIAuthenticationBasic KwURIAuthentication = "basic"

	// KwURIAuthenticationDigest means HTTP digest authentication.
	KwURIAuthenticationDigest KwURIAuthentication = "digest"

	// KwURIAuthenticationCertificate means TLS authentication
	// based on X.509 certificates.
	KwURIAuthenticationCertificate KwURIAuthentication = "certificate"
)

// KwURISecurity represents standard keyword values for
// "uri-security-supported" attribute.
//
// See RFC8011, 5.4.3.
type KwURISecurity string

const (
	// KwURISecurityNone means that there is no secure communication
	// channel in use for given URI
	KwURISecurityNone KwURISecurity = "none"

	// KwURISecurityTLS indicates TLS security
	KwURISecurityTLS KwURISecurity = "tls"
)

// KwMultipleDocumentHandling represents standard keyword values for
// "multiple-document-handling" attribute.
//
// See RFC8011, 5.2.4.
type KwMultipleDocumentHandling string

// KwMultipleDocumentHandling constants indicate, how Printer handles
// Job with multiple documents.
//
// Imagine we have a Job with documents "a" and "b". Also, we may
// request a single copy or multiple copies. If multiple copies of,
// say, document "a" are requested,  we will denote it as a(*).
//
// So, depending of the "multiple-document-handling" parameter,
// Job with multiple documents "a" and "b" will be printed
// as follows (square brakes means finishing process):
//
//   "single-document"
//   Documents are concatenated,  and "b" may start at the last page of "a",
//   but each copy of a+b concatenation starts at its own page.
//     Sungle copy:       [a+b]
//     Multiple copies:   [a+b] [a+b] ... [a+b]
//
//   "single-document-new-sheet"
//   Like "single-document", but "b" starts at its own page:
//     Single copy:       [a b]
//     Multiple copies:   [a b] [a b] ... [a b]
//
//   "separate-documents-uncollated-copies"
//   Each document handled separately. First printed all copies
//   of "a", then all copies of "b":
//     Single copy:       [a] [b]
//     Multiple copies:   [a] [a] ... [a]  [b] [b] ... [b]
//
//   "separate-documents-collated-copies"
//   Like "separate-documents-uncollated-copies", but in a case
//   of multiple copies, ordering is different:
//     Single copy:       [a] [b]
//     Multiple copies:   [a] [b] [a] [b] ... [a] [b]
const (
	KwMultipleDocumentHandlingSingleDocument         KwMultipleDocumentHandling = "single-document"
	KwMultipleDocumentHandlingSingleDocumentNewSheet KwMultipleDocumentHandling = "single-document-new-sheet"

	KwMultipleDocumentHandlingSeparateDocumentsUncollatedCopies KwMultipleDocumentHandling = "separate-documents-uncollated-copies"
	KwMultipleDocumentHandlingSeparateDocumentsCollatedCopies   KwMultipleDocumentHandling = "separate-documents-collated-copies"
)

// kwRegisteredTypes lists all registered keyword types for IPP codec.
var kwRegisteredTypes = map[reflect.Type]struct{}{
	reflect.TypeOf(KwCompression("")):              struct{}{},
	reflect.TypeOf(KwJobDelayOutputUntil("")):      struct{}{},
	reflect.TypeOf(KwJobHoldUntil("")):             struct{}{},
	reflect.TypeOf(KwJobSheets("")):                struct{}{},
	reflect.TypeOf(KwJobSpooling("")):              struct{}{},
	reflect.TypeOf(KwPdlOverride("")):              struct{}{},
	reflect.TypeOf(KwPrinterStateReasons("")):      struct{}{},
	reflect.TypeOf(KwURIAuthentication("")):        struct{}{},
	reflect.TypeOf(KwURISecurity("")):              struct{}{},
	reflect.TypeOf(KwMultipleDocumentHandling("")): struct{}{},
}
