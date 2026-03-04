// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Attribute promotion policy

package iana

import (
	"github.com/OpenPrinting/goipp"
)

// promotionAllowed returns true, if attribute value with the 'from'
// tag can be used where the 'to' tag is expected.
func promotionAllowed(from, to goipp.Tag) bool {
	if from == to {
		return true
	}

	switch from {
	case goipp.TagText, goipp.TagTextLang:
		return to == goipp.TagText || to == goipp.TagTextLang

	case goipp.TagName, goipp.TagNameLang, goipp.TagKeyword:
		return to == goipp.TagName || to == goipp.TagNameLang ||
			to == goipp.TagURIScheme ||
			to == goipp.TagCharset ||
			to == goipp.TagLanguage

	case goipp.TagURIScheme, goipp.TagCharset, goipp.TagLanguage:
		return to == goipp.TagName || to == goipp.TagNameLang ||
			to == goipp.TagKeyword

	case goipp.TagMimeType:
		return to == goipp.TagName || to == goipp.TagNameLang
	}

	return false
}
