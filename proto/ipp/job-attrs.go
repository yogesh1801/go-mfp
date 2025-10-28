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
	Copies                   optional.Val[int]                        `ipp:"copies,(1:MAX)"`
	Finishings               []int                                    `ipp:"finishings,enum"`
	JobHoldUntil             optional.Val[KwJobHoldUntil]             `ipp:"job-hold-until"`
	JobPriority              optional.Val[int]                        `ipp:"job-priority,(1:100)"`
	JobSheets                optional.Val[KwJobSheets]                `ipp:"job-sheets"`
	Media                    optional.Val[KwMedia]                    `ipp:"media"`
	MultipleDocumentHandling optional.Val[KwMultipleDocumentHandling] `ipp:"multiple-document-handling"`
	NumberUp                 optional.Val[int]                        `ipp:"number-up,(1:MAX)"`
	OrientationRequested     optional.Val[int]                        `ipp:"orientation-requested,enum"`
	PageRanges               []goipp.IntegerOrRange                   `ipp:"page-ranges"`
	PrinterResolution        optional.Val[goipp.Resolution]           `ipp:"printer-resolution"`
	PrintQuality             optional.Val[int]                        `ipp:"print-quality,enum"`
	Sides                    optional.Val[KwSides]                    `ipp:"sides"`

	// PWG5100.7: IPP Job Extensions v2.1 (JOBEXT)
	// 6.8 Job Template Attributes
	JobDelayOutputUntil     optional.Val[KwJobDelayOutputUntil] `ipp:"job-delay-output-until"`
	JobDelayOutputUntilTime optional.Val[time.Time]             `ipp:"job-delay-output-until-time"`
	JobHoldUntilTime        optional.Val[time.Time]             `ipp:"job-hold-until-time"`
	JobAccountID            optional.Val[string]                `ipp:"job-account-id,name"`
	JobAccountingUserID     optional.Val[string]                `ipp:"job-accounting-user-id,name"`
	JobCancelAfter          optional.Val[int]                   `ipp:"job-cancel-after,(0:MAX)"`
	JobRetainUntil          optional.Val[string]                `ipp:"job-retain-until,keyword"`
	JobRetainUntilInterval  optional.Val[int]                   `ipp:"job-retain-until-interval,(0:MAX)"`
	JobRetainUntilTime      optional.Val[time.Time]             `ipp:"job-retain-until-time"`
	JobSheetMessage         optional.Val[string]                `ipp:"job-sheet-message,text"`
	JobSheetsCol            []JobSheets                         `ipp:"job-sheets-col"`

	// PWG5100.11: IPP Job and Printer Extensions – Set 2 (JPS2)
	// 7 Job Template Attributes
	FeedOrientation    optional.Val[string] `ipp:"feed-orientation,keyword"`
	FontNameRequested  optional.Val[string] `ipp:"font-name-requested,name"`
	FontSizeRequested  optional.Val[int]    `ipp:"font-size-requested,(1:MAX)"`
	JobPhoneNumber     optional.Val[string] `ipp:"job-phone-number,uri"`
	JobRecipientName   optional.Val[string] `ipp:"job-recipient-name,name"`
	JobSaveDisposition []JobSaveDisposition `ipp:"job-save-disposition"`
	PdlInitFile        []JobPdlInitFile     `ipp:"pdl-init-file"`

	// PWG5100.13: IPP Driver Replacement Extensions v2.0 (NODRIVER)
	// 6.2 Job and Document Template Attributes
	JobErrorAction       optional.Val[string] `ipp:"job-error-action,keyword"`
	MediaOverprint       []JobMediaOverprint  `ipp:"media-overprint"`
	PrintColorMode       optional.Val[string] `ipp:"print-color-mode,keyword"`
	PrintRenderingIntent optional.Val[string] `ipp:"print-rendering-intent,keyword"`
	PrintScaling         optional.Val[string] `ipp:"print-scaling,keyword"`
}

// JobTemplate are attributes, included into the Printer Description and
// describing possible settings for JobAttributes
type JobTemplate struct {
	// RFC8011, Internet Printing Protocol/1.1: Model and Semantics
	// 5.2 Job Template Attributes
	CopiesDefault                     optional.Val[int]                        `ipp:"copies-default,(1:MAX)"`
	CopiesSupported                   optional.Val[goipp.Range]                `ipp:"copies-supported,(1:MAX)"`
	FinishingsDefault                 []int                                    `ipp:"finishings-default,enum"`
	FinishingsSupported               []int                                    `ipp:"finishings-supported,enum"`
	JobHoldUntilDefault               optional.Val[KwJobHoldUntil]             `ipp:"job-hold-until-default"`
	JobHoldUntilSupported             []KwJobHoldUntil                         `ipp:"job-hold-until-supported"`
	JobPriorityDefault                optional.Val[int]                        `ipp:"job-priority-default,(1:100)"`
	JobPrioritySupported              optional.Val[int]                        `ipp:"job-priority-supported,(1:100)"`
	JobSheetsDefault                  optional.Val[KwJobSheets]                `ipp:"job-sheets-default"`
	JobSheetsSupported                []KwJobSheets                            `ipp:"job-sheets-supported"`
	MediaDefault                      optional.Val[KwMedia]                    `ipp:"media-default"`
	MediaReady                        []KwMedia                                `ipp:"media-ready"`
	MediaSupported                    []KwMedia                                `ipp:"media-supported"`
	MultipleDocumentHandlingDefault   optional.Val[KwMultipleDocumentHandling] `ipp:"multiple-document-handling-default"`
	MultipleDocumentHandlingSupported []KwMultipleDocumentHandling             `ipp:"multiple-document-handling-supported"`
	NumberUpDefault                   optional.Val[int]                        `ipp:"number-up-default,(1:MAX)"`
	NumberUpSupported                 []goipp.IntegerOrRange                   `ipp:"number-up-supported,(1:MAX)"`
	OrientationRequestedDefault       optional.Val[int]                        `ipp:"orientation-requested-default,enum"`
	OrientationRequestedSupported     []int                                    `ipp:"orientation-requested-supported,enum"`
	PageRangesSupported               optional.Val[bool]                       `ipp:"page-ranges-supported"`
	PrinterResolutionDefault          optional.Val[goipp.Resolution]           `ipp:"printer-resolution-default"`
	PrinterResolutionSupported        []goipp.Resolution                       `ipp:"printer-resolution-supported"`
	PrintQualityDefault               optional.Val[int]                        `ipp:"print-quality-default,enum"`
	PrintQualitySupported             []int                                    `ipp:"print-quality-supported,enum"`
	SidesDefault                      optional.Val[KwSides]                    `ipp:"sides-default"`
	SidesSupported                    []KwSides                                `ipp:"sides-supported"`

	// PWG5100.7: IPP Job Extensions v2.1 (JOBEXT)
	// 6.9 Printer Description Attributes
	JobAccountIDDefault              optional.Val[string]                `ipp:"job-account-id-default,name|no-value"`
	JobAccountIDSupported            optional.Val[bool]                  `ipp:"job-account-id-supported"`
	JobAccountingUserIDDefault       optional.Val[string]                `ipp:"job-accounting-user-id-default,name|no-value"`
	JobAccountingUserIDSupported     optional.Val[bool]                  `ipp:"job-accounting-user-id-supported"`
	JobCancelAfterDefault            optional.Val[int]                   `ipp:"job-cancel-after-default,(0:MAX)"`
	JobCancelAfterSupported          optional.Val[goipp.Range]           `ipp:"job-cancel-after-supported,(0:MAX)"`
	JobDelayOutputUntilDefault       optional.Val[KwJobDelayOutputUntil] `ipp:"job-delay-output-until-default"`
	JobDelayOutputUntilSupported     []KwJobDelayOutputUntil             `ipp:"job-delay-output-until-supported"`
	JobDelayOutputUntilTimeSupported optional.Val[goipp.Range]           `ipp:"job-delay-output-until-time-supported,(0:MAX)"`
	JobHoldUntilTimeSupported        optional.Val[goipp.Range]           `ipp:"job-hold-until-time-supported,(0:MAX)"`
	JobRetainUntilDefault            optional.Val[string]                `ipp:"job-retain-until-default,keyword"`
	JobRetainUntilIntervalDefault    optional.Val[int]                   `ipp:"job-retain-until-interval-default,(0:MAX)"`
	JobRetainUntilIntervalSupported  optional.Val[goipp.Range]           `ipp:"job-retain-until-interval-supported,(0:MAX)"`
	JobRetainUntilSupported          []string                            `ipp:"job-retain-until-supported,keyword"`
	JobRetainUntilTimeSupported      optional.Val[goipp.Range]           `ipp:"job-retain-until-time-supported,(0:MAX)"`
	JobSheetsColDefault              []JobSheets                         `ipp:"job-sheets-col-default,collection|no-value"`
	JobSheetsColSupported            []string                            `ipp:"job-sheets-col-supported,keyword"`

	// PWG5100.11: IPP Job and Printer Extensions – Set 2 (JPS2)
	// 7 Job Template Attributes
	FeedOrientationDefault               optional.Val[string] `ipp:"feed-orientation-default,keyword"`
	FeedOrientationSupported             []string             `ipp:"feed-orientation-supported,keyword"`
	FontNameRequestedDefault             optional.Val[string] `ipp:"font-name-requested-default,name"`
	FontNameRequestedSupported           []string             `ipp:"font-name-requested-supported,name"`
	FontSizeRequestedDefault             optional.Val[int]    `ipp:"font-size-requested-default,(1:MAX)"`
	FontSizeRequestedSupported           []int                `ipp:"font-size-requested-supported,(1:MAX)"`
	JobPhoneNumberDefault                optional.Val[string] `ipp:"job-phone-number-default,uri"`
	JobPhoneNumberSupported              optional.Val[bool]   `ipp:"job-phone-number-supported"`
	JobRecipientNameDefault              optional.Val[string] `ipp:"job-recipient-name-default,name"`
	JobRecipientNameSupported            optional.Val[bool]   `ipp:"job-recipient-name-supported"`
	JobSaveDispositionDefault            []JobSaveDisposition `ipp:"job-save-disposition-default"`
	JobSaveDispositionSupported          []string             `ipp:"job-save-disposition-supported,keyword"`
	PdlInitFileDefault                   []JobPdlInitFile     `ipp:"pdl-init-file-default"`
	PdlInitFileEntrySupported            []string             `ipp:"pdl-init-file-entry-supported,name"`
	PdlInitFileNameSubdirectorySupported optional.Val[bool]   `ipp:"pdl-init-file-name-subdirectory-supported"`
	PdlInitFileNameSupported             []string             `ipp:"pdl-init-file-name-supported,name"`
	PdlInitFileSupported                 []string             `ipp:" pdl-init-file-supported,name"`
	PrintProcessingAttributesSupported   []string             `ipp:"print-processing-attributes-supported,keyword"`
	SaveDispositionSupported             []string             `ipp:"save-disposition-supported,keyword"`
	SaveDocumentFormatDefault            optional.Val[string] `ipp:"save-document-format-default,mimeMediaType"`
	SaveDocumentFormatSupported          []string             `ipp:"save-document-format-supported,mimeMediaType"`
	SaveInfoSupported                    []string             `ipp:"save-info-supported,keyword"`
	SaveLocationDefault                  optional.Val[string] `ipp:"save-location-default,uri"`
	SaveLocationSupported                []string             `ipp:"save-location-supported,uri"`
	SaveNameSubdirectorySupported        optional.Val[bool]   `ipp:"save-name-subdirectory-supported"`
	SaveNameSupported                    optional.Val[bool]   `ipp:"save-name-supported"`

	// PWG5100.13: IPP Driver Replacement Extensions v2.0 (NODRIVER)
	// 6.2 Job and Document Template Attributes
	// 6.5 Printer Description Attributes
	JobErrorActionDefault           optional.Val[string]      `ipp:"job-error-action-default,keyword"`
	JobErrorActionSupported         []string                  `ipp:"job-error-action-supported,keyword"`
	MediaOverprintDefault           []JobMediaOverprint       `ipp:"media-overprint-default"`
	MediaOverprintDistanceSupported optional.Val[goipp.Range] `ipp:"media-overprint-distance-supported,(0:MAX)"`
	MediaOverprintMethodSupported   []string                  `ipp:"media-overprint-method-supported,keyword"`
	MediaOverprintSupported         []string                  `ipp:"media-overprint-supported,keyword"`
	PrintColorModeDefault           optional.Val[string]      `ipp:"print-color-mode-default,keyword"`
	PrintColorModeSupported         []string                  `ipp:"print-color-mode-supported,keyword"`
	PrinterMandatoryJobAttributes   []string                  `ipp:"printer-mandatory-job-attributes,keyword"`
	PrintRenderingIntentDefault     optional.Val[string]      `ipp:"print-rendering-intent-default,keyword"`
	PrintRenderingIntentSupported   []string                  `ipp:"print-rendering-intent-supported,keyword"`
	PrintScalingDefault             optional.Val[string]      `ipp:"print-scaling-default,keyword"`
	PrintScalingSupported           []string                  `ipp:"print-scaling-supported,keyword"`
}

// MediaCol is the "media-col", "media-col-xxx" collection entry.
// It is used in many places.
//
// PWG5100.7
type MediaCol struct {
	// ----- PWG5100.3 -----
	KwMediaBackCoating optional.Val[KwMediaBackCoating] `ipp:"media-back-coating"`
	MediaColor         optional.Val[KwColor]            `ipp:"media-color"`
	MediaFrontCoating  optional.Val[KwMediaBackCoating] `ipp:"media-front-coating"`
	MediaHoleCount     optional.Val[int]                `ipp:"media-hole-count,(0:MAX)"`
	MediaInfo          optional.Val[string]             `ipp:"media-info,text"`
	MediaKey           optional.Val[KwMedia]            `ipp:"media-key"`
	MediaOrderCount    optional.Val[int]                `ipp:"media-order-count,(1:MAX)"`
	MediaPrePrinted    optional.Val[string]             `ipp:"media-pre-printed,keyword"`
	MediaRecycled      optional.Val[string]             `ipp:"media-recycled,keyword"`
	MediaSize          optional.Val[MediaSize]          `ipp:"media-size"`
	MediaType          optional.Val[string]             `ipp:"media-type,keyword"`
	MediaWeightMetric  optional.Val[int]                `ipp:"media-weight-metric,(0:MAX)"`

	// ----- PWG5100.7 -----
	MediaBottomMargin     optional.Val[int]                   `ipp:"media-bottom-margin,(0:MAX)"`
	MediaGrain            optional.Val[string]                `ipp:"media-grain,keyword"`
	MediaLeftMargin       optional.Val[int]                   `ipp:"media-left-margin,(0:MAX)"`
	MediaRightMargin      optional.Val[int]                   `ipp:"media-right-margin,(0:MAX)"`
	MediaSizeName         optional.Val[string]                `ipp:"media-size-name,keyword"`
	MediaSourceProperties optional.Val[MediaSourceProperties] `ipp:"media-source-properties"`
	MediaSource           optional.Val[string]                `ipp:"media-source,keyword"`
	MediaThickness        optional.Val[int]                   `ipp:"media-thickness,(1:MAX)"`
	MediaTooth            optional.Val[string]                `ipp:"media-tooth,keyword"`
	MediaTopMargin        optional.Val[int]                   `ipp:"media-top-margin,(0:MAX)"`
}

// MediaSize represents media size parameters (which may be either
// pair of integers or pair of ranges) and used in many places
type MediaSize struct {
	XDimension goipp.IntegerOrRange `ipp:"x-dimension,(1:MAX)"`
	YDimension goipp.IntegerOrRange `ipp:"y-dimension,(1:MAX)"`
}

// MediaSourceProperties represents "media-source-properties"
// collectiobn in MediaCol
type MediaSourceProperties struct {
	MediaSourceFeedDirection   string `ipp:"media-source-feed-direction,keyword"`
	MediaSourceFeedOrientation int    `ipp:"media-source-feed-orientation,enum"`
}

// JobSheets represents "job-sheets-col" collection entry in
// JobAttributes
type JobSheets struct {
	JobSheets KwJobSheets `ipp:"job-sheets"`
	Media     string      `ipp:"media,keyword"`
	MediaCol  []MediaCol  `ipp:"media-col"`
}

// JobSaveDisposition represents "job-save-disposition"
// collection entry in JobAttributes and "job-save-disposition-default"
// entry in JobTemplate
type JobSaveDisposition struct {
	SaveDisposition string        `ipp:"save-disposition,keyword"`
	SaveInfo        []JobSaveInfo `ipp:"save-info"`
}

// JobSaveInfo represents "save-info" collection entry
// in JobSaveDisposition
type JobSaveInfo struct {
	SaveLocation       string `ipp:"save-location,uri"`
	SaveName           string `ipp:"save-name,name"`
	SaveDocumentFormat string `ipp:"save-document-format,mimeMediaType"`
}

// JobPdlInitFile represents "pdl-init-file" collection entry
// in JobAttributes
type JobPdlInitFile struct {
	PdlInitFileLocation string `ipp:"pdl-init-file-location,uri"`
	PdlInitFileName     string `ipp:"pdl-init-file-name,name"`
	PdlInitFileEntry    string `ipp:"pdl-init-file-entry,name"`
}

// JobMediaOverprint represents "media-overprint" collection entry
// in JobAttributes
type JobMediaOverprint struct {
	MediaOverprintDistance int    `ipp:"media-overprint-distance,(0:MAX)"`
	MediaOverprintMethod   string `ipp:"media-overprint-method,keyword"`
}

// JobPresets represents "job-presets-supported" collection entry
// in PrinterDescription
type JobPresets struct {
	PresetCategory string `ipp:"preset-category,keyword"`
	PresetName     string `ipp:"preset-name,name"`
	JobAttributes
}
