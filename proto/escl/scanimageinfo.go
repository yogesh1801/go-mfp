// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan Image Info

package escl

import (
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScanImageInfo is the scanner response, that represents the actual
// information on scanned image. The actual image parameters may
// be different from the estimated [ScanBufferInfo] values.
//
// eSCL Technical Specification, 11.2.
//
// GET /{root}/ScanJobs/{job-id}/ScanImageInfo
type ScanImageInfo struct {
	JobURI             string               // Base URL for the job
	JobUUID            optional.Val[string] // Unique, persistent
	ActualWidth        int                  // Actual image width
	ActualHeight       int                  // Actual image height
	ActualBytesPerLine int                  // Actual bytes per line
	BlankPageDetected  optional.Val[bool]   // Blank page detected
}

// DecodeScanImageInfo decodes [ScanImageInfo] from the XML tree.
func DecodeScanImageInfo(root xmldoc.Element) (ret *ScanImageInfo, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	var info ScanImageInfo

	// Lookup relevant XML elements
	jobURI := xmldoc.Lookup{Name: NsPWG + ":JobUri", Required: true}
	jobUUID := xmldoc.Lookup{Name: NsPWG + ":JobUuid"}
	wid := xmldoc.Lookup{Name: NsScan + ":ActualWidth", Required: true}
	hei := xmldoc.Lookup{Name: NsScan + ":ActualHeight", Required: true}
	bpl := xmldoc.Lookup{Name: NsScan + ":ActualBytesPerLine",
		Required: true}
	blank := xmldoc.Lookup{Name: NsScan + ":BlankPageDetected"}

	missed := root.Lookup(&jobURI, &jobUUID, &wid, &hei, &bpl, &blank)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	info.JobURI = jobURI.Elem.Text

	if jobUUID.Found {
		info.JobUUID = optional.New(jobUUID.Elem.Text)
	}

	info.ActualWidth, err = decodeNonNegativeInt(wid.Elem)

	if err == nil {
		info.ActualHeight, err = decodeNonNegativeInt(hei.Elem)
	}
	if err == nil {
		info.ActualBytesPerLine, err = decodeNonNegativeInt(bpl.Elem)
	}

	if err == nil && blank.Found {
		info.BlankPageDetected, err = decodeOptional(
			blank.Elem, decodeBool)
	}

	ret = &info
	return
}

// ToXML generates XML tree for the [ScanImageInfo].
func (info *ScanImageInfo) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsScan + ":ScanImageInfo",
		Children: []xmldoc.Element{
			{Name: NsPWG + ":JobUri", Text: info.JobURI},
		},
	}

	if info.JobUUID != nil {
		chld := xmldoc.Element{
			Name: NsPWG + ":JobUuid",
			Text: *info.JobUUID,
		}
		elm.Children = append(elm.Children, chld)
	}

	elm.Children = append(elm.Children,
		xmldoc.WithText(NsScan+":ActualWidth",
			strconv.FormatUint(uint64(info.ActualWidth), 10)))

	elm.Children = append(elm.Children,
		xmldoc.WithText(NsScan+":ActualHeight",
			strconv.FormatUint(uint64(info.ActualHeight), 10)))

	elm.Children = append(elm.Children,
		xmldoc.WithText(NsScan+":ActualBytesPerLine",
			strconv.FormatUint(
				uint64(info.ActualBytesPerLine), 10)))

	if info.BlankPageDetected != nil {
		elm.Children = append(elm.Children,
			xmldoc.WithText(NsScan+":BlankPageDetected",
				strconv.FormatBool(*info.BlankPageDetected)))
	}

	return elm
}
