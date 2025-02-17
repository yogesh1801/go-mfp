// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package escl

import "github.com/alexpevzner/mfp/xmldoc"

// SettingProfile defines a valid combination of scanning parameters.
//
// eSCL Technical Specification, 8.1.2.
type SettingProfile struct {
	ColorModes           []ColorMode          // Supported color modes
	DocumentFormats      []string             // MIME types of supported formats
	DocumentFormatsExt   []string             // eSCL 2.1+
	SupportedResolutions SupportedResolutions // Supported resolutions
	ColorSpaces          []ColorSpace         // Supported color spaces
	CcdChannels          []CcdChannel         // Supported CCD channels
	BinaryRenderings     []BinaryRendering    // Supported bin renderings
}

// decodeSettingProfile decodes [SettingProfile] from the XML tree.
func decodeSettingProfile(root xmldoc.Element) (
	prof SettingProfile, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	colormodes := xmldoc.Lookup{Name: NsScan + ":ColorModes"}
	formats := xmldoc.Lookup{Name: NsScan + ":DocumentFormats"}
	resolutions := xmldoc.Lookup{Name: NsScan + ":SupportedResolutions",
		Required: true}
	clrSpaces := xmldoc.Lookup{Name: NsScan + ":ColorSpaces"}
	ccdChannels := xmldoc.Lookup{Name: NsScan + ":CcdChannels"}
	binrend := xmldoc.Lookup{Name: NsScan + ":BinaryRenderings"}

	missed := root.Lookup(&colormodes, &formats, &resolutions,
		&clrSpaces, &ccdChannels, &binrend)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	if colormodes.Found {
		for _, elem := range colormodes.Elem.Children {
			if elem.Name == NsScan+":ColorMode" {
				var cm ColorMode
				cm, err = decodeColorMode(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						colormodes.Elem, err)
					return
				}

				prof.ColorModes = append(prof.ColorModes, cm)
			}
		}
	}

	if formats.Found {
		for _, elem := range formats.Elem.Children {
			switch elem.Name {
			case NsPWG + ":DocumentFormat":
				prof.DocumentFormats = append(
					prof.DocumentFormats, elem.Text)
			case NsScan + ":DocumentFormatExt":
				prof.DocumentFormatsExt = append(
					prof.DocumentFormatsExt, elem.Text)
			}
		}
	}

	prof.SupportedResolutions, err = decodeSupportedResolutions(
		resolutions.Elem)
	if err != nil {
		return
	}

	if clrSpaces.Found {
		for _, elem := range formats.Elem.Children {
			if elem.Name == NsScan+":ColorSpace" {
				var sps ColorSpace
				sps, err = decodeColorSpace(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						clrSpaces.Elem, err)
					return
				}

				prof.ColorSpaces = append(prof.ColorSpaces, sps)
			}
		}
	}

	if ccdChannels.Found {
		for _, elem := range formats.Elem.Children {
			if elem.Name == NsScan+":CcdChannel" {
				var sps CcdChannel
				sps, err = decodeCcdChannel(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						ccdChannels.Elem, err)
					return
				}

				prof.CcdChannels = append(prof.CcdChannels, sps)
			}
		}
	}

	if binrend.Found {
		for _, elem := range formats.Elem.Children {
			if elem.Name == NsScan+":BinaryRenderings" {
				var rns BinaryRendering
				rns, err = decodeBinaryRendering(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						binrend.Elem, err)
					return
				}

				prof.BinaryRenderings = append(
					prof.BinaryRenderings, rns)
			}
		}
	}

	return
}
