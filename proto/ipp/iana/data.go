// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP registrations database
//
// THIS IS GENERATED FILE. DON'T EDIT!

package iana

import (
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/goipp"
)

// CUPSDeviceAttributes is the CUPS Device Attributes attributes
var CUPSDeviceAttributes = map[string]*Attribute{
	// CUPS Device Attributes/device-class (CUPS)
	"device-class": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// CUPS Device Attributes/device-id (CUPS)
	"device-id": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// CUPS Device Attributes/device-info (CUPS)
	"device-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// CUPS Device Attributes/device-location (CUPS)
	"device-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// CUPS Device Attributes/device-make-and-model (CUPS)
	"device-make-and-model": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// CUPS Device Attributes/device-uri (CUPS)
	"device-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
}

// CUPSPPDAttributes is the CUPS PPD Attributes attributes
var CUPSPPDAttributes = map[string]*Attribute{
	// CUPS PPD Attributes/ppd-device-id (CUPS)
	"ppd-device-id": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// CUPS PPD Attributes/ppd-make (CUPS)
	"ppd-make": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// CUPS PPD Attributes/ppd-make-and-model (CUPS)
	"ppd-make-and-model": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// CUPS PPD Attributes/ppd-model-number (CUPS)
	"ppd-model-number": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// CUPS PPD Attributes/ppd-name (CUPS)
	"ppd-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// CUPS PPD Attributes/ppd-natural-language (CUPS)
	"ppd-natural-language": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// CUPS PPD Attributes/ppd-product (CUPS)
	"ppd-product": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// CUPS PPD Attributes/ppd-psversion (CUPS)
	"ppd-psversion": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// CUPS PPD Attributes/ppd-type (CUPS)
	"ppd-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
}

// CUPSPrinterClassAttributes is the CUPS Printer Class Attributes attributes
var CUPSPrinterClassAttributes = map[string]*Attribute{
	// CUPS Printer Class Attributes/member-names (CUPS)
	"member-names": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// CUPS Printer Class Attributes/member-uris (CUPS)
	"member-uris": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
}

// DocumentDescription is the Document Description attributes
var DocumentDescription = map[string]*Attribute{
	// Document Description/document-name (PWG5100.5)
	"document-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
}

// DocumentStatus is the Document Status attributes
var DocumentStatus = map[string]*Attribute{
	// Document Status/attributes-charset (PWG5100.5)
	"attributes-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Document Status/attributes-natural-language (PWG5100.5)
	"attributes-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Document Status/chamber-humidity-actual (PWG5100.21)
	"chamber-humidity-actual": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/chamber-temperature-actual (PWG5100.21)
	"chamber-temperature-actual": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/compression (PWG5100.5)
	"compression": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/copies-actual (PWG5100.5)
	"copies-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/cover-back-actual (PWG5100.5)
	"cover-back-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/cover-front-actual (PWG5100.5)
	"cover-front-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/date-time-at-completed (PWG5100.5)
	"date-time-at-completed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Document Status/date-time-at-created (PWG5100.5)
	"date-time-at-created": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Document Status/date-time-at-creation (PWG5100.5)
	"date-time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Document Status/date-time-at-processing (PWG5100.5)
	"date-time-at-processing": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Document Status/detailed-status-messages (PWG5100.5)
	"detailed-status-messages": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-access-errors (PWG5100.5)
	"document-access-errors": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-charset (PWG5100.5)
	"document-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Document Status/document-digital-signature (PWG5100.5)
	"document-digital-signature": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/document-format (PWG5100.5)
	"document-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Document Status/document-format-details (PWG5100.7)
	"document-format-details": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/document-format-details-detected (PWG5100.7)
	"document-format-details-detected": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/document-format-detected (PWG5100.5)
	"document-format-detected": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Document Status/document-format-ready (PWG5100.18)
	"document-format-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Document Status/document-format-supplied (PWG5100.7)
	"document-format-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Document Status/document-format-version (PWG5100.5)
	"document-format-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-format-version-detected (PWG5100.5)
	"document-format-version-detected": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-format-version-supplied (PWG5100.7)
	"document-format-version-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-job-id (PWG5100.5)
	"document-job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/document-job-uri (PWG5100.5)
	"document-job-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/document-message (PWG5100.5)
	"document-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-message-supplied (PWG5100.7)
	"document-message-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-metadata (PWG5100.13)
	"document-metadata": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Document Status/document-name-supplied (PWG5100.7)
	"document-name-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Status/document-natural-language (PWG5100.5)
	"document-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Document Status/document-number (PWG5100.5)
	"document-number": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/document-printer-uri (PWG5100.5)
	"document-printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/document-resource-ids (PWG5100.22)
	"document-resource-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/document-state (PWG5100.5)
	"document-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/document-state-message (PWG5100.5)
	"document-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-state-reasons (PWG5100.5)
	"document-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/document-uri (PWG5100.5)
	"document-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/document-uuid (PWG5100.13)
	"document-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/errors-count (PWG5100.7)
	"errors-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/finishings-actual (PWG5100.5)
	"finishings-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/finishings-col-actual (PWG5100.5)
	"finishings-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/force-front-side-actual (PWG5100.5)
	"force-front-side-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/imposition-template-actual (PWG5100.5)
	"imposition-template-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Status/impressions (PWG5100.5)
	"impressions": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/impressions-col (PWG5100.7)
	"impressions-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Status/impressions-col/blank (PWG5100.7)
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/blank-two-sided (PWG5100.7)
			"blank-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/full-color (PWG5100.7)
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/full-color-two-sided (PWG5100.7)
			"full-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/highlight-color (PWG5100.7)
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/highlight-color-two-sided (PWG5100.7)
			"highlight-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/monochrome (PWG5100.7)
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/monochrome-two-sided (PWG5100.7)
			"monochrome-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Document Status/impressions-completed (PWG5100.5)
	"impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/impressions-completed-col (XEROX20150505)
	"impressions-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/input-attributes-actual (PWG5100.15)
	"input-attributes-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/insert-sheet-actual (PWG5100.5)
	"insert-sheet-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/k-octets (PWG5100.5)
	"k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/k-octets-processed (PWG5100.5)
	"k-octets-processed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/last-document (PWG5100.5)
	"last-document": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Document Status/materials-col-actual (PWG5100.21)
	"materials-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/media-actual (PWG5100.5)
	"media-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Status/media-col-actual (PWG5100.5)
	"media-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/media-sheets (PWG5100.5)
	"media-sheets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/media-sheets-col (PWG5100.7)
	"media-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Status/media-sheets-col/blank (PWG5100.7)
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/media-sheets-col/full-color (PWG5100.7)
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/media-sheets-col/highlight-color (PWG5100.7)
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/media-sheets-col/monochrome (XEROX20150505)
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Document Status/media-sheets-completed (PWG5100.5)
	"media-sheets-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/media-sheets-completed-col (PWG5100.5)
	"media-sheets-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/more-info (PWG5100.5)
	"more-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/multiple-object-handling-actual (PWG5100.21)
	"multiple-object-handling-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/number-up-actual (PWG5100.5)
	"number-up-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/orientation-requested-actual (PWG5100.5)
	"orientation-requested-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/output-attributes-actual (PWG5100.17)
	"output-attributes-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/output-bin-actual (PWG5100.5)
	"output-bin-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Status/output-device-actual (PWG5100.7)
	"output-device-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Status/output-device-assigned (PWG5100.5)
	"output-device-assigned": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Status/output-device-document-state (PWG5100.18)
	"output-device-document-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/output-device-document-state-message (PWG5100.18)
	"output-device-document-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/output-device-document-state-reasons (PWG5100.18)
	"output-device-document-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/page-delivery-actual (PWG5100.5)
	"page-delivery-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/page-order-received-actual (PWG5100.5)
	"page-order-received-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/page-ranges-actual (PWG5100.5)
	"page-ranges-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Document Status/pages (PWG5100.13)
	"pages": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/pages-col (PWG5100.7)
	"pages-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Status/pages-col/full-color (PWG5100.7)
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/pages-col/monochrome (PWG5100.7)
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Document Status/pages-completed (PWG5100.13)
	"pages-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/pages-completed-col (PWG5100.7)
	"pages-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/platform-temperature-actual (PWG5100.21)
	"platform-temperature-actual": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/presentation-direction-number-up-actual (PWG5100.5)
	"presentation-direction-number-up-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/print-accuracy-actual (PWG5100.21)
	"print-accuracy-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/print-base-actual (PWG5100.21)
	"print-base-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/print-color-mode-actual (PWG5100.13)
	"print-color-mode-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/print-content-optimize-actual (PWG5100.7)
	"print-content-optimize-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/print-objects-actual (PWG5100.21)
	"print-objects-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/print-quality-actual (PWG5100.5)
	"print-quality-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/print-supports-actual (PWG5100.21)
	"print-supports-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/printer-resolution-actual (PWG5100.5)
	"printer-resolution-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Document Status/printer-up-time (PWG5100.5)
	"printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/separator-sheets-actual (PWG5100.5)
	"separator-sheets-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/sheet-completed-copy-number (PWG5100.5)
	"sheet-completed-copy-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/sides-actual (PWG5100.5)
	"sides-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/time-at-completed (PWG5100.5)
	"time-at-completed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/time-at-creation (PWG5100.5)
	"time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/time-at-processing (PWG5100.5)
	"time-at-processing": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/warnings-count (PWG5100.7)
	"warnings-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/x-image-position-actual (PWG5100.5)
	"x-image-position-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/x-image-shift-actual (PWG5100.5)
	"x-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/x-side1-image-shift-actual (PWG5100.5)
	"x-side1-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/x-side2-image-shift-actual (PWG5100.5)
	"x-side2-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/y-image-position-actual (PWG5100.5)
	"y-image-position-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/y-image-shift-actual (PWG5100.5)
	"y-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/y-side1-image-shift-actual (PWG5100.5)
	"y-side1-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/y-side2-image-shift-actual (PWG5100.5)
	"y-side2-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// DocumentTemplate is the Document Template attributes
var DocumentTemplate = map[string]*Attribute{
	// Document Template/chamber-humidity (PWG5100.21)
	"chamber-humidity": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/chamber-temperature (PWG5100.21)
	"chamber-temperature": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/copies (PWG5100.5)
	"copies": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/cover-back (PWG5100.5)
	"cover-back": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/cover-front (PWG5100.5)
	"cover-front": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/feed-orientation (PWG5100.5)
	"feed-orientation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/finishings (PWG5100.5)
	"finishings": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Template/finishings-col (PWG5100.5)
	"finishings-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/force-front-side (PWG5100.5)
	"force-front-side": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/imposition-template (PWG5100.5)
	"imposition-template": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Template/insert-sheet (PWG5100.5)
	"insert-sheet": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/materials-col (PWG5100.21)
	"materials-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/materials-col/material-amount (PWG5100.21)
			"material-amount": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-amount-units (PWG5100.21)
			"material-amount-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-color (PWG5100.21)
			"material-color": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-diameter (PWG5100.21)
			"material-diameter": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-diameter-tolerance (PWG5100.21)
			"material-diameter-tolerance": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-fill-density (PWG5100.21)
			"material-fill-density": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-key (PWG5100.21)
			"material-key": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-name (PWG5100.21)
			"material-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Document Template/materials-col/material-nozzle-diameter (PWG5100.21)
			"material-nozzle-diameter": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-purpose (PWG5100.21)
			"material-purpose": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-rate (PWG5100.21)
			"material-rate": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-rate-units (PWG5100.21)
			"material-rate-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-retraction (PWG5100.21)
			"material-retraction": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Document Template/materials-col/material-shell-thickness (PWG5100.21)
			"material-shell-thickness": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-temperature (PWG5100.21)
			"material-temperature": &Attribute{
				SetOf: false,
				Min:   -273,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
			},
			// Document Template/materials-col/material-type (PWG5100.21)
			"material-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Document Template/media (PWG5100.5)
	"media": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Template/media-col (PWG5100.5)
	"media-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/media-col/media-top-offset (IPPLABEL)
			"media-top-offset": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   -2147483648,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/media-col/media-tracking (IPPLABEL)
			"media-tracking": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Document Template/media-input-tray-check (PWG5100.5)
	"media-input-tray-check": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Template/media-overprint (PWG5100.13)
	"media-overprint": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/media-overprint/media-overprint-distance (PWG5100.13)
			"media-overprint-distance": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/media-overprint/media-overprint-method (PWG5100.13)
			"media-overprint-method": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Document Template/multiple-object-handling (PWG5100.21)
	"multiple-object-handling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/number-up (PWG5100.5)
	"number-up": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/orientation-requested (PWG5100.5)
	"orientation-requested": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Template/output-bin (PWG5100.5)
	"output-bin": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Template/output-device (PWG5100.7)
	"output-device": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Template/page-delivery (PWG5100.5)
	"page-delivery": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/page-order-received (IPP20190509B)
	"page-order-received": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/page-ranges (PWG5100.5)
	"page-ranges": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Document Template/platform-temperature (PWG5100.21)
	"platform-temperature": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/presentation-direction-number-up (PWG5100.5)
	"presentation-direction-number-up": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-accuracy (PWG5100.21)
	"print-accuracy": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/print-accuracy/accuracy-units (PWG5100.21)
			"accuracy-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/print-accuracy/x-accuracy (PWG5100.21)
			"x-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/print-accuracy/y-accuracy (PWG5100.21)
			"y-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/print-accuracy/z-accuracy (PWG5100.21)
			"z-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Document Template/print-base (PWG5100.21)
	"print-base": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-color-mode (PWG5100.13)
	"print-color-mode": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-content-optimize (PWG5100.7)
	"print-content-optimize": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-darkness (IPPLABEL)
	"print-darkness": &Attribute{
		SetOf: false,
		Min:   -100,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/print-objects (PWG5100.21)
	"print-objects": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/print-objects/document-number (PWG5100.21)
			"document-number": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/print-objects/object-offset (PWG5100.21)
			"object-offset": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Document Template/print-objects/object-offset/x-offset (PWG5100.21)
					"x-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Document Template/print-objects/object-offset/y-offset (PWG5100.21)
					"y-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Document Template/print-objects/object-offset/z-offset (PWG5100.21)
					"z-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Document Template/print-objects/object-size (PWG5100.21)
			"object-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Document Template/print-objects/object-size/x-dimension (PWG5100.21)
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Document Template/print-objects/object-size/y-dimension (PWG5100.21)
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Document Template/print-objects/object-size/z-dimension (PWG5100.21)
					"z-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Document Template/print-objects/object-uuid (PWG5100.21)
			"object-uuid": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Document Template/print-quality (PWG5100.5)
	"print-quality": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Template/print-rendering-intent (PWG5100.13)
	"print-rendering-intent": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-scaling (PWG5100.13)
	"print-scaling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-speed (IPPLABEL)
	"print-speed": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/print-supports (PWG5100.21)
	"print-supports": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/printer-resolution (PWG5100.5)
	"printer-resolution": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Document Template/separator-sheets (PWG5100.5)
	"separator-sheets": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/sheet-collate (PWG5100.5)
	"sheet-collate": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/sides (PWG5100.5)
	"sides": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/x-image-position (PWG5100.5)
	"x-image-position": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/x-image-shift (PWG5100.5)
	"x-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/x-side1-image-shift (PWG5100.5)
	"x-side1-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/x-side2-image-shift (PWG5100.5)
	"x-side2-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/y-image-position (PWG5100.5)
	"y-image-position": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/y-image-shift (PWG5100.5)
	"y-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/y-side1-image-shift (PWG5100.5)
	"y-side1-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/y-side2-image-shift (PWG5100.5)
	"y-side2-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// EventNotifications is the Event Notifications attributes
var EventNotifications = map[string]*Attribute{
	// Event Notifications/job-id (rfc3996)
	"job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Event Notifications/job-impressions-completed (rfc3996)
	"job-impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Event Notifications/job-state (rfc3996)
	"job-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Event Notifications/job-state-reasons (rfc3996)
	"job-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Event Notifications/job-uuid (PWG5100.13)
	"job-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Event Notifications/notify-charset (rfc3996)
	"notify-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Event Notifications/notify-natural-language (rfc3996)
	"notify-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Event Notifications/notify-printer-uri (rfc3996)
	"notify-printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Event Notifications/notify-sequence-number (rfc3996)
	"notify-sequence-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Event Notifications/notify-subscribed-event (rfc3995)
	"notify-subscribed-event": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Event Notifications/notify-subscription-id (rfc3996)
	"notify-subscription-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Event Notifications/notify-subscription-uuid (PWG5100.13)
	"notify-subscription-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Event Notifications/notify-text (rfc3995)
	"notify-text": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Event Notifications/notify-user-data (rfc3996)
	"notify-user-data": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   63,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Event Notifications/printer-current-time (rfc3996)
	"printer-current-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Event Notifications/printer-is-accepting-jobs (rfc3996)
	"printer-is-accepting-jobs": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Event Notifications/printer-state (rfc3996)
	"printer-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Event Notifications/printer-state-reasons (rfc3996)
	"printer-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Event Notifications/printer-up-time (rfc3996)
	"printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// JobDescription is the Job Description attributes
var JobDescription = map[string]*Attribute{
	// Job Description/current-page-order (IPP20190509B)
	"current-page-order": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Description/job-charge-info (PWG5100.16)
	"job-charge-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Description/job-collation-type (rfc3381)
	"job-collation-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Description/job-impressions-col (PWG5100.7)
	"job-impressions-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Description/job-impressions-col/blank (PWG5100.7)
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/blank-two-sided (PWG5100.7)
			"blank-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/full-color (PWG5100.7)
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/full-color-two-sided (PWG5100.7)
			"full-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/highlight-color (PWG5100.7)
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/highlight-color-two-sided (PWG5100.7)
			"highlight-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/monochrome (PWG5100.7)
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/monochrome-two-sided (PWG5100.7)
			"monochrome-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Description/job-media-sheets-col (PWG5100.7)
	"job-media-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Description/job-media-sheets-col/blank (PWG5100.7)
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/blank-two-sided (PWG5100.7)
			"blank-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/full-color (PWG5100.7)
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/full-color-two-sided (PWG5100.7)
			"full-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/highlight-color (PWG5100.7)
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/highlight-color-two-sided (PWG5100.7)
			"highlight-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/monochrome (PWG5100.7)
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/monochrome-two-sided (PWG5100.7)
			"monochrome-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Description/job-message-from-operator (rfc8011)
	"job-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Description/job-message-to-operator-actual (PWG5100.8)
	"job-message-to-operator-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Description/job-name (rfc8011)
	"job-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Description/job-save-printer-make-and-model (PWG5100.11)
	"job-save-printer-make-and-model": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
}

// JobStatus is the Job Status attributes
var JobStatus = map[string]*Attribute{
	// Job Status/attributes-charset (rfc8011)
	"attributes-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Job Status/attributes-natural-language (rfc8011)
	"attributes-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Job Status/chamber-humidity-actual (PWG5100.21)
	"chamber-humidity-actual": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/chamber-temperature-actual (PWG5100.21)
	"chamber-temperature-actual": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/client-info (PWG5100.7)
	"client-info": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/compression-supplied (PWG5100.7)
	"compression-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/copies-actual (PWG5100.8)
	"copies-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/cover-back-actual (PWG5100.8)
	"cover-back-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/cover-front-actual (PWG5100.8)
	"cover-front-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/current-page-order (PWG5100.3)
	"current-page-order": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/date-time-at-completed (rfc8011)
	"date-time-at-completed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagNoValue},
	},
	// Job Status/date-time-at-completed-estimated (PWG5100.3)
	"date-time-at-completed-estimated": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Job Status/date-time-at-creation (rfc8011)
	"date-time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Status/date-time-at-processing (rfc8011)
	"date-time-at-processing": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagNoValue},
	},
	// Job Status/date-time-at-processing-estimated (PWG5100.3)
	"date-time-at-processing-estimated": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Job Status/destination-statuses (PWG5100.15)
	"destination-statuses": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Status/destination-statuses/destination-uri (PWG5100.15)
			"destination-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Job Status/destination-statuses/images-completed (PWG5100.15)
			"images-completed": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/destination-statuses/transmission-status (PWG5100.15)
			"transmission-status": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
		}},
	},
	// Job Status/document-charset-supplied (PWG5100.7)
	"document-charset-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Job Status/document-digital-signature-supplied (PWG5100.7)
	"document-digital-signature-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/document-format-details-detected (PWG5100.7)
	"document-format-details-detected": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/document-format-details-supplied (PWG5100.7-2003)
	"document-format-details-supplied": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/document-format-ready (PWG5100.18)
	"document-format-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Job Status/document-format-supplied (PWG5100.7)
	"document-format-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Job Status/document-format-version-supplied (PWG5100.7)
	"document-format-version-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/document-message-supplied (PWG5100.7)
	"document-message-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/document-metadata (PWG5100.13)
	"document-metadata": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Job Status/document-name-supplied (PWG5100.7)
	"document-name-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/document-natural-language-supplied (PWG5100.7)
	"document-natural-language-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Job Status/errors-count (PWG5100.7)
	"errors-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/finishings-actual (PWG5100.8)
	"finishings-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/finishings-col-actual (PWG5100.8)
	"finishings-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/force-front-side-actual (PWG5100.8)
	"force-front-side-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/imposition-template-actual (PWG5100.8)
	"imposition-template-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/impressions-completed-current-copy (rfc3381)
	"impressions-completed-current-copy": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/input-attributes-actual (PWG5100.15)
	"input-attributes-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/insert-sheet-actual (PWG5100.8)
	"insert-sheet-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/ipp-attribute-fidelity (PWG5100.7)
	"ipp-attribute-fidelity": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Job Status/job-account-id-actual (PWG5100.8)
	"job-account-id-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/job-account-type-actual (PWG5100.16)
	"job-account-type-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/job-accounting-sheets-actual (PWG5100.8)
	"job-accounting-sheets-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-accounting-user-id-actual (PWG5100.8)
	"job-accounting-user-id-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/job-copies-actual (PWG5100.7)
	"job-copies-actual": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-cover-back-actual (PWG5100.7)
	"job-cover-back-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-cover-front-actual (PWG5100.7)
	"job-cover-front-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-detailed-status-messages (rfc8011)
	"job-detailed-status-messages": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/job-document-access-errors (rfc8011)
	"job-document-access-errors": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/job-error-sheet-actual (PWG5100.8)
	"job-error-sheet-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-finishings-actual (PWG5100.7)
	"job-finishings-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/job-hold-until-actual (PWG5100.8)
	"job-hold-until-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/job-id (rfc8011)
	"job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-impressions (rfc8011)
	"job-impressions": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-impressions-col (XEROX20150505)
	"job-impressions-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Status/job-impressions-col/blank (XEROX20150505)
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/blank-two-sided (XEROX20150505)
			"blank-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/full-color (XEROX20150505)
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/full-color-two-sided (XEROX20150505)
			"full-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/highlight-color (XEROX20150505)
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/highlight-color-two-sided (XEROX20150505)
			"highlight-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/monochrome (XEROX20150505)
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/monochrome-two-sided (XEROX20150505)
			"monochrome-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Status/job-impressions-completed (rfc8011)
	"job-impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-impressions-completed-col (PWG5100.7)
	"job-impressions-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-k-octets (rfc8011)
	"job-k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-k-octets-processed (rfc8011)
	"job-k-octets-processed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-mandatory-attributes (PWG5100.7)
	"job-mandatory-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/job-media-sheets (rfc8011)
	"job-media-sheets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-media-sheets-col (XEROX20150505)
	"job-media-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Status/job-media-sheets-col/blank (XEROX20150505)
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-media-sheets-col/full-color (XEROX20150505)
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-media-sheets-col/highlight-color (XEROX20150505)
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-media-sheets-col/monochrome (XEROX20150505)
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Status/job-media-sheets-completed (rfc8011)
	"job-media-sheets-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-media-sheets-completed-col (PWG5100.7)
	"job-media-sheets-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-more-info (rfc8011)
	"job-more-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/job-originating-user-name (rfc8011)
	"job-originating-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/job-originating-user-uri (PWG5100.13)
	"job-originating-user-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/job-pages (PWG5100.13)
	"job-pages": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-pages-col (PWG5100.7)
	"job-pages-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Status/job-pages-col/blank (PWG5100.7)
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-pages-col/full-color (PWG5100.7)
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-pages-col/monochrome (PWG5100.7)
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Status/job-pages-completed (PWG5100.13)
	"job-pages-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-pages-completed-col (PWG5100.7)
	"job-pages-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-printer-up-time (rfc8011)
	"job-printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-printer-uri (rfc8011)
	"job-printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/job-priority-actual (PWG5100.8)
	"job-priority-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-processing-time (PWG5100.7)
	"job-processing-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-release-action (PWG5100.11)
	"job-release-action": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/job-resource-ids (PWG5100.22)
	"job-resource-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-sheet-message-actual (PWG5100.8)
	"job-sheet-message-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/job-sheets-actual (PWG5100.8)
	"job-sheets-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/job-sheets-col-actual (PWG5100.8)
	"job-sheets-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-state (rfc8011)
	"job-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum, goipp.TagUnknown},
	},
	// Job Status/job-state-message (rfc8011)
	"job-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/job-state-reasons (rfc8011)
	"job-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/job-storage (PWG5100.11)
	"job-storage": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-uri (rfc8011)
	"job-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/job-uuid (PWG5100.13)
	"job-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/materials-col-actual (PWG5100.21)
	"materials-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/media-actual (PWG5100.8)
	"media-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/media-col-actual (PWG5100.8)
	"media-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/media-input-tray-check-actual (PWG5100.8)
	"media-input-tray-check-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/multiple-document-handling-actual (PWG5100.8)
	"multiple-document-handling-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/multiple-object-handling-actual (PWG5100.21)
	"multiple-object-handling-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/number-of-documents (rfc8011)
	"number-of-documents": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/number-of-intervening-jobs (rfc8011)
	"number-of-intervening-jobs": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/number-up-actual (PWG5100.8)
	"number-up-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/orientation-requested-actual (PWG5100.8)
	"orientation-requested-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/original-requesting-user-name (rfc3998)
	"original-requesting-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/output-attributes-actual (PWG5100.17)
	"output-attributes-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/output-bin-actual (PWG5100.8)
	"output-bin-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/output-device-actual (PWG5100.7)
	"output-device-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/output-device-assigned (rfc8011)
	"output-device-assigned": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/output-device-job-state (PWG5100.18)
	"output-device-job-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/output-device-job-state-message (PWG5100.18)
	"output-device-job-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/output-device-job-state-reasons (PWG5100.18)
	"output-device-job-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/output-device-uuid-assigned (PWG5100.18)
	"output-device-uuid-assigned": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/overrides-actual (PWG5100.6)
	"overrides-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/page-delivery-actual (PWG5100.8)
	"page-delivery-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/page-order-received-actual (IPP20190509B)
	"page-order-received-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/page-ranges-actual (PWG5100.8)
	"page-ranges-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Job Status/parent-job-id (PWG5100.11)
	"parent-job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/parent-job-uuid (PWG5100.11)
	"parent-job-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   45,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/platform-temperature-actual (PWG5100.21)
	"platform-temperature-actual": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/presentation-direction-number-up-actual (PWG5100.8)
	"presentation-direction-number-up-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/print-accuracy-actual (PWG5100.21)
	"print-accuracy-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/print-base-actual (PWG5100.21)
	"print-base-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/print-color-mode-actual (PWG5100.13)
	"print-color-mode-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/print-content-optimize-actual (PWG5100.7)
	"print-content-optimize-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/print-objects-actual (PWG5100.21)
	"print-objects-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/print-quality-actual (PWG5100.8)
	"print-quality-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/print-supports-actual (PWG5100.21)
	"print-supports-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/printer-resolution-actual (PWG5100.8)
	"printer-resolution-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Job Status/separator-sheets-actual (PWG5100.8)
	"separator-sheets-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/sheet-collate-actual (PWG5100.8)
	"sheet-collate-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/sheet-completed-copy-number (rfc3381)
	"sheet-completed-copy-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/sheet-completed-document-number (rfc3381)
	"sheet-completed-document-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/sides-actual (PWG5100.8)
	"sides-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/time-at-completed (rfc8011)
	"time-at-completed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Job Status/time-at-completed-estimated (PWG5100.3)
	"time-at-completed-estimated": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagUnknown},
	},
	// Job Status/time-at-creation (rfc8011)
	"time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/time-at-processing (rfc8011)
	"time-at-processing": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Job Status/time-at-processing-estimated (PWG5100.3)
	"time-at-processing-estimated": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagUnknown},
	},
	// Job Status/warnings-count (PWG5100.7)
	"warnings-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/x-image-position-actual (PWG5100.8)
	"x-image-position-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/x-image-shift-actual (PWG5100.8)
	"x-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/x-side1-image-shift-actual (PWG5100.8)
	"x-side1-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/x-side2-image-shift-actual (PWG5100.8)
	"x-side2-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/y-image-position-actual (PWG5100.8)
	"y-image-position-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/y-image-shift-actual (PWG5100.8)
	"y-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/y-side1-image-shift-actual (PWG5100.8)
	"y-side1-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/y-side2-image-shift-actual (PWG5100.8)
	"y-side2-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// JobTemplate is the Job Template attributes
var JobTemplate = map[string]*Attribute{
	// Job Template/auth-info (CUPS)
	"auth-info": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Template/chamber-humidity (PWG5100.21)
	"chamber-humidity": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/chamber-temperature (PWG5100.21)
	"chamber-temperature": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/confirmation-sheet-print (PWG5100.15)
	"confirmation-sheet-print": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Job Template/copies (rfc8011)
	"copies": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/cover-back (PWG5100.3)
	"cover-back": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/cover-back/cover-type (PWG5100.3)
			"cover-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/cover-back/media (rfc8011)
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/cover-back/media-col (PWG5100.3)
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/cover-front (PWG5100.3)
	"cover-front": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/cover-front/cover-type (PWG5100.3)
			"cover-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/cover-front/media (rfc8011)
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/cover-front/media-col (PWG5100.3)
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/cover-sheet-info (PWG5100.15)
	"cover-sheet-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/cover-sheet-info/from-name (PWG5100.15)
			"from-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/cover-sheet-info/logo (PWG5100.15)
			"logo": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Job Template/cover-sheet-info/message (PWG5100.15)
			"message": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/cover-sheet-info/organization-name (PWG5100.15)
			"organization-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/cover-sheet-info/subject (PWG5100.15)
			"subject": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/cover-sheet-info/to-name (PWG5100.15)
			"to-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// Job Template/destination-uris (PWG5100.15)
	"destination-uris": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/destination-uris/destination-attributes (PWG5100.17)
			"destination-attributes": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Job Template/destination-uris/destination-uri (PWG5100.15)
			"destination-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Job Template/destination-uris/post-dial-string (PWG5100.15)
			"post-dial-string": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/destination-uris/pre-dial-string (PWG5100.15)
			"pre-dial-string": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/destination-uris/t33-subaddress (PWG5100.15)
			"t33-subaddress": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Template/feed-orientation (PWG5100.11)
	"feed-orientation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/finishings (rfc8011)
	"finishings": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/finishings-col (PWG5100.1)
	"finishings-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
		Members: []map[string]*Attribute{{
			// Job Template/finishings-col/baling (PWG5100.1)
			"baling": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/baling/baling-type (PWG5100.1)
					"baling-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
					// Job Template/finishings-col/baling/baling-when (PWG5100.1)
					"baling-when": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
				}},
			},
			// Job Template/finishings-col/binding (PWG5100.1)
			"binding": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/binding/binding-reference-edge (PWG5100.1)
					"binding-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/binding/binding-type (PWG5100.1)
					"binding-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
				}},
			},
			// Job Template/finishings-col/coating (PWG5100.1)
			"coating": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/coating/coating-sides (PWG5100.1)
					"coating-sides": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/coating/coating-type (PWG5100.1)
					"coating-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
				}},
			},
			// Job Template/finishings-col/covering (PWG5100.1)
			"covering": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/covering/covering-name (PWG5100.1)
					"covering-name": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
				}},
			},
			// Job Template/finishings-col/finishing-template (PWG5100.1)
			"finishing-template": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/finishings-col/folding (PWG5100.1)
			"folding": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/folding/folding-direction (PWG5100.1)
					"folding-direction": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/folding/folding-offset (PWG5100.1)
					"folding-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/folding/folding-reference-edge (PWG5100.1)
					"folding-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
				}},
			},
			// Job Template/finishings-col/imposition-template (PWG5100.1)
			"imposition-template": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/finishings-col/laminating (PWG5100.1)
			"laminating": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/laminating/laminating-sides (PWG5100.1)
					"laminating-sides": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/laminating/laminating-type (PWG5100.1)
					"laminating-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
				}},
			},
			// Job Template/finishings-col/media-sheets-supported (PWG5100.1)
			"media-sheets-supported": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Job Template/finishings-col/media-size (PWG5100.1)
			"media-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Job Template/finishings-col/media-size-name (PWG5100.1)
			"media-size-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/finishings-col/punching (PWG5100.1)
			"punching": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/punching/punching-locations (PWG5100.1)
					"punching-locations": &Attribute{
						SetOf: true,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/punching/punching-offset (PWG5100.1)
					"punching-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/punching/punching-reference-edge (PWG5100.1)
					"punching-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
				}},
			},
			// Job Template/finishings-col/stitching (PWG5100.1)
			"stitching": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/stitching/stitching-angle (PWG5100.1)
					"stitching-angle": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   359,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/stitching/stitching-locations (PWG5100.1)
					"stitching-locations": &Attribute{
						SetOf: true,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/stitching/stitching-method (PWG5100.1)
					"stitching-method": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/stitching/stitching-offset (PWG5100.1)
					"stitching-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/stitching/stitching-reference-edge (PWG5100.1)
					"stitching-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
				}},
			},
			// Job Template/finishings-col/trimming (PWG5100.1)
			"trimming": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/trimming/trimming-offset (PWG5100.1)
					"trimming-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/trimming/trimming-reference-edge (PWG5100.1)
					"trimming-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/trimming/trimming-type (PWG5100.1)
					"trimming-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
					// Job Template/finishings-col/trimming/trimming-when (PWG5100.1)
					"trimming-when": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
				}},
			},
		}},
	},
	// Job Template/force-front-side (PWG5100.3)
	"force-front-side": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/image-orientation (PWG5100.3)
	"image-orientation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/imposition-template (PWG5100.3)
	"imposition-template": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/insert-sheet (PWG5100.3)
	"insert-sheet": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/insert-sheet/insert-after-page-number (PWG5100.3)
			"insert-after-page-number": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/insert-sheet/insert-count (PWG5100.3)
			"insert-count": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/insert-sheet/media (PWG5100.3)
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/insert-sheet/media-col (PWG5100.3)
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/job-account-id (PWG5100.7)
	"job-account-id": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Template/job-account-type (PWG5100.16)
	"job-account-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-accounting-sheets (PWG5100.3)
	"job-accounting-sheets": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/job-accounting-sheets/job-accounting-sheets-type (PWG5100.3)
			"job-accounting-sheets-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-accounting-sheets/media (PWG5100.3)
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-accounting-sheets/media-col (PWG5100.3)
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/job-accounting-user-id (PWG5100.7)
	"job-accounting-user-id": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Template/job-cancel-after (PWG5100.11)
	"job-cancel-after": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-complete-before (PWG5100.3)
	"job-complete-before": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-complete-before-time (PWG5100.3)
	"job-complete-before-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Template/job-copies (PWG5100.7)
	"job-copies": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-cover-back (PWG5100.7)
	"job-cover-back": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Template/job-cover-front (PWG5100.7)
	"job-cover-front": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Template/job-delay-output-until (PWG5100.7)
	"job-delay-output-until": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-delay-output-until-time (PWG5100.7)
	"job-delay-output-until-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Template/job-error-action (PWG5100.13)
	"job-error-action": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/job-error-sheet (PWG5100.3)
	"job-error-sheet": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/job-error-sheet/job-error-sheet-type (PWG5100.3)
			"job-error-sheet-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-error-sheet/job-error-sheet-when (PWG5100.3)
			"job-error-sheet-when": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/job-error-sheet/media (PWG5100.3)
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-error-sheet/media-col (PWG5100.3)
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/job-finishings (PWG5100.7)
	"job-finishings": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/job-hold-until (rfc8011)
	"job-hold-until": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-hold-until-time (PWG5100.7)
	"job-hold-until-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Template/job-media-progress (CUPS)
	"job-media-progress": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-message-to-operator (PWG5100.3)
	"job-message-to-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Template/job-originating-host-name (CUPS)
	"job-originating-host-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Template/job-pages-per-set (PWG5100.1)
	"job-pages-per-set": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-phone-number (PWG5100.3)
	"job-phone-number": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Template/job-printer-state-message (CUPS)
	"job-printer-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Template/job-printer-state-reasons (CUPS)
	"job-printer-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/job-priority (rfc8011)
	"job-priority": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-recipient-name (PWG5100.3)
	"job-recipient-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Template/job-retain-until (PWG5100.7)
	"job-retain-until": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-retain-until-interval (PWG5100.7)
	"job-retain-until-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-retain-until-time (PWG5100.7)
	"job-retain-until-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Template/job-sheet-message (PWG5100.3)
	"job-sheet-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Template/job-sheets (rfc8011)
	"job-sheets": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-sheets-col (PWG5100.7)
	"job-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/job-sheets-col/job-sheets (PWG5100.7)
			"job-sheets": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-sheets-col/media (PWG5100.7)
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-sheets-col/media-col (PWG5100.7)
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/materials-col (PWG5100.21)
	"materials-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/materials-col/material-amount (PWG5100.21)
			"material-amount": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-amount-units (PWG5100.21)
			"material-amount-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-color (PWG5100.21)
			"material-color": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-diameter (PWG5100.21)
			"material-diameter": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-diameter-tolerance (PWG5100.21)
			"material-diameter-tolerance": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-fill-density (PWG5100.21)
			"material-fill-density": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-key (PWG5100.21)
			"material-key": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-name (PWG5100.21)
			"material-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Job Template/materials-col/material-nozzle-diameter (PWG5100.21)
			"material-nozzle-diameter": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-purpose (PWG5100.21)
			"material-purpose": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-rate (PWG5100.21)
			"material-rate": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-rate-units (PWG5100.21)
			"material-rate-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-retraction (PWG5100.21)
			"material-retraction": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Job Template/materials-col/material-shell-thickness (PWG5100.21)
			"material-shell-thickness": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-temperature (PWG5100.21)
			"material-temperature": &Attribute{
				SetOf: false,
				Min:   -273,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
			},
			// Job Template/materials-col/material-type (PWG5100.21)
			"material-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Job Template/media (rfc8011)
	"media": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/media-col (PWG5100.7)
	"media-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/media-col/media-back-coating (PWG5100.7)
			"media-back-coating": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-bottom-margin (PWG5100.7)
			"media-bottom-margin": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-color (PWG5100.7)
			"media-color": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-front-coating (PWG5100.7)
			"media-front-coating": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-grain (PWG5100.7)
			"media-grain": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-hole-count (PWG5100.7)
			"media-hole-count": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-info (PWG5100.7)
			"media-info": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/media-col/media-key (PWG5100.7)
			"media-key": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-left-margin (PWG5100.7)
			"media-left-margin": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-order-count (PWG5100.7)
			"media-order-count": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-pre-printed (PWG5100.7)
			"media-pre-printed": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-recycled (PWG5100.7)
			"media-recycled": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-right-margin (PWG5100.7)
			"media-right-margin": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-size (PWG5100.7)
			"media-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/media-col/media-size/x-dimension (PWG5100.7)
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/media-col/media-size/y-dimension (PWG5100.7)
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Job Template/media-col/media-size-name (PWG5100.7)
			"media-size-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-source (PWG5100.7)
			"media-source": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-thickness (PWG5100.7)
			"media-thickness": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-tooth (PWG5100.7)
			"media-tooth": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-top-margin (PWG5100.7)
			"media-top-margin": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-top-offset (IPPLABEL)
			"media-top-offset": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   -2147483648,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-tracking (IPPLABEL)
			"media-tracking": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/media-col/media-type (PWG5100.7)
			"media-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-weight-metric (PWG5100.7)
			"media-weight-metric": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Template/media-input-tray-check (PWG5100.3)
	"media-input-tray-check": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/media-overprint (PWG5100.13)
	"media-overprint": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/media-overprint/media-overprint-distance (PWG5100.13)
			"media-overprint-distance": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-overprint/media-overprint-method (PWG5100.13)
			"media-overprint-method": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Job Template/multiple-document-handling (rfc8011)
	"multiple-document-handling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/multiple-object-handling (PWG5100.21)
	"multiple-object-handling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/number-of-retries (PWG5100.15)
	"number-of-retries": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/number-up (rfc8011)
	"number-up": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/orientation-requested (rfc8011)
	"orientation-requested": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/output-bin (PWG5100.2)
	"output-bin": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/output-device (PWG5100.7)
	"output-device": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Template/overrides (PWG5100.6)
	"overrides": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/overrides/document-copies (PWG5100.6)
			"document-copies": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Job Template/overrides/document-numbers (PWG5100.6)
			"document-numbers": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Job Template/overrides/pages (PWG5100.6)
			"pages": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
		}},
	},
	// Job Template/page-border (CUPS)
	"page-border": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/page-delivery (PWG5100.3)
	"page-delivery": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/page-order-received (PWG5100.3)
	"page-order-received": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/page-ranges (rfc8011)
	"page-ranges": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Job Template/page-set (CUPS)
	"page-set": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/pages-per-subset (PWG5100.13)
	"pages-per-subset": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/pclm-source-resolution (HP20180907)
	"pclm-source-resolution": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Job Template/platform-temperature (PWG5100.21)
	"platform-temperature": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/presentation-direction-number-up (PWG5100.3)
	"presentation-direction-number-up": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-accuracy (PWG5100.21)
	"print-accuracy": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/print-accuracy/accuracy-units (PWG5100.21)
			"accuracy-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/print-accuracy/x-accuracy (PWG5100.21)
			"x-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/print-accuracy/y-accuracy (PWG5100.21)
			"y-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/print-accuracy/z-accuracy (PWG5100.21)
			"z-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Template/print-base (PWG5100.21)
	"print-base": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-color-mode (PWG5100.13)
	"print-color-mode": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-content-optimize (PWG5100.7)
	"print-content-optimize": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-darkness (IPPLABEL)
	"print-darkness": &Attribute{
		SetOf: false,
		Min:   -100,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/print-objects (PWG5100.21)
	"print-objects": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/print-objects/document-number (PWG5100.21)
			"document-number": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/print-objects/object-offset (PWG5100.21)
			"object-offset": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/print-objects/object-offset/x-offset (PWG5100.21)
					"x-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/print-objects/object-offset/y-offset (PWG5100.21)
					"y-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/print-objects/object-offset/z-offset (PWG5100.21)
					"z-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Job Template/print-objects/object-size (PWG5100.21)
			"object-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/print-objects/object-size/x-dimension (PWG5100.21)
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/print-objects/object-size/y-dimension (PWG5100.21)
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/print-objects/object-size/z-dimension (PWG5100.21)
					"z-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Job Template/print-objects/object-uuid (PWG5100.21)
			"object-uuid": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Job Template/print-quality (rfc8011)
	"print-quality": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/print-rendering-intent (PWG5100.13)
	"print-rendering-intent": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-scaling (PWG5100.13)
	"print-scaling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-speed (IPPLABEL)
	"print-speed": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/print-supports (PWG5100.21)
	"print-supports": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/printer-resolution (rfc8011)
	"printer-resolution": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Job Template/proof-copies (PWG5100.11)
	"proof-copies": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/proof-print (PWG5100.11)
	"proof-print": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/proof-print/media (PWG5100.11)
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/proof-print/media-col (PWG5100.11)
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Job Template/proof-print/proof-print-copies (PWG5100.11)
			"proof-print-copies": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Template/retry-interval (PWG5100.15)
	"retry-interval": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/retry-time-out (PWG5100.15)
	"retry-time-out": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/separator-sheets (PWG5100.3)
	"separator-sheets": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/separator-sheets/media (rfc8011)
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/separator-sheets/media-col (PWG5100.3)
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Job Template/separator-sheets/separator-sheets-type (PWG5100.3)
			"separator-sheets-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Job Template/sheet-collate (rfc3381)
	"sheet-collate": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/sides (rfc8011)
	"sides": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/x-image-position (PWG5100.3)
	"x-image-position": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/x-image-shift (PWG5100.3)
	"x-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/x-side1-image-shift (PWG5100.3)
	"x-side1-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/x-side2-image-shift (PWG5100.3)
	"x-side2-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/y-image-position (PWG5100.3)
	"y-image-position": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/y-image-shift (PWG5100.3)
	"y-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/y-side1-image-shift (PWG5100.3)
	"y-side1-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/y-side2-image-shift (PWG5100.3)
	"y-side2-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// Operation is the Operation attributes
var Operation = map[string]*Attribute{
	// Operation/attributes-charset (rfc8011)
	"attributes-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Operation/attributes-natural-language (rfc8011)
	"attributes-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Operation/charge-info-message (PWG5100.16)
	"charge-info-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/client-info (PWG5100.7)
	"client-info": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/client-info/client-name (PWG5100.7)
			"client-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Operation/client-info/client-patches (PWG5100.7)
			"client-patches": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
			},
			// Operation/client-info/client-string-version (PWG5100.7)
			"client-string-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/client-info/client-type (PWG5100.7)
			"client-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
			// Operation/client-info/client-version (PWG5100.7)
			"client-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   64,
				Tags:  []goipp.Tag{goipp.TagString, goipp.TagNoValue},
			},
		}},
	},
	// Operation/compression (rfc8011)
	"compression": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/compression-accepted (PWG5100.17)
	"compression-accepted": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/destination-accesses (PWG5100.17)
	"destination-accesses": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
		Members: []map[string]*Attribute{{
			// Operation/destination-accesses/access-oauth-token (PWG5100.17)
			"access-oauth-token": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
			// Operation/destination-accesses/access-oauth-uri (PWG5100.17)
			"access-oauth-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Operation/destination-accesses/access-password (PWG5100.17)
			"access-password": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/destination-accesses/access-pin (PWG5100.17)
			"access-pin": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/destination-accesses/access-user-name (PWG5100.17)
			"access-user-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/destination-accesses/access-x509-certificate (IPPWG20180620)
			"access-x509-certificate": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
		}},
	},
	// Operation/detailed-status-message (rfc8011)
	"detailed-status-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/device-class (CUPS)
	"device-class": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/document-access (PWG5100.18)
	"document-access": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
		Members: []map[string]*Attribute{{
			// Operation/document-access/access-oauth-token (PWG5100.18)
			"access-oauth-token": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
			// Operation/document-access/access-oauth-uri (PWG5100.18)
			"access-oauth-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Operation/document-access/access-password (PWG5100.18)
			"access-password": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-access/access-pin (PWG5100.18)
			"access-pin": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-access/access-user-name (PWG5100.18)
			"access-user-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-access/access-x509-certificate (IPPWG20180620)
			"access-x509-certificate": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
		}},
	},
	// Operation/document-access-error (rfc8011)
	"document-access-error": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/document-charset (PWG5100.5)
	"document-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Operation/document-data-get-interval (PWG5100.17)
	"document-data-get-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/document-data-wait (PWG5100.17)
	"document-data-wait": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/document-digital-signature (PWG5100.7)
	"document-digital-signature": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/document-format (rfc8011)
	"document-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/document-format-accepted (PWG5100.18)
	"document-format-accepted": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/document-format-details (PWG5100.7-2003)
	"document-format-details": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/document-format-details/document-format (PWG5100.7-2003)
			"document-format": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagMimeType},
			},
			// Operation/document-format-details/document-format-device-id (PWG5100.7-2003)
			"document-format-device-id": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-format-details/document-format-version (PWG5100.7-2003)
			"document-format-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-format-details/document-natural-language (PWG5100.7-2003)
			"document-natural-language": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagLanguage},
			},
			// Operation/document-format-details/document-source-application-name (PWG5100.7-2003)
			"document-source-application-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Operation/document-format-details/document-source-application-version (PWG5100.7-2003)
			"document-source-application-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-format-details/document-source-os-name (PWG5100.7-2003)
			"document-source-os-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   40,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Operation/document-format-details/document-source-os-version (PWG5100.7-2003)
			"document-source-os-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   40,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// Operation/document-format-version (PWG5100.7)
	"document-format-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/document-message (PWG5100.5)
	"document-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/document-metadata (PWG5100.13)
	"document-metadata": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Operation/document-name (rfc8011)
	"document-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/document-natural-language (rfc8011)
	"document-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Operation/document-number (PWG5100.5)
	"document-number": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/document-password (PWG5100.13)
	"document-password": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Operation/document-preprocessed (PWG5100.18)
	"document-preprocessed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/document-uri (rfc8011)
	"document-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/encrypted-job-request-format (PWG5100.TRUSTNOONE)
	"encrypted-job-request-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/encrypted-job-request-id (PWG5100.TRUSTNOONE)
	"encrypted-job-request-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/exclude-schemes (CUPS)
	"exclude-schemes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/fetch-status-code (PWG5100.18)
	"fetch-status-code": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/fetch-status-message (PWG5100.18)
	"fetch-status-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/first-index (PWG5100.13)
	"first-index": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/first-printer-name (CUPS)
	"first-printer-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/identify-actions (PWG5100.13)
	"identify-actions": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/include-schemes (CUPS)
	"include-schemes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/input-attributes (PWG5100.15)
	"input-attributes": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/input-attributes/input-auto-scaling (PWG5100.15)
			"input-auto-scaling": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Operation/input-attributes/input-auto-skew-correction (PWG5100.15)
			"input-auto-skew-correction": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Operation/input-attributes/input-brightness (PWG5100.15)
			"input-brightness": &Attribute{
				SetOf: false,
				Min:   -100,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-color-mode (PWG5100.15)
			"input-color-mode": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-content-type (PWG5100.15)
			"input-content-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-contrast (PWG5100.15)
			"input-contrast": &Attribute{
				SetOf: false,
				Min:   -100,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-film-scan-mode (PWG5100.15)
			"input-film-scan-mode": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-images-to-transfer (PWG5100.15)
			"input-images-to-transfer": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-media (PWG5100.15)
			"input-media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Operation/input-attributes/input-orientation-requested (PWG5100.15)
			"input-orientation-requested": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-quality (PWG5100.15)
			"input-quality": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
			// Operation/input-attributes/input-resolution (PWG5100.15)
			"input-resolution": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagResolution},
			},
			// Operation/input-attributes/input-scaling-height (PWG5100.15)
			"input-scaling-height": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   1000,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-scaling-width (PWG5100.15)
			"input-scaling-width": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   1000,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-scan-regions (PWG5100.15)
			"input-scan-regions": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Operation/input-attributes/input-scan-regions/x-dimension (PWG5100.15)
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Operation/input-attributes/input-scan-regions/x-origin (PWG5100.15)
					"x-origin": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Operation/input-attributes/input-scan-regions/y-dimension (PWG5100.15)
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Operation/input-attributes/input-scan-regions/y-origin (PWG5100.15)
					"y-origin": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Operation/input-attributes/input-sharpness (PWG5100.15)
			"input-sharpness": &Attribute{
				SetOf: false,
				Min:   -100,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-sides (PWG5100.15)
			"input-sides": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-source (PWG5100.15)
			"input-source": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Operation/ipp-attribute-fidelity (rfc8011)
	"ipp-attribute-fidelity": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/job-authorization-uri (PWG5100.16)
	"job-authorization-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/job-hold-until (rfc8011)
	"job-hold-until": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Operation/job-hold-until-time (PWG5100.7)
	"job-hold-until-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Operation/job-id (rfc8011)
	"job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-ids (PWG5100.7)
	"job-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-impressions (rfc8011)
	"job-impressions": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-impressions-col (XEROX20150505)
	"job-impressions-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Operation/job-impressions-estimated (PWG5100.16)
	"job-impressions-estimated": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-k-octets (rfc8011)
	"job-k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-mandatory-attributes (PWG5100.7)
	"job-mandatory-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/job-media-sheets (rfc8011)
	"job-media-sheets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-media-sheets-col (XEROX20150505)
	"job-media-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Operation/job-message-from-operator (rfc3380)
	"job-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/job-name (rfc8011)
	"job-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/job-pages (PWG5100.7)
	"job-pages": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-pages-col (PWG5100.7)
	"job-pages-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Operation/job-password (PWG5100.11)
	"job-password": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Operation/job-password-encryption (PWG5100.11)
	"job-password-encryption": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/job-release-action (PWG5100.11)
	"job-release-action": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/job-state (rfc8011)
	"job-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/job-state-message (rfc8011)
	"job-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/job-state-reasons (rfc8011)
	"job-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/job-storage (PWG5100.11)
	"job-storage": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/job-storage/job-storage-access (PWG5100.11)
			"job-storage-access": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/job-storage/job-storage-disposition (PWG5100.11)
			"job-storage-disposition": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/job-storage/job-storage-group (PWG5100.11)
			"job-storage-group": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
		}},
	},
	// Operation/job-uri (rfc8011)
	"job-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/last-document (rfc8011)
	"last-document": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/limit (rfc8011)
	"limit": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/message (rfc8011)
	"message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/my-jobs (rfc8011)
	"my-jobs": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/notify-get-interval (rfc3996)
	"notify-get-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-printer-ids (PWG5100.22)
	"notify-printer-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-resource-id (PWG5100.22)
	"notify-resource-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-sequence-numbers (rfc3996)
	"notify-sequence-numbers": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-subscription-ids (rfc3996)
	"notify-subscription-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-wait (rfc3996)
	"notify-wait": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/original-requesting-user-name (rfc3998)
	"original-requesting-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/output-attributes (PWG5100.17)
	"output-attributes": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/output-attributes/noise-removal (PWG5100.17)
			"noise-removal": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/output-attributes/output-compression-quality-factor (PWG5100.17)
			"output-compression-quality-factor": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Operation/output-device-job-states (PWG5100.18)
	"output-device-job-states": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/output-device-uuid (PWG5100.18)
	"output-device-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/output-device-x509-certificate (PWG5100.22)
	"output-device-x509-certificate": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/output-device-x509-request (PWG5100.22)
	"output-device-x509-request": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/preferred-attributes (PWG5100.13)
	"preferred-attributes": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Operation/printer-geo-location (PWG5100.22)
	"printer-geo-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/printer-id (PWG5100.22)
	"printer-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/printer-ids (PWG5100.22)
	"printer-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/printer-location (PWG5100.22)
	"printer-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/printer-message-from-operator (rfc3380)
	"printer-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/printer-service-type (PWG5100.22)
	"printer-service-type": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/printer-type (CUPS)
	"printer-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/printer-type-mask (CUPS)
	"printer-type-mask": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/printer-up-time (rfc3996)
	"printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/printer-uri (rfc8011)
	"printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/printer-xri-requested (PWG5100.22)
	"printer-xri-requested": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/printer-xri-requested/xri-authentication (PWG5100.22)
			"xri-authentication": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/printer-xri-requested/xri-security (PWG5100.22)
			"xri-security": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Operation/profile-uri-actual (PWG5100.16)
	"profile-uri-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/requested-attributes (rfc8011)
	"requested-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/requesting-user-name (rfc8011)
	"requesting-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/requesting-user-pkcs7-public-key (PWG5100.TRUSTNOONE)
	"requesting-user-pkcs7-public-key": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/requesting-user-uri (PWG5100.13)
	"requesting-user-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/resource-format (PWG5100.22)
	"resource-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/resource-format-accepted (PWG5100.22)
	"resource-format-accepted": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/resource-formats (PWG5100.22)
	"resource-formats": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/resource-id (PWG5100.22)
	"resource-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/resource-ids (PWG5100.22)
	"resource-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/resource-k-octets (PWG5100.22)
	"resource-k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/resource-natural-language (PWG5100.22)
	"resource-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Operation/resource-patches (PWG5100.22)
	"resource-patches": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
	},
	// Operation/resource-signature (PWG5100.22)
	"resource-signature": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Operation/resource-states (PWG5100.22)
	"resource-states": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/resource-string-version (PWG5100.22)
	"resource-string-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
	},
	// Operation/resource-type (PWG5100.22)
	"resource-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/resource-types (PWG5100.22)
	"resource-types": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/resource-version (PWG5100.22)
	"resource-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString, goipp.TagNoValue},
	},
	// Operation/restart-get-interval (PWG5100.22)
	"restart-get-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/status-message (rfc8011)
	"status-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/system-uri (PWG5100.22)
	"system-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/timeout (CUPS)
	"timeout": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/which-jobs (rfc8011)
	"which-jobs": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/which-printers (PWG5100.22)
	"which-printers": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
}

// PrinterDescription is the Printer Description attributes
var PrinterDescription = map[string]*Attribute{
	// Printer Description/accuracy-units-supported (PWG5100.21)
	"accuracy-units-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/auth-info-required (CUPS)
	"auth-info-required": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/baling-type-supported (PWG5100.1)
	"baling-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/baling-when-supported (PWG5100.1)
	"baling-when-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/binding-reference-edge-supported (PWG5100.1)
	"binding-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/binding-type-supported (PWG5100.1)
	"binding-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/chamber-humidity-default (PWG5100.21)
	"chamber-humidity-default": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Printer Description/chamber-humidity-supported (PWG5100.21)
	"chamber-humidity-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/chamber-temperature-default (PWG5100.21)
	"chamber-temperature-default": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Printer Description/chamber-temperature-supported (PWG5100.21)
	"chamber-temperature-supported": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/charset-configured (rfc8011)
	"charset-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Printer Description/charset-supported (rfc8011)
	"charset-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Printer Description/client-info-supported (PWG5100.7)
	"client-info-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/coating-sides-supported (PWG5100.1)
	"coating-sides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/coating-type-supported (PWG5100.1)
	"coating-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/color-supported (rfc8011)
	"color-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/compression-supported (rfc8011)
	"compression-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/confirmation-sheet-print-default (PWG5100.15)
	"confirmation-sheet-print-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/copies-default (rfc8011)
	"copies-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/copies-supported (rfc8011)
	"copies-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/cover-back-default (PWG5100.3)
	"cover-back-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/cover-back-supported (PWG5100.3)
	"cover-back-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/cover-front-default (PWG5100.3)
	"cover-front-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/cover-front-supported (PWG5100.3)
	"cover-front-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/cover-sheet-info-default (PWG5100.15)
	"cover-sheet-info-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/cover-sheet-info-supported (PWG5100.15)
	"cover-sheet-info-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/cover-type-supported (PWG5100.3)
	"cover-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/covering-name-supported (PWG5100.1)
	"covering-name-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/destination-accesses-supported (PWG5100.17)
	"destination-accesses-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/destination-uri-ready (PWG5100.17)
	"destination-uri-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/destination-uri-ready/destination-attributes (PWG5100.17)
			"destination-attributes": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Printer Description/destination-uri-ready/destination-attributes-supported (PWG5100.17)
			"destination-attributes-supported": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/destination-uri-ready/destination-info (PWG5100.17)
			"destination-info": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Printer Description/destination-uri-ready/destination-is-directory (PWG5100.17)
			"destination-is-directory": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Printer Description/destination-uri-ready/destination-mandatory-access-attributes (PWG5100.17)
			"destination-mandatory-access-attributes": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/destination-uri-ready/destination-name (PWG5100.17)
			"destination-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Printer Description/destination-uri-ready/destination-oauth-scope (PWG5100.17)
			"destination-oauth-scope": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
			// Printer Description/destination-uri-ready/destination-oauth-token (PWG5100.17)
			"destination-oauth-token": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
			// Printer Description/destination-uri-ready/destination-oauth-uri (PWG5100.17)
			"destination-oauth-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Printer Description/destination-uri-ready/destination-uri (PWG5100.17)
			"destination-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Printer Description/destination-uri-schemes-supported (PWG5100.15)
	"destination-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/destination-uris-supported (PWG5100.15)
	"destination-uris-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/device-uri (CUPS)
	"device-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/document-access-supported (PWG5100.18)
	"document-access-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-charset-default (PWG5100.7)
	"document-charset-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Printer Description/document-charset-supported (PWG5100.7)
	"document-charset-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Printer Description/document-creation-attributes-supported (PWG5100.5)
	"document-creation-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-digital-signature-default (PWG5100.7)
	"document-digital-signature-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-digital-signature-supported (PWG5100.7)
	"document-digital-signature-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-format-default (rfc8011)
	"document-format-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/document-format-details-default (PWG5100.7)
	"document-format-details-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/document-format-details-supported (PWG5100.7)
	"document-format-details-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-format-supported (rfc8011)
	"document-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/document-format-version-default (PWG5100.7)
	"document-format-version-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/document-format-version-supported (PWG5100.7)
	"document-format-version-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/document-natural-language-default (PWG5100.7)
	"document-natural-language-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/document-natural-language-supported (PWG5100.7)
	"document-natural-language-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/document-password-supported (PWG5100.13)
	"document-password-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/document-privacy-attributes (IPPPRIVACY10)
	"document-privacy-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-privacy-scope (IPPPRIVACY10)
	"document-privacy-scope": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/feed-orientation-default (PWG5100.11)
	"feed-orientation-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/feed-orientation-supported (PWG5100.11)
	"feed-orientation-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/fetch-document-attributes-supported (PWG5100.18)
	"fetch-document-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/finishing-template-supported (PWG5100.1)
	"finishing-template-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/finishings-col-database (PWG5100.1)
	"finishings-col-database": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/finishings-col-default (PWG5100.1)
	"finishings-col-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/finishings-col-ready (PWG5100.1)
	"finishings-col-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/finishings-col-supported (PWG5100.1)
	"finishings-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/finishings-default (rfc8011)
	"finishings-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/finishings-ready (PWG5100.1)
	"finishings-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/finishings-supported (rfc8011)
	"finishings-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/folding-direction-supported (PWG5100.1)
	"folding-direction-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/folding-offset-supported (PWG5100.1)
	"folding-offset-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/folding-reference-edge-supported (PWG5100.1)
	"folding-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/force-front-side-default  (PWG5100.3)
	"force-front-side-default ": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/force-front-side-supported (PWG5100.3)
	"force-front-side-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/force-front-side-supported  (PWG5100.3)
	"force-front-side-supported ": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/from-name-supported (PWG5100.15)
	"from-name-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/generated-natural-language-supported (rfc8011)
	"generated-natural-language-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/identify-actions-default (PWG5100.13)
	"identify-actions-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/identify-actions-supported (PWG5100.13)
	"identify-actions-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/image-orientation-default (PWG5100.3)
	"image-orientation-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/image-orientation-supported (PWG5100.3)
	"image-orientation-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/imposition-template-default (PWG5100.3)
	"imposition-template-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/imposition-template-supported (PWG5100.3)
	"imposition-template-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/input-attributes-default (PWG5100.15)
	"input-attributes-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/input-attributes-supported (PWG5100.15)
	"input-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-color-mode-supported (PWG5100.15)
	"input-color-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-content-type-supported (PWG5100.15)
	"input-content-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-film-scan-mode-supported (PWG5100.15)
	"input-film-scan-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-media-supported (PWG5100.15)
	"input-media-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/input-orientation-requested-supported (PWG5100.15)
	"input-orientation-requested-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/input-quality-supported (PWG5100.15)
	"input-quality-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/input-resolution-supported (PWG5100.15)
	"input-resolution-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/input-scan-regions-supported (PWG5100.15)
	"input-scan-regions-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/input-scan-regions-supported/x-dimension (PWG5100.15)
			"x-dimension": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Printer Description/input-scan-regions-supported/x-origin (PWG5100.15)
			"x-origin": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Printer Description/input-scan-regions-supported/y-dimension (PWG5100.15)
			"y-dimension": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Printer Description/input-scan-regions-supported/y-origin (PWG5100.15)
			"y-origin": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
		}},
	},
	// Printer Description/input-sides-supported (PWG5100.15)
	"input-sides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-source-supported (PWG5100.15)
	"input-source-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/insert-after-page-number-supported (PWG5100.3)
	"insert-after-page-number-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/insert-count-supported (PWG5100.3)
	"insert-count-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/insert-sheet-default (PWG5100.3)
	"insert-sheet-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/insert-sheet-supported (PWG5100.3)
	"insert-sheet-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/ipp-features-supported (PWG5100.13)
	"ipp-features-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/ipp-versions-supported (rfc8011)
	"ipp-versions-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/ippget-event-life (rfc3996)
	"ippget-event-life": &Attribute{
		SetOf: false,
		Min:   15,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-account-id-default (PWG5100.3)
	"job-account-id-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/job-account-id-supported (PWG5100.3)
	"job-account-id-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-account-type-default (PWG5100.16)
	"job-account-type-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-account-type-supported (PWG5100.16)
	"job-account-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-accounting-output-bin-default (PWG5100.3)
	"job-accounting-output-bin-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-accounting-output-bin-supported (PWG5100.3)
	"job-accounting-output-bin-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-accounting-sheets-default (PWG5100.3)
	"job-accounting-sheets-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-accounting-sheets-supported (PWG5100.3)
	"job-accounting-sheets-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-accounting-sheets-type-supported (PWG5100.3)
	"job-accounting-sheets-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-accounting-user-id-default (PWG5100.3)
	"job-accounting-user-id-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/job-accounting-user-id-supported (PWG5100.3)
	"job-accounting-user-id-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-authorization-uri-supported (PWG5100.16)
	"job-authorization-uri-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-cancel-after-default (PWG5100.11)
	"job-cancel-after-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-cancel-after-supported (PWG5100.7)
	"job-cancel-after-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-complete-before-supported (PWG5100.3)
	"job-complete-before-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-constraints-supported (PWG5100.13)
	"job-constraints-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/job-constraints-supported/resolver-name (PWG5100.13)
			"resolver-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
		}},
	},
	// Printer Description/job-copies-supported (PWG5100.7)
	"job-copies-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-cover-back-default (PWG5100.7)
	"job-cover-back-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-cover-back-supported (PWG5100.7)
	"job-cover-back-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-cover-front-default (PWG5100.7)
	"job-cover-front-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-cover-front-supported (PWG5100.7)
	"job-cover-front-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-creation-attributes-supported (PWG5100.7)
	"job-creation-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-delay-output-until-default (PWG5100.7)
	"job-delay-output-until-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-delay-output-until-interval-supported (PWG5100.7)
	"job-delay-output-until-interval-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-delay-output-until-supported (PWG5100.7)
	"job-delay-output-until-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-delay-output-until-time-supported (PWG5100.7)
	"job-delay-output-until-time-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-destination-spooling-supported (PWG5100.17)
	"job-destination-spooling-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-error-action-default (PWG5100.13)
	"job-error-action-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-error-action-supported (PWG5100.13)
	"job-error-action-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-error-sheet-default (PWG5100.3)
	"job-error-sheet-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-error-sheet-supported (PWG5100.3)
	"job-error-sheet-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-error-sheet-type-supported (PWG5100.3)
	"job-error-sheet-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-error-sheet-when-supported (PWG5100.3)
	"job-error-sheet-when-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-finishings-col-supported (PWG5100.7)
	"job-finishings-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-finishings-default (PWG5100.7)
	"job-finishings-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/job-finishings-ready (PWG5100.7)
	"job-finishings-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/job-finishings-supported (PWG5100.7)
	"job-finishings-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/job-history-attributes-configured (PWG5100.7)
	"job-history-attributes-configured": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-history-attributes-supported (PWG5100.7)
	"job-history-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-history-interval-configured (PWG5100.7)
	"job-history-interval-configured": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-history-interval-supported (PWG5100.7)
	"job-history-interval-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-hold-until-default (rfc8011)
	"job-hold-until-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-hold-until-supported (rfc8011)
	"job-hold-until-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-hold-until-time-supported (PWG5100.7)
	"job-hold-until-time-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-ids-supported (PWG5100.7)
	"job-ids-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-impressions-supported (rfc8011)
	"job-impressions-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-k-limit (CUPS)
	"job-k-limit": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-k-octets-supported (rfc8011)
	"job-k-octets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-mandatory-attributes-supported (PWG5100.7)
	"job-mandatory-attributes-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-media-sheets-supported (rfc8011)
	"job-media-sheets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-message-to-operator-default (PWG5100.3)
	"job-message-to-operator-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/job-message-to-operator-supported (PWG5100.3)
	"job-message-to-operator-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-page-limit (CUPS)
	"job-page-limit": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-pages-per-set-supported (PWG5100.1)
	"job-pages-per-set-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-password-encryption-supported (PWG5100.11)
	"job-password-encryption-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-password-length-supported (PWG5100.11)
	"job-password-length-supported": &Attribute{
		SetOf: false,
		Min:   4,
		Max:   765,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-password-repertoire-configured (PWG5100.11)
	"job-password-repertoire-configured": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-password-repertoire-supported (PWG5100.11)
	"job-password-repertoire-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-password-supported (PWG5100.11)
	"job-password-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-phone-number-default (PWG5100.3)
	"job-phone-number-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// Printer Description/job-phone-number-scheme-supported (PWG5100.3)
	"job-phone-number-scheme-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/job-phone-number-supported (PWG5100.3)
	"job-phone-number-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-presets-supported (PWG5100.13)
	"job-presets-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/job-presets-supported/preset-category (PWG5100.13)
			"preset-category": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/job-presets-supported/preset-name (PWG5100.13)
			"preset-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Printer Description/job-priority-default (rfc8011)
	"job-priority-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-priority-supported (rfc8011)
	"job-priority-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-privacy-attributes (IPPPRIVACY10)
	"job-privacy-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-privacy-scope (IPPPRIVACY10)
	"job-privacy-scope": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-quota-period (CUPS)
	"job-quota-period": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-recipient-name-default (PWG5100.3)
	"job-recipient-name-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/job-recipient-name-supported (PWG5100.3)
	"job-recipient-name-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-release-action-default (PWG5100.11)
	"job-release-action-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-release-action-supported (PWG5100.11)
	"job-release-action-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-resolvers-supported (PWG5100.13)
	"job-resolvers-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/job-resolvers-supported/resolver-name (PWG5100.13)
			"resolver-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
		}},
	},
	// Printer Description/job-retain-until-default (PWG5100.7)
	"job-retain-until-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-retain-until-interval-supported (PWG5100.7)
	"job-retain-until-interval-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-retain-until-supported (PWG5100.7)
	"job-retain-until-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-retain-until-time-supported (PWG5100.7)
	"job-retain-until-time-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-sheet-message-default (PWG5100.3)
	"job-sheet-message-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/job-sheet-message-supported (PWG5100.3)
	"job-sheet-message-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-sheets-col-default (PWG5100.3)
	"job-sheets-col-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-sheets-col-supported (PWG5100.3)
	"job-sheets-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-sheets-default (CUPS)
	"job-sheets-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-sheets-supported (rfc8011)
	"job-sheets-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-spooling-supported (PWG5100.7)
	"job-spooling-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-storage-access-supported (PWG5100.11)
	"job-storage-access-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-storage-disposition-supported (PWG5100.11)
	"job-storage-disposition-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-storage-group-supported (PWG5100.11)
	"job-storage-group-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/job-storage-supported (PWG5100.11)
	"job-storage-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-triggers-supported (PWG5100.13)
	"job-triggers-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/job-triggers-supported/preset-name (PWG5100.13)
			"preset-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Printer Description/jpeg-features-supported (PWG5100.13)
	"jpeg-features-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/jpeg-k-octets-supported (PWG5100.13)
	"jpeg-k-octets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/jpeg-x-dimension-supported (PWG5100.13)
	"jpeg-x-dimension-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/jpeg-y-dimension-supported (PWG5100.13)
	"jpeg-y-dimension-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/label-mode-configured (IPPLABEL)
	"label-mode-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/label-mode-supported (IPPLABEL)
	"label-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/label-tear-offset-configured (IPPLABEL)
	"label-tear-offset-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/label-tear-offset-supported (IPPLABEL)
	"label-tear-offset-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/laminating-sides-supported (PWG5100.1)
	"laminating-sides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/laminating-type-supported (PWG5100.1)
	"laminating-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/logo-uri-formats-supported (PWG5100.15)
	"logo-uri-formats-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/logo-uri-schemes-supported (PWG5100.15)
	"logo-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/marker-change-time (CUPS)
	"marker-change-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/marker-colors (CUPS)
	"marker-colors": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/marker-high-levels (CUPS)
	"marker-high-levels": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/marker-levels (CUPS)
	"marker-levels": &Attribute{
		SetOf: false,
		Min:   -3,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/marker-low-levels (CUPS)
	"marker-low-levels": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/marker-message (CUPS)
	"marker-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/marker-names (CUPS)
	"marker-names": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/marker-types (CUPS)
	"marker-types": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/material-amount-units-supported (PWG5100.21)
	"material-amount-units-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/material-diameter-supported (PWG5100.21)
	"material-diameter-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-nozzle-diameter-supported (PWG5100.21)
	"material-nozzle-diameter-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-purpose-supported (PWG5100.21)
	"material-purpose-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/material-rate-supported (PWG5100.21)
	"material-rate-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-rate-units-supported (PWG5100.21)
	"material-rate-units-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/material-shell-thickness-supported (PWG5100.21)
	"material-shell-thickness-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-temperature-supported (PWG5100.21)
	"material-temperature-supported": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-type-supported (PWG5100.21)
	"material-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/materials-col-database (PWG5100.21)
	"materials-col-database": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/materials-col-default (PWG5100.21)
	"materials-col-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/materials-col-ready (PWG5100.21)
	"materials-col-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/materials-col-supported (PWG5100.21)
	"materials-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/max-client-info-supported (PWG5100.7)
	"max-client-info-supported": &Attribute{
		SetOf: false,
		Min:   4,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/max-materials-col-supported (PWG5100.21)
	"max-materials-col-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/max-page-ranges-supported (PWG5100.7)
	"max-page-ranges-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/max-save-info-supported (PWG5100.11)
	"max-save-info-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/max-stitching-locations-supported (PWG5100.1)
	"max-stitching-locations-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-back-coating-supported (PWG5100.7)
	"media-back-coating-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-bottom-margin-supported (PWG5100.7)
	"media-bottom-margin-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-col-database (PWG5100.7)
	"media-col-database": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/media-col-database/media-size (PWG5100.7)
			"media-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Printer Description/media-col-database/media-size/x-dimension (PWG5100.7)
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
					},
					// Printer Description/media-col-database/media-size/y-dimension (PWG5100.7)
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
					},
				}},
			},
			// Printer Description/media-col-database/media-source-properties (PWG5100.7)
			"media-source-properties": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Printer Description/media-col-database/media-source-properties/media-source-feed-direction (PWG5100.7)
					"media-source-feed-direction": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Printer Description/media-col-database/media-source-properties/media-source-feed-orientation (PWG5100.7)
					"media-source-feed-orientation": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagEnum},
					},
				}},
			},
		}},
	},
	// Printer Description/media-col-default (PWG5100.7)
	"media-col-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/media-col-ready (PWG5100.7)
	"media-col-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/media-col-ready/media-size (PWG5100.7)
			"media-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Printer Description/media-col-ready/media-size/x-dimension (PWG5100.7)
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Printer Description/media-col-ready/media-size/y-dimension (PWG5100.7)
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Printer Description/media-col-ready/media-source-properties (PWG5100.7)
			"media-source-properties": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Printer Description/media-col-ready/media-source-properties/media-source-feed-direction (PWG5100.7)
					"media-source-feed-direction": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Printer Description/media-col-ready/media-source-properties/media-source-feed-orientation (PWG5100.7)
					"media-source-feed-orientation": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagEnum},
					},
				}},
			},
		}},
	},
	// Printer Description/media-col-supported (PWG5100.7)
	"media-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/media-color-supported (PWG5100.7)
	"media-color-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-default (rfc8011)
	"media-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/media-front-coating-supported (PWG5100.7)
	"media-front-coating-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-grain-supported (PWG5100.7)
	"media-grain-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-hole-count-supported (PWG5100.7)
	"media-hole-count-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/media-info-supported (PWG5100.7)
	"media-info-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/media-key-supported (PWG5100.7)
	"media-key-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-left-margin-supported (PWG5100.7)
	"media-left-margin-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-order-count-supported (PWG5100.7)
	"media-order-count-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/media-overprint-default (PWG5100.13)
	"media-overprint-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
	},
	// Printer Description/media-overprint-distance-supported (PWG5100.13)
	"media-overprint-distance-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/media-overprint-method-supported (PWG5100.13)
	"media-overprint-method-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/media-overprint-supported (PWG5100.13)
	"media-overprint-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/media-pre-printed-supported (PWG5100.7)
	"media-pre-printed-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-ready (rfc8011)
	"media-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-recycled-supported (PWG5100.7)
	"media-recycled-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-right-margin-supported (PWG5100.7)
	"media-right-margin-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-size-supported (PWG5100.7)
	"media-size-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/media-size-supported/x-dimension (PWG5100.7)
			"x-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
			},
			// Printer Description/media-size-supported/y-dimension (PWG5100.7)
			"y-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
			},
		}},
	},
	// Printer Description/media-source-supported (PWG5100.7)
	"media-source-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-supported (rfc8011)
	"media-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-thickness-supported (PWG5100.7)
	"media-thickness-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/media-tooth-supported (PWG5100.7)
	"media-tooth-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-top-margin-supported (PWG5100.7)
	"media-top-margin-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-top-offset-supported (IPPLABEL)
	"media-top-offset-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   -2147483648,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/media-tracking-supported (IPPLABEL)
	"media-tracking-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/media-type-supported (PWG5100.7)
	"media-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-weight-metric-supported (PWG5100.7)
	"media-weight-metric-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/message-supported (PWG5100.15)
	"message-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/multiple-destination-uris-supported (PWG5100.15)
	"multiple-destination-uris-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/multiple-document-handling-default (rfc8011)
	"multiple-document-handling-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/multiple-document-handling-supported (rfc8011)
	"multiple-document-handling-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/multiple-document-jobs-supported (rfc8011)
	"multiple-document-jobs-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/multiple-object-handling-default (PWG5100.21)
	"multiple-object-handling-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/multiple-object-handling-supported (PWG5100.21)
	"multiple-object-handling-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/multiple-operation-time-out (rfc8011)
	"multiple-operation-time-out": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/multiple-operation-time-out-action (PWG5100.13)
	"multiple-operation-time-out-action": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/natural-language-configured (rfc8011)
	"natural-language-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/notify-attributes-supported (rfc3995)
	"notify-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/notify-events-default (rfc3995)
	"notify-events-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/notify-events-supported (rfc3995)
	"notify-events-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/notify-lease-duration-default (rfc3995)
	"notify-lease-duration-default": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/notify-lease-duration-supported (rfc3995)
	"notify-lease-duration-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/notify-pull-method-supported (rfc3995)
	"notify-pull-method-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/notify-schemes-supported (rfc3995)
	"notify-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/number-of-retries-default (PWG5100.15)
	"number-of-retries-default": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/number-of-retries-supported (PWG5100.15)
	"number-of-retries-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/number-up-default (rfc8011)
	"number-up-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/number-up-supported (rfc8011)
	"number-up-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/oauth-authorization-scope (PWG5100.23)
	"oauth-authorization-scope": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/oauth-authorization-server-uri (PWG5100.23)
	"oauth-authorization-server-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// Printer Description/operations-supported (rfc8011)
	"operations-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/organization-name-supported (PWG5100.15)
	"organization-name-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/orientation-requested-default (rfc8011)
	"orientation-requested-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum, goipp.TagNoValue},
	},
	// Printer Description/orientation-requested-supported (rfc8011)
	"orientation-requested-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/output-attributes-default (PWG5100.17)
	"output-attributes-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
	},
	// Printer Description/output-attributes-supported (PWG5100.17)
	"output-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/output-bin-default (PWG5100.2)
	"output-bin-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/output-bin-supported (PWG5100.2)
	"output-bin-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/output-device-supported (PWG5100.7)
	"output-device-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/output-device-uuid-supported (PWG5100.18)
	"output-device-uuid-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/overrides-supported (PWG5100.6)
	"overrides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-delivery-default (PWG5100.3)
	"page-delivery-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-delivery-supported (PWG5100.3)
	"page-delivery-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-order-received-default (PWG5100.3)
	"page-order-received-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-order-received-supported (PWG5100.3)
	"page-order-received-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-ranges-supported (rfc8011)
	"page-ranges-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/pages-per-subset-supported (PWG5100.13)
	"pages-per-subset-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/parent-printers-supported (rfc3998)
	"parent-printers-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/pclm-raster-back-side (HP20180907)
	"pclm-raster-back-side": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pclm-source-resolution-supported (HP20180907)
	"pclm-source-resolution-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/pclm-strip-height-preferred (HP20180907)
	"pclm-strip-height-preferred": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/pclm-strip-height-supported (HP20180907)
	"pclm-strip-height-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/pdf-features-supported (PWG5100.21)
	"pdf-features-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pdf-k-octets-supported (PWG5100.13)
	"pdf-k-octets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/pdf-versions-supported (PWG5100.13)
	"pdf-versions-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pdl-init-file-entry-supported (PWG5100.11)
	"pdl-init-file-entry-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/pdl-init-file-location-supported (PWG5100.11)
	"pdl-init-file-location-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/pdl-init-file-name-subdirectory-supported (PWG5100.11)
	"pdl-init-file-name-subdirectory-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/pdl-init-file-name-supported (PWG5100.11)
	"pdl-init-file-name-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/pdl-init-file-supported (PWG5100.11)
	"pdl-init-file-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pdl-override-guaranteed-supported (IPPWG20151019)
	"pdl-override-guaranteed-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pdl-override-supported (rfc8011)
	"pdl-override-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pkcs7-document-format-supported (PWG5100.TRUSTNOONE)
	"pkcs7-document-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/platform-shape (PWG5100.21)
	"platform-shape": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/platform-temperature-default (PWG5100.21)
	"platform-temperature-default": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/platform-temperature-supported (PWG5100.21)
	"platform-temperature-supported": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/port-monitor (CUPS)
	"port-monitor": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/port-monitor-supported (CUPS)
	"port-monitor-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/ppd-name (CUPS)
	"ppd-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/preferred-attributes-supported (PWG5100.13)
	"preferred-attributes-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/presentation-direction-number-up-default (PWG5100.3)
	"presentation-direction-number-up-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/presentation-direction-number-up-supported (PWG5100.3)
	"presentation-direction-number-up-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-accuracy-supported (PWG5100.21)
	"print-accuracy-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/print-accuracy-supported/accuracy-units (PWG5100.21)
			"accuracy-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/print-accuracy-supported/x-accuracy (PWG5100.21)
			"x-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Printer Description/print-accuracy-supported/y-accuracy (PWG5100.21)
			"y-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Printer Description/print-accuracy-supported/z-accuracy (PWG5100.21)
			"z-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Printer Description/print-base-default (PWG5100.21)
	"print-base-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-base-supported (PWG5100.21)
	"print-base-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-color-mode-default (PWG5100.13)
	"print-color-mode-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-color-mode-icc-profiles (PWG5100.13)
	"print-color-mode-icc-profiles": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/print-color-mode-icc-profiles/print-color-mode (PWG5100.13)
			"print-color-mode": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/print-color-mode-icc-profiles/profile-uri (PWG5100.13)
			"profile-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Printer Description/print-color-mode-supported (PWG5100.13)
	"print-color-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-content-optimize-default (PWG5100.7)
	"print-content-optimize-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-content-optimize-supported (PWG5100.7)
	"print-content-optimize-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-darkness-default (IPPLABEL)
	"print-darkness-default": &Attribute{
		SetOf: false,
		Min:   -100,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/print-darkness-supported (IPPLABEL)
	"print-darkness-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/print-objects-supported (PWG5100.21)
	"print-objects-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-processing-attributes-supported (PWG5100.13)
	"print-processing-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-quality-default (rfc8011)
	"print-quality-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/print-quality-supported (rfc8011)
	"print-quality-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/print-rendering-intent-default (PWG5100.13)
	"print-rendering-intent-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-rendering-intent-supported (PWG5100.13)
	"print-rendering-intent-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-scaling-default (PWG5100.13)
	"print-scaling-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-scaling-supported (PWG5100.13)
	"print-scaling-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-speed-default (IPPLABEL)
	"print-speed-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/print-speed-supported (IPPLABEL)
	"print-speed-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/print-supports-default (PWG5100.21)
	"print-supports-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-supports-supported (PWG5100.21)
	"print-supports-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-asset-tag (PWG5100.11)
	"printer-asset-tag": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Description/printer-camera-image-uri (PWG5100.21)
	"printer-camera-image-uri": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-charge-info (PWG5100.16)
	"printer-charge-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-charge-info-uri (PWG5100.16)
	"printer-charge-info-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-commands (CUPS)
	"printer-commands": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-contact-col (PWG5100.22)
	"printer-contact-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagUnknown},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-contact-col/contact-name (PWG5100.22)
			"contact-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Printer Description/printer-contact-col/contact-uri (PWG5100.22)
			"contact-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Printer Description/printer-contact-col/contact-vcard (PWG5100.22)
			"contact-vcard": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// Printer Description/printer-current-time (rfc8011)
	"printer-current-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Printer Description/printer-darkness-configured (IPPLABEL)
	"printer-darkness-configured": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/printer-darkness-supported (IPPLABEL)
	"printer-darkness-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/printer-detailed-status-messages (PWG5100.11)
	"printer-detailed-status-messages": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-device-id (PWG5107.2)
	"printer-device-id": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-dns-sd-name (PWG5100.13)
	"printer-dns-sd-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   63,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/printer-driver-installer (rfc8011)
	"printer-driver-installer": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-fax-log-uri (PWG5100.15)
	"printer-fax-log-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-fax-modem-info (PWG5100.15)
	"printer-fax-modem-info": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-fax-modem-name (PWG5100.15)
	"printer-fax-modem-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/printer-fax-modem-number (PWG5100.15)
	"printer-fax-modem-number": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-geo-location (PWG5100.13)
	"printer-geo-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagUnknown},
	},
	// Printer Description/printer-get-attributes-supported (PWG5100.13)
	"printer-get-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-icc-profiles (PWG5100.13)
	"printer-icc-profiles": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-icc-profiles/profile-name (PWG5100.13)
			"profile-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Printer Description/printer-icc-profiles/profile-url (PWG5100.13)
			"profile-url": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Printer Description/printer-icons (PWG5100.13)
	"printer-icons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-id (CUPS)
	"printer-id": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/printer-info (rfc8011)
	"printer-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-is-accepting-jobs (CUPS)
	"printer-is-accepting-jobs": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/printer-is-shared (CUPS)
	"printer-is-shared": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/printer-is-temporary (CUPS)
	"printer-is-temporary": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/printer-kind (PWG5100.13)
	"printer-kind": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/printer-location (rfc8011)
	"printer-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-make-and-model (rfc8011)
	"printer-make-and-model": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-mandatory-job-attributes (PWG5100.13)
	"printer-mandatory-job-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-mode-configured (PWG5100.18)
	"printer-mode-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-mode-supported (PWG5100.18)
	"printer-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-more-info (CUPS)
	"printer-more-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-more-info-manufacturer (rfc8011)
	"printer-more-info-manufacturer": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-name (rfc8011)
	"printer-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/printer-organization (PWG5100.13)
	"printer-organization": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-organizational-unit (PWG5100.13)
	"printer-organizational-unit": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-pkcs7-public-key (PWG5100.TRUSTNOONE)
	"printer-pkcs7-public-key": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-pkcs7-repertoire-configured (PWG5100.TRUSTNOONE)
	"printer-pkcs7-repertoire-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-pkcs7-repertoire-supported (PWG5100.TRUSTNOONE)
	"printer-pkcs7-repertoire-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-privacy-policy-uri (IPPPRIVACY10)
	"printer-privacy-policy-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-requested-job-attributes (PWG5100.16)
	"printer-requested-job-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-resolution-default (rfc8011)
	"printer-resolution-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/printer-resolution-supported (rfc8011)
	"printer-resolution-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/printer-service-contact-col (PWG5100.11)
	"printer-service-contact-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-service-contact-col/contact-name (PWG5100.11)
			"contact-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Printer Description/printer-service-contact-col/contact-uri (PWG5100.11)
			"contact-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Printer Description/printer-service-contact-col/contact-vcard (PWG5100.11)
			"contact-vcard": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// Printer Description/printer-state (CUPS)
	"printer-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/printer-state-message (CUPS)
	"printer-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-static-resource-directory-uri (PWG5100.18)
	"printer-static-resource-directory-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-static-resource-k-octets-supported (PWG5100.18)
	"printer-static-resource-k-octets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/printer-strings-languages-supported (PWG5100.13)
	"printer-strings-languages-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/printer-strings-uri (PWG5100.13)
	"printer-strings-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// Printer Description/printer-type (CUPS)
	"printer-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/printer-type-mask (CUPS)
	"printer-type-mask": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/printer-volume-supported (PWG5100.21)
	"printer-volume-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-volume-supported/x-dimension (PWG5100.21)
			"x-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Printer Description/printer-volume-supported/y-dimension (PWG5100.21)
			"y-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Printer Description/printer-volume-supported/z-dimension (PWG5100.21)
			"z-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Printer Description/printer-wifi-password (IPPWIFI)
	"printer-wifi-password": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Description/printer-wifi-ssid (IPPWIFI)
	"printer-wifi-ssid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/printer-xri-supported (rfc3380)
	"printer-xri-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-xri-supported/xri-authentication (rfc3380)
			"xri-authentication": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/printer-xri-supported/xri-security (rfc3380)
			"xri-security": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/printer-xri-supported/xri-uri (rfc3380)
			"xri-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Printer Description/proof-copies-supported (PWG5100.11)
	"proof-copies-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/proof-print-copies-supported (PWG5100.11)
	"proof-print-copies-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/proof-print-default (PWG5100.11)
	"proof-print-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
	},
	// Printer Description/proof-print-supported (PWG5100.11)
	"proof-print-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/punching-hole-diameter-configured (PWG5100.1)
	"punching-hole-diameter-configured": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/punching-locations-supported (PWG5100.1)
	"punching-locations-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/punching-offset-supported (PWG5100.1)
	"punching-offset-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/punching-reference-edge-supported (PWG5100.1)
	"punching-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pwg-raster-document-resolution-supported (PWG5102.4)
	"pwg-raster-document-resolution-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/pwg-raster-document-sheet-back (PWG5102.4)
	"pwg-raster-document-sheet-back": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pwg-raster-document-type-supported (PWG5102.4)
	"pwg-raster-document-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pwg-safe-gcode-supported (PWG5199.7)
	"pwg-safe-gcode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/reference-uri-schemes-supported (rfc8011)
	"reference-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/repertoire-supported (PWG5101.2)
	"repertoire-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/requesting-user-name-allowed (CUPS)
	"requesting-user-name-allowed": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagDeleteAttr},
	},
	// Printer Description/requesting-user-name-denied (CUPS)
	"requesting-user-name-denied": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagDeleteAttr},
	},
	// Printer Description/requesting-user-uri-schemes-supported (PWG5100.13)
	"requesting-user-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/requesting-user-uri-supported (PWG5100.13)
	"requesting-user-uri-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/retry-interval-default (PWG5100.15)
	"retry-interval-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/retry-interval-supported (PWG5100.15)
	"retry-interval-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/retry-time-out-default (PWG5100.15)
	"retry-time-out-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/retry-time-out-supported (PWG5100.15)
	"retry-time-out-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/save-disposition-supported (PWG5100.11)
	"save-disposition-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/save-document-format-default (PWG5100.11)
	"save-document-format-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/save-document-format-supported (PWG5100.11)
	"save-document-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/save-location-default (PWG5100.11)
	"save-location-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/save-location-supported (PWG5100.11)
	"save-location-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/save-name-subdirectory-supported (PWG5100.11)
	"save-name-subdirectory-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/save-name-supported (PWG5100.11)
	"save-name-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/separator-sheets-default (PWG5100.3)
	"separator-sheets-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/separator-sheets-supported (PWG5100.3)
	"separator-sheets-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/sheet-collate-default (rfc3381)
	"sheet-collate-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/sheet-collate-supported (rfc3381)
	"sheet-collate-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/sides-default (rfc8011)
	"sides-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/sides-supported (rfc8011)
	"sides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/smi2699-auth-print-group (IPPSERVER)
	"smi2699-auth-print-group": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/smi2699-auth-proxy-group (IPPSERVER)
	"smi2699-auth-proxy-group": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/smi2699-device-command (IPPSERVER)
	"smi2699-device-command": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/smi2699-device-format (IPPSERVER)
	"smi2699-device-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/smi2699-device-name (IPPSERVER)
	"smi2699-device-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/smi2699-device-uri (IPPSERVER)
	"smi2699-device-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/stitching-angle-supported (PWG5100.1)
	"stitching-angle-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   359,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/stitching-locations-supported (PWG5100.1)
	"stitching-locations-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/stitching-method-supported (PWG5100.1)
	"stitching-method-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/stitching-offset-supported (PWG5100.1)
	"stitching-offset-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/stitching-reference-edge-supported (PWG5100.1)
	"stitching-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/subject-supported (PWG5100.15)
	"subject-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/subordinate-printers-supported (rfc3998)
	"subordinate-printers-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/subscription-privacy-attributes (IPPPRIVACY10)
	"subscription-privacy-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/subscription-privacy-scope (IPPPRIVACY10)
	"subscription-privacy-scope": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/to-name-supported (PWG5100.15)
	"to-name-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/trimming-offset-supported (PWG5100.1)
	"trimming-offset-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/trimming-reference-edge-supported (PWG5100.1)
	"trimming-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/trimming-type-supported (PWG5100.1)
	"trimming-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/trimming-when-supported (PWG5100.1)
	"trimming-when-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/urf-supported (CUPS)
	"urf-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/uri-authentication-supported (rfc8011)
	"uri-authentication-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/uri-security-supported (rfc8011)
	"uri-security-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/user-defined-values-supported (PWG5100.3)
	"user-defined-values-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/which-jobs-supported (PWG5100.7)
	"which-jobs-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/x-image-position-default (PWG5100.3)
	"x-image-position-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/x-image-position-supported (PWG5100.3)
	"x-image-position-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/x-image-shift-default (PWG5100.3)
	"x-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/x-image-shift-supported (PWG5100.3)
	"x-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/x-side1-image-shift-default (PWG5100.3)
	"x-side1-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/x-side1-image-shift-supported (PWG5100.3)
	"x-side1-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/x-side2-image-shift-default (PWG5100.3)
	"x-side2-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/x-side2-image-shift-supported (PWG5100.3)
	"x-side2-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/y-image-position-default (PWG5100.3)
	"y-image-position-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/y-image-position-supported (PWG5100.3)
	"y-image-position-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/y-image-shift-default (PWG5100.3)
	"y-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/y-image-shift-supported (PWG5100.3)
	"y-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/y-side1-image-shift-default (PWG5100.3)
	"y-side1-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/y-side1-image-shift-supported (PWG5100.3)
	"y-side1-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/y-side2-image-shift-default (PWG5100.3)
	"y-side2-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/y-side2-image-shift-supported (PWG5100.3)
	"y-side2-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
}

// PrinterStatus is the Printer Status attributes
var PrinterStatus = map[string]*Attribute{
	// Printer Status/chamber-humidity-current (PWG5100.21)
	"chamber-humidity-current": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagUnknown},
	},
	// Printer Status/chamber-temperature-current (PWG5100.21)
	"chamber-temperature-current": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagUnknown},
	},
	// Printer Status/device-service-count (PWG5100.13)
	"device-service-count": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/device-uuid (PWG5100.13)
	"device-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/document-format-varying-attributes (rfc3380)
	"document-format-varying-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/job-settable-attributes-supported (rfc3380)
	"job-settable-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/pages-per-minute (rfc8011)
	"pages-per-minute": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/pages-per-minute-color (rfc8011)
	"pages-per-minute-color": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-alert (PWG5100.9)
	"printer-alert": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-alert-description (PWG5100.9)
	"printer-alert-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-camera-image-uri (PWG5100.21)
	"printer-camera-image-uri": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-config-change-date-time (PWG5100.13)
	"printer-config-change-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Printer Status/printer-config-change-time (PWG5100.13)
	"printer-config-change-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-config-changes (PWG5100.22)
	"printer-config-changes": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-console-display (IPPCONSOLE)
	"printer-console-display": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-console-light (IPPCONSOLE)
	"printer-console-light": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-console-light-description (IPPCONSOLE)
	"printer-console-light-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-cover (IPP20210223)
	"printer-cover": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-cover-description (IPP20210223)
	"printer-cover-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-detailed-status-messages (PWG5100.7)
	"printer-detailed-status-messages": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-finisher (PWG5100.1)
	"printer-finisher": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-finisher-description (PWG5100.1)
	"printer-finisher-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-finisher-supplies (PWG5100.1)
	"printer-finisher-supplies": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-finisher-supplies-description (PWG5100.1)
	"printer-finisher-supplies-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-firmware-name (PWG5100.13)
	"printer-firmware-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Status/printer-firmware-patches (PWG5100.13)
	"printer-firmware-patches": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-firmware-string-version (PWG5100.13)
	"printer-firmware-string-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-firmware-version (PWG5100.13)
	"printer-firmware-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-home-page-uri (IPPCONSOLE)
	"printer-home-page-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-id (PWG5100.22)
	"printer-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-impressions-completed (PWG5100.22)
	"printer-impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-impressions-completed-col (PWG5100.22)
	"printer-impressions-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Status/printer-input-tray (PWG5100.13)
	"printer-input-tray": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-is-accepting-jobs (rfc8011)
	"printer-is-accepting-jobs": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Status/printer-media-sheets-completed (PWG5100.22)
	"printer-media-sheets-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-media-sheets-completed-col (PWG5100.22)
	"printer-media-sheets-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Status/printer-message-date-time (rfc3380)
	"printer-message-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Printer Status/printer-message-from-operator (rfc8011)
	"printer-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-message-time (rfc3380)
	"printer-message-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-more-info (rfc8011)
	"printer-more-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-output-tray (PWG5100.13)
	"printer-output-tray": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-pages-completed (PWG5100.22)
	"printer-pages-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-pages-completed-col (PWG5100.22)
	"printer-pages-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Status/printer-serial-number (PWG5100.11)
	"printer-serial-number": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-service-type (PWG5100.22)
	"printer-service-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/printer-settable-attributes-supported (rfc3380)
	"printer-settable-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/printer-state (rfc8011)
	"printer-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Status/printer-state-change-date-time (rfc3995)
	"printer-state-change-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Printer Status/printer-state-change-time (rfc3995)
	"printer-state-change-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-state-message (rfc8011)
	"printer-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-state-reasons (rfc8011)
	"printer-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/printer-static-resource-k-octets-free (PWG5100.18)
	"printer-static-resource-k-octets-free": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-storage (PWG5100.11)
	"printer-storage": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-storage-description (PWG5100.11)
	"printer-storage-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-supply (PWG5100.13)
	"printer-supply": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-supply-description (PWG5100.13)
	"printer-supply-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-supply-info-uri (PWG5100.13)
	"printer-supply-info-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-up-time (rfc8011)
	"printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-uri-supported (rfc8011)
	"printer-uri-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-uuid (PWG5100.13)
	"printer-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-wifi-state (IPPWIFI)
	"printer-wifi-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Status/queued-job-count (rfc8011)
	"queued-job-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/xri-authentication-supported (rfc3380)
	"xri-authentication-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/xri-security-supported (rfc3380)
	"xri-security-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/xri-uri-scheme-supported (rfc3380)
	"xri-uri-scheme-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
}

// ResourceDescription is the Resource Description attributes
var ResourceDescription = map[string]*Attribute{
	// Resource Description/resource-info (PWG5100.22)
	"resource-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Resource Description/resource-name (PWG5100.22)
	"resource-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
}

// ResourceStatus is the Resource Status attributes
var ResourceStatus = map[string]*Attribute{
	// Resource Status/date-time-at-canceled (PWG5100.22)
	"date-time-at-canceled": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagNoValue},
	},
	// Resource Status/date-time-at-creation (PWG5100.22)
	"date-time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Resource Status/date-time-at-installed (PWG5100.22)
	"date-time-at-installed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagNoValue},
	},
	// Resource Status/resource-data-uri (PWG5100.22)
	"resource-data-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// Resource Status/resource-format (PWG5100.22)
	"resource-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Resource Status/resource-id (PWG5100.22)
	"resource-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Resource Status/resource-k-octets (PWG5100.22)
	"resource-k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Resource Status/resource-natural-language (PWG5100.22)
	"resource-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Resource Status/resource-patches (PWG5100.22)
	"resource-patches": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
	},
	// Resource Status/resource-signature (PWG5100.22)
	"resource-signature": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Resource Status/resource-state (PWG5100.22)
	"resource-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Resource Status/resource-state-message (PWG5100.22)
	"resource-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Resource Status/resource-state-reasons (PWG5100.22)
	"resource-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Resource Status/resource-string-version (PWG5100.22)
	"resource-string-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
	},
	// Resource Status/resource-type (PWG5100.22)
	"resource-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Resource Status/resource-use-count (PWG5100.22)
	"resource-use-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Resource Status/resource-uuid (PWG5100.22)
	"resource-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   45,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Resource Status/resource-version (PWG5100.22)
	"resource-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString, goipp.TagNoValue},
	},
	// Resource Status/time-at-canceled (PWG5100.22)
	"time-at-canceled": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Resource Status/time-at-creation (PWG5100.22)
	"time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Resource Status/time-at-installed (PWG5100.22)
	"time-at-installed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
}

// SubscriptionStatus is the Subscription Status attributes
var SubscriptionStatus = map[string]*Attribute{
	// Subscription Status/notify-job-id (rfc3995)
	"notify-job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-lease-expiration-time (rfc3995)
	"notify-lease-expiration-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-printer-up-time (rfc3995)
	"notify-printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-printer-uri (rfc3995)
	"notify-printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Subscription Status/notify-resource-id (PWG5100.22)
	"notify-resource-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-sequence-number (rfc3995)
	"notify-sequence-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-status-code (rfc3995)
	"notify-status-code": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Subscription Status/notify-subscriber-user-name (rfc3995)
	"notify-subscriber-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Subscription Status/notify-subscriber-user-uri (PWG5100.13)
	"notify-subscriber-user-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Subscription Status/notify-subscription-id (rfc3995)
	"notify-subscription-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-subscription-uuid (PWG5100.13)
	"notify-subscription-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Subscription Status/notify-system-up-time (PWG5100.22)
	"notify-system-up-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-system-uri (PWG5100.22)
	"notify-system-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
}

// SubscriptionTemplate is the Subscription Template attributes
var SubscriptionTemplate = map[string]*Attribute{
	// Subscription Template/notify-attributes (rfc3995)
	"notify-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Subscription Template/notify-charset (rfc3995)
	"notify-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Subscription Template/notify-events (rfc3995)
	"notify-events": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Subscription Template/notify-lease-duration (rfc3995)
	"notify-lease-duration": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Template/notify-max-events-supported (rfc3995)
	"notify-max-events-supported": &Attribute{
		SetOf: false,
		Min:   2,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Template/notify-natural-language (rfc3995)
	"notify-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Subscription Template/notify-pull-method (rfc3995)
	"notify-pull-method": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Subscription Template/notify-recipient-uri (rfc3995)
	"notify-recipient-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Subscription Template/notify-time-interval (rfc3995)
	"notify-time-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Template/notify-user-data (rfc3995)
	"notify-user-data": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   63,
		Tags:  []goipp.Tag{goipp.TagString},
	},
}

// SystemDescription is the System Description attributes
var SystemDescription = map[string]*Attribute{
	// System Description/charset-configured (PWG5100.22)
	"charset-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// System Description/charset-supported (PWG5100.22)
	"charset-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// System Description/document-format-supported (PWG5100.22)
	"document-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// System Description/generated-natural-language-supported (PWG5100.22)
	"generated-natural-language-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// System Description/ipp-features-supported (PWG5100.22)
	"ipp-features-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/ipp-versions-supported (PWG5100.22)
	"ipp-versions-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/ippget-event-life (PWG5100.22)
	"ippget-event-life": &Attribute{
		SetOf: false,
		Min:   15,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Description/multiple-document-printers-supported (PWG5100.22)
	"multiple-document-printers-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// System Description/natural-language-configured (PWG5100.22)
	"natural-language-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// System Description/notify-attributes-supported (PWG5100.22)
	"notify-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/notify-events-default (PWG5100.22)
	"notify-events-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/notify-events-supported (PWG5100.22)
	"notify-events-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/notify-lease-duration-default (PWG5100.22)
	"notify-lease-duration-default": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Description/notify-lease-duration-supported (PWG5100.22)
	"notify-lease-duration-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// System Description/notify-max-events-supported (PWG5100.22)
	"notify-max-events-supported": &Attribute{
		SetOf: false,
		Min:   2,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Description/notify-pull-method-supported (PWG5100.22)
	"notify-pull-method-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/notify-schemes-supported (PWG5100.22)
	"notify-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// System Description/oauth-authorization-scope (PWG5100.23)
	"oauth-authorization-scope": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// System Description/oauth-authorization-server-uri (PWG5100.23)
	"oauth-authorization-server-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// System Description/operations-supported (PWG5100.22)
	"operations-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// System Description/output-device-x509-type-supported (PWG5100.22)
	"output-device-x509-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/power-calendar-policy-col (PWG5100.22)
	"power-calendar-policy-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Description/power-calendar-policy-col/calendar-id (PWG5100.22)
			"calendar-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/day-of-month (PWG5100.22)
			"day-of-month": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   31,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/day-of-week (PWG5100.22)
			"day-of-week": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   7,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/hour (PWG5100.22)
			"hour": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   23,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/minute (PWG5100.22)
			"minute": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   59,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/month (PWG5100.22)
			"month": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   12,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/request-power-state (PWG5100.22)
			"request-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Description/power-calendar-policy-col/run-once (PWG5100.22)
			"run-once": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
		}},
	},
	// System Description/power-event-policy-col (PWG5100.22)
	"power-event-policy-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Description/power-event-policy-col/event-id (PWG5100.22)
			"event-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-event-policy-col/event-name (PWG5100.22)
			"event-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// System Description/power-event-policy-col/request-power-state (PWG5100.22)
			"request-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// System Description/power-timeout-policy-col (PWG5100.22)
	"power-timeout-policy-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Description/power-timeout-policy-col/start-power-state (PWG5100.22)
			"start-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Description/power-timeout-policy-col/timeout-id (PWG5100.22)
			"timeout-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-timeout-policy-col/timeout-predicate (PWG5100.22)
			"timeout-predicate": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Description/power-timeout-policy-col/timeout-seconds (PWG5100.22)
			"timeout-seconds": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// System Description/printer-creation-attributes-supported (PWG5100.22)
	"printer-creation-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/resource-format-supported (PWG5100.22)
	"resource-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// System Description/resource-settable-attributes-supported (PWG5100.22)
	"resource-settable-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/resource-type-supported (PWG5100.22)
	"resource-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/smi2699-auth-group-supported (IPPSERVER)
	"smi2699-auth-group-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Description/smi2699-device-command-supported (IPPSERVER)
	"smi2699-device-command-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Description/smi2699-device-format-supported (IPPSERVER)
	"smi2699-device-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// System Description/smi2699-device-uri-schemes-supported (IPPSERVER)
	"smi2699-device-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// System Description/system-asset-tag (PWG5100.22)
	"system-asset-tag": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// System Description/system-contact-col (PWG5100.22)
	"system-contact-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagUnknown},
	},
	// System Description/system-current-time (PWG5100.22)
	"system-current-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// System Description/system-default-printer-id (PWG5100.22)
	"system-default-printer-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// System Description/system-dns-sd-name (PWG5100.22)
	"system-dns-sd-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   63,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Description/system-geo-location (PWG5100.22)
	"system-geo-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagUnknown},
	},
	// System Description/system-info (PWG5100.22)
	"system-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Description/system-location (PWG5100.22)
	"system-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Description/system-make-and-model (PWG5100.22)
	"system-make-and-model": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Description/system-mandatory-printer-attributes (PWG5100.22)
	"system-mandatory-printer-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/system-mandatory-registration-attributes (PWG5100.22)
	"system-mandatory-registration-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/system-message-from-operator (PWG5100.22)
	"system-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Description/system-name (PWG5100.22)
	"system-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Description/system-service-contact-col (PWG5100.22)
	"system-service-contact-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagUnknown},
	},
	// System Description/system-settable-attributes-supported (PWG5100.22)
	"system-settable-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/system-strings-languages-supported (PWG5100.22)
	"system-strings-languages-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// System Description/system-strings-uri (PWG5100.22)
	"system-strings-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// System Description/system-xri-supported (PWG5100.22)
	"system-xri-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
}

// SystemStatus is the System Status attributes
var SystemStatus = map[string]*Attribute{
	// System Status/power-log-col (PWG5100.22)
	"power-log-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-log-col/log-id (PWG5100.22)
			"log-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-log-col/power-state (PWG5100.22)
			"power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/power-log-col/power-state-date-time (PWG5100.22)
			"power-state-date-time": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagDateTime},
			},
			// System Status/power-log-col/power-state-message (PWG5100.22)
			"power-state-message": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// System Status/power-state-capabilities-col (PWG5100.22)
	"power-state-capabilities-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-state-capabilities-col/can-accept-jobs (PWG5100.22)
			"can-accept-jobs": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// System Status/power-state-capabilities-col/can-process-jobs (PWG5100.22)
			"can-process-jobs": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// System Status/power-state-capabilities-col/power-active-watts (PWG5100.22)
			"power-active-watts": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-capabilities-col/power-inactive-watts (PWG5100.22)
			"power-inactive-watts": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-capabilities-col/power-state (PWG5100.22)
			"power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// System Status/power-state-counters-col (PWG5100.22)
	"power-state-counters-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-state-counters-col/hibernate-transitions (PWG5100.22)
			"hibernate-transitions": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-counters-col/on-transitions (PWG5100.22)
			"on-transitions": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-counters-col/standby-transitions (PWG5100.22)
			"standby-transitions": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-counters-col/suspend-transitions (PWG5100.22)
			"suspend-transitions": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// System Status/power-state-monitor-col (PWG5100.22)
	"power-state-monitor-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-state-monitor-col/current-month-kwh (PWG5100.22)
			"current-month-kwh": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-monitor-col/current-watts (PWG5100.22)
			"current-watts": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-monitor-col/lifetime-kwh (PWG5100.22)
			"lifetime-kwh": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-monitor-col/meters-are-actual (PWG5100.22)
			"meters-are-actual": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// System Status/power-state-monitor-col/power-state (PWG5100.22)
			"power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/power-state-monitor-col/power-state-message (PWG5100.22)
			"power-state-message": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// System Status/power-state-monitor-col/power-usage-is-rms-watts (PWG5100.22)
			"power-usage-is-rms-watts": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
		}},
	},
	// System Status/power-state-transitions-col (PWG5100.22)
	"power-state-transitions-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-state-transitions-col/end-power-state (PWG5100.22)
			"end-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/power-state-transitions-col/start-power-state (PWG5100.22)
			"start-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/power-state-transitions-col/state-transition-seconds (PWG5100.22)
			"state-transition-seconds": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// System Status/system-config-change-date-time (PWG5100.22)
	"system-config-change-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// System Status/system-config-change-time (PWG5100.22)
	"system-config-change-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-config-changes (PWG5100.22)
	"system-config-changes": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-configured-printers (PWG5100.22)
	"system-configured-printers": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/system-configured-printers/printer-id (PWG5100.22)
			"printer-id": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   65535,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/system-configured-printers/printer-info (PWG5100.22)
			"printer-info": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// System Status/system-configured-printers/printer-is-accepting-jobs (PWG5100.22)
			"printer-is-accepting-jobs": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// System Status/system-configured-printers/printer-name (PWG5100.22)
			"printer-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// System Status/system-configured-printers/printer-service-type (PWG5100.22)
			"printer-service-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/system-configured-printers/printer-state (PWG5100.22)
			"printer-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
			// System Status/system-configured-printers/printer-state-reasons (PWG5100.22)
			"printer-state-reasons": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/system-configured-printers/printer-xri-supported (PWG5100.22)
			"printer-xri-supported": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// System Status/system-configured-resources (PWG5100.22)
	"system-configured-resources": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/system-configured-resources/resource-format (PWG5100.22)
			"resource-format": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagMimeType},
			},
			// System Status/system-configured-resources/resource-id (PWG5100.22)
			"resource-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/system-configured-resources/resource-info (PWG5100.22)
			"resource-info": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// System Status/system-configured-resources/resource-name (PWG5100.22)
			"resource-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// System Status/system-configured-resources/resource-state (PWG5100.22)
			"resource-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
			// System Status/system-configured-resources/resource-type (PWG5100.22)
			"resource-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// System Status/system-firmware-name (PWG5100.22)
	"system-firmware-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Status/system-firmware-patches (PWG5100.22)
	"system-firmware-patches": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-firmware-string-version (PWG5100.22)
	"system-firmware-string-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-firmware-version (PWG5100.22)
	"system-firmware-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// System Status/system-impressions-completed (PWG5100.22)
	"system-impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-impressions-completed-col (PWG5100.22)
	"system-impressions-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// System Status/system-media-sheets-completed (PWG5100.22)
	"system-media-sheets-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-media-sheets-completed-col (PWG5100.22)
	"system-media-sheets-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// System Status/system-pages-completed (PWG5100.22)
	"system-pages-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-pages-completed-col (PWG5100.22)
	"system-pages-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// System Status/system-resident-application-name (PWG5100.22)
	"system-resident-application-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Status/system-resident-application-patches (PWG5100.22)
	"system-resident-application-patches": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-resident-application-string-version (PWG5100.22)
	"system-resident-application-string-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-resident-application-version (PWG5100.22)
	"system-resident-application-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// System Status/system-serial-number (PWG5100.22)
	"system-serial-number": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-state (PWG5100.22)
	"system-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// System Status/system-state-change-date-time (PWG5100.22)
	"system-state-change-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// System Status/system-state-change-time (PWG5100.22)
	"system-state-change-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-state-message (PWG5100.22)
	"system-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-state-reasons (PWG5100.22)
	"system-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Status/system-time-source (PWG5100.22)
	"system-time-source": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// System Status/system-up-time (PWG5100.22)
	"system-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-user-application-name (PWG5100.22)
	"system-user-application-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Status/system-user-application-patches (PWG5100.22)
	"system-user-application-patches": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-user-application-string-version (PWG5100.22)
	"system-user-application-string-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-user-application-version (PWG5100.22)
	"system-user-application-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// System Status/system-uuid (PWG5100.22)
	"system-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   45,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// System Status/xri-authentication-supported (PWG5100.22)
	"xri-authentication-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Status/xri-security-supported (PWG5100.22)
	"xri-security-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Status/xri-uri-scheme-supported (PWG5100.22)
	"xri-uri-scheme-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
}

// Collections contains all top-level collections (groups) of
// attributes, indexed by name
var Collections = map[string]map[string]*Attribute{
	"CUPS Device Attributes":        CUPSDeviceAttributes,
	"CUPS PPD Attributes":           CUPSPPDAttributes,
	"CUPS Printer Class Attributes": CUPSPrinterClassAttributes,
	"Document Description":          DocumentDescription,
	"Document Status":               DocumentStatus,
	"Document Template":             DocumentTemplate,
	"Event Notifications":           EventNotifications,
	"Job Description":               JobDescription,
	"Job Status":                    JobStatus,
	"Job Template":                  JobTemplate,
	"Operation":                     Operation,
	"Printer Description":           PrinterDescription,
	"Printer Status":                PrinterStatus,
	"Resource Description":          ResourceDescription,
	"Resource Status":               ResourceStatus,
	"Subscription Status":           SubscriptionStatus,
	"Subscription Template":         SubscriptionTemplate,
	"System Description":            SystemDescription,
	"System Status":                 SystemStatus,
}

// borrowings contains a table of attributes borrowing
// between collections.
var borrowings = []borrowing{
	{"Document Status/cover-back-actual", "Job Template/cover-back"},
	{"Document Status/cover-front-actual", "Job Template/cover-front"},
	{"Document Status/document-format-details", "Operation/document-format-details"},
	{"Document Status/document-format-details-detected", "Operation/document-format-details"},
	{"Document Status/finishings-col-actual", "Job Template/finishings-col"},
	{"Document Status/impressions-completed-col", "Document Status/impressions-col"},
	{"Document Status/input-attributes-actual", "Operation/input-attributes"},
	{"Document Status/insert-sheet-actual", "Job Template/insert-sheet"},
	{"Document Status/materials-col-actual", "Document Template/materials-col"},
	{"Document Status/media-col-actual", "Job Template/media-col"},
	{"Document Status/media-sheets-completed-col", "Document Status/media-sheets-col"},
	{"Document Status/output-attributes-actual", "Operation/output-attributes"},
	{"Document Status/pages-completed-col", "Document Status/pages-col"},
	{"Document Status/print-accuracy-actual", "Document Template/print-accuracy"},
	{"Document Status/print-objects-actual", "Document Template/print-objects"},
	{"Document Status/separator-sheets-actual", "Job Template/separator-sheets"},
	{"Document Template/cover-back", "Job Template/cover-back"},
	{"Document Template/cover-front", "Job Template/cover-front"},
	{"Document Template/finishings-col", "Job Template/finishings-col"},
	{"Document Template/insert-sheet", "Job Template/insert-sheet"},
	{"Document Template/media-col", "Job Template/media-col"},
	{"Document Template/separator-sheets", "Job Template/separator-sheets"},
	{"Job Status/client-info", "Operation/client-info"},
	{"Job Status/cover-back-actual", "Job Template/cover-back"},
	{"Job Status/cover-front-actual", "Job Template/cover-front"},
	{"Job Status/document-format-details-detected", "Operation/document-format-details"},
	{"Job Status/document-format-details-supplied", "Operation/document-format-details"},
	{"Job Status/finishings-col-actual", "Job Template/finishings-col"},
	{"Job Status/input-attributes-actual", "Operation/input-attributes"},
	{"Job Status/insert-sheet-actual", "Job Template/insert-sheet"},
	{"Job Status/job-accounting-sheets-actual", "Job Template/job-accounting-sheets"},
	{"Job Status/job-cover-back-actual", "Job Template/cover-back"},
	{"Job Status/job-cover-front-actual", "Job Template/cover-front"},
	{"Job Status/job-error-sheet-actual", "Job Template/job-error-sheet"},
	{"Job Status/job-impressions-completed-col", "Job Status/job-impressions-col"},
	{"Job Status/job-media-sheets-completed-col", "Job Status/job-media-sheets-col"},
	{"Job Status/job-pages-completed-col", "Job Status/job-pages-col"},
	{"Job Status/job-sheets-col-actual", "Job Template/job-sheets-col"},
	{"Job Status/job-storage", "Operation/job-storage"},
	{"Job Status/materials-col-actual", "Document Template/materials-col"},
	{"Job Status/media-col-actual", "Job Template/media-col"},
	{"Job Status/output-attributes-actual", "Operation/output-attributes"},
	{"Job Status/overrides-actual", "Job Template/overrides"},
	{"Job Status/print-accuracy-actual", "Document Template/print-accuracy"},
	{"Job Status/print-objects-actual", "Document Template/print-objects"},
	{"Job Status/separator-sheets-actual", "Job Template/separator-sheets"},
	{"Job Template/job-cover-back", "Job Template/cover-back"},
	{"Job Template/job-cover-front", "Job Template/cover-front"},
	{"Job Template/overrides", "Job Template"},
	{"Job Template/cover-back/media-col", "Job Template/media-col"},
	{"Job Template/cover-front/media-col", "Job Template/media-col"},
	{"Job Template/destination-uris/destination-attributes", "Document Template"},
	{"Job Template/destination-uris/destination-attributes", "Job Template"},
	{"Job Template/destination-uris/destination-attributes", "Operation"},
	{"Job Template/finishings-col/media-size", "Job Template/media-col/media-size"},
	{"Job Template/insert-sheet/media-col", "Job Template/media-col"},
	{"Job Template/job-accounting-sheets/media-col", "Job Template/media-col"},
	{"Job Template/job-error-sheet/media-col", "Job Template/media-col"},
	{"Job Template/job-sheets-col/media-col", "Job Template/media-col"},
	{"Job Template/proof-print/media-col", "Job Template/media-col"},
	{"Job Template/separator-sheets/media-col", "Job Template/media-col"},
	{"Operation/job-impressions-col", "Job Description/job-impressions-col"},
	{"Operation/job-media-sheets-col", "Job Description/job-media-sheets-col"},
	{"Operation/job-pages-col", "Job Status/job-pages-col"},
	{"Operation/preferred-attributes", "Document Template"},
	{"Operation/preferred-attributes", "Job Template"},
	{"Operation/preferred-attributes", "Operation"},
	{"Printer Description/cover-back-default", "Job Template/cover-back"},
	{"Printer Description/cover-front-default", "Job Template/cover-front"},
	{"Printer Description/cover-sheet-info-default", "Job Template/cover-sheet-info"},
	{"Printer Description/document-format-details-default", "Operation/document-format-details"},
	{"Printer Description/finishings-col-database", "Job Template/finishings-col"},
	{"Printer Description/finishings-col-default", "Job Template/finishings-col"},
	{"Printer Description/finishings-col-ready", "Job Template/finishings-col"},
	{"Printer Description/input-attributes-default", "Operation/input-attributes"},
	{"Printer Description/insert-sheet-default", "Job Template/insert-sheet"},
	{"Printer Description/job-accounting-sheets-default", "Job Template/job-accounting-sheets"},
	{"Printer Description/job-constraints-supported", "Job Template"},
	{"Printer Description/job-cover-back-default", "Job Template/cover-back"},
	{"Printer Description/job-cover-front-default", "Job Template/cover-front"},
	{"Printer Description/job-error-sheet-default", "Job Template/job-error-sheet"},
	{"Printer Description/job-resolvers-supported", "Job Template"},
	{"Printer Description/job-sheets-col-default", "Job Template/job-sheets-col"},
	{"Printer Description/materials-col-database", "Document Template/materials-col"},
	{"Printer Description/materials-col-default", "Document Template/materials-col"},
	{"Printer Description/materials-col-ready", "Document Template/materials-col"},
	{"Printer Description/media-col-database", "Job Template/media-col"},
	{"Printer Description/media-col-default", "Job Template/media-col"},
	{"Printer Description/media-col-ready", "Job Template/media-col"},
	{"Printer Description/media-overprint-default", "Document Template/media-overprint"},
	{"Printer Description/output-attributes-default", "Operation/output-attributes"},
	{"Printer Description/proof-print-default", "Job Template/proof-print"},
	{"Printer Description/separator-sheets-default", "Job Template/separator-sheets"},
	{"Printer Description/destination-uri-ready/destination-attributes", "Document Template"},
	{"Printer Description/destination-uri-ready/destination-attributes", "Job Template"},
	{"Printer Description/destination-uri-ready/destination-attributes", "Operation"},
	{"Printer Status/printer-impressions-completed-col", "Job Description/job-impressions-col"},
	{"Printer Status/printer-media-sheets-completed-col", "Job Description/job-media-sheets-col"},
	{"Printer Status/printer-pages-completed-col", "Document Status/pages-col"},
	{"System Description/system-contact-col", "Printer Description/printer-contact-col"},
	{"System Description/system-service-contact-col", "Printer Description/printer-contact-col"},
	{"System Description/system-xri-supported", "Printer Description/printer-xri-supported"},
	{"System Status/system-impressions-completed-col", "Job Description/job-media-sheets-col"},
	{"System Status/system-media-sheets-completed-col", "Job Description/job-media-sheets-col"},
	{"System Status/system-pages-completed-col", "Document Status/pages-col"},
	{"System Status/system-configured-printers/printer-xri-supported", "Printer Description/printer-xri-supported"},
}

// exceptions contains member attributes that doesn't exist even if borrowed.
var exceptions = generic.NewSetOf(
	"Job Template/destination-uris/destination-attributes/document-password",
	"Job Template/destination-uris/destination-attributes/job-password",
	"Job Template/destination-uris/destination-attributes/job-password-encryption",
	"Printer Description/destination-uri-ready/destination-attributes/document-password",
	"Printer Description/destination-uri-ready/destination-attributes/job-password",
	"Printer Description/destination-uri-ready/destination-attributes/job-password-encryption",
)
