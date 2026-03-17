// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// MediaSideImageInfo: shared image information for a single scan side,
// used by both MediaBackImageInfo and MediaFrontImageInfo

package wsscan

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// MediaSideImageInfo contains image dimension information for one side of a
// scanned document. It is the common structure behind MediaBackImageInfo and
// MediaFrontImageInfo.
type MediaSideImageInfo struct {
	BytesPerLine  int
	NumberOfLines int
	PixelsPerLine int
}

// toXML creates an XML element for MediaSideImageInfo using the given element
// name (e.g. "wscn:MediaBackImageInfo" or "wscn:MediaFrontImageInfo").
func (m MediaSideImageInfo) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":BytesPerLine",
				Text: strconv.Itoa(m.BytesPerLine)},
			{Name: NsWSCN + ":NumberOfLines",
				Text: strconv.Itoa(m.NumberOfLines)},
			{Name: NsWSCN + ":PixelsPerLine",
				Text: strconv.Itoa(m.PixelsPerLine)},
		},
	}
}

// decodeMediaSideImageInfo decodes a MediaSideImageInfo from an XML element.
func decodeMediaSideImageInfo(root xmldoc.Element) (MediaSideImageInfo, error) {
	var m MediaSideImageInfo

	bytesPerLine := xmldoc.Lookup{Name: NsWSCN +
		":BytesPerLine", Required: true}
	numberOfLines := xmldoc.Lookup{Name: NsWSCN +
		":NumberOfLines", Required: true}
	pixelsPerLine := xmldoc.Lookup{Name: NsWSCN +
		":PixelsPerLine", Required: true}

	if missed := root.Lookup(
		&bytesPerLine, &numberOfLines, &pixelsPerLine); missed != nil {
		return m, xmldoc.XMLErrMissed(missed.Name)
	}

	var err error

	if m.BytesPerLine, err = decodeNonNegativeInt(
		bytesPerLine.Elem); err != nil {
		return m, fmt.Errorf("BytesPerLine: %w", err)
	}

	if m.NumberOfLines, err = decodeNonNegativeInt(
		numberOfLines.Elem); err != nil {
		return m, fmt.Errorf("NumberOfLines: %w", err)
	}
	if m.NumberOfLines < 1 {
		return m, fmt.Errorf("NumberOfLines: must be at least 1, got %d",
			m.NumberOfLines)
	}

	if m.PixelsPerLine, err = decodeNonNegativeInt(
		pixelsPerLine.Elem); err != nil {
		return m, fmt.Errorf("PixelsPerLine: %w", err)
	}
	if m.PixelsPerLine < 1 {
		return m, fmt.Errorf("PixelsPerLine: must be at least 1, got %d",
			m.PixelsPerLine)
	}

	return m, nil
}
