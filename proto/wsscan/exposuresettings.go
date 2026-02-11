// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ExposureSettings: contains individual adjustment values for image data

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ExposureSettings contains individual adjustment values that the
// WSD Scan Service should apply to the image data after acquisition.
type ExposureSettings struct {
	Brightness optional.Val[Brightness]
	Contrast   optional.Val[Contrast]
	Sharpness  optional.Val[Sharpness]
}

// decodeExposureSettings decodes an ExposureSettings from an XML element.
func decodeExposureSettings(root xmldoc.Element) (ExposureSettings, error) {
	var es ExposureSettings

	// Setup lookups for optional child elements
	brightnessLookup := xmldoc.Lookup{
		Name: NsWSCN + ":Brightness",
	}
	contrastLookup := xmldoc.Lookup{
		Name: NsWSCN + ":Contrast",
	}
	sharpnessLookup := xmldoc.Lookup{
		Name: NsWSCN + ":Sharpness",
	}

	root.Lookup(&brightnessLookup, &contrastLookup, &sharpnessLookup)

	// Decode Brightness if present
	if brightnessLookup.Elem.Name != "" {
		brightness, err := decodeBrightness(brightnessLookup.Elem)
		if err != nil {
			return es, err
		}
		es.Brightness = optional.New(brightness)
	}

	// Decode Contrast if present
	if contrastLookup.Elem.Name != "" {
		contrast, err := decodeContrast(contrastLookup.Elem)
		if err != nil {
			return es, err
		}
		es.Contrast = optional.New(contrast)
	}

	// Decode Sharpness if present
	if sharpnessLookup.Elem.Name != "" {
		sharpness, err := decodeSharpness(sharpnessLookup.Elem)
		if err != nil {
			return es, err
		}
		es.Sharpness = optional.New(sharpness)
	}

	return es, nil
}

// toXML converts an ExposureSettings to an XML element.
func (es ExposureSettings) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}
	var children []xmldoc.Element

	// Add Brightness if present
	if es.Brightness != nil {
		children = append(children, optional.Get(es.Brightness).toXML(
			NsWSCN+":Brightness"))
	}

	// Add Contrast if present
	if es.Contrast != nil {
		children = append(children, optional.Get(es.Contrast).toXML(
			NsWSCN+":Contrast"))
	}

	// Add Sharpness if present
	if es.Sharpness != nil {
		children = append(children, optional.Get(es.Sharpness).toXML(
			NsWSCN+":Sharpness"))
	}

	elm.Children = children
	return elm
}
