// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Film: describes the capabilities of the film scanning option

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Film describes the capabilities of the film scanning option
// attached to the scanner. It is an optional element.
type Film struct {
	FilmColor              ColorEntry
	FilmMaximumSize        Dimensions
	FilmMinimumSize        Dimensions
	FilmOpticalResolution  Dimensions
	FilmResolutions        Resolutions
	FilmScanModesSupported []FilmScanMode
}

// toXML creates an XML element for Film.
func (f Film) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	elm.Children = append(elm.Children,
		f.FilmColor.toXML(NsWSCN+":FilmColor"))

	elm.Children = append(elm.Children,
		f.FilmMaximumSize.toXML(NsWSCN+":FilmMaximumSize"))

	elm.Children = append(elm.Children,
		f.FilmMinimumSize.toXML(NsWSCN+":FilmMinimumSize"))

	elm.Children = append(elm.Children,
		f.FilmOpticalResolution.toXML(NsWSCN+":FilmOpticalResolution"))

	elm.Children = append(elm.Children,
		f.FilmResolutions.toXML(NsWSCN+":FilmResolutions"))

	if len(f.FilmScanModesSupported) > 0 {
		modes := xmldoc.Element{Name: NsWSCN + ":FilmScanModesSupported"}
		for _, mode := range f.FilmScanModesSupported {
			modes.Children = append(modes.Children,
				mode.toXML(NsWSCN+":FilmScanModeValue"))
		}
		elm.Children = append(elm.Children, modes)
	}

	return elm
}

// decodeFilm decodes a Film from an XML element.
func decodeFilm(root xmldoc.Element) (Film, error) {
	var f Film

	filmColor := xmldoc.Lookup{
		Name:     NsWSCN + ":FilmColor",
		Required: true,
	}
	filmMaxSize := xmldoc.Lookup{
		Name:     NsWSCN + ":FilmMaximumSize",
		Required: true,
	}
	filmMinSize := xmldoc.Lookup{
		Name:     NsWSCN + ":FilmMinimumSize",
		Required: true,
	}
	filmOptRes := xmldoc.Lookup{
		Name:     NsWSCN + ":FilmOpticalResolution",
		Required: true,
	}
	filmRes := xmldoc.Lookup{
		Name:     NsWSCN + ":FilmResolutions",
		Required: true,
	}
	filmModes := xmldoc.Lookup{
		Name:     NsWSCN + ":FilmScanModesSupported",
		Required: true,
	}

	missed := root.Lookup(&filmColor, &filmMaxSize, &filmMinSize,
		&filmOptRes, &filmRes, &filmModes)
	if missed != nil {
		return f, xmldoc.XMLErrMissed(missed.Name)
	}

	color, err := decodeColorEntry(filmColor.Elem)
	if err != nil {
		return f, fmt.Errorf("FilmColor: %w", err)
	}
	f.FilmColor = color

	maxSize, err := decodeDimensions(filmMaxSize.Elem)
	if err != nil {
		return f, fmt.Errorf("FilmMaximumSize: %w", err)
	}
	f.FilmMaximumSize = maxSize

	minSize, err := decodeDimensions(filmMinSize.Elem)
	if err != nil {
		return f, fmt.Errorf("FilmMinimumSize: %w", err)
	}
	f.FilmMinimumSize = minSize

	optRes, err := decodeDimensions(filmOptRes.Elem)
	if err != nil {
		return f, fmt.Errorf("FilmOpticalResolution: %w", err)
	}
	f.FilmOpticalResolution = optRes

	res, err := decodeResolutions(filmRes.Elem)
	if err != nil {
		return f, fmt.Errorf("FilmResolutions: %w", err)
	}
	f.FilmResolutions = res

	for _, child := range filmModes.Elem.Children {
		if child.Name != NsWSCN+":FilmScanModeValue" {
			continue
		}
		mode, err := decodeFilmScanMode(child)
		if err != nil {
			return f, fmt.Errorf("FilmScanModeValue: %w", err)
		}
		f.FilmScanModesSupported = append(f.FilmScanModesSupported, mode)
	}

	if len(f.FilmScanModesSupported) == 0 {
		return f, fmt.Errorf("FilmScanModesSupported: at least one mode is required")
	}

	return f, nil
}
