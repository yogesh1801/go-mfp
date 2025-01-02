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
)

// Version represents a protocol version using the following encoding:
//
// [byte 2][byte 2][byte 1][byte 0]
// [     major    ][     minor    ]
type Version uint32

// MakeVersion makes a [Version] from bajor and minor parts, i.e.:
//
//	v := MakeVersion(2,0)
func MakeVersion(major, minor int) Version {
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

	return MakeVersion(int(major), int(minor)), nil

ERROR:
	return 0, fmt.Errorf("%q: invalid eSCL version", s)
}

// Major returns major part of the [Version]
func (ver Version) Major() int {
	return int(ver >> 16)
}

// Minor returns minor part of the [Version]
func (ver Version) Minor() int {
	return int(ver & 0xffff)
}

// String returns string representation of the [Version] (e.g., "2.0")
func (ver Version) String() string {
	return strconv.Itoa(ver.Major()) + "." + strconv.Itoa(ver.Minor())
}
