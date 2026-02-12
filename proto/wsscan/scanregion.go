// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ScanRegion: specifies the area to scan within the input document boundaries

package wsscan

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScanRegion specifies the area to scan within the input document boundaries.
// It contains required Height and Width,
// and optional XOffset and YOffset child elements.
type ScanRegion struct {
	ScanRegionHeight  ValWithOptions[int]
	ScanRegionWidth   ValWithOptions[int]
	ScanRegionXOffset optional.Val[ValWithOptions[int]]
	ScanRegionYOffset optional.Val[ValWithOptions[int]]
}

// decodeScanRegion decodes a ScanRegion from an XML element.
func decodeScanRegion(root xmldoc.Element) (ScanRegion, error) {
	var sr ScanRegion

	// Setup lookups for required child elements
	height := xmldoc.Lookup{
		Name:     NsWSCN + ":ScanRegionHeight",
		Required: true,
	}
	width := xmldoc.Lookup{
		Name:     NsWSCN + ":ScanRegionWidth",
		Required: true,
	}
	// Optional child elements
	xOffset := xmldoc.Lookup{
		Name:     NsWSCN + ":ScanRegionXOffset",
		Required: false,
	}
	yOffset := xmldoc.Lookup{
		Name:     NsWSCN + ":ScanRegionYOffset",
		Required: false,
	}

	missed := root.Lookup(&height, &width, &xOffset, &yOffset)
	if missed != nil {
		return sr, xmldoc.XMLErrMissed(missed.Name)
	}

	// Decoder function for integer values
	intDecoder := func(s string) (int, error) {
		return strconv.Atoi(s)
	}

	// Decode ScanRegionHeight
	h, err := sr.ScanRegionHeight.decodeValWithOptions(height.Elem, intDecoder)
	if err != nil {
		return sr, fmt.Errorf("ScanRegionHeight: %w", err)
	}
	sr.ScanRegionHeight = h

	// Decode ScanRegionWidth
	w, err := sr.ScanRegionWidth.decodeValWithOptions(width.Elem, intDecoder)
	if err != nil {
		return sr, fmt.Errorf("ScanRegionWidth: %w", err)
	}
	sr.ScanRegionWidth = w

	// Decode ScanRegionXOffset if present
	if xOffset.Found {
		var x ValWithOptions[int]
		decoded, err := x.decodeValWithOptions(xOffset.Elem, intDecoder)
		if err != nil {
			return sr, fmt.Errorf("ScanRegionXOffset: %w", err)
		}
		sr.ScanRegionXOffset = optional.New(decoded)
	}

	// Decode ScanRegionYOffset if present
	if yOffset.Found {
		var y ValWithOptions[int]
		decoded, err := y.decodeValWithOptions(yOffset.Elem, intDecoder)
		if err != nil {
			return sr, fmt.Errorf("ScanRegionYOffset: %w", err)
		}
		sr.ScanRegionYOffset = optional.New(decoded)
	}

	return sr, nil
}

// toXML creates an XML element for ScanRegion.
func (sr ScanRegion) toXML(name string) xmldoc.Element {
	intToString := func(i int) string {
		return strconv.Itoa(i)
	}

	elm := xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			sr.ScanRegionHeight.toXML(NsWSCN+":ScanRegionHeight", intToString),
			sr.ScanRegionWidth.toXML(NsWSCN+":ScanRegionWidth", intToString),
		},
	}

	// Add optional XOffset if present
	if sr.ScanRegionXOffset != nil {
		elm.Children = append(elm.Children,
			optional.Get(sr.ScanRegionXOffset).toXML(
				NsWSCN+":ScanRegionXOffset", intToString))
	}

	// Add optional YOffset if present
	if sr.ScanRegionYOffset != nil {
		elm.Children = append(elm.Children,
			optional.Get(sr.ScanRegionYOffset).toXML(
				NsWSCN+":ScanRegionYOffset", intToString))
	}

	return elm
}
