// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan region.

package escl

import (
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScanRegion defines the desired scan region
type ScanRegion struct {
	XOffset            int   // Horizontal offset, 0-based
	YOffset            int   // Vertical offset, 0-based
	Width              int   // Region width
	Height             int   // Region height
	ContentRegionUnits Units // Always ThreeHundredthsOfInches
}

// decodeScanRegion decodes [ScanRegion] from the XML tree
func decodeScanRegion(root xmldoc.Element) (reg ScanRegion, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	xoff := xmldoc.Lookup{Name: NsPWG + ":XOffset", Required: true}
	yoff := xmldoc.Lookup{Name: NsPWG + ":YOffset", Required: true}
	wid := xmldoc.Lookup{Name: NsPWG + ":Width", Required: true}
	hei := xmldoc.Lookup{Name: NsPWG + ":Height", Required: true}
	units := xmldoc.Lookup{Name: NsPWG + ":ContentRegionUnits",
		Required: true}

	missed := root.Lookup(&xoff, &yoff, &wid, &hei, &units)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	reg.XOffset, err = decodeNonNegativeInt(xoff.Elem)
	if err == nil {
		reg.YOffset, err = decodeNonNegativeInt(yoff.Elem)
	}
	if err == nil {
		reg.Width, err = decodeNonNegativeInt(wid.Elem)
	}
	if err == nil {
		reg.Height, err = decodeNonNegativeInt(hei.Elem)
	}
	if err == nil {
		reg.ContentRegionUnits, err = decodeUnits(units.Elem)
	}

	return
}

// toXML generates XML tree for the [ScanRegion].
func (r ScanRegion) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{
				Name: NsPWG + ":XOffset",
				Text: strconv.FormatUint(uint64(r.XOffset), 10),
			},
			{
				Name: NsPWG + ":YOffset",
				Text: strconv.FormatUint(uint64(r.YOffset), 10),
			},
			{
				Name: NsPWG + ":Width",
				Text: strconv.FormatUint(uint64(r.Width), 10),
			},
			{
				Name: NsPWG + ":Height",
				Text: strconv.FormatUint(uint64(r.Height), 10),
			},
			ThreeHundredthsOfInches.toXML(
				NsPWG + ":ContentRegionUnits"),
		},
	}

	return elm
}
