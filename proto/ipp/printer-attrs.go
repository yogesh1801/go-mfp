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

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/goipp"
)

// Commonly used Operation Attributes.
const (
	// AttrOperationAttributesCharset specifies character set
	// for all attributes in the message.
	//
	// It must be very first attribute in the message (RFC8011, 4.1.4.).
	AttrOperationAttributesCharset = "attributes-charset"

	// AttrOperationAttributesNaturalLanguage specifies natural
	// language for all textWithoutLanguage and nameWithoutLanguage
	// attributes in the message.
	//
	// It must be second attribute in the message (RFC8011, 4.1.4.).
	AttrOperationAttributesNaturalLanguage = "attributes-natural-language"

	// AttrOperationPrinterURI uses in requests to specify Printer URI,
	// where appropriate.
	AttrOperationPrinterURI = "printer-uri"

	// AttrOperationRequestedAttributes used in Get-Printer-Attributes
	// request to specify list of requested attributes.
	AttrOperationRequestedAttributes = "requested-attributes"

	// AttrOperationStatusMessage provides a short textual description
	// of the status of the operation.
	AttrOperationStatusMessage = "status-message"
)

// PrinterAttributes represents IPP Printer Attributes
type PrinterAttributes struct {
	ObjectRawAttrs
	PrinterDescriptionGroup

	PrinterDescription
	JobTemplate
	MediaColDatabase
}

// PrinterDescription contains Printer Description and Status Attributes
type PrinterDescription struct {
	// RFC8011, 5.4: Printer Description and Status Attributes
	CharsetConfigured                 optional.Val[string]        `ipp:"charset-configured,charset"`
	CharsetSupported                  []string                    `ipp:"charset-supported,1setOf charset"`
	ColorSupported                    optional.Val[bool]          `ipp:"color-supported"`
	CompressionSupported              []KwCompression             `ipp:"compression-supported"`
	DocumentFormatDefault             optional.Val[string]        `ipp:"document-format-default,mimeMediaType"`
	DocumentFormatSupported           []string                    `ipp:"document-format-supported,1setOf mimeMediaType"`
	GeneratedNaturalLanguageSupported []string                    `ipp:"generated-natural-language-supported,1setOf naturalLanguage"`
	IppVersionsSupported              []goipp.Version             `ipp:"ipp-versions-supported"`
	JobImpressionsSupported           optional.Val[goipp.Range]   `ipp:"job-impressions-supported"`
	JobKOctetsSupported               optional.Val[goipp.Range]   `ipp:"job-k-octets-supported"`
	JobMediaSheetsSupported           optional.Val[goipp.Range]   `ipp:"job-media-sheets-supported"`
	MediaSizeSupported                []MediaSize                 `ipp:"media-size-supported"`
	MultipleDocumentJobsSupported     optional.Val[bool]          `ipp:"multiple-document-jobs-supported"`
	MultipleOperationTimeOut          optional.Val[int]           `ipp:"multiple-operation-time-out"`
	NaturalLanguageConfigured         optional.Val[string]        `ipp:"natural-language-configured,naturalLanguage"`
	OperationsSupported               []goipp.Op                  `ipp:"operations-supported,1setOf enum"`
	PagesPerMinuteColor               optional.Val[int]           `ipp:"pages-per-minute-color"`
	PagesPerMinute                    optional.Val[int]           `ipp:"pages-per-minute"`
	PdlOverrideSupported              optional.Val[KwPdlOverride] `ipp:"pdl-override-supported"`
	PrinterDriverInstaller            optional.Val[string]        `ipp:"printer-driver-installer,uri"`
	PrinterDeviceID                   optional.Val[string]        `ipp:"printer-device-id,text"`
	PrinterInfo                       optional.Val[string]        `ipp:"printer-info,text"`
	PrinterIsAcceptingJobs            optional.Val[bool]          `ipp:"printer-is-accepting-jobs"`
	PrinterLocation                   optional.Val[string]        `ipp:"printer-location,text"`
	PrinterMakeAndModel               optional.Val[string]        `ipp:"printer-make-and-model,text"`
	PrinterMessageFromOperator        optional.Val[string]        `ipp:"printer-message-from-operator,text"`
	PrinterMoreInfoManufacturer       optional.Val[string]        `ipp:"printer-more-info-manufacturer,uri"`
	PrinterMoreInfo                   optional.Val[string]        `ipp:"printer-more-info,uri"`
	PrinterName                       optional.Val[string]        `ipp:"printer-name,name"`
	PrinterStateMessage               optional.Val[string]        `ipp:"printer-state-message,text"`
	PrinterState                      optional.Val[int]           `ipp:"printer-state,enum"`
	PrinterStateReasons               []KwPrinterStateReasons     `ipp:"printer-state-reasons"`
	PrinterUpTime                     optional.Val[int]           `ipp:"printer-up-time"`
	PrinterURISupported               []string                    `ipp:"printer-uri-supported,1setOf uri"`
	QueuedJobCount                    optional.Val[int]           `ipp:"queued-job-count"`
	ReferenceURISchemesSupported      []string                    `ipp:"reference-uri-schemes-supported,1setOf uriScheme"`
	URIAuthenticationSupported        []KwURIAuthentication       `ipp:"uri-authentication-supported"`
	URISecuritySupported              []KwURISecurity             `ipp:"uri-security-supported"`

	// PWG5100.7: IPP Job Extensions v2.1 (JOBEXT)
	// 6.9 Printer Description Attributes
	ClientInfoSupported              []string                    `ipp:"client-info-supported,1setOf keyword"`
	DocumentCharsetDefault           optional.Val[string]        `ipp:"document-charset-default,charset"`
	DocumentCharsetSupported         []string                    `ipp:"document-charset-supported,1setOf charset"`
	DocumentFormatDetailsSupported   []string                    `ipp:"document-format-details-supported,1setOf keyword"`
	DocumentNaturalLanguageDefault   optional.Val[string]        `ipp:"document-natural-language-default,naturalLanguage"`
	DocumentNaturalLanguageSupported []string                    `ipp:"document-natural-language-supported,1setOf naturalLanguage"`
	JobCreationAttributesSupported   []string                    `ipp:"job-creation-attributes-supported,1setOf keyword"`
	JobHistoryAttributesConfigured   []string                    `ipp:"job-history-attributes-configured,1setOf keyword"`
	JobHistoryAttributesSupported    []string                    `ipp:"job-history-attributes-supported,1setOf keyword"`
	JobHistoryIntervalConfigured     optional.Val[int]           `ipp:"job-history-interval-configured,integer(0:MAX)"`
	JobHistoryIntervalSupported      optional.Val[goipp.Range]   `ipp:"job-history-interval-supported,rangeOfInteger(0:MAX)"`
	JobMandatoryAttributesSupported  optional.Val[bool]          `ipp:"job-mandatory-attributes-supported"`
	JobSpoolingSupported             optional.Val[KwJobSpooling] `ipp:"job-spooling-supported"`
	MediaBackCoatingSupported        []KwMediaBackCoating        `ipp:"media-back-coating-supported"`
	MediaBottomMarginSupported       []int                       `ipp:"media-bottom-margin-supported,1setOf integer(0:MAX)"`
	MediaColDefault                  optional.Val[MediaCol]      `ipp:"media-col-default"`
	MediaColorSupported              []string                    `ipp:"media-color-supported,1setOf keyword"`
	MediaColReady                    []MediaCol                  `ipp:"media-col-ready"`
	MediaColSupported                []string                    `ipp:"media-col-supported,1setOf keyword"`
	MediaFrontCoatingSupported       []KwMediaBackCoating        `ipp:"media-front-coating-supported"`
	MediaGrainSupported              []string                    `ipp:"media-grain-supported,1setOf keyword"`
	MediaHoleCountSupported          []goipp.Range               `ipp:"media-hole-count-supported,1setOf rangeOfInteger(0:MAX)"`
	MediaKeySupported                []KwMedia                   `ipp:"media-key-supported"`
	MediaLeftMarginSupported         []int                       `ipp:"media-left-margin-supported,1setOf integer(0:MAX)"`
	MediaOrderCountSupported         []goipp.Range               `ipp:"media-order-count-supported,1setOf rangeOfInteger(1:MAX)"`
	MediaPrePrintedSupported         []string                    `ipp:"media-pre-printed-supported,1setOf keyword"`
	MediaRecycledSupported           []string                    `ipp:"media-recycled-supported,1setOf keyword"`
	MediaRightMarginSupported        []int                       `ipp:"media-right-margin-supported,1setOf integer(0:MAX)"`
	MediaSourceSupported             []string                    `ipp:"media-source-supported,1setOf keyword"`
	MediaThicknessSupported          []goipp.Range               `ipp:"media-thickness-supported,1setOf rangeOfInteger(1:MAX)"`
	MediaToothSupported              []string                    `ipp:"media-tooth-supported,1setOf keyword"`
	MediaTopMarginSupported          []int                       `ipp:"media-top-margin-supported,1setOf integer(0:MAX)"`
	MediaTypeSupported               []string                    `ipp:"media-type-supported,1setOf keyword"`
	MediaWeightMetricSupported       []goipp.Range               `ipp:"media-weight-metric-supported,1setOf rangeOfInteger(1:MAX)"`

	// PWG5100.13: IPP Driver Replacement Extensions v2.0 (NODRIVER)
	// 6.5 Printer Description Attributes
	DocumentPasswordSupported         optional.Val[int]         `ipp:"document-password-supported,integer(0:1023)"`
	IdentifyActionsDefault            []string                  `ipp:"identify-actions-default,1setOf keyword"`
	IdentifyActionsSupported          []string                  `ipp:"identify-actions-supported,1setOf keyword"`
	IppFeaturesSupported              []string                  `ipp:"ipp-features-supported,1setOf keyword"`
	JobPresetsSupported               []JobPresets              `ipp:"job-presets-supported"`
	JpegFeaturesSupported             []string                  `ipp:"jpeg-features-supported,1setOf keyword"`
	JpegKOctetsSupported              optional.Val[goipp.Range] `ipp:"jpeg-k-octets-supported,rangeOfinteger(0:MAX)"`
	JpegXDimensionSupported           optional.Val[goipp.Range] `ipp:"jpeg-x-dimension-supported,rangeOfinteger(0:65535)"`
	JpegYDimensionSupported           optional.Val[goipp.Range] `ipp:"jpeg-y-dimension-supported,rangeOfinteger(0:65535)"`
	MultipleOperationTimeOutAction    optional.Val[string]      `ipp:"multiple-operation-time-out-action,keyword"`
	PdfKOctetsSupported               optional.Val[goipp.Range] `ipp:"pdf-k-octets-supported,rangeOfinteger(0:MAX)"`
	PdfVersionsSupported              []string                  `ipp:"pdf-versions-supported,1setOf keyword"`
	PreferredAttributesSupported      optional.Val[bool]        `ipp:"preferred-attributes-supported"`
	PrinterDNSSdName                  optional.Val[string]      `ipp:"printer-dns-sd-name,name"`
	PrinterGeoLocation                optional.Val[string]      `ipp:"printer-geo-location,uri|unknown"`
	PrinterGetAttributesSupported     []string                  `ipp:"printer-get-attributes-supported,1setOf keyword"`
	PrinterIcons                      []string                  `ipp:"printer-icons,1setOf uri"`
	PrinterKind                       []string                  `ipp:"printer-kind,1setOf keyword"`
	PrinterOrganization               []string                  `ipp:"printer-organization,1setOf text"`
	PrinterOrganizationalUnit         []string                  `ipp:"printer-organizational-unit,1setOf text"`
	PrinterStringsLanguagesSupported  []string                  `ipp:"printer-strings-languages-supported,1setOf naturalLanguage"`
	PrinterStringsURI                 optional.Val[string]      `ipp:"printer-strings-uri,uri"`
	RequestingUserURISupported        optional.Val[bool]        `ipp:"requesting-user-uri-supported"`
	RequestingUserURISchemesSupported optional.Val[string]      `ipp:"requesting-user-uri-schemes-supported,uriScheme"`

	// PWG5100.13: IPP Driver Replacement Extensions v2.0 (NODRIVER)
	// 6.6 Printer Status Attributes
	DeviceServiceCount           optional.Val[int]       `ipp:"device-service-count,integer(1:MAX)"`
	DeviceUUID                   optional.Val[string]    `ipp:"device-uuid,uri"`
	PrinterConfigChangeDateTime  optional.Val[time.Time] `ipp:"printer-config-change-date-time"`
	PrinterConfigChangeTime      optional.Val[int]       `ipp:"printer-config-change-time,integer(1:MAX)"`
	PrinterFirmwareName          optional.Val[string]    `ipp:"printer-firmware-name,name"`
	PrinterFirmwarePatches       []string                `ipp:"printer-firmware-patches,1setOf text"`
	PrinterFirmwareStringVersion []string                `ipp:"printer-firmware-string-version,1setOf text"`
	PrinterFirmwareVersion       []string                `ipp:"printer-firmware-version,1setOf string"`
	PrinterInputTray             []string                `ipp:"printer-input-tray,1setOf string"`
	PrinterOutputTray            []string                `ipp:"printer-output-tray,1setOf string"`
	PrinterSupplyDescription     []goipp.TextWithLang    `ipp:"printer-supply-description,1setOf text"`
	PrinterSupplyInfoURI         optional.Val[string]    `ipp:"printer-supply-info-uri,uri"`
	PrinterSupply                []string                `ipp:"printer-supply,1setOf string"`
	PrinterUUID                  optional.Val[string]    `ipp:"printer-uuid,uri"`

	// Wi-Fi Peer-to-Peer Services Print (P2Ps-Print)
	// Technical Specification
	// (for Wi-Fi Direct® services certification)
	PclmRasterBackSide       optional.Val[string] `ipp:"pclm-raster-back-side,keyword"`
	PclmStripHeightPreferred optional.Val[int]    `ipp:"pclm-strip-height-preferred"`
	PclmStripHeightSupported []int                `ipp:"pclm-strip-height-supported"`

	// CUPS extensions
	DeviceURI          []string                    `ipp:"device-uri,1setOf uri"`
	MarkerChangeTime   optional.Val[int]           `ipp:"marker-change-time,integer(0:MAX)"`
	MarkerColors       []string                    `ipp:"marker-colors,1setOf name"`
	MarkerHighLevels   []int                       `ipp:"marker-high-levels,1setOf integer(0:100)"`
	MarkerLevels       []int                       `ipp:"marker-levels,1setOf integer(-3:100)"`
	MarkerLowLevels    []int                       `ipp:"marker-low-levels,1setOf integer(0:100)"`
	MarkerMessage      optional.Val[string]        `ipp:"marker-message,text"`
	MarkerNames        []string                    `ipp:"marker-names,1setOf name"`
	MarkerTypes        []string                    `ipp:"marker-types,1setOf keyword"`
	PrinterID          optional.Val[int]           `ipp:"printer-id"`
	PrinterIsShared    optional.Val[bool]          `ipp:"printer-is-shared"`
	PrinterIsTemporary optional.Val[bool]          `ipp:"printer-is-temporary"`
	PrinterType        optional.Val[EnPrinterType] `ipp:"printer-type"`
	UrfSupported       []string                    `ipp:"urf-supported,1setOf keyword"`
}

// PrinterJobSaveDisposition represents "job-save-disposition-default"
// collection entry in PrinterAttributes
type PrinterJobSaveDisposition struct {
	SaveDisposition optional.Val[string] `ipp:"save-disposition,keyword"`
	SaveInfo        []PrinterSaveInfo    `ipp:"save-info"`
}

// PrinterSaveInfo represents "save-info" collection entry
// in PrinterJobSaveDisposition
type PrinterSaveInfo struct {
	SaveLocation       optional.Val[string] `ipp:"save-location,uri"`
	SaveName           optional.Val[string] `ipp:"save-name,name"`
	SaveDocumentFormat optional.Val[string] `ipp:"save-document-format,mimeMediaType"`
}

// DecodePrinterAttributes decodes [PrinterAttributes] from
// [goipp.Attributes].
func DecodePrinterAttributes(attrs goipp.Attributes, opt DecodeOptions) (
	*PrinterAttributes, error) {

	pa := &PrinterAttributes{}
	dec := ippDecoder{opt: opt}
	err := dec.Decode(pa, attrs)
	if err != nil {
		return nil, err
	}
	return pa, nil
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
