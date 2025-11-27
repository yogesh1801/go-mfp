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
	"github.com/OpenPrinting/goipp"
)

// DocumentDescription is the Document Description attributes
var DocumentDescription = map[string]*Attribute{
	// Document Description/document-name
	"document-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
}

// DocumentStatus is the Document Status attributes
var DocumentStatus = map[string]*Attribute{
	// Document Status/attributes-charset
	"attributes-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Document Status/attributes-natural-language
	"attributes-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Document Status/chamber-humidity-actual
	"chamber-humidity-actual": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/chamber-temperature-actual
	"chamber-temperature-actual": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/compression
	"compression": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/copies-actual
	"copies-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/cover-back-actual
	"cover-back-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/cover-front-actual
	"cover-front-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/date-time-at-completed
	"date-time-at-completed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Document Status/date-time-at-created
	"date-time-at-created": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Document Status/date-time-at-creation
	"date-time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Document Status/date-time-at-processing
	"date-time-at-processing": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Document Status/detailed-status-messages
	"detailed-status-messages": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-access-errors
	"document-access-errors": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-charset
	"document-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Document Status/document-digital-signature
	"document-digital-signature": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/document-format
	"document-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Document Status/document-format-details
	"document-format-details": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/document-format-details-detected
	"document-format-details-detected": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/document-format-detected
	"document-format-detected": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Document Status/document-format-ready
	"document-format-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Document Status/document-format-supplied
	"document-format-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Document Status/document-format-version
	"document-format-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-format-version-detected
	"document-format-version-detected": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-format-version-supplied
	"document-format-version-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-job-id
	"document-job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/document-job-uri
	"document-job-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/document-message
	"document-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-message-supplied
	"document-message-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-metadata
	"document-metadata": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Document Status/document-name-supplied
	"document-name-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Status/document-natural-language
	"document-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Document Status/document-number
	"document-number": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/document-printer-uri
	"document-printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/document-resource-ids
	"document-resource-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/document-state
	"document-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/document-state-message
	"document-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/document-state-reasons
	"document-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/document-uri
	"document-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/document-uuid
	"document-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/errors-count
	"errors-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/finishings-actual
	"finishings-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/finishings-col-actual
	"finishings-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/force-front-side-actual
	"force-front-side-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/imposition-template-actual
	"imposition-template-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Status/impressions
	"impressions": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/impressions-col
	"impressions-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Status/impressions-col/blank
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/blank-two-sided
			"blank-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/full-color
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/full-color-two-sided
			"full-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/highlight-color
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/highlight-color-two-sided
			"highlight-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/monochrome
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/impressions-col/monochrome-two-sided
			"monochrome-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Document Status/impressions-completed
	"impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/impressions-completed-col
	"impressions-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/input-attributes-actual
	"input-attributes-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/insert-sheet-actual
	"insert-sheet-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/k-octets
	"k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/k-octets-processed
	"k-octets-processed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/last-document
	"last-document": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Document Status/materials-col-actual
	"materials-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/media-actual
	"media-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Status/media-col-actual
	"media-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/media-sheets
	"media-sheets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/media-sheets-col
	"media-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Status/media-sheets-col/blank
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/media-sheets-col/full-color
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/media-sheets-col/highlight-color
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/media-sheets-col/monochrome
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Document Status/media-sheets-completed
	"media-sheets-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/media-sheets-completed-col
	"media-sheets-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/more-info
	"more-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Document Status/multiple-object-handling-actual
	"multiple-object-handling-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/number-up-actual
	"number-up-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/orientation-requested-actual
	"orientation-requested-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/output-attributes-actual
	"output-attributes-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/output-bin-actual
	"output-bin-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Status/output-device-actual
	"output-device-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Status/output-device-assigned
	"output-device-assigned": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Status/output-device-document-state
	"output-device-document-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/output-device-document-state-message
	"output-device-document-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Document Status/output-device-document-state-reasons
	"output-device-document-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/page-delivery-actual
	"page-delivery-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/page-order-received-actual
	"page-order-received-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/page-ranges-actual
	"page-ranges-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Document Status/pages
	"pages": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/pages-col
	"pages-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Status/pages-col/full-color
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Status/pages-col/monochrome
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Document Status/pages-completed
	"pages-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/pages-completed-col
	"pages-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/platform-temperature-actual
	"platform-temperature-actual": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/presentation-direction-number-up-actual
	"presentation-direction-number-up-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/print-accuracy-actual
	"print-accuracy-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/print-base-actual
	"print-base-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/print-color-mode-actual
	"print-color-mode-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/print-content-optimize-actual
	"print-content-optimize-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/print-objects-actual
	"print-objects-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/print-quality-actual
	"print-quality-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Status/print-supports-actual
	"print-supports-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/printer-resolution-actual
	"printer-resolution-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Document Status/printer-up-time
	"printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/separator-sheets-actual
	"separator-sheets-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Status/sheet-completed-copy-number
	"sheet-completed-copy-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/sides-actual
	"sides-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/time-at-completed
	"time-at-completed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/time-at-creation
	"time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/time-at-processing
	"time-at-processing": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/warnings-count
	"warnings-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/x-image-position-actual
	"x-image-position-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/x-image-shift-actual
	"x-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/x-side1-image-shift-actual
	"x-side1-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/x-side2-image-shift-actual
	"x-side2-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/y-image-position-actual
	"y-image-position-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Status/y-image-shift-actual
	"y-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/y-side1-image-shift-actual
	"y-side1-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Status/y-side2-image-shift-actual
	"y-side2-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// DocumentTemplate is the Document Template attributes
var DocumentTemplate = map[string]*Attribute{
	// Document Template/chamber-humidity
	"chamber-humidity": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/chamber-temperature
	"chamber-temperature": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/copies
	"copies": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/cover-back
	"cover-back": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/cover-front
	"cover-front": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/feed-orientation
	"feed-orientation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/finishings
	"finishings": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Template/finishings-col
	"finishings-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/force-front-side
	"force-front-side": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/imposition-template
	"imposition-template": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Template/insert-sheet
	"insert-sheet": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/materials-col
	"materials-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/materials-col/material-amount
			"material-amount": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-amount-units
			"material-amount-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-color
			"material-color": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-diameter
			"material-diameter": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-diameter-tolerance
			"material-diameter-tolerance": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-fill-density
			"material-fill-density": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-key
			"material-key": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-name
			"material-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Document Template/materials-col/material-nozzle-diameter
			"material-nozzle-diameter": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-purpose
			"material-purpose": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-rate
			"material-rate": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-rate-units
			"material-rate-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/materials-col/material-retraction
			"material-retraction": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Document Template/materials-col/material-shell-thickness
			"material-shell-thickness": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/materials-col/material-temperature
			"material-temperature": &Attribute{
				SetOf: false,
				Min:   -273,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
			},
			// Document Template/materials-col/material-type
			"material-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Document Template/media
	"media": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Template/media-col
	"media-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/media-col/media-top-offset
			"media-top-offset": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   -2147483648,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/media-col/media-tracking
			"media-tracking": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Document Template/media-input-tray-check
	"media-input-tray-check": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Template/media-overprint
	"media-overprint": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/media-overprint/media-overprint-distance
			"media-overprint-distance": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/media-overprint/media-overprint-method
			"media-overprint-method": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Document Template/multiple-object-handling
	"multiple-object-handling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/number-up
	"number-up": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/orientation-requested
	"orientation-requested": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Template/output-bin
	"output-bin": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Document Template/output-device
	"output-device": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Document Template/page-delivery
	"page-delivery": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/page-order-received
	"page-order-received": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/page-ranges
	"page-ranges": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Document Template/platform-temperature
	"platform-temperature": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/presentation-direction-number-up
	"presentation-direction-number-up": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-accuracy
	"print-accuracy": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/print-accuracy/accuracy-units
			"accuracy-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Document Template/print-accuracy/x-accuracy
			"x-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/print-accuracy/y-accuracy
			"y-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/print-accuracy/z-accuracy
			"z-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Document Template/print-base
	"print-base": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-color-mode
	"print-color-mode": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-content-optimize
	"print-content-optimize": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-darkness
	"print-darkness": &Attribute{
		SetOf: false,
		Min:   -100,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/print-objects
	"print-objects": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Document Template/print-objects/document-number
			"document-number": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Document Template/print-objects/object-offset
			"object-offset": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Document Template/print-objects/object-offset/x-offset
					"x-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Document Template/print-objects/object-offset/y-offset
					"y-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Document Template/print-objects/object-offset/z-offset
					"z-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Document Template/print-objects/object-size
			"object-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Document Template/print-objects/object-size/x-dimension
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Document Template/print-objects/object-size/y-dimension
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Document Template/print-objects/object-size/z-dimension
					"z-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Document Template/print-objects/object-uuid
			"object-uuid": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Document Template/print-quality
	"print-quality": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Document Template/print-rendering-intent
	"print-rendering-intent": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-scaling
	"print-scaling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/print-speed
	"print-speed": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/print-supports
	"print-supports": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/printer-resolution
	"printer-resolution": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Document Template/separator-sheets
	"separator-sheets": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Document Template/sheet-collate
	"sheet-collate": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/sides
	"sides": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/x-image-position
	"x-image-position": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/x-image-shift
	"x-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/x-side1-image-shift
	"x-side1-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/x-side2-image-shift
	"x-side2-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/y-image-position
	"y-image-position": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Document Template/y-image-shift
	"y-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/y-side1-image-shift
	"y-side1-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Document Template/y-side2-image-shift
	"y-side2-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// EventNotifications is the Event Notifications attributes
var EventNotifications = map[string]*Attribute{
	// Event Notifications/job-id
	"job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Event Notifications/job-impressions-completed
	"job-impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Event Notifications/job-state
	"job-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Event Notifications/job-state-reasons
	"job-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Event Notifications/job-uuid
	"job-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Event Notifications/notify-charset
	"notify-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Event Notifications/notify-natural-language
	"notify-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Event Notifications/notify-printer-uri
	"notify-printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Event Notifications/notify-sequence-number
	"notify-sequence-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Event Notifications/notify-subscribed-event
	"notify-subscribed-event": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Event Notifications/notify-subscription-id
	"notify-subscription-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Event Notifications/notify-subscription-uuid
	"notify-subscription-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Event Notifications/notify-text
	"notify-text": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Event Notifications/notify-user-data
	"notify-user-data": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   63,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Event Notifications/printer-current-time
	"printer-current-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Event Notifications/printer-is-accepting-jobs
	"printer-is-accepting-jobs": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Event Notifications/printer-state
	"printer-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Event Notifications/printer-state-reasons
	"printer-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Event Notifications/printer-up-time
	"printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// JobDescription is the Job Description attributes
var JobDescription = map[string]*Attribute{
	// Job Description/current-page-order
	"current-page-order": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Description/job-charge-info
	"job-charge-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Description/job-collation-type
	"job-collation-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Description/job-impressions-col
	"job-impressions-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Description/job-impressions-col/blank
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/blank-two-sided
			"blank-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/full-color
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/full-color-two-sided
			"full-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/highlight-color
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/highlight-color-two-sided
			"highlight-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/monochrome
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-impressions-col/monochrome-two-sided
			"monochrome-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Description/job-media-sheets-col
	"job-media-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Description/job-media-sheets-col/blank
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/blank-two-sided
			"blank-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/full-color
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/full-color-two-sided
			"full-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/highlight-color
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/highlight-color-two-sided
			"highlight-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/monochrome
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Description/job-media-sheets-col/monochrome-two-sided
			"monochrome-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Description/job-message-from-operator
	"job-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Description/job-message-to-operator-actual
	"job-message-to-operator-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Description/job-name
	"job-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Description/job-save-printer-make-and-model
	"job-save-printer-make-and-model": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
}

// JobStatus is the Job Status attributes
var JobStatus = map[string]*Attribute{
	// Job Status/attributes-charset
	"attributes-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Job Status/attributes-natural-language
	"attributes-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Job Status/chamber-humidity-actual
	"chamber-humidity-actual": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/chamber-temperature-actual
	"chamber-temperature-actual": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/client-info
	"client-info": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/compression-supplied
	"compression-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/copies-actual
	"copies-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/cover-back-actual
	"cover-back-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/cover-front-actual
	"cover-front-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/current-page-order
	"current-page-order": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/date-time-at-completed
	"date-time-at-completed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagNoValue},
	},
	// Job Status/date-time-at-completed-estimated
	"date-time-at-completed-estimated": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Job Status/date-time-at-creation
	"date-time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Status/date-time-at-processing
	"date-time-at-processing": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagNoValue},
	},
	// Job Status/date-time-at-processing-estimated
	"date-time-at-processing-estimated": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Job Status/destination-statuses
	"destination-statuses": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Status/destination-statuses/destination-uri
			"destination-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Job Status/destination-statuses/images-completed
			"images-completed": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/destination-statuses/transmission-status
			"transmission-status": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
		}},
	},
	// Job Status/document-charset-supplied
	"document-charset-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Job Status/document-digital-signature-supplied
	"document-digital-signature-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/document-format-details-detected
	"document-format-details-detected": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/document-format-details-supplied
	"document-format-details-supplied": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/document-format-ready
	"document-format-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Job Status/document-format-supplied
	"document-format-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Job Status/document-format-version-supplied
	"document-format-version-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/document-message-supplied
	"document-message-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/document-metadata
	"document-metadata": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Job Status/document-name-supplied
	"document-name-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/document-natural-language-supplied
	"document-natural-language-supplied": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Job Status/errors-count
	"errors-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/finishings-actual
	"finishings-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/finishings-col-actual
	"finishings-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/force-front-side-actual
	"force-front-side-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/imposition-template-actual
	"imposition-template-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/impressions-completed-current-copy
	"impressions-completed-current-copy": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/input-attributes-actual
	"input-attributes-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/insert-sheet-actual
	"insert-sheet-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/ipp-attribute-fidelity
	"ipp-attribute-fidelity": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Job Status/job-account-id-actual
	"job-account-id-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/job-account-type-actual
	"job-account-type-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/job-accounting-sheets-actual
	"job-accounting-sheets-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-accounting-user-id-actual
	"job-accounting-user-id-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/job-copies-actual
	"job-copies-actual": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-cover-back-actual
	"job-cover-back-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-cover-front-actual
	"job-cover-front-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-detailed-status-messages
	"job-detailed-status-messages": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/job-document-access-errors
	"job-document-access-errors": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/job-error-sheet-actual
	"job-error-sheet-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-finishings-actual
	"job-finishings-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/job-hold-until-actual
	"job-hold-until-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/job-id
	"job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-impressions
	"job-impressions": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-impressions-col
	"job-impressions-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Status/job-impressions-col/blank
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/blank-two-sided
			"blank-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/full-color
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/full-color-two-sided
			"full-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/highlight-color
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/highlight-color-two-sided
			"highlight-color-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/monochrome
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-impressions-col/monochrome-two-sided
			"monochrome-two-sided": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Status/job-impressions-completed
	"job-impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-impressions-completed-col
	"job-impressions-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-k-octets
	"job-k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-k-octets-processed
	"job-k-octets-processed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-mandatory-attributes
	"job-mandatory-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/job-media-sheets
	"job-media-sheets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-media-sheets-col
	"job-media-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Status/job-media-sheets-col/blank
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-media-sheets-col/full-color
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-media-sheets-col/highlight-color
			"highlight-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-media-sheets-col/monochrome
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Status/job-media-sheets-completed
	"job-media-sheets-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-media-sheets-completed-col
	"job-media-sheets-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-more-info
	"job-more-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/job-originating-user-name
	"job-originating-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/job-originating-user-uri
	"job-originating-user-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/job-pages
	"job-pages": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-pages-col
	"job-pages-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Status/job-pages-col/blank
			"blank": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-pages-col/full-color
			"full-color": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Status/job-pages-col/monochrome
			"monochrome": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Status/job-pages-completed
	"job-pages-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-pages-completed-col
	"job-pages-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-printer-up-time
	"job-printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-printer-uri
	"job-printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/job-priority-actual
	"job-priority-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-processing-time
	"job-processing-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-release-action
	"job-release-action": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/job-resource-ids
	"job-resource-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/job-sheet-message-actual
	"job-sheet-message-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/job-sheets-actual
	"job-sheets-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/job-sheets-col-actual
	"job-sheets-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-state
	"job-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum, goipp.TagUnknown},
	},
	// Job Status/job-state-message
	"job-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/job-state-reasons
	"job-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/job-storage
	"job-storage": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/job-uri
	"job-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/job-uuid
	"job-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/materials-col-actual
	"materials-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/media-actual
	"media-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/media-col-actual
	"media-col-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/media-input-tray-check-actual
	"media-input-tray-check-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/multiple-document-handling-actual
	"multiple-document-handling-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/multiple-object-handling-actual
	"multiple-object-handling-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/number-of-documents
	"number-of-documents": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/number-of-intervening-jobs
	"number-of-intervening-jobs": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/number-up-actual
	"number-up-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/orientation-requested-actual
	"orientation-requested-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/original-requesting-user-name
	"original-requesting-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/output-attributes-actual
	"output-attributes-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/output-bin-actual
	"output-bin-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Status/output-device-actual
	"output-device-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/output-device-assigned
	"output-device-assigned": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Status/output-device-job-state
	"output-device-job-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/output-device-job-state-message
	"output-device-job-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Status/output-device-job-state-reasons
	"output-device-job-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/output-device-uuid-assigned
	"output-device-uuid-assigned": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/overrides-actual
	"overrides-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/page-delivery-actual
	"page-delivery-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/page-order-received-actual
	"page-order-received-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/page-ranges-actual
	"page-ranges-actual": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Job Status/parent-job-id
	"parent-job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/parent-job-uuid
	"parent-job-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   45,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Status/platform-temperature-actual
	"platform-temperature-actual": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/presentation-direction-number-up-actual
	"presentation-direction-number-up-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/print-accuracy-actual
	"print-accuracy-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/print-base-actual
	"print-base-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/print-color-mode-actual
	"print-color-mode-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/print-content-optimize-actual
	"print-content-optimize-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/print-objects-actual
	"print-objects-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/print-quality-actual
	"print-quality-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Status/print-supports-actual
	"print-supports-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/printer-resolution-actual
	"printer-resolution-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Job Status/separator-sheets-actual
	"separator-sheets-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Status/sheet-collate-actual
	"sheet-collate-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/sheet-completed-copy-number
	"sheet-completed-copy-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/sheet-completed-document-number
	"sheet-completed-document-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/sides-actual
	"sides-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/time-at-completed
	"time-at-completed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Job Status/time-at-completed-estimated
	"time-at-completed-estimated": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagUnknown},
	},
	// Job Status/time-at-creation
	"time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/time-at-processing
	"time-at-processing": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Job Status/time-at-processing-estimated
	"time-at-processing-estimated": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagUnknown},
	},
	// Job Status/warnings-count
	"warnings-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/x-image-position-actual
	"x-image-position-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/x-image-shift-actual
	"x-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/x-side1-image-shift-actual
	"x-side1-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/x-side2-image-shift-actual
	"x-side2-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/y-image-position-actual
	"y-image-position-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Status/y-image-shift-actual
	"y-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/y-side1-image-shift-actual
	"y-side1-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Status/y-side2-image-shift-actual
	"y-side2-image-shift-actual": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// JobTemplate is the Job Template attributes
var JobTemplate = map[string]*Attribute{
	// Job Template/chamber-humidity
	"chamber-humidity": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/chamber-temperature
	"chamber-temperature": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/confirmation-sheet-print
	"confirmation-sheet-print": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Job Template/copies
	"copies": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/cover-back
	"cover-back": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/cover-back/cover-type
			"cover-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/cover-back/media
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/cover-back/media-col
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/cover-front
	"cover-front": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/cover-front/cover-type
			"cover-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/cover-front/media
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/cover-front/media-col
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/cover-sheet-info
	"cover-sheet-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/cover-sheet-info/from-name
			"from-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/cover-sheet-info/logo
			"logo": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Job Template/cover-sheet-info/message
			"message": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/cover-sheet-info/organization-name
			"organization-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/cover-sheet-info/subject
			"subject": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/cover-sheet-info/to-name
			"to-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// Job Template/destination-uris
	"destination-uris": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/destination-uris/destination-attributes
			"destination-attributes": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Job Template/destination-uris/destination-uri
			"destination-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Job Template/destination-uris/post-dial-string
			"post-dial-string": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/destination-uris/pre-dial-string
			"pre-dial-string": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/destination-uris/t33-subaddress
			"t33-subaddress": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Template/feed-orientation
	"feed-orientation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/finishings
	"finishings": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/finishings-col
	"finishings-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
		Members: []map[string]*Attribute{{
			// Job Template/finishings-col/baling
			"baling": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/baling/baling-type
					"baling-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
					// Job Template/finishings-col/baling/baling-when
					"baling-when": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
				}},
			},
			// Job Template/finishings-col/binding
			"binding": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/binding/binding-reference-edge
					"binding-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/binding/binding-type
					"binding-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
				}},
			},
			// Job Template/finishings-col/coating
			"coating": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/coating/coating-sides
					"coating-sides": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/coating/coating-type
					"coating-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
				}},
			},
			// Job Template/finishings-col/covering
			"covering": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/covering/covering-name
					"covering-name": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
				}},
			},
			// Job Template/finishings-col/finishing-template
			"finishing-template": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/finishings-col/folding
			"folding": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/folding/folding-direction
					"folding-direction": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/folding/folding-offset
					"folding-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/folding/folding-reference-edge
					"folding-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
				}},
			},
			// Job Template/finishings-col/imposition-template
			"imposition-template": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/finishings-col/laminating
			"laminating": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/laminating/laminating-sides
					"laminating-sides": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/laminating/laminating-type
					"laminating-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
				}},
			},
			// Job Template/finishings-col/media-sheets-supported
			"media-sheets-supported": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Job Template/finishings-col/media-size
			"media-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Job Template/finishings-col/media-size-name
			"media-size-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/finishings-col/punching
			"punching": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/punching/punching-locations
					"punching-locations": &Attribute{
						SetOf: true,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/punching/punching-offset
					"punching-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/punching/punching-reference-edge
					"punching-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
				}},
			},
			// Job Template/finishings-col/stitching
			"stitching": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/stitching/stitching-angle
					"stitching-angle": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   359,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/stitching/stitching-locations
					"stitching-locations": &Attribute{
						SetOf: true,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/stitching/stitching-method
					"stitching-method": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/stitching/stitching-offset
					"stitching-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/stitching/stitching-reference-edge
					"stitching-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
				}},
			},
			// Job Template/finishings-col/trimming
			"trimming": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/finishings-col/trimming/trimming-offset
					"trimming-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/finishings-col/trimming/trimming-reference-edge
					"trimming-reference-edge": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Job Template/finishings-col/trimming/trimming-type
					"trimming-type": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
					},
					// Job Template/finishings-col/trimming/trimming-when
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
	// Job Template/force-front-side
	"force-front-side": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/image-orientation
	"image-orientation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/imposition-template
	"imposition-template": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/insert-sheet
	"insert-sheet": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/insert-sheet/insert-after-page-number
			"insert-after-page-number": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/insert-sheet/insert-count
			"insert-count": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/insert-sheet/media
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/insert-sheet/media-col
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/job-account-id
	"job-account-id": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Template/job-account-type
	"job-account-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-accounting-sheets
	"job-accounting-sheets": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/job-accounting-sheets/job-accounting-sheets-type
			"job-accounting-sheets-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-accounting-sheets/media
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-accounting-sheets/media-col
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/job-accounting-user-id
	"job-accounting-user-id": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Template/job-cancel-after
	"job-cancel-after": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-complete-before
	"job-complete-before": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-complete-before-time
	"job-complete-before-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Template/job-copies
	"job-copies": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-cover-back
	"job-cover-back": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Template/job-cover-front
	"job-cover-front": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Job Template/job-delay-output-until
	"job-delay-output-until": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-delay-output-until-time
	"job-delay-output-until-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Template/job-error-action
	"job-error-action": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/job-error-sheet
	"job-error-sheet": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/job-error-sheet/job-error-sheet-type
			"job-error-sheet-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-error-sheet/job-error-sheet-when
			"job-error-sheet-when": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/job-error-sheet/media
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-error-sheet/media-col
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/job-finishings
	"job-finishings": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/job-hold-until
	"job-hold-until": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-hold-until-time
	"job-hold-until-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Template/job-message-to-operator
	"job-message-to-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Template/job-pages-per-set
	"job-pages-per-set": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-phone-number
	"job-phone-number": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Job Template/job-priority
	"job-priority": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-recipient-name
	"job-recipient-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Template/job-retain-until
	"job-retain-until": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-retain-until-interval
	"job-retain-until-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/job-retain-until-time
	"job-retain-until-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Job Template/job-sheet-message
	"job-sheet-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Job Template/job-sheets
	"job-sheets": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/job-sheets-col
	"job-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/job-sheets-col/job-sheets
			"job-sheets": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-sheets-col/media
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/job-sheets-col/media-col
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// Job Template/materials-col
	"materials-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/materials-col/material-amount
			"material-amount": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-amount-units
			"material-amount-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-color
			"material-color": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-diameter
			"material-diameter": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-diameter-tolerance
			"material-diameter-tolerance": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-fill-density
			"material-fill-density": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-key
			"material-key": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-name
			"material-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Job Template/materials-col/material-nozzle-diameter
			"material-nozzle-diameter": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-purpose
			"material-purpose": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-rate
			"material-rate": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-rate-units
			"material-rate-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/materials-col/material-retraction
			"material-retraction": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Job Template/materials-col/material-shell-thickness
			"material-shell-thickness": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/materials-col/material-temperature
			"material-temperature": &Attribute{
				SetOf: false,
				Min:   -273,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
			},
			// Job Template/materials-col/material-type
			"material-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Job Template/media
	"media": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/media-col
	"media-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/media-col/media-back-coating
			"media-back-coating": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-bottom-margin
			"media-bottom-margin": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-color
			"media-color": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-front-coating
			"media-front-coating": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-grain
			"media-grain": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-hole-count
			"media-hole-count": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-info
			"media-info": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Job Template/media-col/media-key
			"media-key": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-left-margin
			"media-left-margin": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-order-count
			"media-order-count": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-pre-printed
			"media-pre-printed": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-recycled
			"media-recycled": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-right-margin
			"media-right-margin": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-size
			"media-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/media-col/media-size/x-dimension
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/media-col/media-size/y-dimension
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Job Template/media-col/media-size-name
			"media-size-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-source
			"media-source": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-thickness
			"media-thickness": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-tooth
			"media-tooth": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-top-margin
			"media-top-margin": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-top-offset
			"media-top-offset": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   -2147483648,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-col/media-tracking
			"media-tracking": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/media-col/media-type
			"media-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/media-col/media-weight-metric
			"media-weight-metric": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Template/media-input-tray-check
	"media-input-tray-check": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/media-overprint
	"media-overprint": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/media-overprint/media-overprint-distance
			"media-overprint-distance": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/media-overprint/media-overprint-method
			"media-overprint-method": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Job Template/multiple-document-handling
	"multiple-document-handling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/multiple-object-handling
	"multiple-object-handling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/number-of-retries
	"number-of-retries": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/number-up
	"number-up": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/orientation-requested
	"orientation-requested": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/output-bin
	"output-bin": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Job Template/output-device
	"output-device": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Job Template/overrides
	"overrides": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/overrides/document-copies
			"document-copies": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Job Template/overrides/document-numbers
			"document-numbers": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Job Template/overrides/pages
			"pages": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
		}},
	},
	// Job Template/page-delivery
	"page-delivery": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/page-order-received
	"page-order-received": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/page-ranges
	"page-ranges": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Job Template/pages-per-subset
	"pages-per-subset": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/pclm-source-resolution
	"pclm-source-resolution": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Job Template/platform-temperature
	"platform-temperature": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/presentation-direction-number-up
	"presentation-direction-number-up": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-accuracy
	"print-accuracy": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/print-accuracy/accuracy-units
			"accuracy-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Job Template/print-accuracy/x-accuracy
			"x-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/print-accuracy/y-accuracy
			"y-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/print-accuracy/z-accuracy
			"z-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Template/print-base
	"print-base": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-color-mode
	"print-color-mode": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-content-optimize
	"print-content-optimize": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-darkness
	"print-darkness": &Attribute{
		SetOf: false,
		Min:   -100,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/print-objects
	"print-objects": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/print-objects/document-number
			"document-number": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Job Template/print-objects/object-offset
			"object-offset": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/print-objects/object-offset/x-offset
					"x-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/print-objects/object-offset/y-offset
					"y-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/print-objects/object-offset/z-offset
					"z-offset": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Job Template/print-objects/object-size
			"object-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Job Template/print-objects/object-size/x-dimension
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/print-objects/object-size/y-dimension
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Job Template/print-objects/object-size/z-dimension
					"z-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Job Template/print-objects/object-uuid
			"object-uuid": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Job Template/print-quality
	"print-quality": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Job Template/print-rendering-intent
	"print-rendering-intent": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-scaling
	"print-scaling": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/print-speed
	"print-speed": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/print-supports
	"print-supports": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/printer-resolution
	"printer-resolution": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Job Template/proof-copies
	"proof-copies": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/proof-print
	"proof-print": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/proof-print/media
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/proof-print/media-col
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Job Template/proof-print/proof-print-copies
			"proof-print-copies": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Job Template/retry-interval
	"retry-interval": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/retry-time-out
	"retry-time-out": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/separator-sheets
	"separator-sheets": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Job Template/separator-sheets/media
			"media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Job Template/separator-sheets/media-col
			"media-col": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Job Template/separator-sheets/separator-sheets-type
			"separator-sheets-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Job Template/sheet-collate
	"sheet-collate": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/sides
	"sides": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/x-image-position
	"x-image-position": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/x-image-shift
	"x-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/x-side1-image-shift
	"x-side1-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/x-side2-image-shift
	"x-side2-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/y-image-position
	"y-image-position": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Job Template/y-image-shift
	"y-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/y-side1-image-shift
	"y-side1-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Job Template/y-side2-image-shift
	"y-side2-image-shift": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
}

// Operation is the Operation attributes
var Operation = map[string]*Attribute{
	// Operation/attributes-charset
	"attributes-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Operation/attributes-natural-language
	"attributes-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Operation/charge-info-message
	"charge-info-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/client-info
	"client-info": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/client-info/client-name
			"client-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Operation/client-info/client-patches
			"client-patches": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
			},
			// Operation/client-info/client-string-version
			"client-string-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/client-info/client-type
			"client-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
			// Operation/client-info/client-version
			"client-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   64,
				Tags:  []goipp.Tag{goipp.TagString, goipp.TagNoValue},
			},
		}},
	},
	// Operation/compression
	"compression": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/compression-accepted
	"compression-accepted": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/destination-accesses
	"destination-accesses": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
		Members: []map[string]*Attribute{{
			// Operation/destination-accesses/access-oauth-token
			"access-oauth-token": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
			// Operation/destination-accesses/access-oauth-uri
			"access-oauth-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Operation/destination-accesses/access-password
			"access-password": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/destination-accesses/access-pin
			"access-pin": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/destination-accesses/access-user-name
			"access-user-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/destination-accesses/access-x509-certificate
			"access-x509-certificate": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
		}},
	},
	// Operation/detailed-status-message
	"detailed-status-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/document-access
	"document-access": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
		Members: []map[string]*Attribute{{
			// Operation/document-access/access-oauth-token
			"access-oauth-token": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
			// Operation/document-access/access-oauth-uri
			"access-oauth-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Operation/document-access/access-password
			"access-password": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-access/access-pin
			"access-pin": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-access/access-user-name
			"access-user-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-access/access-x509-certificate
			"access-x509-certificate": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
		}},
	},
	// Operation/document-access-error
	"document-access-error": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/document-charset
	"document-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Operation/document-data-get-interval
	"document-data-get-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/document-data-wait
	"document-data-wait": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/document-digital-signature
	"document-digital-signature": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/document-format
	"document-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/document-format-accepted
	"document-format-accepted": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/document-format-details
	"document-format-details": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/document-format-details/document-format
			"document-format": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagMimeType},
			},
			// Operation/document-format-details/document-format-device-id
			"document-format-device-id": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-format-details/document-format-version
			"document-format-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-format-details/document-natural-language
			"document-natural-language": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagLanguage},
			},
			// Operation/document-format-details/document-source-application-name
			"document-source-application-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Operation/document-format-details/document-source-application-version
			"document-source-application-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Operation/document-format-details/document-source-os-name
			"document-source-os-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   40,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Operation/document-format-details/document-source-os-version
			"document-source-os-version": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   40,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// Operation/document-format-version
	"document-format-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/document-message
	"document-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/document-metadata
	"document-metadata": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Operation/document-name
	"document-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/document-natural-language
	"document-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Operation/document-number
	"document-number": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/document-password
	"document-password": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Operation/document-preprocessed
	"document-preprocessed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/document-uri
	"document-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/encrypted-job-request-format
	"encrypted-job-request-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/encrypted-job-request-id
	"encrypted-job-request-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/fetch-status-code
	"fetch-status-code": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/fetch-status-message
	"fetch-status-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/first-index
	"first-index": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/identify-actions
	"identify-actions": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/input-attributes
	"input-attributes": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/input-attributes/input-auto-scaling
			"input-auto-scaling": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Operation/input-attributes/input-auto-skew-correction
			"input-auto-skew-correction": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Operation/input-attributes/input-brightness
			"input-brightness": &Attribute{
				SetOf: false,
				Min:   -100,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-color-mode
			"input-color-mode": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-content-type
			"input-content-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-contrast
			"input-contrast": &Attribute{
				SetOf: false,
				Min:   -100,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-film-scan-mode
			"input-film-scan-mode": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-images-to-transfer
			"input-images-to-transfer": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-media
			"input-media": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
			// Operation/input-attributes/input-orientation-requested
			"input-orientation-requested": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-quality
			"input-quality": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
			// Operation/input-attributes/input-resolution
			"input-resolution": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagResolution},
			},
			// Operation/input-attributes/input-scaling-height
			"input-scaling-height": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   1000,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-scaling-width
			"input-scaling-width": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   1000,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-scan-regions
			"input-scan-regions": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Operation/input-attributes/input-scan-regions/x-dimension
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Operation/input-attributes/input-scan-regions/x-origin
					"x-origin": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Operation/input-attributes/input-scan-regions/y-dimension
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Operation/input-attributes/input-scan-regions/y-origin
					"y-origin": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Operation/input-attributes/input-sharpness
			"input-sharpness": &Attribute{
				SetOf: false,
				Min:   -100,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/input-attributes/input-sides
			"input-sides": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/input-attributes/input-source
			"input-source": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Operation/ipp-attribute-fidelity
	"ipp-attribute-fidelity": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/job-authorization-uri
	"job-authorization-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/job-hold-until
	"job-hold-until": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Operation/job-hold-until-time
	"job-hold-until-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Operation/job-id
	"job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-ids
	"job-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-impressions
	"job-impressions": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-impressions-col
	"job-impressions-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Operation/job-impressions-estimated
	"job-impressions-estimated": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-k-octets
	"job-k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-mandatory-attributes
	"job-mandatory-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/job-media-sheets
	"job-media-sheets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-media-sheets-col
	"job-media-sheets-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Operation/job-message-from-operator
	"job-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/job-name
	"job-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/job-pages
	"job-pages": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/job-pages-col
	"job-pages-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Operation/job-password
	"job-password": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Operation/job-password-encryption
	"job-password-encryption": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/job-release-action
	"job-release-action": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/job-state
	"job-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/job-state-message
	"job-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/job-state-reasons
	"job-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/job-storage
	"job-storage": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/job-storage/job-storage-access
			"job-storage-access": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/job-storage/job-storage-disposition
			"job-storage-disposition": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/job-storage/job-storage-group
			"job-storage-group": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
		}},
	},
	// Operation/job-uri
	"job-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/last-document
	"last-document": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/limit
	"limit": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/message
	"message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/my-jobs
	"my-jobs": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/notify-get-interval
	"notify-get-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-printer-ids
	"notify-printer-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-resource-id
	"notify-resource-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-sequence-numbers
	"notify-sequence-numbers": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-subscription-ids
	"notify-subscription-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/notify-wait
	"notify-wait": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Operation/original-requesting-user-name
	"original-requesting-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/output-attributes
	"output-attributes": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/output-attributes/noise-removal
			"noise-removal": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Operation/output-attributes/output-compression-quality-factor
			"output-compression-quality-factor": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   100,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Operation/output-device-job-states
	"output-device-job-states": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/output-device-uuid
	"output-device-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/output-device-x509-certificate
	"output-device-x509-certificate": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/output-device-x509-request
	"output-device-x509-request": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/preferred-attributes
	"preferred-attributes": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Operation/printer-geo-location
	"printer-geo-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/printer-id
	"printer-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/printer-ids
	"printer-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/printer-location
	"printer-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/printer-message-from-operator
	"printer-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/printer-service-type
	"printer-service-type": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/printer-up-time
	"printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/printer-uri
	"printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/printer-xri-requested
	"printer-xri-requested": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Operation/printer-xri-requested/xri-authentication
			"xri-authentication": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Operation/printer-xri-requested/xri-security
			"xri-security": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// Operation/profile-uri-actual
	"profile-uri-actual": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/requested-attributes
	"requested-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/requesting-user-name
	"requesting-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Operation/requesting-user-pkcs7-public-key
	"requesting-user-pkcs7-public-key": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/requesting-user-uri
	"requesting-user-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/resource-format
	"resource-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/resource-format-accepted
	"resource-format-accepted": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/resource-formats
	"resource-formats": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Operation/resource-id
	"resource-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/resource-ids
	"resource-ids": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/resource-k-octets
	"resource-k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/resource-natural-language
	"resource-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Operation/resource-patches
	"resource-patches": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
	},
	// Operation/resource-signature
	"resource-signature": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Operation/resource-states
	"resource-states": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Operation/resource-string-version
	"resource-string-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
	},
	// Operation/resource-type
	"resource-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/resource-types
	"resource-types": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/resource-version
	"resource-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString, goipp.TagNoValue},
	},
	// Operation/restart-get-interval
	"restart-get-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Operation/status-message
	"status-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Operation/system-uri
	"system-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Operation/which-jobs
	"which-jobs": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Operation/which-printers
	"which-printers": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
}

// PrinterDescription is the Printer Description attributes
var PrinterDescription = map[string]*Attribute{
	// Printer Description/accuracy-units-supported
	"accuracy-units-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/baling-type-supported
	"baling-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/baling-when-supported
	"baling-when-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/binding-reference-edge-supported
	"binding-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/binding-type-supported
	"binding-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/chamber-humidity-default
	"chamber-humidity-default": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Printer Description/chamber-humidity-supported
	"chamber-humidity-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/chamber-temperature-default
	"chamber-temperature-default": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Printer Description/chamber-temperature-supported
	"chamber-temperature-supported": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/charset-configured
	"charset-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Printer Description/charset-supported
	"charset-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Printer Description/client-info-supported
	"client-info-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/coating-sides-supported
	"coating-sides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/coating-type-supported
	"coating-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/color-supported
	"color-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/compression-supported
	"compression-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/confirmation-sheet-print-default
	"confirmation-sheet-print-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/copies-default
	"copies-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/copies-supported
	"copies-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/cover-back-default
	"cover-back-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/cover-back-supported
	"cover-back-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/cover-front-default
	"cover-front-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/cover-front-supported
	"cover-front-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/cover-sheet-info-default
	"cover-sheet-info-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/cover-sheet-info-supported
	"cover-sheet-info-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/cover-type-supported
	"cover-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/covering-name-supported
	"covering-name-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/destination-accesses-supported
	"destination-accesses-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/destination-uri-ready
	"destination-uri-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/destination-uri-ready/destination-attributes
			"destination-attributes": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
			// Printer Description/destination-uri-ready/destination-attributes-supported
			"destination-attributes-supported": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/destination-uri-ready/destination-info
			"destination-info": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// Printer Description/destination-uri-ready/destination-is-directory
			"destination-is-directory": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// Printer Description/destination-uri-ready/destination-mandatory-access-attributes
			"destination-mandatory-access-attributes": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/destination-uri-ready/destination-name
			"destination-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Printer Description/destination-uri-ready/destination-oauth-scope
			"destination-oauth-scope": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
			// Printer Description/destination-uri-ready/destination-oauth-token
			"destination-oauth-token": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagString},
			},
			// Printer Description/destination-uri-ready/destination-oauth-uri
			"destination-oauth-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Printer Description/destination-uri-ready/destination-uri
			"destination-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Printer Description/destination-uri-schemes-supported
	"destination-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/destination-uris-supported
	"destination-uris-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-access-supported
	"document-access-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-charset-default
	"document-charset-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Printer Description/document-charset-supported
	"document-charset-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Printer Description/document-creation-attributes-supported
	"document-creation-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-digital-signature-default
	"document-digital-signature-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-digital-signature-supported
	"document-digital-signature-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-format-default
	"document-format-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/document-format-details-default
	"document-format-details-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/document-format-details-supported
	"document-format-details-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-format-supported
	"document-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/document-format-version-default
	"document-format-version-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/document-format-version-supported
	"document-format-version-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/document-natural-language-default
	"document-natural-language-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/document-natural-language-supported
	"document-natural-language-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/document-password-supported
	"document-password-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/document-privacy-attributes
	"document-privacy-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/document-privacy-scope
	"document-privacy-scope": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/feed-orientation-default
	"feed-orientation-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/feed-orientation-supported
	"feed-orientation-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/fetch-document-attributes-supported
	"fetch-document-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/finishing-template-supported
	"finishing-template-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/finishings-col-database
	"finishings-col-database": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/finishings-col-default
	"finishings-col-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/finishings-col-ready
	"finishings-col-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/finishings-col-supported
	"finishings-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/finishings-default
	"finishings-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/finishings-ready
	"finishings-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/finishings-supported
	"finishings-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/folding-direction-supported
	"folding-direction-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/folding-offset-supported
	"folding-offset-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/folding-reference-edge-supported
	"folding-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/force-front-side-default
	"force-front-side-default ": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/force-front-side-supported
	"force-front-side-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/force-front-side-supported
	"force-front-side-supported ": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/from-name-supported
	"from-name-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/generated-natural-language-supported
	"generated-natural-language-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/identify-actions-default
	"identify-actions-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/identify-actions-supported
	"identify-actions-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/image-orientation-default
	"image-orientation-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/image-orientation-supported
	"image-orientation-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/imposition-template-default
	"imposition-template-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/imposition-template-supported
	"imposition-template-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/input-attributes-default
	"input-attributes-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/input-attributes-supported
	"input-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-color-mode-supported
	"input-color-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-content-type-supported
	"input-content-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-film-scan-mode-supported
	"input-film-scan-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-media-supported
	"input-media-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/input-orientation-requested-supported
	"input-orientation-requested-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/input-quality-supported
	"input-quality-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/input-resolution-supported
	"input-resolution-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/input-scan-regions-supported
	"input-scan-regions-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/input-scan-regions-supported/x-dimension
			"x-dimension": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Printer Description/input-scan-regions-supported/x-origin
			"x-origin": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Printer Description/input-scan-regions-supported/y-dimension
			"y-dimension": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
			// Printer Description/input-scan-regions-supported/y-origin
			"y-origin": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagRange},
			},
		}},
	},
	// Printer Description/input-sides-supported
	"input-sides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/input-source-supported
	"input-source-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/insert-after-page-number-supported
	"insert-after-page-number-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/insert-count-supported
	"insert-count-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/insert-sheet-default
	"insert-sheet-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/insert-sheet-supported
	"insert-sheet-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/ipp-features-supported
	"ipp-features-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/ipp-versions-supported
	"ipp-versions-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/ippget-event-life
	"ippget-event-life": &Attribute{
		SetOf: false,
		Min:   15,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-account-id-default
	"job-account-id-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/job-account-id-supported
	"job-account-id-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-account-type-default
	"job-account-type-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-account-type-supported
	"job-account-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-accounting-output-bin-default
	"job-accounting-output-bin-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-accounting-output-bin-supported
	"job-accounting-output-bin-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-accounting-sheets-default
	"job-accounting-sheets-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-accounting-sheets-supported
	"job-accounting-sheets-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-accounting-sheets-type-supported
	"job-accounting-sheets-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-accounting-user-id-default
	"job-accounting-user-id-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/job-accounting-user-id-supported
	"job-accounting-user-id-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-authorization-uri-supported
	"job-authorization-uri-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-cancel-after-default
	"job-cancel-after-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-cancel-after-supported
	"job-cancel-after-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-complete-before-supported
	"job-complete-before-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-constraints-supported
	"job-constraints-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/job-constraints-supported/resolver-name
			"resolver-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
		}},
	},
	// Printer Description/job-copies-supported
	"job-copies-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-cover-back-default
	"job-cover-back-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-cover-back-supported
	"job-cover-back-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-cover-front-default
	"job-cover-front-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-cover-front-supported
	"job-cover-front-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-creation-attributes-supported
	"job-creation-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-delay-output-until-default
	"job-delay-output-until-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-delay-output-until-interval-supported
	"job-delay-output-until-interval-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-delay-output-until-supported
	"job-delay-output-until-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-delay-output-until-time-supported
	"job-delay-output-until-time-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-destination-spooling-supported
	"job-destination-spooling-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-error-action-default
	"job-error-action-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-error-action-supported
	"job-error-action-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-error-sheet-default
	"job-error-sheet-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-error-sheet-supported
	"job-error-sheet-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-error-sheet-type-supported
	"job-error-sheet-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-error-sheet-when-supported
	"job-error-sheet-when-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-finishings-col-supported
	"job-finishings-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-finishings-default
	"job-finishings-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/job-finishings-ready
	"job-finishings-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/job-finishings-supported
	"job-finishings-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/job-history-attributes-configured
	"job-history-attributes-configured": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-history-attributes-supported
	"job-history-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-history-interval-configured
	"job-history-interval-configured": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-history-interval-supported
	"job-history-interval-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-hold-until-default
	"job-hold-until-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-hold-until-supported
	"job-hold-until-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-hold-until-time-supported
	"job-hold-until-time-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-ids-supported
	"job-ids-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-impressions-supported
	"job-impressions-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-k-octets-supported
	"job-k-octets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-mandatory-attributes-supported
	"job-mandatory-attributes-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-media-sheets-supported
	"job-media-sheets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-message-to-operator-default
	"job-message-to-operator-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/job-message-to-operator-supported
	"job-message-to-operator-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-pages-per-set-supported
	"job-pages-per-set-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-password-encryption-supported
	"job-password-encryption-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-password-length-supported
	"job-password-length-supported": &Attribute{
		SetOf: false,
		Min:   4,
		Max:   765,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-password-repertoire-configured
	"job-password-repertoire-configured": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-password-repertoire-supported
	"job-password-repertoire-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-password-supported
	"job-password-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-phone-number-default
	"job-phone-number-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// Printer Description/job-phone-number-scheme-supported
	"job-phone-number-scheme-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/job-phone-number-supported
	"job-phone-number-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-presets-supported
	"job-presets-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/job-presets-supported/preset-category
			"preset-category": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/job-presets-supported/preset-name
			"preset-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Printer Description/job-priority-default
	"job-priority-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-priority-supported
	"job-priority-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/job-privacy-attributes
	"job-privacy-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-privacy-scope
	"job-privacy-scope": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-recipient-name-default
	"job-recipient-name-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/job-recipient-name-supported
	"job-recipient-name-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-release-action-default
	"job-release-action-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-release-action-supported
	"job-release-action-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-resolvers-supported
	"job-resolvers-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/job-resolvers-supported/resolver-name
			"resolver-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
		}},
	},
	// Printer Description/job-retain-until-default
	"job-retain-until-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-retain-until-interval-supported
	"job-retain-until-interval-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-retain-until-supported
	"job-retain-until-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-retain-until-time-supported
	"job-retain-until-time-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/job-sheet-message-default
	"job-sheet-message-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/job-sheet-message-supported
	"job-sheet-message-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/job-sheets-col-default
	"job-sheets-col-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/job-sheets-col-supported
	"job-sheets-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-sheets-default
	"job-sheets-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-sheets-supported
	"job-sheets-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/job-spooling-supported
	"job-spooling-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-storage-access-supported
	"job-storage-access-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-storage-disposition-supported
	"job-storage-disposition-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-storage-group-supported
	"job-storage-group-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/job-storage-supported
	"job-storage-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/job-triggers-supported
	"job-triggers-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/job-triggers-supported/preset-name
			"preset-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
			},
		}},
	},
	// Printer Description/jpeg-features-supported
	"jpeg-features-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/jpeg-k-octets-supported
	"jpeg-k-octets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/jpeg-x-dimension-supported
	"jpeg-x-dimension-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/jpeg-y-dimension-supported
	"jpeg-y-dimension-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/label-mode-configured
	"label-mode-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/label-mode-supported
	"label-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/label-tear-offset-configured
	"label-tear-offset-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/label-tear-offset-supported
	"label-tear-offset-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/laminating-sides-supported
	"laminating-sides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/laminating-type-supported
	"laminating-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/logo-uri-formats-supported
	"logo-uri-formats-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/logo-uri-schemes-supported
	"logo-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/material-amount-units-supported
	"material-amount-units-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/material-diameter-supported
	"material-diameter-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-nozzle-diameter-supported
	"material-nozzle-diameter-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-purpose-supported
	"material-purpose-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/material-rate-supported
	"material-rate-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-rate-units-supported
	"material-rate-units-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/material-shell-thickness-supported
	"material-shell-thickness-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-temperature-supported
	"material-temperature-supported": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/material-type-supported
	"material-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/materials-col-database
	"materials-col-database": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/materials-col-default
	"materials-col-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/materials-col-ready
	"materials-col-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/materials-col-supported
	"materials-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/max-client-info-supported
	"max-client-info-supported": &Attribute{
		SetOf: false,
		Min:   4,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/max-materials-col-supported
	"max-materials-col-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/max-page-ranges-supported
	"max-page-ranges-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/max-save-info-supported
	"max-save-info-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/max-stitching-locations-supported
	"max-stitching-locations-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-back-coating-supported
	"media-back-coating-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-bottom-margin-supported
	"media-bottom-margin-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-col-database
	"media-col-database": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/media-col-database/media-size
			"media-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Printer Description/media-col-database/media-size/x-dimension
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
					},
					// Printer Description/media-col-database/media-size/y-dimension
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
					},
				}},
			},
			// Printer Description/media-col-database/media-source-properties
			"media-source-properties": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Printer Description/media-col-database/media-source-properties/media-source-feed-direction
					"media-source-feed-direction": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Printer Description/media-col-database/media-source-properties/media-source-feed-orientation
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
	// Printer Description/media-col-default
	"media-col-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/media-col-ready
	"media-col-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/media-col-ready/media-size
			"media-size": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Printer Description/media-col-ready/media-size/x-dimension
					"x-dimension": &Attribute{
						SetOf: false,
						Min:   1,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
					// Printer Description/media-col-ready/media-size/y-dimension
					"y-dimension": &Attribute{
						SetOf: false,
						Min:   0,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagInteger},
					},
				}},
			},
			// Printer Description/media-col-ready/media-source-properties
			"media-source-properties": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
				Members: []map[string]*Attribute{{
					// Printer Description/media-col-ready/media-source-properties/media-source-feed-direction
					"media-source-feed-direction": &Attribute{
						SetOf: false,
						Min:   MIN,
						Max:   MAX,
						Tags:  []goipp.Tag{goipp.TagKeyword},
					},
					// Printer Description/media-col-ready/media-source-properties/media-source-feed-orientation
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
	// Printer Description/media-col-supported
	"media-col-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/media-color-supported
	"media-color-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-default
	"media-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/media-front-coating-supported
	"media-front-coating-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-grain-supported
	"media-grain-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-hole-count-supported
	"media-hole-count-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/media-info-supported
	"media-info-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/media-key-supported
	"media-key-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-left-margin-supported
	"media-left-margin-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-order-count-supported
	"media-order-count-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/media-overprint-default
	"media-overprint-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
	},
	// Printer Description/media-overprint-distance-supported
	"media-overprint-distance-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/media-overprint-method-supported
	"media-overprint-method-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/media-overprint-supported
	"media-overprint-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/media-pre-printed-supported
	"media-pre-printed-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-ready
	"media-ready": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-recycled-supported
	"media-recycled-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-right-margin-supported
	"media-right-margin-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-size-supported
	"media-size-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/media-size-supported/x-dimension
			"x-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
			},
			// Printer Description/media-size-supported/y-dimension
			"y-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
			},
		}},
	},
	// Printer Description/media-source-supported
	"media-source-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-supported
	"media-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-thickness-supported
	"media-thickness-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/media-tooth-supported
	"media-tooth-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-top-margin-supported
	"media-top-margin-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/media-top-offset-supported
	"media-top-offset-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   -2147483648,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/media-tracking-supported
	"media-tracking-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/media-type-supported
	"media-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/media-weight-metric-supported
	"media-weight-metric-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/message-supported
	"message-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/multiple-destination-uris-supported
	"multiple-destination-uris-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/multiple-document-handling-default
	"multiple-document-handling-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/multiple-document-handling-supported
	"multiple-document-handling-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/multiple-document-jobs-supported
	"multiple-document-jobs-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/multiple-object-handling-default
	"multiple-object-handling-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/multiple-object-handling-supported
	"multiple-object-handling-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/multiple-operation-time-out
	"multiple-operation-time-out": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/multiple-operation-time-out-action
	"multiple-operation-time-out-action": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/natural-language-configured
	"natural-language-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/notify-attributes-supported
	"notify-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/notify-events-default
	"notify-events-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/notify-events-supported
	"notify-events-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/notify-lease-duration-default
	"notify-lease-duration-default": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/notify-lease-duration-supported
	"notify-lease-duration-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/notify-pull-method-supported
	"notify-pull-method-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/notify-schemes-supported
	"notify-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/number-of-retries-default
	"number-of-retries-default": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/number-of-retries-supported
	"number-of-retries-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/number-up-default
	"number-up-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/number-up-supported
	"number-up-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/oauth-authorization-scope
	"oauth-authorization-scope": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// Printer Description/oauth-authorization-server-uri
	"oauth-authorization-server-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// Printer Description/operations-supported
	"operations-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/organization-name-supported
	"organization-name-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/orientation-requested-default
	"orientation-requested-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum, goipp.TagNoValue},
	},
	// Printer Description/orientation-requested-supported
	"orientation-requested-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/output-attributes-default
	"output-attributes-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
	},
	// Printer Description/output-attributes-supported
	"output-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/output-bin-default
	"output-bin-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/output-bin-supported
	"output-bin-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/output-device-supported
	"output-device-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/output-device-uuid-supported
	"output-device-uuid-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/overrides-supported
	"overrides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-delivery-default
	"page-delivery-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-delivery-supported
	"page-delivery-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-order-received-default
	"page-order-received-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-order-received-supported
	"page-order-received-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/page-ranges-supported
	"page-ranges-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/pages-per-subset-supported
	"pages-per-subset-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/parent-printers-supported
	"parent-printers-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/pclm-raster-back-side
	"pclm-raster-back-side": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pclm-source-resolution-supported
	"pclm-source-resolution-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/pclm-strip-height-preferred
	"pclm-strip-height-preferred": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/pclm-strip-height-supported
	"pclm-strip-height-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/pdf-features-supported
	"pdf-features-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pdf-k-octets-supported
	"pdf-k-octets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/pdf-versions-supported
	"pdf-versions-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pdl-init-file-entry-supported
	"pdl-init-file-entry-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/pdl-init-file-location-supported
	"pdl-init-file-location-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/pdl-init-file-name-subdirectory-supported
	"pdl-init-file-name-subdirectory-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/pdl-init-file-name-supported
	"pdl-init-file-name-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/pdl-init-file-supported
	"pdl-init-file-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pdl-override-guaranteed-supported
	"pdl-override-guaranteed-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pdl-override-supported
	"pdl-override-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pkcs7-document-format-supported
	"pkcs7-document-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/platform-shape
	"platform-shape": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/platform-temperature-default
	"platform-temperature-default": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/platform-temperature-supported
	"platform-temperature-supported": &Attribute{
		SetOf: true,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/preferred-attributes-supported
	"preferred-attributes-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/presentation-direction-number-up-default
	"presentation-direction-number-up-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/presentation-direction-number-up-supported
	"presentation-direction-number-up-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-accuracy-supported
	"print-accuracy-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/print-accuracy-supported/accuracy-units
			"accuracy-units": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/print-accuracy-supported/x-accuracy
			"x-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Printer Description/print-accuracy-supported/y-accuracy
			"y-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Printer Description/print-accuracy-supported/z-accuracy
			"z-accuracy": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Printer Description/print-base-default
	"print-base-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-base-supported
	"print-base-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-color-mode-default
	"print-color-mode-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-color-mode-icc-profiles
	"print-color-mode-icc-profiles": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/print-color-mode-icc-profiles/print-color-mode
			"print-color-mode": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/print-color-mode-icc-profiles/profile-uri
			"profile-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Printer Description/print-color-mode-supported
	"print-color-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-content-optimize-default
	"print-content-optimize-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-content-optimize-supported
	"print-content-optimize-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-darkness-default
	"print-darkness-default": &Attribute{
		SetOf: false,
		Min:   -100,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/print-darkness-supported
	"print-darkness-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/print-objects-supported
	"print-objects-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-processing-attributes-supported
	"print-processing-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-quality-default
	"print-quality-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/print-quality-supported
	"print-quality-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Description/print-rendering-intent-default
	"print-rendering-intent-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-rendering-intent-supported
	"print-rendering-intent-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-scaling-default
	"print-scaling-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-scaling-supported
	"print-scaling-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-speed-default
	"print-speed-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/print-speed-supported
	"print-speed-supported": &Attribute{
		SetOf: true,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/print-supports-default
	"print-supports-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/print-supports-supported
	"print-supports-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-asset-tag
	"printer-asset-tag": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Description/printer-camera-image-uri
	"printer-camera-image-uri": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-charge-info
	"printer-charge-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-charge-info-uri
	"printer-charge-info-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-contact-col
	"printer-contact-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagUnknown},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-contact-col/contact-name
			"contact-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Printer Description/printer-contact-col/contact-uri
			"contact-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Printer Description/printer-contact-col/contact-vcard
			"contact-vcard": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// Printer Description/printer-current-time
	"printer-current-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Printer Description/printer-darkness-configured
	"printer-darkness-configured": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/printer-darkness-supported
	"printer-darkness-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/printer-detailed-status-messages
	"printer-detailed-status-messages": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-device-id
	"printer-device-id": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-dns-sd-name
	"printer-dns-sd-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   63,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/printer-driver-installer
	"printer-driver-installer": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-fax-log-uri
	"printer-fax-log-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-fax-modem-info
	"printer-fax-modem-info": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-fax-modem-name
	"printer-fax-modem-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/printer-fax-modem-number
	"printer-fax-modem-number": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-geo-location
	"printer-geo-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagUnknown},
	},
	// Printer Description/printer-get-attributes-supported
	"printer-get-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-icc-profiles
	"printer-icc-profiles": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-icc-profiles/profile-name
			"profile-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Printer Description/printer-icc-profiles/profile-url
			"profile-url": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Printer Description/printer-icons
	"printer-icons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-info
	"printer-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-kind
	"printer-kind": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/printer-location
	"printer-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-make-and-model
	"printer-make-and-model": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-mandatory-job-attributes
	"printer-mandatory-job-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-mode-configured
	"printer-mode-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-mode-supported
	"printer-mode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-more-info-manufacturer
	"printer-more-info-manufacturer": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-name
	"printer-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/printer-organization
	"printer-organization": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-organizational-unit
	"printer-organizational-unit": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-pkcs7-public-key
	"printer-pkcs7-public-key": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/printer-pkcs7-repertoire-configured
	"printer-pkcs7-repertoire-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-pkcs7-repertoire-supported
	"printer-pkcs7-repertoire-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-privacy-policy-uri
	"printer-privacy-policy-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-requested-job-attributes
	"printer-requested-job-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/printer-resolution-default
	"printer-resolution-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/printer-resolution-supported
	"printer-resolution-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/printer-service-contact-col
	"printer-service-contact-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-service-contact-col/contact-name
			"contact-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// Printer Description/printer-service-contact-col/contact-uri
			"contact-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
			// Printer Description/printer-service-contact-col/contact-vcard
			"contact-vcard": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// Printer Description/printer-static-resource-directory-uri
	"printer-static-resource-directory-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/printer-static-resource-k-octets-supported
	"printer-static-resource-k-octets-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/printer-strings-languages-supported
	"printer-strings-languages-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Printer Description/printer-strings-uri
	"printer-strings-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// Printer Description/printer-volume-supported
	"printer-volume-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-volume-supported/x-dimension
			"x-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Printer Description/printer-volume-supported/y-dimension
			"y-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// Printer Description/printer-volume-supported/z-dimension
			"z-dimension": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// Printer Description/printer-wifi-password
	"printer-wifi-password": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Description/printer-wifi-ssid
	"printer-wifi-ssid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/printer-xri-supported
	"printer-xri-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// Printer Description/printer-xri-supported/xri-authentication
			"xri-authentication": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/printer-xri-supported/xri-security
			"xri-security": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// Printer Description/printer-xri-supported/xri-uri
			"xri-uri": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagURI},
			},
		}},
	},
	// Printer Description/proof-copies-supported
	"proof-copies-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/proof-print-copies-supported
	"proof-print-copies-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/proof-print-default
	"proof-print-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagNoValue},
	},
	// Printer Description/proof-print-supported
	"proof-print-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/punching-hole-diameter-configured
	"punching-hole-diameter-configured": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/punching-locations-supported
	"punching-locations-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/punching-offset-supported
	"punching-offset-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/punching-reference-edge-supported
	"punching-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pwg-raster-document-resolution-supported
	"pwg-raster-document-resolution-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagResolution},
	},
	// Printer Description/pwg-raster-document-sheet-back
	"pwg-raster-document-sheet-back": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pwg-raster-document-type-supported
	"pwg-raster-document-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/pwg-safe-gcode-supported
	"pwg-safe-gcode-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Description/reference-uri-schemes-supported
	"reference-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/repertoire-supported
	"repertoire-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// Printer Description/requesting-user-uri-schemes-supported
	"requesting-user-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// Printer Description/requesting-user-uri-supported
	"requesting-user-uri-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/retry-interval-default
	"retry-interval-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/retry-interval-supported
	"retry-interval-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/retry-time-out-default
	"retry-time-out-default": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/retry-time-out-supported
	"retry-time-out-supported": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/save-disposition-supported
	"save-disposition-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/save-document-format-default
	"save-document-format-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/save-document-format-supported
	"save-document-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/save-location-default
	"save-location-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/save-location-supported
	"save-location-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/save-name-subdirectory-supported
	"save-name-subdirectory-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/save-name-supported
	"save-name-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Description/separator-sheets-default
	"separator-sheets-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Description/separator-sheets-supported
	"separator-sheets-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/sheet-collate-default
	"sheet-collate-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/sheet-collate-supported
	"sheet-collate-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/sides-default
	"sides-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/sides-supported
	"sides-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/smi2699-auth-print-group
	"smi2699-auth-print-group": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/smi2699-auth-proxy-group
	"smi2699-auth-proxy-group": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/smi2699-device-command
	"smi2699-device-command": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/smi2699-device-format
	"smi2699-device-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Printer Description/smi2699-device-name
	"smi2699-device-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Description/smi2699-device-uri
	"smi2699-device-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/stitching-angle-supported
	"stitching-angle-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   359,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/stitching-locations-supported
	"stitching-locations-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/stitching-method-supported
	"stitching-method-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/stitching-offset-supported
	"stitching-offset-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/stitching-reference-edge-supported
	"stitching-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/subject-supported
	"subject-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/subordinate-printers-supported
	"subordinate-printers-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Description/subscription-privacy-attributes
	"subscription-privacy-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/subscription-privacy-scope
	"subscription-privacy-scope": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/to-name-supported
	"to-name-supported": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   1023,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/trimming-offset-supported
	"trimming-offset-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// Printer Description/trimming-reference-edge-supported
	"trimming-reference-edge-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/trimming-type-supported
	"trimming-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/trimming-when-supported
	"trimming-when-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/uri-authentication-supported
	"uri-authentication-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/uri-security-supported
	"uri-security-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/user-defined-values-supported
	"user-defined-values-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/which-jobs-supported
	"which-jobs-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/x-image-position-default
	"x-image-position-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/x-image-position-supported
	"x-image-position-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/x-image-shift-default
	"x-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/x-image-shift-supported
	"x-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/x-side1-image-shift-default
	"x-side1-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/x-side1-image-shift-supported
	"x-side1-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/x-side2-image-shift-default
	"x-side2-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/x-side2-image-shift-supported
	"x-side2-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/y-image-position-default
	"y-image-position-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/y-image-position-supported
	"y-image-position-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Description/y-image-shift-default
	"y-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/y-image-shift-supported
	"y-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/y-side1-image-shift-default
	"y-side1-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/y-side1-image-shift-supported
	"y-side1-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
	// Printer Description/y-side2-image-shift-default
	"y-side2-image-shift-default": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Description/y-side2-image-shift-supported
	"y-side2-image-shift-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagRange},
	},
}

// PrinterStatus is the Printer Status attributes
var PrinterStatus = map[string]*Attribute{
	// Printer Status/chamber-humidity-current
	"chamber-humidity-current": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   100,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagUnknown},
	},
	// Printer Status/chamber-temperature-current
	"chamber-temperature-current": &Attribute{
		SetOf: false,
		Min:   -273,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagUnknown},
	},
	// Printer Status/device-service-count
	"device-service-count": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/device-uuid
	"device-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/document-format-varying-attributes
	"document-format-varying-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/job-settable-attributes-supported
	"job-settable-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/pages-per-minute
	"pages-per-minute": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/pages-per-minute-color
	"pages-per-minute-color": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-alert
	"printer-alert": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-alert-description
	"printer-alert-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-camera-image-uri
	"printer-camera-image-uri": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-config-change-date-time
	"printer-config-change-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagUnknown},
	},
	// Printer Status/printer-config-change-time
	"printer-config-change-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-config-changes
	"printer-config-changes": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-console-display
	"printer-console-display": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-console-light
	"printer-console-light": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-console-light-description
	"printer-console-light-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-cover
	"printer-cover": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-cover-description
	"printer-cover-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-detailed-status-messages
	"printer-detailed-status-messages": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-finisher
	"printer-finisher": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-finisher-description
	"printer-finisher-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-finisher-supplies
	"printer-finisher-supplies": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-finisher-supplies-description
	"printer-finisher-supplies-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-firmware-name
	"printer-firmware-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Printer Status/printer-firmware-patches
	"printer-firmware-patches": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-firmware-string-version
	"printer-firmware-string-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-firmware-version
	"printer-firmware-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-home-page-uri
	"printer-home-page-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-id
	"printer-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-impressions-completed
	"printer-impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-impressions-completed-col
	"printer-impressions-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Status/printer-input-tray
	"printer-input-tray": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-is-accepting-jobs
	"printer-is-accepting-jobs": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// Printer Status/printer-media-sheets-completed
	"printer-media-sheets-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-media-sheets-completed-col
	"printer-media-sheets-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Status/printer-message-date-time
	"printer-message-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Printer Status/printer-message-from-operator
	"printer-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-message-time
	"printer-message-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-more-info
	"printer-more-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-output-tray
	"printer-output-tray": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-pages-completed
	"printer-pages-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-pages-completed-col
	"printer-pages-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// Printer Status/printer-serial-number
	"printer-serial-number": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-service-type
	"printer-service-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/printer-settable-attributes-supported
	"printer-settable-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/printer-state
	"printer-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Status/printer-state-change-date-time
	"printer-state-change-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Printer Status/printer-state-change-time
	"printer-state-change-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-state-message
	"printer-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-state-reasons
	"printer-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/printer-static-resource-k-octets-free
	"printer-static-resource-k-octets-free": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-storage
	"printer-storage": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-storage-description
	"printer-storage-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-supply
	"printer-supply": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Printer Status/printer-supply-description
	"printer-supply-description": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Printer Status/printer-supply-info-uri
	"printer-supply-info-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-up-time
	"printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/printer-uri-supported
	"printer-uri-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-uuid
	"printer-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Printer Status/printer-wifi-state
	"printer-wifi-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Printer Status/queued-job-count
	"queued-job-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Printer Status/xri-authentication-supported
	"xri-authentication-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/xri-security-supported
	"xri-security-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Printer Status/xri-uri-scheme-supported
	"xri-uri-scheme-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
}

// ResourceDescription is the Resource Description attributes
var ResourceDescription = map[string]*Attribute{
	// Resource Description/resource-info
	"resource-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Resource Description/resource-name
	"resource-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
}

// ResourceStatus is the Resource Status attributes
var ResourceStatus = map[string]*Attribute{
	// Resource Status/date-time-at-canceled
	"date-time-at-canceled": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagNoValue},
	},
	// Resource Status/date-time-at-creation
	"date-time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// Resource Status/date-time-at-installed
	"date-time-at-installed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime, goipp.TagNoValue},
	},
	// Resource Status/resource-data-uri
	"resource-data-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// Resource Status/resource-format
	"resource-format": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// Resource Status/resource-id
	"resource-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Resource Status/resource-k-octets
	"resource-k-octets": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Resource Status/resource-natural-language
	"resource-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Resource Status/resource-patches
	"resource-patches": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
	},
	// Resource Status/resource-signature
	"resource-signature": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// Resource Status/resource-state
	"resource-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Resource Status/resource-state-message
	"resource-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// Resource Status/resource-state-reasons
	"resource-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Resource Status/resource-string-version
	"resource-string-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang, goipp.TagNoValue},
	},
	// Resource Status/resource-type
	"resource-type": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Resource Status/resource-use-count
	"resource-use-count": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Resource Status/resource-uuid
	"resource-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   45,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Resource Status/resource-version
	"resource-version": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString, goipp.TagNoValue},
	},
	// Resource Status/time-at-canceled
	"time-at-canceled": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// Resource Status/time-at-creation
	"time-at-creation": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Resource Status/time-at-installed
	"time-at-installed": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
}

// SubscriptionStatus is the Subscription Status attributes
var SubscriptionStatus = map[string]*Attribute{
	// Subscription Status/notify-job-id
	"notify-job-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-lease-expiration-time
	"notify-lease-expiration-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-printer-up-time
	"notify-printer-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-printer-uri
	"notify-printer-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Subscription Status/notify-resource-id
	"notify-resource-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-sequence-number
	"notify-sequence-number": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-status-code
	"notify-status-code": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// Subscription Status/notify-subscriber-user-name
	"notify-subscriber-user-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// Subscription Status/notify-subscriber-user-uri
	"notify-subscriber-user-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Subscription Status/notify-subscription-id
	"notify-subscription-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-subscription-uuid
	"notify-subscription-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Subscription Status/notify-system-up-time
	"notify-system-up-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Status/notify-system-uri
	"notify-system-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
}

// SubscriptionTemplate is the Subscription Template attributes
var SubscriptionTemplate = map[string]*Attribute{
	// Subscription Template/notify-attributes
	"notify-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Subscription Template/notify-charset
	"notify-charset": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// Subscription Template/notify-events
	"notify-events": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Subscription Template/notify-lease-duration
	"notify-lease-duration": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Template/notify-max-events-supported
	"notify-max-events-supported": &Attribute{
		SetOf: false,
		Min:   2,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Template/notify-natural-language
	"notify-natural-language": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// Subscription Template/notify-pull-method
	"notify-pull-method": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// Subscription Template/notify-recipient-uri
	"notify-recipient-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// Subscription Template/notify-time-interval
	"notify-time-interval": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// Subscription Template/notify-user-data
	"notify-user-data": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   63,
		Tags:  []goipp.Tag{goipp.TagString},
	},
}

// SystemDescription is the System Description attributes
var SystemDescription = map[string]*Attribute{
	// System Description/charset-configured
	"charset-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// System Description/charset-supported
	"charset-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagCharset},
	},
	// System Description/document-format-supported
	"document-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// System Description/generated-natural-language-supported
	"generated-natural-language-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// System Description/ipp-features-supported
	"ipp-features-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/ipp-versions-supported
	"ipp-versions-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/ippget-event-life
	"ippget-event-life": &Attribute{
		SetOf: false,
		Min:   15,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Description/multiple-document-printers-supported
	"multiple-document-printers-supported": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBoolean},
	},
	// System Description/natural-language-configured
	"natural-language-configured": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// System Description/notify-attributes-supported
	"notify-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/notify-events-default
	"notify-events-default": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/notify-events-supported
	"notify-events-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/notify-lease-duration-default
	"notify-lease-duration-default": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Description/notify-lease-duration-supported
	"notify-lease-duration-supported": &Attribute{
		SetOf: true,
		Min:   0,
		Max:   67108863,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagRange},
	},
	// System Description/notify-max-events-supported
	"notify-max-events-supported": &Attribute{
		SetOf: false,
		Min:   2,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Description/notify-pull-method-supported
	"notify-pull-method-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/notify-schemes-supported
	"notify-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// System Description/oauth-authorization-scope
	"oauth-authorization-scope": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang, goipp.TagNoValue},
	},
	// System Description/oauth-authorization-server-uri
	"oauth-authorization-server-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// System Description/operations-supported
	"operations-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// System Description/output-device-x509-type-supported
	"output-device-x509-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/power-calendar-policy-col
	"power-calendar-policy-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Description/power-calendar-policy-col/calendar-id
			"calendar-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/day-of-month
			"day-of-month": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   31,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/day-of-week
			"day-of-week": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   7,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/hour
			"hour": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   23,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/minute
			"minute": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   59,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/month
			"month": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   12,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-calendar-policy-col/request-power-state
			"request-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Description/power-calendar-policy-col/run-once
			"run-once": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
		}},
	},
	// System Description/power-event-policy-col
	"power-event-policy-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Description/power-event-policy-col/event-id
			"event-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-event-policy-col/event-name
			"event-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// System Description/power-event-policy-col/request-power-state
			"request-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// System Description/power-timeout-policy-col
	"power-timeout-policy-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Description/power-timeout-policy-col/start-power-state
			"start-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Description/power-timeout-policy-col/timeout-id
			"timeout-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Description/power-timeout-policy-col/timeout-predicate
			"timeout-predicate": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Description/power-timeout-policy-col/timeout-seconds
			"timeout-seconds": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// System Description/printer-creation-attributes-supported
	"printer-creation-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/resource-format-supported
	"resource-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// System Description/resource-settable-attributes-supported
	"resource-settable-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/resource-type-supported
	"resource-type-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/smi2699-auth-group-supported
	"smi2699-auth-group-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Description/smi2699-device-command-supported
	"smi2699-device-command-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Description/smi2699-device-format-supported
	"smi2699-device-format-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagMimeType},
	},
	// System Description/smi2699-device-uri-schemes-supported
	"smi2699-device-uri-schemes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURIScheme},
	},
	// System Description/system-asset-tag
	"system-asset-tag": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// System Description/system-contact-col
	"system-contact-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagUnknown},
	},
	// System Description/system-current-time
	"system-current-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// System Description/system-default-printer-id
	"system-default-printer-id": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   65535,
		Tags:  []goipp.Tag{goipp.TagInteger, goipp.TagNoValue},
	},
	// System Description/system-dns-sd-name
	"system-dns-sd-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   63,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Description/system-geo-location
	"system-geo-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagUnknown},
	},
	// System Description/system-info
	"system-info": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Description/system-location
	"system-location": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Description/system-make-and-model
	"system-make-and-model": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Description/system-mandatory-printer-attributes
	"system-mandatory-printer-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/system-mandatory-registration-attributes
	"system-mandatory-registration-attributes": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/system-message-from-operator
	"system-message-from-operator": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Description/system-name
	"system-name": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   127,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Description/system-service-contact-col
	"system-service-contact-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection, goipp.TagUnknown},
	},
	// System Description/system-settable-attributes-supported
	"system-settable-attributes-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Description/system-strings-languages-supported
	"system-strings-languages-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagLanguage},
	},
	// System Description/system-strings-uri
	"system-strings-uri": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
	},
	// System Description/system-xri-supported
	"system-xri-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
}

// SystemStatus is the System Status attributes
var SystemStatus = map[string]*Attribute{
	// System Status/power-log-col
	"power-log-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-log-col/log-id
			"log-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-log-col/power-state
			"power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/power-log-col/power-state-date-time
			"power-state-date-time": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagDateTime},
			},
			// System Status/power-log-col/power-state-message
			"power-state-message": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
		}},
	},
	// System Status/power-state-capabilities-col
	"power-state-capabilities-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-state-capabilities-col/can-accept-jobs
			"can-accept-jobs": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// System Status/power-state-capabilities-col/can-process-jobs
			"can-process-jobs": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// System Status/power-state-capabilities-col/power-active-watts
			"power-active-watts": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-capabilities-col/power-inactive-watts
			"power-inactive-watts": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-capabilities-col/power-state
			"power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// System Status/power-state-counters-col
	"power-state-counters-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-state-counters-col/hibernate-transitions
			"hibernate-transitions": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-counters-col/on-transitions
			"on-transitions": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-counters-col/standby-transitions
			"standby-transitions": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-counters-col/suspend-transitions
			"suspend-transitions": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// System Status/power-state-monitor-col
	"power-state-monitor-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-state-monitor-col/current-month-kwh
			"current-month-kwh": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-monitor-col/current-watts
			"current-watts": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-monitor-col/lifetime-kwh
			"lifetime-kwh": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/power-state-monitor-col/meters-are-actual
			"meters-are-actual": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// System Status/power-state-monitor-col/power-state
			"power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/power-state-monitor-col/power-state-message
			"power-state-message": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// System Status/power-state-monitor-col/power-usage-is-rms-watts
			"power-usage-is-rms-watts": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
		}},
	},
	// System Status/power-state-transitions-col
	"power-state-transitions-col": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/power-state-transitions-col/end-power-state
			"end-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/power-state-transitions-col/start-power-state
			"start-power-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/power-state-transitions-col/state-transition-seconds
			"state-transition-seconds": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
		}},
	},
	// System Status/system-config-change-date-time
	"system-config-change-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// System Status/system-config-change-time
	"system-config-change-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-config-changes
	"system-config-changes": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-configured-printers
	"system-configured-printers": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/system-configured-printers/printer-id
			"printer-id": &Attribute{
				SetOf: false,
				Min:   0,
				Max:   65535,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/system-configured-printers/printer-info
			"printer-info": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// System Status/system-configured-printers/printer-is-accepting-jobs
			"printer-is-accepting-jobs": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBoolean},
			},
			// System Status/system-configured-printers/printer-name
			"printer-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// System Status/system-configured-printers/printer-service-type
			"printer-service-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/system-configured-printers/printer-state
			"printer-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
			// System Status/system-configured-printers/printer-state-reasons
			"printer-state-reasons": &Attribute{
				SetOf: true,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			// System Status/system-configured-printers/printer-xri-supported
			"printer-xri-supported": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagBeginCollection},
			},
		}},
	},
	// System Status/system-configured-resources
	"system-configured-resources": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
		Members: []map[string]*Attribute{{
			// System Status/system-configured-resources/resource-format
			"resource-format": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagMimeType},
			},
			// System Status/system-configured-resources/resource-id
			"resource-id": &Attribute{
				SetOf: false,
				Min:   1,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			// System Status/system-configured-resources/resource-info
			"resource-info": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagTextLang},
			},
			// System Status/system-configured-resources/resource-name
			"resource-name": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   127,
				Tags:  []goipp.Tag{goipp.TagNameLang},
			},
			// System Status/system-configured-resources/resource-state
			"resource-state": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagEnum},
			},
			// System Status/system-configured-resources/resource-type
			"resource-type": &Attribute{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
		}},
	},
	// System Status/system-firmware-name
	"system-firmware-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Status/system-firmware-patches
	"system-firmware-patches": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-firmware-string-version
	"system-firmware-string-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-firmware-version
	"system-firmware-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// System Status/system-impressions-completed
	"system-impressions-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-impressions-completed-col
	"system-impressions-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// System Status/system-media-sheets-completed
	"system-media-sheets-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-media-sheets-completed-col
	"system-media-sheets-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// System Status/system-pages-completed
	"system-pages-completed": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-pages-completed-col
	"system-pages-completed-col": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagBeginCollection},
	},
	// System Status/system-resident-application-name
	"system-resident-application-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Status/system-resident-application-patches
	"system-resident-application-patches": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-resident-application-string-version
	"system-resident-application-string-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-resident-application-version
	"system-resident-application-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// System Status/system-serial-number
	"system-serial-number": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   255,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-state
	"system-state": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagEnum},
	},
	// System Status/system-state-change-date-time
	"system-state-change-date-time": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagDateTime},
	},
	// System Status/system-state-change-time
	"system-state-change-time": &Attribute{
		SetOf: false,
		Min:   0,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-state-message
	"system-state-message": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-state-reasons
	"system-state-reasons": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Status/system-time-source
	"system-time-source": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNameLang},
	},
	// System Status/system-up-time
	"system-up-time": &Attribute{
		SetOf: false,
		Min:   1,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagInteger},
	},
	// System Status/system-user-application-name
	"system-user-application-name": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagNameLang},
	},
	// System Status/system-user-application-patches
	"system-user-application-patches": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-user-application-string-version
	"system-user-application-string-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagTextLang},
	},
	// System Status/system-user-application-version
	"system-user-application-version": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   64,
		Tags:  []goipp.Tag{goipp.TagString},
	},
	// System Status/system-uuid
	"system-uuid": &Attribute{
		SetOf: false,
		Min:   MIN,
		Max:   45,
		Tags:  []goipp.Tag{goipp.TagURI},
	},
	// System Status/xri-authentication-supported
	"xri-authentication-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Status/xri-security-supported
	"xri-security-supported": &Attribute{
		SetOf: true,
		Min:   MIN,
		Max:   MAX,
		Tags:  []goipp.Tag{goipp.TagKeyword},
	},
	// System Status/xri-uri-scheme-supported
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
	"Document Description":  DocumentDescription,
	"Document Status":       DocumentStatus,
	"Document Template":     DocumentTemplate,
	"Event Notifications":   EventNotifications,
	"Job Description":       JobDescription,
	"Job Status":            JobStatus,
	"Job Template":          JobTemplate,
	"Operation":             Operation,
	"Printer Description":   PrinterDescription,
	"Printer Status":        PrinterStatus,
	"Resource Description":  ResourceDescription,
	"Resource Status":       ResourceStatus,
	"Subscription Status":   SubscriptionStatus,
	"Subscription Template": SubscriptionTemplate,
	"System Description":    SystemDescription,
	"System Status":         SystemStatus,
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
