// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// SettingProfile defines a valid combination of scanning parameters.

package escl

import (
	"github.com/alexpevzner/mfp/util/generic"
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// SettingProfile defines a valid combination of scanning parameters.
//
// eSCL Technical Specification, 8.1.2.
type SettingProfile struct {
	ColorModes           []ColorMode          // Supported color modes
	ContentTypes         []ContentType        // Supported content types
	DocumentFormats      []string             // MIME types of supported formats
	DocumentFormatsExt   []string             // eSCL 2.1+
	SupportedResolutions SupportedResolutions // Supported resolutions
	ColorSpaces          []ColorSpace         // Supported color spaces
	CCDChannels          []CCDChannel         // Supported CCD channels
	BinaryRenderings     []BinaryRendering    // Supported bin renderings
}

// decodeSettingProfile decodes [SettingProfile] from the XML tree.
func decodeSettingProfile(root xmldoc.Element) (
	prof SettingProfile, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	colormodes := xmldoc.Lookup{Name: NsScan + ":ColorModes"}
	contenttypes := xmldoc.Lookup{Name: NsScan + ":ContentTypes"}
	formats := xmldoc.Lookup{Name: NsScan + ":DocumentFormats"}
	resolutions := xmldoc.Lookup{Name: NsScan + ":SupportedResolutions",
		Required: true}
	clrSpaces := xmldoc.Lookup{Name: NsScan + ":ColorSpaces"}
	ccdChannels := xmldoc.Lookup{Name: NsScan + ":CcdChannels"}
	binrend := xmldoc.Lookup{Name: NsScan + ":BinaryRenderings"}

	missed := root.Lookup(&colormodes, &contenttypes, &formats,
		&resolutions, &clrSpaces, &ccdChannels, &binrend)
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

	if contenttypes.Found {
		for _, elem := range contenttypes.Elem.Children {
			if elem.Name == NsPWG+":ContentType" {
				var ct ContentType
				ct, err = decodeContentType(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						contenttypes.Elem, err)
					return
				}

				prof.ContentTypes = append(
					prof.ContentTypes, ct)
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
		for _, elem := range clrSpaces.Elem.Children {
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
		for _, elem := range ccdChannels.Elem.Children {
			if elem.Name == NsScan+":CcdChannel" {
				var chn CCDChannel
				chn, err = decodeCCDChannel(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						ccdChannels.Elem, err)
					return
				}

				prof.CCDChannels = append(prof.CCDChannels, chn)
			}
		}
	}

	if binrend.Found {
		for _, elem := range binrend.Elem.Children {
			if elem.Name == NsScan+":BinaryRendering" {
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

// toXML generates XML tree for the [SettingProfile].
func (prof SettingProfile) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}
	var chld xmldoc.Element

	if prof.ColorModes != nil {
		chld = xmldoc.Element{Name: NsScan + ":ColorModes"}
		for _, cm := range prof.ColorModes {
			chld2 := cm.toXML(NsScan + ":ColorMode")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	if prof.ContentTypes != nil {
		chld = xmldoc.Element{Name: NsScan + ":ContentTypes"}
		for _, ct := range prof.ContentTypes {
			chld2 := ct.toXML(NsPWG + ":ContentType")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	if prof.DocumentFormats != nil || prof.DocumentFormatsExt != nil {
		// pwd:DocumentFormat and scan:DocumentFormatsExt both
		// goes as children of the scan:DocumentFormats element,
		// and our representation doesn't preserve the ordering
		// between them.
		//
		// So here we behave as most scanners do: put elements
		// with the same MIME type together.
		//
		// After that, we dump all remaining scan:DocumentFormatExt
		// elements, if any.
		chld = xmldoc.Element{Name: NsScan + ":DocumentFormats"}

		ext := generic.NewSetOf(prof.DocumentFormatsExt...)

		for _, fmt := range prof.DocumentFormats {
			chld2 := xmldoc.WithText(NsPWG+":DocumentFormat", fmt)
			chld.Children = append(chld.Children, chld2)

			if ext.Contains(fmt) {
				ext.Del(fmt)
				chld2 = xmldoc.WithText(
					NsScan+":DocumentFormatExt", fmt)
				chld.Children = append(chld.Children, chld2)
			}
		}

		for _, fmt := range prof.DocumentFormatsExt {
			if ext.Contains(fmt) {
				chld2 := xmldoc.WithText(
					NsScan+":DocumentFormatExt", fmt)
				chld.Children = append(chld.Children, chld2)
			}
		}

		elm.Children = append(elm.Children, chld)
	}

	chld = prof.SupportedResolutions.toXML(NsScan + ":SupportedResolutions")
	elm.Children = append(elm.Children, chld)

	if prof.ColorSpaces != nil {
		chld = xmldoc.Element{Name: NsScan + ":ColorSpaces"}
		for _, sps := range prof.ColorSpaces {
			chld2 := sps.toXML(NsScan + ":ColorSpace")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	if prof.CCDChannels != nil {
		chld = xmldoc.Element{Name: NsScan + ":CcdChannels"}
		for _, chn := range prof.CCDChannels {
			chld2 := chn.toXML(NsScan + ":CcdChannel")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	if prof.BinaryRenderings != nil {
		chld = xmldoc.Element{Name: NsScan + ":BinaryRenderings"}
		for _, rnd := range prof.BinaryRenderings {
			chld2 := rnd.toXML(NsScan + ":BinaryRendering")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	return elm
}
