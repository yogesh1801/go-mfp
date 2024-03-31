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
	CharsetConfigured                 string                         `ipp:"!charset-configured,charset"`
	CharsetSupported                  []string                       `ipp:"!charset-supported,charset"`
	ColorSupported                    bool                           `ipp:"color-supported"`
	CompressionSupported              []string                       `ipp:"!compression-supported,keyword"`
	DocumentFormatDefault             string                         `ipp:"!document-format-default,mimeMediaType"`
	DocumentFormatSupported           []string                       `ipp:"!document-format-supported,mimeMediaType"`
	GeneratedNaturalLanguageSupported []string                       `ipp:"!generated-natural-language-supported,naturalLanguage"`
	IppFeaturesSupported              []string                       `ipp:"?ipp-features-supported,keyword"`
	IppVersionsSupported              []goipp.Version                `ipp:"!ipp-versions-supported"`
	JobImpressionsSupported           goipp.Range                    `ipp:"?job-impressions-supported"`
	JobKOctetsSupported               goipp.Range                    `ipp:"?job-k-octets-supported"`
	JobMediaSheetsSupported           goipp.Range                    `ipp:"?job-media-sheets-supported"`
	MediaSizeSupported                []PrinterMediaSizeSupported    `ipp:"!media-size-supported,norange"`
	MediaSizeSupportedRange           PrinterMediaSizeSupportedRange `ipp:"!media-size-supported,range"`
	MultipleDocumentJobsSupported     bool                           `ipp:"?multiple-document-jobs-supported"`
	MultipleOperationTimeOut          int                            `ipp:"?multiple-operation-time-out"`
	NaturalLanguageConfigured         string                         `ipp:"!natural-language-configured,naturalLanguage"`
	OperationsSupported               []goipp.Op                     `ipp:"!operations-supported,enum"`
	PagesPerMinuteColor               int                            `ipp:"?pages-per-minute-color"`
	PagesPerMinute                    int                            `ipp:"?pages-per-minute"`
	PdlOverrideSupported              string                         `ipp:"!pdl-override-supported,keyword"`
	PrinterDriverInstaller            string                         `ipp:"?printer-driver-installer,uri"`
	PrinterInfo                       string                         `ipp:"?printer-info,text"`
	PrinterIsAcceptingJobs            bool                           `ipp:"!printer-is-accepting-jobs"`
	PrinterLocation                   string                         `ipp:"?printer-location,text"`
	PrinterMakeAndModel               string                         `ipp:"?printer-make-and-model,text"`
	PrinterMessageFromOperator        string                         `ipp:"?printer-message-from-operator,text"`
	PrinterMoreInfoManufacturer       string                         `ipp:"?printer-more-info-manufacturer,uri"`
	PrinterMoreInfo                   string                         `ipp:"?printer-more-info,uri"`
	PrinterName                       string                         `ipp:"printer-name,name"`
	PrinterState                      int                            `ipp:"!printer-state,enum"`
	PrinterStateMessage               string                         `ipp:"printer-state-message,text"`
	PrinterStateReasons               []string                       `ipp:"!printer-state-reasons,keyword"`
	PrinterUpTime                     int                            `ipp:"!printer-up-time"`
	PrinterURISupported               []string                       `ipp:"!printer-uri-supported,uri"`
	QueuedJobCount                    int                            `ipp:"!queued-job-count"`
	ReferenceURISchemesSupported      []string                       `ipp:"?reference-uri-schemes-supported,uriScheme"`
	URIAuthenticationSupported        []string                       `ipp:"uri-authentication-supported,keyword"`
	URISecuritySupported              []string                       `ipp:"uri-security-supported,keyword"`

	// RFC8011, 5.2: Job Template Attributes
	CopiesDefault                     int                `ipp:"?copies-default,>0"`
	CopiesSupported                   goipp.Range        `ipp:"?copies-supported,>0"`
	FinishingsDefault                 []int              `ipp:"?finishings-default,enum"`
	FinishingsSupported               []int              `ipp:"?finishings-supported,enum"`
	JobHoldUntilDefault               string             `ipp:"?job-hold-until-default,keyword"`
	JobHoldUntilSupported             []string           `ipp:"?job-hold-until-supported,keyword"`
	JobPriorityDefault                int                `ipp:"?job-priority-default,1:100"`
	JobPrioritySupported              int                `ipp:"?job-priority-supported,1:100"`
	JobSheetsDefault                  string             `ipp:"?job-sheets-default,keyword"`
	JobSheetsSupported                []string           `ipp:"?job-sheets-supported"`
	MediaDefault                      string             `ipp:"?media-default,keyword"`
	MediaReady                        []string           `ipp:"?media-ready,keyword"`
	MediaSupported                    []string           `ipp:"?media-supported,keyword"`
	MultipleDocumentHandlingDefault   string             `ipp:"?multiple-document-handling-default,keyword"`
	MultipleDocumentHandlingSupported []string           `ipp:"?multiple-document-handling-supported,keyword"`
	NumberUpDefault                   int                `ipp:"?number-up-default,>0"`
	NumberUpSupported                 []int              `ipp:"?number-up-supported,>0,norange"`
	NumberUpSupportedRange            []int              `ipp:"?number-up-supported,>0,range"`
	OrientationRequestedDefault       int                `ipp:"?orientation-requested-default,enum"`
	OrientationRequestedSupported     []int              `ipp:"?orientation-requested-supported,enum"`
	PageRangesSupported               bool               `ipp:"?page-ranges-supported"`
	PrinterResolutionDefault          goipp.Resolution   `ipp:"?printer-resolution-default"`
	PrinterResolutionSupported        []goipp.Resolution `ipp:"?printer-resolution-supported"`
	PrintQualityDefault               int                `ipp:"?print-quality-default,enum"`
	PrintQualitySupported             []int              `ipp:"?print-quality-supported,enum"`
	SidesDefault                      string             `ipp:"?sides-default,keyword"`
	SidesSupported                    []string           `ipp:"?sides-supported,keyword"`

	// PWG5100.11, 7: Job Template Attributes
	FeedOrientationDefault           string                      `ipp:"?feed-orientation-default,keyword"`
	FeedOrientationSupported         []string                    `ipp:"?feed-orientation-supported,keyword"`
	FontNameRequestedDefault         string                      `ipp:"?font-name-requested-default,name"`
	FontNameRequestedSupported       []string                    `ipp:"?font-name-requested-supported,name"`
	FontSizeRequestedDefault         int                         `ipp:"?font-size-requested-default,>0"`
	FontSizeRequestedSupported       []int                       `ipp:"?font-size-requested-supported,>0"`
	JobDelayOutputUntilDefault       string                      `ipp:"?job-delay-output-until-default,keyword"`
	JobDelayOutputUntilSupported     []string                    `ipp:"?job-delay-output-until-supported,keyword"`
	JobDelayOutputUntilTimeSupported goipp.Range                 `ipp:"?job-delay-output-until-time-supported,>-1"`
	JobHoldUntilTimeSupported        goipp.Range                 `ipp:"?job-hold-until-time-supported,>-1"`
	JobPhoneNumberDefault            string                      `ipp:"?job-phone-number-default,uri"`
	JobPhoneNumberSupported          bool                        `ipp:"?job-phone-number-supported"`
	JobRecipientNameDefault          string                      `ipp:"?job-recipient-name-default,name"`
	JobRecipientNameSupported        bool                        `ipp:"?job-recipient-name-supported"`
	JobSaveDispositionDefault        []PrinterJobSaveDisposition `ipp:"?job-save-disposition-default"`
	JobSaveDispositionSupported      []string                    `ipp:"?job-save-disposition-supported,keyword"`
	SaveDispositionSupported         string                      `ipp:"?save-disposition-supported,keyword"`
	SaveDocumentFormatDefault        string                      `ipp:"?save-document-format-default,mimeMediaType"`
	SaveDocumentFormatSupported      []string                    `ipp:"?save-document-format-supported,mimeMediaType"`
	SaveInfoSupported                []string                    `ipp:"?save-info-supported,keyword"`
	SaveLocationDefault              string                      `ipp:"?save-location-default,uri"`
	SaveLocationSupported            []string                    `ipp:"?save-location-supported,uri"`
	SaveNameSubdirectorySupported    bool                        `ipp:"?save-name-subdirectory-supported"`
	SaveNameSupported                bool                        `ipp:"?save-name-supported"`

	// Other
	MarkerChangeTime int      `ipp:"?marker-change-time,>-1"`
	MarkerColors     []string `ipp:"?marker-colors,name"`
	MarkerHighLevels []int    `ipp:"?marker-high-levels,0:100"`
	MarkerLevels     []int    `ipp:"?marker-levels,-3:100"`
	MarkerLowLevels  []int    `ipp:"?marker-low-levels,0:100"`
	MarkerMessage    string   `ipp:"?marker-message,text"`
	MarkerNames      []string `ipp:"?marker-names,name"`
	MarkerTypes      []string `ipp:"?marker-types,keyword"`
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

// PrinterJobSaveDisposition represents "job-save-disposition-default"
// collection entry in PrinterAttributes
type PrinterJobSaveDisposition struct {
	SaveDisposition string            `ipp:"save-disposition,keyword"`
	SaveInfo        []PrinterSaveInfo `ipp:"?save-info"`
}

// PrinterSaveInfo represents "save-info" collection entry
// in PrinterJobSaveDisposition
type PrinterSaveInfo struct {
	SaveLocation       string `ipp:"?save-location,uri"`
	SaveName           string `ipp:"?save-name,name"`
	SaveDocumentFormat string `ipp:"?save-document-format,mimeMediaType"`
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
