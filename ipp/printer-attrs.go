// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Printer Attributes

package ipp

import (
	"github.com/OpenPrinting/goipp"
)

// Default values for common attributes
var (
	// DefaultCharsetConfigured is the default value for
	// ""charset-configured" printer attribute
	DefaultCharsetConfigured = "utf-8"

	// DefaultCharsetSupported is the default value for
	// ""charset-supported" printer attribute
	DefaultCharsetSupported = []string{DefaultCharsetConfigured}

	// DefaultIppVersionsSupported is the default value for
	// "ipp-features-supported" printer attribute
	DefaultIppVersionsSupported = []goipp.Version{
		goipp.MakeVersion(2, 0),
		goipp.MakeVersion(1, 0),
		goipp.MakeVersion(1, 1),
	}
)

// Standard keywords for CompressionSupported attribute
const (
	CompressionDeflate  = "deflate"  // RFC 1951 ZIP inflate/deflate
	CompressionGzip     = "gzip"     // RFC 1952 GNU zip
	CompressionCompress = "compress" // RFC 1977 UNIX compression
)

// PrinterAttributes represents IPP Printer Attributes
type PrinterAttributes struct {
	// RFC8011, 5.4: Printer Description and Status Attributes
	CharsetConfigured                 string                         `ipp:"charset-configured,charset"`
	CharsetSupported                  []string                       `ipp:"charset-supported,charset"`
	ColorSupported                    bool                           `ipp:"?color-supported"`
	CompressionSupported              []string                       `ipp:"compression-supported,keyword"`
	DocumentFormatDefault             string                         `ipp:"document-format-default,mimeMediaType"`
	DocumentFormatSupported           []string                       `ipp:"document-format-supported,mimeMediaType"`
	GeneratedNaturalLanguageSupported []string                       `ipp:"generated-natural-language-supported,naturalLanguage"`
	IppFeaturesSupported              []string                       `ipp:"ipp-features-supported,keyword"`
	IppVersionsSupported              []goipp.Version                `ipp:"ipp-versions-supported"`
	JobImpressionsSupported           goipp.Range                    `ipp:"?job-impressions-supported"`
	JobKOctetsSupported               goipp.Range                    `ipp:"?job-k-octets-supported"`
	JobMediaSheetsSupported           goipp.Range                    `ipp:"?job-media-sheets-supported"`
	MediaSizeSupported                []PrinterMediaSizeSupported    `ipp:"media-size-supported,norange"`
	MediaSizeSupportedRange           PrinterMediaSizeSupportedRange `ipp:"media-size-supported,range"`
	MultipleDocumentJobsSupported     bool                           `ipp:"?multiple-document-jobs-supported"`
	MultipleOperationTimeOut          int                            `ipp:"?multiple-operation-time-out"`
	NaturalLanguageConfigured         string                         `ipp:"natural-language-configured,naturalLanguage"`
	OperationsSupported               []goipp.Op                     `ipp:"operations-supported,enum"`
	PagesPerMinuteColor               int                            `ipp:"?pages-per-minute-color"`
	PagesPerMinute                    int                            `ipp:"?pages-per-minute"`
	PdlOverrideSupported              string                         `ipp:"pdl-override-supported,keyword"`
	PrinterDriverInstaller            string                         `ipp:"?printer-driver-installer,uri"`
	PrinterInfo                       string                         `ipp:"?printer-info,text"`
	PrinterIsAcceptingJobs            bool                           `ipp:"printer-is-accepting-jobs"`
	PrinterLocation                   string                         `ipp:"?printer-location,text"`
	PrinterMakeAndModel               string                         `ipp:"?printer-make-and-model,text"`
	PrinterMessageFromOperator        string                         `ipp:"?printer-message-from-operator,text"`
	PrinterMoreInfoManufacturer       string                         `ipp:"?printer-more-info-manufacturer,uri"`
	PrinterMoreInfo                   string                         `ipp:"?printer-more-info,uri"`
	PrinterName                       string                         `ipp:"printer-name,name"`
	PrinterState                      int                            `ipp:"printer-state,enum"`
	PrinterStateMessage               string                         `ipp:"?printer-state-message,text"`
	PrinterStateReasons               []string                       `ipp:"printer-state-reasons,keyword"`
	PrinterUpTime                     int                            `ipp:"printer-up-time"`
	PrinterUriSupported               []string                       `ipp:"printer-uri-supported,uri"`
	QueuedJobCount                    int                            `ipp:"queued-job-count"`
	ReferenceUriSchemesSupported      []string                       `ipp:"?reference-uri-schemes-supported,uriScheme"`
	UriAuthenticationSupported        []string                       `ipp:"uri-authentication-supported,keyword"`
	UriSecuritySupported              []string                       `ipp:"uri-security-supported,keyword"`
}

// PrinterMediaSizeSupported represents "media-size-supported"
// collection entry, Integer variant
type PrinterMediaSizeSupported struct {
	XDimension int `ipp:"x-dimension"`
	YDimension int `ipp:"y-dimension"`
}

// PrinterMediaSizeSupportedRange represents "media-size-supported"
// collection entry, rangeOfInteger variant
type PrinterMediaSizeSupportedRange struct {
	XDimension goipp.Range `ipp:"x-dimension"`
	YDimension goipp.Range `ipp:"y-dimension"`
}

// EncodeAttrs encodes printer attributes into goipp.Attributes
func (pa *PrinterAttributes) EncodeAttrs() goipp.Attributes {
	var attrs goipp.Attributes
	ippCodecPrinterAttributes.encode(pa, &attrs)
	return attrs
}

// DecodeAttrs decodes printer attributes from goipp.Attributes
func (pa *PrinterAttributes) DecodeAttrs(attrs goipp.Attributes) error {
	return ippCodecPrinterAttributes.decode(pa, attrs)
}

// EncodeMsg encodes printer attributes into the appropriate group
// of attributes of the IPP message
func (pa *PrinterAttributes) EncodeMsg(msg *goipp.Message) {
	msg.Printer = pa.EncodeAttrs()
}

// DecodeMsg decodes printer attributes from the appropriate group
// of attributes of the IPP message
func (pa *PrinterAttributes) DecodeMsg(msg *goipp.Message) error {
	return pa.DecodeAttrs(msg.Printer)
}

// IsCharsetSupported tells if charset is supported
func (pa *PrinterAttributes) IsCharsetSupported(cs string) bool {
	for _, supp := range pa.CharsetSupported {
		if cs == supp {
			return true
		}
	}
	return false
}

// IsOperationSupported tells if operation is supported
func (pa *PrinterAttributes) IsOperationSupported(op goipp.Op) bool {
	for _, supp := range pa.OperationsSupported {
		if op == supp {
			return true
		}
	}
	return false
}
