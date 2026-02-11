// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Exposure: specifies the exposure settings of the document

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Exposure specifies the exposure settings of the document.
// It has an optional MustHonor boolean attribute and contains
// either AutoExposure or ExposureSettings child elements.
type Exposure struct {
	MustHonor        optional.Val[BooleanElement]
	AutoExposure     optional.Val[BooleanElement]
	ExposureSettings optional.Val[ExposureSettings]
}

// decodeExposure decodes an Exposure from an XML element.
func decodeExposure(root xmldoc.Element) (Exposure, error) {
	var exp Exposure

	// Decode MustHonor attribute if present
	if attr, found := root.AttrByName(NsWSCN + ":MustHonor"); found {
		val := BooleanElement(attr.Value)
		if err := val.Validate(); err != nil {
			return exp, err
		}
		exp.MustHonor = optional.New(val)
	}

	// Setup lookups for optional child elements
	autoExposureLookup := xmldoc.Lookup{
		Name: NsWSCN + ":AutoExposure",
	}
	exposureSettingsLookup := xmldoc.Lookup{
		Name: NsWSCN + ":ExposureSettings",
	}

	root.Lookup(&autoExposureLookup, &exposureSettingsLookup)

	// Decode AutoExposure if present
	if autoExposureLookup.Elem.Name != "" {
		autoExp, err := decodeBooleanElement(autoExposureLookup.Elem)
		if err != nil {
			return exp, err
		}
		exp.AutoExposure = optional.New(autoExp)
	}

	// Decode ExposureSettings if present
	if exposureSettingsLookup.Elem.Name != "" {
		expSettings, err := decodeExposureSettings(exposureSettingsLookup.Elem)
		if err != nil {
			return exp, err
		}
		exp.ExposureSettings = optional.New(expSettings)
	}

	return exp, nil
}

// toXML converts an Exposure to an XML element.
func (exp Exposure) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	// Add MustHonor attribute if present
	if exp.MustHonor != nil {
		elm.Attrs = []xmldoc.Attr{
			{
				Name:  NsWSCN + ":MustHonor",
				Value: string(optional.Get(exp.MustHonor)),
			},
		}
	}

	var children []xmldoc.Element

	// Add AutoExposure if present
	if exp.AutoExposure != nil {
		children = append(children, optional.Get(exp.AutoExposure).toXML(
			NsWSCN+":AutoExposure"))
	}

	// Add ExposureSettings if present
	if exp.ExposureSettings != nil {
		children = append(children, optional.Get(exp.ExposureSettings).toXML(
			NsWSCN+":ExposureSettings"))
	}

	elm.Children = children
	return elm
}
