// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Job and Job Template Attributes

package ipp

import (
	"time"

	"github.com/OpenPrinting/goipp"
)

// JobAttributes are attributes, supplied with Job creation request
type JobAttributes struct {
	// RFC8011, Internet Printing Protocol/1.1: Model and Semantics
	// 5.2 Job Template Attributes
	Copies                   int                    `ipp:"?copies,>0"`
	Finishings               []int                  `ipp:"?finishings,enum"`
	JobHoldUntil             string                 `ipp:"?job-hold-until,keyword"`
	JobPriority              int                    `ipp:"?job-priority,1:100"`
	JobSheets                string                 `ipp:"?job-sheets,keyword"`
	Media                    string                 `ipp:"?media,keyword"`
	MediaReady               []string               `ipp:"?media-ready,keyword"`
	MultipleDocumentHandling string                 `ipp:"?multiple-document-handling,keyword"`
	NumberUp                 int                    `ipp:"?number-up,>0"`
	OrientationRequested     int                    `ipp:"?orientation-requested,enum"`
	PageRanges               []goipp.IntegerOrRange `ipp:"?page-ranges"`
	PrinterResolution        goipp.Resolution       `ipp:"?printer-resolution"`
	PrintQuality             int                    `ipp:"?print-quality,enum"`
	Sides                    string                 `ipp:"?sides,keyword"`

	// PWG5100.11: IPP Job and Printer Extensions – Set 2 (JPS2)
	// 7 Job Template Attributes
	FeedOrientation         string               `ipp:"?feed-orientation,keyword"`
	FontNameRequested       string               `ipp:"?font-name-requested,name"`
	FontSizeRequested       int                  `ipp:"?font-size-requested,>0"`
	JobDelayOutputUntil     string               `ipp:"?job-delay-output-until,keyword"`
	JobDelayOutputUntilTime time.Time            `ipp:"?job-delay-output-until-time"`
	JobHoldUntilTime        time.Time            `ipp:"?job-hold-until-time"`
	JobPhoneNumber          string               `ipp:"?job-phone-number,uri"`
	JobRecipientName        string               `ipp:"?job-recipient-name,name"`
	JobSaveDisposition      []JobSaveDisposition `ipp:"?job-save-disposition"`
	PdlInitFile             []JobPdlInitFile     `ipp:"?pdl-init-file"`
}

// JobTemplate are attributes, included into the Printer Description and
// describing possible settings for JobAttributes
type JobTemplate struct {
	// RFC8011, Internet Printing Protocol/1.1: Model and Semantics
	// 5.2 Job Template Attributes
	CopiesDefault                     int                    `ipp:"?copies-default,>0"`
	CopiesSupported                   goipp.Range            `ipp:"?copies-supported,>0"`
	FinishingsDefault                 []int                  `ipp:"?finishings-default,enum"`
	FinishingsSupported               []int                  `ipp:"?finishings-supported,enum"`
	JobHoldUntilDefault               string                 `ipp:"?job-hold-until-default,keyword"`
	JobHoldUntilSupported             []string               `ipp:"?job-hold-until-supported,keyword"`
	JobPriorityDefault                int                    `ipp:"?job-priority-default,1:100"`
	JobPrioritySupported              int                    `ipp:"?job-priority-supported,1:100"`
	JobSheetsDefault                  string                 `ipp:"?job-sheets-default,keyword"`
	JobSheetsSupported                []string               `ipp:"?job-sheets-supported"`
	MediaDefault                      string                 `ipp:"?media-default,keyword"`
	MediaReady                        []string               `ipp:"?media-ready,keyword"`
	MediaSupported                    []string               `ipp:"?media-supported,keyword"`
	MultipleDocumentHandlingDefault   string                 `ipp:"?multiple-document-handling-default,keyword"`
	MultipleDocumentHandlingSupported []string               `ipp:"?multiple-document-handling-supported,keyword"`
	NumberUpDefault                   int                    `ipp:"?number-up-default,>0"`
	NumberUpSupported                 []goipp.IntegerOrRange `ipp:"?number-up-supported,>0"`
	OrientationRequestedDefault       int                    `ipp:"?orientation-requested-default,enum"`
	OrientationRequestedSupported     []int                  `ipp:"?orientation-requested-supported,enum"`
	PageRangesSupported               bool                   `ipp:"?page-ranges-supported"`
	PrinterResolutionDefault          goipp.Resolution       `ipp:"?printer-resolution-default"`
	PrinterResolutionSupported        []goipp.Resolution     `ipp:"?printer-resolution-supported"`
	PrintQualityDefault               int                    `ipp:"?print-quality-default,enum"`
	PrintQualitySupported             []int                  `ipp:"?print-quality-supported,enum"`
	SidesDefault                      string                 `ipp:"?sides-default,keyword"`
	SidesSupported                    []string               `ipp:"?sides-supported,keyword"`

	// PWG5100.11: IPP Job and Printer Extensions – Set 2 (JPS2)
	// 7 Job Template Attributes
	FeedOrientationDefault               string               `ipp:"?feed-orientation-default,keyword"`
	FeedOrientationSupported             []string             `ipp:"?feed-orientation-supported,keyword"`
	FontNameRequestedDefault             string               `ipp:"?font-name-requested-default,name"`
	FontNameRequestedSupported           []string             `ipp:"?font-name-requested-supported,name"`
	FontSizeRequestedDefault             int                  `ipp:"?font-size-requested-default,>0"`
	FontSizeRequestedSupported           []int                `ipp:"?font-size-requested-supported,>0"`
	JobDelayOutputUntilDefault           string               `ipp:"?job-delay-output-until-default,keyword"`
	JobDelayOutputUntilSupported         []string             `ipp:"?job-delay-output-until-supported,keyword"`
	JobDelayOutputUntilTimeSupported     goipp.Range          `ipp:"?job-delay-output-until-time-supported,>-1"`
	JobHoldUntilTimeSupported            goipp.Range          `ipp:"?job-hold-until-time-supported,>-1"`
	JobPhoneNumberDefault                string               `ipp:"?job-phone-number-default,uri"`
	JobPhoneNumberSupported              bool                 `ipp:"?job-phone-number-supported"`
	JobRecipientNameDefault              string               `ipp:"?job-recipient-name-default,name"`
	JobRecipientNameSupported            bool                 `ipp:"?job-recipient-name-supported"`
	JobSaveDispositionDefault            []JobSaveDisposition `ipp:"?job-save-disposition-default"`
	JobSaveDispositionSupported          []string             `ipp:"?job-save-disposition-supported,keyword"`
	PdlInitFileDefault                   []JobPdlInitFile     `ipp:"?pdl-init-file-default"`
	PdlInitFileEntrySupported            []string             `ipp:"?pdl-init-file-entry-supported,name"`
	PdlInitFileNameSubdirectorySupported bool                 `ipp:"?pdl-init-file-name-subdirectory-supported"`
	PdlInitFileNameSupported             []string             `ipp:"?pdl-init-file-name-supported,name"`
	PdlInitFileSupported                 []string             `ipp:"? pdl-init-file-supported,name"`
	SaveDispositionSupported             []string             `ipp:"?save-disposition-supported,keyword"`
	SaveDocumentFormatDefault            string               `ipp:"?save-document-format-default,mimeMediaType"`
	SaveDocumentFormatSupported          []string             `ipp:"?save-document-format-supported,mimeMediaType"`
	SaveInfoSupported                    []string             `ipp:"?save-info-supported,keyword"`
	SaveLocationDefault                  string               `ipp:"?save-location-default,uri"`
	SaveLocationSupported                []string             `ipp:"?save-location-supported,uri"`
	SaveNameSubdirectorySupported        bool                 `ipp:"?save-name-subdirectory-supported"`
	SaveNameSupported                    bool                 `ipp:"?save-name-supported"`
}

// JobSaveDisposition represents "job-save-disposition"
// collection entry in JobAttributes and "job-save-disposition-default"
// entry in JobTemplate
type JobSaveDisposition struct {
	SaveDisposition string        `ipp:"save-disposition,keyword"`
	SaveInfo        []JobSaveInfo `ipp:"?save-info"`
}

// JobSaveInfo represents "save-info" collection entry
// in JobSaveDisposition
type JobSaveInfo struct {
	SaveLocation       string `ipp:"?save-location,uri"`
	SaveName           string `ipp:"?save-name,name"`
	SaveDocumentFormat string `ipp:"?save-document-format,mimeMediaType"`
}

// JobPdlInitFile represents "pdl-init-file" collection entry
// in JobAttributes
type JobPdlInitFile struct {
	PdlInitFileLocation string `ipp:"?pdl-init-file-location,uri"`
	PdlInitFileName     string `ipp:"?pdl-init-file-name,name"`
	PdlInitFileEntry    string `ipp:"?pdl-init-file-entry,name"`
}
