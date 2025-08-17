// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color space

package escl

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ColorSpace defines the color space used for scanning.
//
// The Mopria eSCL specification doesn't provide detailed information
// of this type.
//
// So here we represent it as a string and list known values as
// string constants.
type ColorSpace string

// Known values for ColorSpace:
const (
	UnknownColorSpace ColorSpace = ""
	SRGB              ColorSpace = "sRGB"
)

// decodeColorSpace decodes [ColorSpace] from the XML tree.
func decodeColorSpace(root xmldoc.Element) (sps ColorSpace, err error) {
	var v string
	v, err = decodeNMTOKEN(root)

	if err != nil {
		err = fmt.Errorf("invalid ColorSpace: %q", root.Text)
		err = xmldoc.XMLErrWrap(root, err)
		return
	}

	sps = ColorSpace(v)
	return
}

// toXML generates XML tree for the [ColorSpace].
func (sps ColorSpace) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: sps.String(),
	}
}

// String returns a string representation of the [ColorSpace]
func (sps ColorSpace) String() string {
	return string(sps)
}

// DecodeColorSpace decodes [ColorSpace] out of its XML string representation.
func DecodeColorSpace(s string) ColorSpace {
	return ColorSpace(s)
}
