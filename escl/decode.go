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
	"strconv"

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
