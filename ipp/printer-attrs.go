// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Printer Attributes

package ipp

import (
	"time"

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
	PrinterDescription
	JobTemplate
}

// PrinterDescription contains Printer Description and Status Attributes
type PrinterDescription struct {
	// RFC8011, 5.4: Printer Description and Status Attributes
	CharsetConfigured                 string          `ipp:"!charset-configured,charset"`
	CharsetSupported                  []string        `ipp:"!charset-supported,charset"`
	ColorSupported                    bool            `ipp:"color-supported"`
	CompressionSupported              []string        `ipp:"!compression-supported,keyword"`
	DocumentFormatDefault             string          `ipp:"!document-format-default,mimeMediaType"`
	DocumentFormatSupported           []string        `ipp:"!document-format-supported,mimeMediaType"`
	GeneratedNaturalLanguageSupported []string        `ipp:"!generated-natural-language-supported,naturalLanguage"`
	IppVersionsSupported              []goipp.Version `ipp:"!ipp-versions-supported"`
	JobImpressionsSupported           goipp.Range     `ipp:"?job-impressions-supported"`
	JobKOctetsSupported               goipp.Range     `ipp:"?job-k-octets-supported"`
	JobMediaSheetsSupported           goipp.Range     `ipp:"?job-media-sheets-supported"`
	MediaSizeSupported                []MediaSize     `ipp:"!media-size-supported"`
	MultipleDocumentJobsSupported     bool            `ipp:"?multiple-document-jobs-supported"`
	MultipleOperationTimeOut          int             `ipp:"?multiple-operation-time-out"`
	NaturalLanguageConfigured         string          `ipp:"!natural-language-configured,naturalLanguage"`
	OperationsSupported               []goipp.Op      `ipp:"!operations-supported,enum"`
	PagesPerMinuteColor               int             `ipp:"?pages-per-minute-color"`
	PagesPerMinute                    int             `ipp:"?pages-per-minute"`
	PdlOverrideSupported              string          `ipp:"!pdl-override-supported,keyword"`
	PrinterDriverInstaller            string          `ipp:"?printer-driver-installer,uri"`
	PrinterInfo                       string          `ipp:"?printer-info,text"`
	PrinterIsAcceptingJobs            bool            `ipp:"!printer-is-accepting-jobs"`
	PrinterLocation                   string          `ipp:"?printer-location,text"`
	PrinterMakeAndModel               string          `ipp:"?printer-make-and-model,text"`
	PrinterMessageFromOperator        string          `ipp:"?printer-message-from-operator,text"`
	PrinterMoreInfoManufacturer       string          `ipp:"?printer-more-info-manufacturer,uri"`
	PrinterMoreInfo                   string          `ipp:"?printer-more-info,uri"`
	PrinterName                       string          `ipp:"printer-name,name"`
	PrinterState                      int             `ipp:"!printer-state,enum"`
	PrinterStateMessage               string          `ipp:"printer-state-message,text"`
	PrinterStateReasons               []string        `ipp:"!printer-state-reasons,keyword"`
	PrinterUpTime                     int             `ipp:"!printer-up-time"`
	PrinterURISupported               []string        `ipp:"!printer-uri-supported,uri"`
	QueuedJobCount                    int             `ipp:"!queued-job-count"`
	ReferenceURISchemesSupported      []string        `ipp:"?reference-uri-schemes-supported,uriScheme"`
	URIAuthenticationSupported        []string        `ipp:"uri-authentication-supported,keyword"`
	URISecuritySupported              []string        `ipp:"uri-security-supported,keyword"`

	// PWG5100.7: IPP Job Extensions v2.1 (JOBEXT)
	// 6.9 Printer Description Attributes
	ClientInfoSupported              []string    `ipp:"?client-info-supported,keyword"`
	DocumentCharsetDefault           string      `ipp:"?document-charset-default,charset"`
	DocumentCharsetSupported         []string    `ipp:"?document-charset-supported,charset"`
	DocumentFormatDetailsSupported   []string    `ipp:"?document-format-details-supported,keyword"`
	DocumentNaturalLanguageDefault   string      `ipp:"?document-natural-language-default,naturalLanguage"`
	DocumentNaturalLanguageSupported []string    `ipp:"?document-natural-language-supported,naturalLanguage"`
	JobCreationAttributesSupported   []string    `ipp:"?job-creation-attributes-supported,keyword"`
	JobHistoryAttributesConfigured   []string    `ipp:"?job-history-attributes-configured,keyword"`
	JobHistoryAttributesSupported    []string    `ipp:"?job-history-attributes-supported,keyword"`
	JobHistoryIntervalConfigured     int         `ipp:"?job-history-interval-configured,0:MAX"`
	JobHistoryIntervalSupported      goipp.Range `ipp:"?job-history-interval-supported,0:MAX"`
	JobMandatoryAttributesSupported  bool        `ipp:"?job-mandatory-attributes-supported"`
	JobSpoolingSupported             string      `ipp:"? job-spooling-supported,keyword"`

	// PWG5100.13: IPP Driver Replacement Extensions v2.0 (NODRIVER)
	// 6.5 Printer Description Attributes
	DocumentPasswordSupported         int          `ipp:"?document-password-supported,0:1023"`
	IdentifyActionsDefault            []string     `ipp:"?identify-actions-default,keyword"`
	IdentifyActionsSupported          []string     `ipp:"?identify-actions-supported,keyword"`
	IppFeaturesSupported              []string     `ipp:"?ipp-features-supported,keyword"`
	JobPresetsSupported               []JobPresets `ipp:"?job-presets-supported"`
	JpegFeaturesSupported             []string     `ipp:"?jpeg-features-supported,keyword"`
	JpegKOctetsSupported              goipp.Range  `ipp:"?jpeg-k-octets-supported,0:MAX"`
	JpegXDimensionSupported           goipp.Range  `ipp:"?jpeg-x-dimension-supported,0:65535"`
	JpegYDimensionSupported           goipp.Range  `ipp:"?jpeg-y-dimension-supported,0:65535"`
	MultipleOperationTimeOutAction    string       `ipp:"?multiple-operation-time-out-action,keyword"`
	PdfKOctetsSupported               goipp.Range  `ipp:"?pdf-k-octets-supported,0:MAX"`
	PdfVersionsSupported              []string     `ipp:"?pdf-versions-supported,keyword"`
	PreferredAttributesSupported      bool         `ipp:"?preferred-attributes-supported"`
	PrinterDNSSdName                  string       `ipp:"?printer-dns-sd-name,name"`
	PrinterGeoLocation                string       `ipp:"?printer-geo-location,uri|unknown"`
	PrinterGetAttributesSupported     []string     `ipp:"?printer-get-attributes-supported,keyword"`
	PrinterIcons                      []string     `ipp:"?printer-icons,uri"`
	PrinterKind                       []string     `ipp:"?printer-kind,keyword"`
	PrinterOrganization               []string     `ipp:"?printer-organization,text"`
	PrinterOrganizationalUnit         []string     `ipp:"?printer-organizational-unit,text"`
	PrinterStringsLanguagesSupported  []string     `ipp:"?printer-strings-languages-supported,naturalLanguage"`
	PrinterStringsURI                 string       `ipp:"?printer-strings-uri,uri"`
	RequestingUserURISupported        bool         `ipp:"?requesting-user-uri-supported"`
	RequestingUserURISchemesSupported string       `ipp:"?requesting-user-uri-schemes-supported,uriScheme"`

	// PWG5100.13: IPP Driver Replacement Extensions v2.0 (NODRIVER)
	// 6.6 Printer Status Attributes
	DeviceServiceCount           int                  `ipp:"?device-service-count,1:MAX"`
	DeviceUUID                   string               `ipp:"?device-uuid,uri"`
	PrinterConfigChangeDateTime  time.Time            `ipp:"?printer-config-change-date-time"`
	PrinterConfigChangeTime      int                  `ipp:"?printer-config-change-time,1:MAX"`
	PrinterFirmwareName          string               `ipp:"?printer-firmware-name,name"`
	PrinterFirmwarePatches       []string             `ipp:"?printer-firmware-patches,text"`
	PrinterFirmwareStringVersion []string             `ipp:"?printer-firmware-string-version,text"`
	PrinterFirmwareVersion       []string             `ipp:"?printer-firmware-version,string"`
	PrinterInputTray             []string             `ipp:"?printer-input-tray,string"`
	PrinterOutputTray            []string             `ipp:"?printer-output-tray,string"`
	PrinterSupplyDescription     []goipp.TextWithLang `ipp:"?printer-supply-description"`
	PrinterSupplyInfoURI         string               `ipp:"?printer-supply-info-uri,uri"`
	PrinterSupply                []string             `ipp:"?printer-supply.string"`
	PrinterUUID                  string               `ipp:"?printer-uuid,uri"`

	// These seems to be originated from CUPS. I was unable to
	// find any RFC or PWG standard describing these attrubutes
	//
	// Anyway, these attributes are widely supported by hardware
	// printers
	MarkerChangeTime int      `ipp:"?marker-change-time,0:MAX"`
	MarkerColors     []string `ipp:"?marker-colors,name"`
	MarkerHighLevels []int    `ipp:"?marker-high-levels,0:100"`
	MarkerLevels     []int    `ipp:"?marker-levels,-3:100"`
	MarkerLowLevels  []int    `ipp:"?marker-low-levels,0:100"`
	MarkerMessage    string   `ipp:"?marker-message,text"`
	MarkerNames      []string `ipp:"?marker-names,name"`
	MarkerTypes      []string `ipp:"?marker-types,keyword"`
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
