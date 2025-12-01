// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common definitions

package iana

import (
	"math"
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

// Attribute defines attribute syntax, range and members.
type Attribute struct {
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
	Members []map[string]*Attribute // Collection members
}

// IsCollection reports if attribute is collection or 1SetOf collection
func (attr *Attribute) IsCollection() bool {
	return attr.Tags[0] == goipp.TagBeginCollection
}

// OOBTag returns IPP tag that Attribute defines to represent the
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
func (attr *Attribute) OOBTag() goipp.Tag {
	for _, tag := range attr.Tags {
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

// Member returns Attribute's member by name.
func (attr *Attribute) Member(name string) *Attribute {
	for _, mbr := range attr.Members {
		if attr2 := mbr[name]; attr2 != nil {
			// Attribute may donor the entire collection
			// that it belongs to, but at this case it is
			// not considered to be member of itself.
			//
			// Handle this case here.
			if attr == attr2 {
				return nil
			}
			return attr2
		}
	}

	return nil
}

// LookupAttribute returns Attribute by its full path.
// The full path looks as follows:
//
//	"Job Template/cover-back/media-col"
func LookupAttribute(path string) *Attribute {
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

	attr := col[splitPath[1]]
	for _, next := range splitPath[2:] {
		attr = attr.Member(next)
	}

	return attr
}

func init() {
	for _, b := range borrowings {
		rcpt := LookupAttribute(b.recipient)
		assert.MustMsg(rcpt != nil, "recipient not found: %q", b.recipient)

		var members []map[string]*Attribute
		if col := Collections[b.donor]; col != nil {
			members = []map[string]*Attribute{col}
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
