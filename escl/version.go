// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Protocol version

package escl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexpevzner/mfp/util/xmldoc"
)

// DefaultVersion is the default [Version] used by [AbstractServer].
var DefaultVersion = MakeVersion(2, 9)

// Version represents the eSCL protocol version.
//
// eSCL uses the following syntax for version "2.85" and it should be
// interpreted as a decimal fraction. Hence, 2.35 is greater that 2.3
// and less that 2.4
//
// We use the following representation:
//
//	[byte 2][byte 2][byte 1][byte 0]
//	[     major    ][     minor    ]
//
// The minor part is normalized to fit the 0...9999 rage
type Version uint32

// MakeVersion makes a [Version] from bajor and minor parts, i.e.:
//
//	MakeVersion(2,0)   -> "2.0"
//	MakeVersion(2,12)  -> "2.12"
//	MakeVersion(2,123) -> "2.123"
func MakeVersion(major, minor int) Version {
	switch {
	case minor < 10:
		minor *= 1000
	case minor < 100:
		minor *= 100
	case minor < 1000:
		minor *= 10
	default:
		for minor >= 10000 {
			minor /= 10
		}
	}

	ver := Version(major<<16) & 0xffff0000
	ver += Version(minor & 0xffff)

	return ver
}

// DecodeVersion decodes version out of its XML string representation.
func DecodeVersion(s string) (Version, error) {
	var major, minor uint64
	var err error

	dot := strings.IndexByte(s, '.')
	if dot == -1 {
		goto ERROR
	}

	major, err = strconv.ParseUint(s[:dot], 10, 16)
	if err != nil {
		goto ERROR
	}

	minor, err = strconv.ParseUint(s[dot+1:], 10, 16)
	if err != nil {
		goto ERROR
	}

	if minor >= 10000 {
		goto ERROR
	}

	return MakeVersion(int(major), int(minor)), nil

ERROR:
	return 0, fmt.Errorf("%q: invalid eSCL version", s)
}

// decodeVersion decodes version out of its XML string representation.
func decodeVersion(root xmldoc.Element) (ver Version, err error) {
	ver, err = DecodeVersion(root.Text)
	if err != nil {
		err = xmldoc.XMLErrWrap(root, err)
	}
	return
}

// toXML generates XML tree for the [Version].
func (ver Version) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: ver.String(),
	}
}

// Major returns major part of the [Version]
func (ver Version) Major() int {
	return int(ver >> 16)
}

// Minor returns minor part of the [Version]
func (ver Version) Minor() int {
	minor := int(ver & 0xffff)
	if minor != 0 {
		for minor%10 == 0 {
			minor /= 10
		}
	}
	return minor
}

// String returns string representation of the [Version] (e.g., "2.0")
func (ver Version) String() string {
	return strconv.Itoa(ver.Major()) + "." + strconv.Itoa(ver.Minor())
}
