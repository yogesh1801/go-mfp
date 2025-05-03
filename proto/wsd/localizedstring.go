// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Localized strings

package wsd

import (
	"strings"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// LocalizedString represents a string with language
type LocalizedString struct {
	String string // String body
	Lang   string // ISO language code
}

// decodeLocalizedString decodes LocalizedString from the XML
func decodeLocalizedString(root xmldoc.Element) LocalizedString {
	ls := LocalizedString{String: root.Text}
	if attr, found := root.AttrByName("xml:lang"); found {
		ls.Lang = attr.Value
	}
	return ls
}

// ToXML generates XML for the LocalizedString
func (ls LocalizedString) ToXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name, Text: ls.String}
	if ls.Lang != "" {
		elm.Attrs = []xmldoc.Attr{{Name: "xml:lang", Value: ls.Lang}}
	}
	return elm
}

// IsZero reports if LocalizedString has zero value
func (ls LocalizedString) IsZero() bool {
	return ls == LocalizedString{}
}

// LocalizedStringList represents a list of localized strings
type LocalizedStringList []LocalizedString

// Contains contains reports if [LocalizedStringList] contains
// the specified [LocalizedString].
func (lsl LocalizedStringList) Contains(s LocalizedString) bool {
	for _, ls := range lsl {
		if s == ls {
			return true
		}
	}

	return false
}

// NeutralLang returns a neutral-language version of the localized
// string.
//
// It uses the following preferences:
//   - version without Lang is the best match
//   - if not found, search for the "en" version
//   - if not found, search for the "en-US"
//   - if not found, search for the first entry, starting with "en-"
//   - if not found yet, return the first entry from the list
//
// If list is empty, it returns LocalizedString{}
func (lsl LocalizedStringList) NeutralLang() LocalizedString {
	var en, enUS, enAny LocalizedString

	for _, ls := range lsl {
		lang := strings.ToLower(ls.Lang)
		switch lang {
		case "":
			return ls // The best match; return immediately
		case "en":
			if en.IsZero() {
				en = ls
			}
		case "en-us":
			if enUS.IsZero() {
				enUS = ls
			}
		default:
			if strings.HasPrefix(lang, "en-") && enAny.IsZero() {
				enAny = ls
			}
		}
	}

	switch {
	case !en.IsZero():
		return en
	case !enUS.IsZero():
		return enUS
	case !enAny.IsZero():
		return enAny
	case len(lsl) != 0:
		return lsl[0]
	}

	return LocalizedString{}
}
