// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ExposureSettings: contains individual adjustment values for image data

package wsscan

import (
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Brightness specifies the relative amount to reduce or enhance the brightness
// of the scanned document.
// The optional attributes Override and UsedDefault are Boolean values.
type Brightness struct {
	TextWithBoolAttrs[int]
}

// Contrast specifies the relative amount to reduce or enhance the contrast
// of the scanned document.
// The optional attributes Override and UsedDefault are Boolean values.
type Contrast struct {
	TextWithBoolAttrs[int]
}

// Sharpness specifies the relative amount to reduce or enhance the sharpness
// of the scanned document.
// The optional attributes Override and UsedDefault are Boolean values.
type Sharpness struct {
	TextWithBoolAttrs[int]
}

// ExposureSettings contains individual adjustment values that the WSD Scan Service
// should apply to the image data after acquisition.
type ExposureSettings struct {
	Brightness optional.Val[Brightness]
	Contrast   optional.Val[Contrast]
	Sharpness  optional.Val[Sharpness]
}

// Common encoder/decoder functions for int values
func intValueDecoder(s string) (int, error) {
	return strconv.Atoi(s)
}

func intValueEncoder(i int) string {
	return strconv.Itoa(i)
}

// decodeBrightness decodes a Brightness from an XML element.
func decodeBrightness(root xmldoc.Element) (Brightness, error) {
	var b Brightness
	decoded, err := b.TextWithBoolAttrs.decodeTextWithBoolAttrs(root, intValueDecoder)
	if err != nil {
		return b, err
	}
	b.TextWithBoolAttrs = decoded
	return b, nil
}

// toXML converts a Brightness to an XML element.
func (b Brightness) toXML(name string) xmldoc.Element {
	return b.TextWithBoolAttrs.toXML(name, intValueEncoder)
}

// decodeContrast decodes a Contrast from an XML element.
func decodeContrast(root xmldoc.Element) (Contrast, error) {
	var c Contrast
	decoded, err := c.TextWithBoolAttrs.decodeTextWithBoolAttrs(root, intValueDecoder)
	if err != nil {
		return c, err
	}
	c.TextWithBoolAttrs = decoded
	return c, nil
}

// toXML converts a Contrast to an XML element.
func (c Contrast) toXML(name string) xmldoc.Element {
	return c.TextWithBoolAttrs.toXML(name, intValueEncoder)
}

// decodeSharpness decodes a Sharpness from an XML element.
func decodeSharpness(root xmldoc.Element) (Sharpness, error) {
	var s Sharpness
	decoded, err := s.TextWithBoolAttrs.decodeTextWithBoolAttrs(root, intValueDecoder)
	if err != nil {
		return s, err
	}
	s.TextWithBoolAttrs = decoded
	return s, nil
}

// toXML converts a Sharpness to an XML element.
func (s Sharpness) toXML(name string) xmldoc.Element {
	return s.TextWithBoolAttrs.toXML(name, intValueEncoder)
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
		children = append(children, optional.Get(es.Brightness).toXML(NsWSCN+":Brightness"))
	}

	// Add Contrast if present
	if es.Contrast != nil {
		children = append(children, optional.Get(es.Contrast).toXML(NsWSCN+":Contrast"))
	}

	// Add Sharpness if present
	if es.Sharpness != nil {
		children = append(children, optional.Get(es.Sharpness).toXML(NsWSCN+":Sharpness"))
	}

	elm.Children = children
	return elm
}
