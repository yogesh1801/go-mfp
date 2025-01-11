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
	RiskyLeftMargins      optional.Val[int] // Risky left margins
	RiskyRightMargins     optional.Val[int] // Risky right margins
	RiskyTopMargins       optional.Val[int] // Risky top margins
	RiskyBottomMargins    optional.Val[int] // Risky bottom margins
	MaxPhysicalWidth      optional.Val[int] // Max physical width
	MaxPhysicalHeight     optional.Val[int] // Max physical height
	SupportedIntents      Intents           // Supported intents
	EdgeAutoDetection     SupportedEdges    // Supported edges detection
	SettingProfiles       []SettingProfile  // Supported scan profiles
}
