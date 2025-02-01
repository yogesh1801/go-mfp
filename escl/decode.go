// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Functions for XML decoding

package escl

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/alexpevzner/mfp/xmldoc"
)

// decodeNonNegativeInt decodes non-negative integer from the XML tree.
func decodeNonNegativeInt(root xmldoc.Element) (v int, err error) {
	var v64 uint64
	v64, err = strconv.ParseUint(root.Text, 10, 32)

	if err != nil || v > math.MaxInt32 {
		err = fmt.Errorf("invalid int: %q", root.Text)
		err = xmldoc.XMLErrWrap(root, err)
		v64 = 0
	}

	return int(v64), err
}

// decodeEnum decodes value of enum-alike type T from the XML tree
//
// decode is the type-specific function that decodes T from string
// (i.e., DecodeColorMode for ColorMode).
//
// ns is the XML namespace prefix (i.e., "scan" or "pwd").
func decodeEnum[T ~int](root xmldoc.Element,
	decode func(string) T, ns string) (val T, err error) {

	ns += ":"
	if strings.HasPrefix(root.Text, ns) {
		val = decode(root.Text[len(ns):])
		if val != 0 {
			return
		}
	}

	typeName := reflect.TypeOf(T(0)).String()
	if i := strings.LastIndexByte(typeName, '.'); i >= 0 {
		typeName = typeName[i+1:]
	}

	err = fmt.Errorf("invalid %s: %q", typeName, root.Text)
	err = xmldoc.XMLErrWrap(root, err)

	return
}
