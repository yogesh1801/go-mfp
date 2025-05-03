// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner resolutions

package escl

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// SupportedResolutions defines the set of resolutions,
// supported by the scanner.
//
// It may optionally contain the ColorMode element, if supported resolutions
// depend on a color mode. At this case, [SettingProfile.SupportedResolutions]
// will contain one or more entries with specified ColorMode and may also
// contain a default entry with ColorMode unset.
//
// eSCL Technical Specification, 8.1.1.
type SupportedResolutions struct {
	ColorMode           optional.Val[ColorMode]       // If depends on color
	DiscreteResolutions DiscreteResolutions           // Discrete res
	ResolutionRange     optional.Val[ResolutionRange] // Res range
}

// decodeSupportedResolutions decodes [SupportedResolutions] from the
// XML tree.
func decodeSupportedResolutions(root xmldoc.Element) (
	supp SupportedResolutions, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	colormode := xmldoc.Lookup{Name: NsScan + ":ColorMode"}
	discrete := xmldoc.Lookup{Name: NsScan + ":DiscreteResolutions"}
	ranges := xmldoc.Lookup{Name: NsScan + ":ResolutionRange"}

	root.Lookup(&colormode, &discrete, &ranges)

	if !discrete.Found && !ranges.Found {
		err = fmt.Errorf("missed %s and %s", discrete.Name, ranges.Name)
		return
	}

	if colormode.Found {
		var cm ColorMode
		cm, err = decodeColorMode(colormode.Elem)
		if err != nil {
			err = xmldoc.XMLErrWrap(colormode.Elem, err)
			return
		}

		supp.ColorMode = optional.New(cm)
	}

	if discrete.Found {
		for _, elem := range discrete.Elem.Children {
			var res DiscreteResolution
			res, err = decodeDiscreteResolution(elem)
			if err != nil {
				err = xmldoc.XMLErrWrap(discrete.Elem, err)
				return
			}

			supp.DiscreteResolutions = append(
				supp.DiscreteResolutions, res)
		}
	}

	if ranges.Found {
		var rng ResolutionRange
		rng, err = decodeResolutionRange(ranges.Elem)
		if err != nil {
			return
		}
		supp.ResolutionRange = optional.New(rng)
	}

	return
}

// toXML generates XML tree for the [SupportedResolutions].
func (supp SupportedResolutions) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	if supp.ColorMode != nil {
		cm := *supp.ColorMode
		chld := cm.toXML(NsScan + ":ColorMode")
		elm.Children = append(elm.Children, chld)
	}

	if supp.DiscreteResolutions != nil {
		chld := xmldoc.Element{Name: NsScan + ":DiscreteResolutions"}
		for _, r := range supp.DiscreteResolutions {
			chld.Children = append(chld.Children,
				r.toXML(NsScan+":DiscreteResolution"))
		}
		elm.Children = append(elm.Children, chld)
	}

	if supp.ResolutionRange != nil {
		rng := *supp.ResolutionRange
		chld := rng.toXML(NsScan + ":ResolutionRange")
		elm.Children = append(elm.Children, chld)
	}

	return elm
}

// DiscreteResolutions define a set of discrete resolutions,
// supported by the scanner.
type DiscreteResolutions []DiscreteResolution

// ResolutionRange defines a set of resolutions range,
// supported by the scanner.
type ResolutionRange struct {
	XResolutionRange Range // Horizontal range
	YResolutionRange Range // Vertical range
}

// decodeResolutionRange, decodes [ResolutionRange] from the XML tree.
func decodeResolutionRange(root xmldoc.Element) (
	rng ResolutionRange, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	x := xmldoc.Lookup{Name: NsScan + ":XResolutionRange", Required: true}
	y := xmldoc.Lookup{Name: NsScan + ":YResolutionRange", Required: true}

	missed := root.Lookup(&x, &y)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	rng.XResolutionRange, err = decodeRange(x.Elem)
	if err == nil {
		rng.YResolutionRange, err = decodeRange(y.Elem)
	}

	return
}

// toXML generates XML tree for the [ResolutionRange].
func (rng ResolutionRange) toXML(name string) xmldoc.Element {
	x := rng.XResolutionRange.toXML(NsScan + ":XResolutionRange")
	y := rng.YResolutionRange.toXML(NsScan + ":YResolutionRange")

	return xmldoc.Element{
		Name:     name,
		Children: []xmldoc.Element{x, y},
	}
}

// DiscreteResolution defines a discrete resolution, supported by the scanner.
type DiscreteResolution struct {
	XResolution int // Horizontal resolution, DPI
	YResolution int // Vertical resolution, DPI
}

// decodeDiscreteResolution decodes [DiscreteResolution] from the
// XML tree.
func decodeDiscreteResolution(root xmldoc.Element) (
	res DiscreteResolution, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	x := xmldoc.Lookup{Name: NsScan + ":XResolution", Required: true}
	y := xmldoc.Lookup{Name: NsScan + ":YResolution", Required: true}

	missed := root.Lookup(&x, &y)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	res.XResolution, err = decodeNonNegativeInt(x.Elem)
	if err == nil {
		res.YResolution, err = decodeNonNegativeInt(y.Elem)
	}

	return
}

// toXML generates XML tree for the [DiscreteResolution].
func (res DiscreteResolution) toXML(name string) xmldoc.Element {
	x := strconv.FormatUint(uint64(res.XResolution), 10)
	y := strconv.FormatUint(uint64(res.YResolution), 10)

	elm := xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{
				Name: NsScan + ":" + "XResolution",
				Text: x,
			},
			{
				Name: NsScan + ":" + "YResolution",
				Text: y,
			},
		},
	}

	return elm
}
