// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package escl

import (
	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/xmldoc"
)

// InputSourceCaps specifies capabilities of each input source
// (Platen, ADF and ADF Duplex).
//
// eSCL Technical Specification, 8.1.3.
type InputSourceCaps struct {
	MaxWidth              int               // Max scan width
	MinWidth              int               // Min scan width
	MaxHeight             int               // Max scan height
	MinHeight             int               // Min scan height
	MaxXOffset            optional.Val[int] // Max XOffset
	MaxYOffset            optional.Val[int] // Max YOffset
	MaxOpticalXResolution optional.Val[int] // Max optical X resolution
	MaxOpticalYResolution optional.Val[int] // Max optical Y resolution
	MaxScanRegions        optional.Val[int] // Max number of scan regions
	RiskyLeftMargins      optional.Val[int] // Risky left margins
	RiskyRightMargins     optional.Val[int] // Risky right margins
	RiskyTopMargins       optional.Val[int] // Risky top margins
	RiskyBottomMargins    optional.Val[int] // Risky bottom margins
	MaxPhysicalWidth      optional.Val[int] // Max physical width
	MaxPhysicalHeight     optional.Val[int] // Max physical height
	SupportedIntents      []Intent          // Supported intents
	EdgeAutoDetection     []SupportedEdge   // Supported edges detection
	SettingProfiles       []SettingProfile  // Supported scan profiles
	FeedDirections        []FeedDirection   // Available feed directions
}

// decodeInputSourceCaps decodes [InputSourceCaps] from the XML tree.
func decodeInputSourceCaps(root xmldoc.Element) (
	caps InputSourceCaps, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	maxWidth := xmldoc.Lookup{Name: NsScan + ":MaxWidth", Required: true}
	minWidth := xmldoc.Lookup{Name: NsScan + ":MinWidth", Required: true}
	maxHeight := xmldoc.Lookup{Name: NsScan + ":MaxHeight", Required: true}
	minHeight := xmldoc.Lookup{Name: NsScan + ":MinHeight", Required: true}
	maxXOff := xmldoc.Lookup{Name: NsScan + ":MaxXOffset"}
	maxYOff := xmldoc.Lookup{Name: NsScan + ":MaxYOffset"}
	maxOptXRes := xmldoc.Lookup{Name: NsScan + ":MaxOpticalXResolution"}
	maxOptYRes := xmldoc.Lookup{Name: NsScan + ":MaxOpticalYResolution"}
	maxRegs := xmldoc.Lookup{Name: NsScan + ":MaxScanRegions"}
	riskyLeft := xmldoc.Lookup{Name: NsScan + ":RiskyLeftMargins"}
	riskyRight := xmldoc.Lookup{Name: NsScan + ":RiskyRightMargins"}
	riskyTop := xmldoc.Lookup{Name: NsScan + ":RiskyTopMargins"}
	riskyBottom := xmldoc.Lookup{Name: NsScan + ":RiskyBottomMargins"}
	maxPhysWidth := xmldoc.Lookup{Name: NsScan + ":MaxPhysicalWidth"}
	maxPhysHeight := xmldoc.Lookup{Name: NsScan + ":MaxPhysicalHeight"}
	intents := xmldoc.Lookup{Name: NsScan + ":SupportedIntents"}
	edges := xmldoc.Lookup{Name: NsScan + ":EdgeAutoDetection"}
	profiles := xmldoc.Lookup{Name: NsScan + ":SettingProfiles"}
	feeds := xmldoc.Lookup{Name: NsScan + ":FeedDirections"}

	missed := root.Lookup(&maxWidth, &minWidth, &maxHeight, &minHeight,
		&maxXOff, &maxYOff, &maxOptXRes, &maxOptYRes, &maxRegs,
		&riskyLeft, &riskyRight, &riskyTop, &riskyBottom,
		&maxPhysWidth, &maxPhysHeight,
		&intents, &edges, &profiles, &feeds)

	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements (oh, there are a lot of them here...)
	caps.MaxWidth, err = decodeNonNegativeInt(maxWidth.Elem)
	if err == nil {
		caps.MinWidth, err = decodeNonNegativeInt(minWidth.Elem)
	}
	if err == nil {
		caps.MaxHeight, err = decodeNonNegativeInt(maxHeight.Elem)
	}
	if err == nil {
		caps.MinHeight, err = decodeNonNegativeInt(minHeight.Elem)
	}

	if err != nil {
		return
	}

	var tmp int

	if maxXOff.Found {
		tmp, err = decodeNonNegativeInt(maxXOff.Elem)
		if err != nil {
			return
		}

		caps.MaxXOffset = optional.New(tmp)
	}

	if maxYOff.Found {
		tmp, err = decodeNonNegativeInt(maxYOff.Elem)
		if err != nil {
			return
		}

		caps.MaxYOffset = optional.New(tmp)
	}

	if maxOptXRes.Found {
		tmp, err = decodeNonNegativeInt(maxOptXRes.Elem)
		if err != nil {
			return
		}

		caps.MaxOpticalXResolution = optional.New(tmp)
	}

	if maxOptYRes.Found {
		tmp, err = decodeNonNegativeInt(maxOptYRes.Elem)
		if err != nil {
			return
		}

		caps.MaxOpticalYResolution = optional.New(tmp)
	}

	if maxRegs.Found {
		tmp, err = decodeNonNegativeInt(maxRegs.Elem)
		if err != nil {
			return
		}

		caps.MaxScanRegions = optional.New(tmp)
	}

	if riskyLeft.Found {
		tmp, err = decodeNonNegativeInt(riskyLeft.Elem)
		if err != nil {
			return
		}

		caps.RiskyLeftMargins = optional.New(tmp)
	}

	if riskyRight.Found {
		tmp, err = decodeNonNegativeInt(riskyRight.Elem)
		if err != nil {
			return
		}

		caps.RiskyRightMargins = optional.New(tmp)
	}

	if riskyTop.Found {
		tmp, err = decodeNonNegativeInt(riskyTop.Elem)
		if err != nil {
			return
		}

		caps.RiskyTopMargins = optional.New(tmp)
	}

	if riskyBottom.Found {
		tmp, err = decodeNonNegativeInt(riskyBottom.Elem)
		if err != nil {
			return
		}

		caps.RiskyBottomMargins = optional.New(tmp)
	}

	if maxPhysWidth.Found {
		tmp, err = decodeNonNegativeInt(maxPhysWidth.Elem)
		if err != nil {
			return
		}

		caps.MaxPhysicalWidth = optional.New(tmp)
	}

	if maxPhysHeight.Found {
		tmp, err = decodeNonNegativeInt(maxPhysHeight.Elem)
		if err != nil {
			return
		}

		caps.MaxPhysicalHeight = optional.New(tmp)
	}

	if intents.Found {
		for _, elem := range intents.Elem.Children {
			if elem.Name == NsScan+":SupportedIntent" {
				var intent Intent
				intent, err = decodeIntent(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						intents.Elem, err)
					return
				}

				caps.SupportedIntents = append(
					caps.SupportedIntents, intent)
			}
		}
	}

	if edges.Found {
		for _, elem := range edges.Elem.Children {
			if elem.Name == NsScan+":SupportedEdge" {
				var edge SupportedEdge
				edge, err = decodeSupportedEdge(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						edges.Elem, err)
					return
				}

				caps.EdgeAutoDetection = append(
					caps.EdgeAutoDetection, edge)
			}
		}
	}

	if profiles.Found {
		for _, elem := range profiles.Elem.Children {
			if elem.Name == NsScan+":SupportedEdge" {
				var prof SettingProfile
				prof, err = decodeSettingProfile(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						profiles.Elem, err)
					return
				}

				caps.SettingProfiles = append(
					caps.SettingProfiles, prof)
			}
		}
	}

	if feeds.Found {
		for _, elem := range feeds.Elem.Children {
			if elem.Name == NsScan+":FeedDirection" {
				var feed FeedDirection
				feed, err = decodeFeedDirection(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(
						feeds.Elem, err)
					return
				}

				caps.FeedDirections = append(
					caps.FeedDirections, feed)
			}
		}
	}

	return
}
