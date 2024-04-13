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

	// KwCompressionDeflate is RFC 1951 ZIP inflate/deflate
	KwCompressionDeflate KwCompression = "deflate"

	// KwCompressionGzip is RFC 1952 GNU zip
	KwCompressionGzip KwCompression = "gzip"

	// KwCompressionCompress is RFC 1977 UNIX compression
	KwCompressionCompress KwCompression = "compress"
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

// KwJobSpooling represents standard keyword values for
// "job-spooling-supported" attribute.
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

// kwRegisteredTypes lists all registered keyword types for IPP codec.
var kwRegisteredTypes = map[reflect.Type]struct{}{
	reflect.TypeOf(KwCompression("")):         struct{}{},
	reflect.TypeOf(KwPdlOverride("")):         struct{}{},
	reflect.TypeOf(KwPrinterStateReasons("")): struct{}{},
	reflect.TypeOf(KwURIAuthentication("")):   struct{}{},
	reflect.TypeOf(KwURISecurity("")):         struct{}{},
}
