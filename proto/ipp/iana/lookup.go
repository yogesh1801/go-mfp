// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
//

package iana

import (
	"strings"
)

// LookupAttribute returns attribute by its full path.
// The full path looks as follows:
//
//	"Job Template/cover-back/media-col"
func LookupAttribute(path string) *DefAttr {
	if exceptions.Contains(path) {
		return nil
	}

	splitPath := strings.Split(path, "/")
	if len(splitPath) < 2 {
		return nil
	}

	col := Collections[splitPath[0]]
	if col == nil {
		return nil
	}

	def := col[splitPath[1]]
	for _, next := range splitPath[2:] {
		def = def.Member(next)
	}

	return def
}

// LookupSet searches attribute by name within set of members.
func LookupSet(set []map[string]*DefAttr, name string) *DefAttr {
	for _, m := range set {
		if def := m[name]; def != nil {
			return def
		}
	}
	return nil
}
