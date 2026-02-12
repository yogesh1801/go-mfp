// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// MediaSide: common type for MediaFront and MediaBack elements

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// MediaSide contains all parameters that are specific
// to the scanning of one side of the physical media.
// This is a common type used for both MediaFront and MediaBack elements.
type MediaSide struct {
	ColorProcessing optional.Val[ColorProcessing]
	Resolution      optional.Val[Resolution]
	ScanRegion      optional.Val[ScanRegion]
}

// decodeMediaSide decodes a MediaSide from an XML element.
func decodeMediaSide(root xmldoc.Element) (MediaSide, error) {
	var ms MediaSide

	// All child elements are optional
	colorProcessing := xmldoc.Lookup{
		Name:     NsWSCN + ":ColorProcessing",
		Required: false,
	}
	resolution := xmldoc.Lookup{
		Name:     NsWSCN + ":Resolution",
		Required: false,
	}
	scanRegion := xmldoc.Lookup{
		Name:     NsWSCN + ":ScanRegion",
		Required: false,
	}

	root.Lookup(&colorProcessing, &resolution, &scanRegion)

	// Decode ColorProcessing if present
	if colorProcessing.Found {
		cp, err := decodeColorProcessing(colorProcessing.Elem)
		if err != nil {
			return ms, fmt.Errorf("ColorProcessing: %w", err)
		}
		ms.ColorProcessing = optional.New(cp)
	}

	// Decode Resolution if present
	if resolution.Found {
		res, err := decodeResolution(resolution.Elem)
		if err != nil {
			return ms, fmt.Errorf("Resolution: %w", err)
		}
		ms.Resolution = optional.New(res)
	}

	// Decode ScanRegion if present
	if scanRegion.Found {
		sr, err := decodeScanRegion(scanRegion.Elem)
		if err != nil {
			return ms, fmt.Errorf("ScanRegion: %w", err)
		}
		ms.ScanRegion = optional.New(sr)
	}

	return ms, nil
}

// toXML creates an XML element for MediaSide.
func (ms MediaSide) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	// Add ColorProcessing if present
	if ms.ColorProcessing != nil {
		elm.Children = append(elm.Children,
			optional.Get(ms.ColorProcessing).toXML(NsWSCN+":ColorProcessing"))
	}

	// Add Resolution if present
	if ms.Resolution != nil {
		elm.Children = append(elm.Children,
			optional.Get(ms.Resolution).toXML(NsWSCN+":Resolution"))
	}

	// Add ScanRegion if present
	if ms.ScanRegion != nil {
		elm.Children = append(elm.Children,
			optional.Get(ms.ScanRegion).toXML(NsWSCN+":ScanRegion"))
	}

	return elm
}
