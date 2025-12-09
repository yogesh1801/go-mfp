// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common definitions

package iana

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/goipp"
)

const (
	// MIN is the lower bound for attribute value range
	MIN = math.MinInt32

	// MAX is the upper bound for attribute value range
	MAX = math.MaxInt32
)

// DefAttr defines attribute syntax, range and members.
type DefAttr struct {
	// Syntax definition
	SetOf    bool        // 1SetOf attribute
	Min, Max int32       // Allowed range of values
	Tags     []goipp.Tag // Allowed value tags

	// Members used only with collection attributes
	// and lists its members.
	//
	// Attribute may either define its own members, or
	// borrow members from some other attributes or even
	// top-level collections. Sometimes, attribute may borrow
	// members from multiple donors, hence the slice of maps.
	Members []map[string]*DefAttr // Collection members
}

// IsCollection reports if attribute is collection or 1SetOf collection
func (def *DefAttr) IsCollection() bool {
	return def.Tags[0] == goipp.TagBeginCollection
}

// OOBTag returns IPP tag that attribute defines to represent the
// Out-of-Band Values.
//
// Out-of-Band values, like 'unknown', 'unsupported', and 'no-value',
// allow to represent the situation, when attribute is supported by
// the IPP object, but has no meaningful value.
//
// See [RFC8011, 5.1.1.] for details.
//
// This function returns goipp.TagZero, if attribute definition doesn't
// provide an OOB tag.
//
// [RFC8011, 5.1.1.]: https://datatracker.ietf.org/doc/html/rfc8011#section-5.1.1
func (def *DefAttr) OOBTag() goipp.Tag {
	for _, tag := range def.Tags {
		switch tag {
		case goipp.TagUnsupportedValue,
			goipp.TagDefault,
			goipp.TagUnknown,
			goipp.TagNoValue,
			goipp.TagNotSettable,
			goipp.TagDeleteAttr,
			goipp.TagAdminDefine:
			return tag
		}
	}

	return goipp.TagZero
}

// HasTag reports if attribute syntax allows the specified tag.
func (def *DefAttr) HasTag(tag goipp.Tag) bool {
	for _, tag2 := range def.Tags {
		if tag == tag2 {
			return true
		}
	}
	return false
}

// EqualSyntax reports if attributes, defined by def and def2
// have equal syntax.
func (def *DefAttr) EqualSyntax(def2 *DefAttr) bool {
	return def.SetOf == def2.SetOf &&
		def.Min == def2.Min &&
		def.Max == def2.Max &&
		slices.Equal(def.Tags, def2.Tags)
}

// Member returns attribute's member by name.
func (def *DefAttr) Member(name string) *DefAttr {
	for _, mbr := range def.Members {
		if def2 := mbr[name]; def2 != nil {
			// Attribute may donor the entire collection
			// that it belongs to, but at this case it is
			// not considered to be member of itself.
			//
			// Handle this case here.
			if def == def2 {
				return nil
			}
			return def2
		}
	}

	return nil
}

// String formats attribute syntax as string, for debugging
func (def *DefAttr) String() string {
	val := []string{}
	noval := []string{}

	for _, tag := range def.Tags {
		if tag.Type() == goipp.TypeVoid {
			noval = append(noval, tag.String())
			continue
		}

		name := tag.String()
		if tag == goipp.TagName {
			name = "name"
		}

		min, max := tag.Limits()

		switch {
		case def.Min == min && def.Max == max:
		case def.Min == MIN && def.Max == MAX:
		case def.Min == min || def.Min == MIN:
			name = fmt.Sprintf("%s(%d)", name, def.Max)
		case def.Max == max || def.Max == MAX:
			name = fmt.Sprintf("%s(%d:MAX)", name, def.Min)
		default:
			name = fmt.Sprintf("%s(%d:%d)", name, def.Min, def.Max)
		}
		val = append(val, name)
	}

	switch {
	case !def.SetOf:
		return strings.Join(append(val, noval...), " | ")
	case len(val) == 1:
		return "1setOf " + strings.Join(append(val, noval...), " | ")
	}

	s := "1setOf (" + strings.Join(val, " | ") + ")"
	if len(noval) > 0 {
		s += " | " + strings.Join(noval, " | ")
	}

	return s
}

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

func init() {
	for _, b := range borrowings {
		rcpt := LookupAttribute(b.recipient)
		assert.MustMsg(rcpt != nil, "recipient not found: %q", b.recipient)

		var members []map[string]*DefAttr
		if col := Collections[b.donor]; col != nil {
			members = []map[string]*DefAttr{col}
		} else {
			donor := LookupAttribute(b.donor)
			assert.MustMsg(donor != nil, "donor not found: %q", b.donor)
			members = donor.Members
		}

		assert.MustMsg(len(members) > 0,
			"donor has no members: %q", b.donor)

		rcpt.Members = append(rcpt.Members, members...)
	}
}

// borrowing represents relation between collection attributes,
// where borrowing.recipient attribute borrows members from the
// borrowing.donor attribute
//
// Both borrowing.recipient and borrowing.donor are full path
// to the attributes:
//
//	"Job Template/cover-back/media-col"
//	"Job Template/media-col"
type borrowing struct {
	recipient, donor string
}
