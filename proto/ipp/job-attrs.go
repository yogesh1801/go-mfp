// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Job and Job Template Attributes

package ipp

import (
	"time"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/goipp"
)

// JobAttributes are attributes, supplied with Job creation request
type JobAttributes struct {
	// RFC8011, Internet Printing Protocol/1.1: Model and Semantics
	// 5.2 Job Template Attributes
	Copies                   optional.Val[int]                        `ipp:"copies,integer(1:MAX)"`
	Finishings               []int                                    `ipp:"finishings,1setOf enum"`
	JobHoldUntil             optional.Val[KwJobHoldUntil]             `ipp:"job-hold-until"`
	JobPriority              optional.Val[int]                        `ipp:"job-priority,integer(1:100)"`
	JobSheets                optional.Val[KwJobSheets]                `ipp:"job-sheets"`
	Media                    optional.Val[KwMedia]                    `ipp:"media"`
	MultipleDocumentHandling optional.Val[KwMultipleDocumentHandling] `ipp:"multiple-document-handling"`
	NumberUp                 optional.Val[int]                        `ipp:"number-up,integer(1:MAX)"`
	OrientationRequested     optional.Val[int]                        `ipp:"orientation-requested,enum"`
	PageRanges               []goipp.Range                            `ipp:"page-ranges"`
	PrinterResolution        optional.Val[goipp.Resolution]           `ipp:"printer-resolution"`
	PrintQuality             optional.Val[int]                        `ipp:"print-quality,enum"`
	Sides                    optional.Val[KwSides]                    `ipp:"sides"`

	// PWG5100.2: IPP “output-bin” attribute extension
	OutputBin optional.Val[string] `ipp:"output-bin"`

	// PWG5100.7: IPP Job Extensions v2.1 (JOBEXT)
	// 6.8 Job Template Attributes
	JobDelayOutputUntil     optional.Val[KwJobDelayOutputUntil] `ipp:"job-delay-output-until"`
	JobDelayOutputUntilTime optional.Val[time.Time]             `ipp:"job-delay-output-until-time"`
	JobHoldUntilTime        optional.Val[time.Time]             `ipp:"job-hold-until-time"`
	JobAccountID            optional.Val[string]                `ipp:"job-account-id,name"`
	JobAccountingUserID     optional.Val[string]                `ipp:"job-accounting-user-id,name"`
	JobCancelAfter          optional.Val[int]                   `ipp:"job-cancel-after"`
	JobRetainUntil          optional.Val[string]                `ipp:"job-retain-until"`
	JobRetainUntilInterval  optional.Val[int]                   `ipp:"job-retain-until-interval,integer(0:MAX)"`
	JobRetainUntilTime      optional.Val[time.Time]             `ipp:"job-retain-until-time"`
	JobSheetMessage         optional.Val[string]                `ipp:"job-sheet-message,text"`
	JobSheetsCol            JobSheets                           `ipp:"job-sheets-col"`
	PrintContentOptimize    optional.Val[string]                `ipp:"print-content-optimize,keyword"`

	// PWG5100.11: IPP Job and Printer Extensions – Set 2 (JPS2)
	// 7 Job Template Attributes
	FeedOrientation  optional.Val[string] `ipp:"feed-orientation,keyword"`
	JobPhoneNumber   optional.Val[string] `ipp:"job-phone-number,uri"`
	JobRecipientName optional.Val[string] `ipp:"job-recipient-name,name"`

	// PWG5100.13: IPP Driver Replacement Extensions v2.0 (NODRIVER)
	// 6.2 Job and Document Template Attributes
	JobErrorAction       optional.Val[string] `ipp:"job-error-action,keyword"`
	MediaOverprint       MediaOverprint       `ipp:"media-overprint"`
	PrintColorMode       optional.Val[string] `ipp:"print-color-mode,keyword"`
	PrintRenderingIntent optional.Val[string] `ipp:"print-rendering-intent,keyword"`
	PrintScaling         optional.Val[string] `ipp:"print-scaling,keyword"`
}

// JobTemplate are attributes, included into the Printer Description and
// describing possible settings for JobAttributes
type JobTemplate struct {
	// RFC8011, Internet Printing Protocol/1.1: Model and Semantics
	// 5.2 Job Template Attributes
	CopiesDefault                     optional.Val[int]                        `ipp:"copies-default,integer(1:MAX)"`
	CopiesSupported                   optional.Val[goipp.Range]                `ipp:"copies-supported,rangeOfinteger(1:MAX)"`
	FinishingsDefault                 []int                                    `ipp:"finishings-default,1setOf enum"`
	FinishingsSupported               []int                                    `ipp:"finishings-supported,1setOf enum"`
	JobHoldUntilDefault               optional.Val[KwJobHoldUntil]             `ipp:"job-hold-until-default"`
	JobHoldUntilSupported             []KwJobHoldUntil                         `ipp:"job-hold-until-supported"`
	JobPriorityDefault                optional.Val[int]                        `ipp:"job-priority-default,integer(1:100)"`
	JobPrioritySupported              optional.Val[int]                        `ipp:"job-priority-supported,integer(1:100)"`
	JobSheetsDefault                  []KwJobSheets                            `ipp:"job-sheets-default"`
	JobSheetsSupported                []KwJobSheets                            `ipp:"job-sheets-supported"`
	MediaDefault                      optional.Val[KwMedia]                    `ipp:"media-default"`
	MediaReady                        []KwMedia                                `ipp:"media-ready"`
	MediaSupported                    []KwMedia                                `ipp:"media-supported"`
	MultipleDocumentHandlingDefault   optional.Val[KwMultipleDocumentHandling] `ipp:"multiple-document-handling-default"`
	MultipleDocumentHandlingSupported []KwMultipleDocumentHandling             `ipp:"multiple-document-handling-supported"`
	NumberUpDefault                   optional.Val[int]                        `ipp:"number-up-default,integer(1:MAX)"`
	NumberUpSupported                 []goipp.IntegerOrRange                   `ipp:"number-up-supported,1setOf integer|rangeOfinteger(1:MAX)"`
	OrientationRequestedDefault       optional.Val[int]                        `ipp:"orientation-requested-default"`
	OrientationRequestedSupported     []int                                    `ipp:"orientation-requested-supported,1setOf enum"`
	PageRangesSupported               optional.Val[bool]                       `ipp:"page-ranges-supported"`
	PrinterResolutionDefault          optional.Val[goipp.Resolution]           `ipp:"printer-resolution-default"`
	PrinterResolutionSupported        []goipp.Resolution                       `ipp:"printer-resolution-supported"`
	PrintQualityDefault               optional.Val[int]                        `ipp:"print-quality-default,enum"`
	PrintQualitySupported             []int                                    `ipp:"print-quality-supported,1setOf enum"`
	SidesDefault                      optional.Val[KwSides]                    `ipp:"sides-default"`
	SidesSupported                    []KwSides                                `ipp:"sides-supported"`

	// PWG5100.2: IPP “output-bin” attribute extension
	OutputBinDefault   optional.Val[string] `ipp:"output-bin-default"`
	OutputBinSupported []string             `ipp:"output-bin-supported"`

	// PWG5100.7: IPP Job Extensions v2.1 (JOBEXT)
	// 6.9 Printer Description Attributes
	JobAccountIDDefault              optional.Val[string]                `ipp:"job-account-id-default,name|no-value"`
	JobAccountIDSupported            optional.Val[bool]                  `ipp:"job-account-id-supported"`
	JobAccountingUserIDDefault       optional.Val[string]                `ipp:"job-accounting-user-id-default,name|no-value"`
	JobAccountingUserIDSupported     optional.Val[bool]                  `ipp:"job-accounting-user-id-supported"`
	JobCancelAfterDefault            optional.Val[int]                   `ipp:"job-cancel-after-default"`
	JobCancelAfterSupported          optional.Val[goipp.Range]           `ipp:"job-cancel-after-supported,rangeOfInteger(0:MAX)"`
	JobDelayOutputUntilDefault       optional.Val[KwJobDelayOutputUntil] `ipp:"job-delay-output-until-default"`
	JobDelayOutputUntilSupported     []KwJobDelayOutputUntil             `ipp:"job-delay-output-until-supported"`
	JobDelayOutputUntilTimeSupported optional.Val[goipp.Range]           `ipp:"job-delay-output-until-time-supported,rangeOfInteger(0:MAX)"`
	JobHoldUntilTimeSupported        optional.Val[bool]                  `ipp:"job-hold-until-time-supported"`
	JobRetainUntilDefault            optional.Val[string]                `ipp:"job-retain-until-default"`
	JobRetainUntilIntervalDefault    optional.Val[int]                   `ipp:"job-retain-until-interval-default"`
	JobRetainUntilIntervalSupported  optional.Val[goipp.Range]           `ipp:"job-retain-until-interval-supported,rangeOfInteger(0:MAX)"`
	JobRetainUntilSupported          []string                            `ipp:"job-retain-until-supported"`
	JobRetainUntilTimeSupported      optional.Val[goipp.Range]           `ipp:"job-retain-until-time-supported,rangeOfInteger(0:MAX)"`
	JobSheetsColDefault              optional.Val[JobSheets]             `ipp:"job-sheets-col-default"`
	JobSheetsColSupported            []string                            `ipp:"job-sheets-col-supported,1setOf keyword"`
	PrintContentOptimizeDefault      optional.Val[string]                `ipp:"print-content-optimize-default,keyword"`
	PrintContentOptimizeSupported    []string                            `ipp:"print-content-optimize-supported,1setOf keyword"`

	// PWG5100.11: IPP Job and Printer Extensions – Set 2 (JPS2)
	// 7 Job Template Attributes
	FeedOrientationDefault               optional.Val[string] `ipp:"feed-orientation-default"`
	FeedOrientationSupported             string               `ipp:"feed-orientation-supported"`
	JobPhoneNumberDefault                optional.Val[string] `ipp:"job-phone-number-default"`
	JobPhoneNumberSupported              optional.Val[bool]   `ipp:"job-phone-number-supported"`
	JobRecipientNameDefault              optional.Val[string] `ipp:"job-recipient-name-default"`
	JobRecipientNameSupported            optional.Val[bool]   `ipp:"job-recipient-name-supported"`
	PdlInitFileEntrySupported            []string             `ipp:"pdl-init-file-entry-supported,1setOf name"`
	PdlInitFileNameSubdirectorySupported optional.Val[bool]   `ipp:"pdl-init-file-name-subdirectory-supported"`
	PdlInitFileNameSupported             []string             `ipp:"pdl-init-file-name-supported,1setOf name"`
	PdlInitFileSupported                 []string             `ipp:"pdl-init-file-supported"`
	PrintProcessingAttributesSupported   []string             `ipp:"print-processing-attributes-supported,1setOf keyword"`
	SaveDispositionSupported             []string             `ipp:"save-disposition-supported,1setOf keyword"`
	SaveDocumentFormatDefault            optional.Val[string] `ipp:"save-document-format-default,mimeMediaType"`
	SaveDocumentFormatSupported          []string             `ipp:"save-document-format-supported,1setOf mimeMediaType"`
	SaveLocationDefault                  optional.Val[string] `ipp:"save-location-default,uri"`
	SaveLocationSupported                []string             `ipp:"save-location-supported,1setOf uri"`
	SaveNameSubdirectorySupported        optional.Val[bool]   `ipp:"save-name-subdirectory-supported"`
	SaveNameSupported                    optional.Val[bool]   `ipp:"save-name-supported"`

	// PWG5100.13: IPP Driver Replacement Extensions v2.0 (NODRIVER)
	// 6.2 Job and Document Template Attributes
	// 6.5 Printer Description Attributes
	JobErrorActionDefault           optional.Val[string]         `ipp:"job-error-action-default,keyword"`
	JobErrorActionSupported         []string                     `ipp:"job-error-action-supported,1setOf keyword"`
	MediaOverprintDefault           optional.Val[MediaOverprint] `ipp:"media-overprint-default"`
	MediaOverprintDistanceSupported optional.Val[goipp.Range]    `ipp:"media-overprint-distance-supported,rangeOfInteger(0:MAX)"`
	MediaOverprintMethodSupported   []string                     `ipp:"media-overprint-method-supported,1setOf keyword"`
	MediaOverprintSupported         []string                     `ipp:"media-overprint-supported,1setOf keyword"`
	PrintColorModeDefault           optional.Val[string]         `ipp:"print-color-mode-default,keyword"`
	PrintColorModeSupported         []string                     `ipp:"print-color-mode-supported,1setOf keyword"`
	PrinterMandatoryJobAttributes   []string                     `ipp:"printer-mandatory-job-attributes,1setOf keyword"`
	PrintRenderingIntentDefault     optional.Val[string]         `ipp:"print-rendering-intent-default,keyword"`
	PrintRenderingIntentSupported   []string                     `ipp:"print-rendering-intent-supported,1setOf keyword"`
	PrintScalingDefault             optional.Val[string]         `ipp:"print-scaling-default,keyword"`
	PrintScalingSupported           []string                     `ipp:"print-scaling-supported,1setOf keyword"`

	// Wi-Fi Peer-to-Peer Services Print (P2Ps-Print)
	// Technical Specification
	// (for Wi-Fi Direct® services certification)
	PclmSourceResolution          optional.Val[goipp.Resolution] `ipp:"pclm-source-resolution"`
	PclmSourceResolutionSupported []goipp.Resolution             `ipp:"pclm-source-resolution-supported"`
}

// JobSheets represents "job-sheets-col" collection entry in
// JobAttributes
type JobSheets struct {
	JobSheets KwJobSheets `ipp:"job-sheets"`
	Media     string      `ipp:"media"`
	MediaCol  MediaCol    `ipp:"media-col"`
}

// JobPresets represents "job-presets-supported" collection entry
// in PrinterDescription
type JobPresets struct {
	PresetCategory string `ipp:"preset-category,keyword"`
	PresetName     string `ipp:"preset-name"`
	JobAttributes
}
