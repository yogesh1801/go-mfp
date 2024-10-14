// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Functions for XML decoding

package wsd

import (
	"fmt"
	"strconv"

	"github.com/alexpevzner/mfp/xmldoc"
)

// decodeUint64 decodes uint64 from the XML tree
func decodeUint64(root xmldoc.Element) (v uint64, err error) {
	v, err = strconv.ParseUint(root.Text, 10, 64)
	if err != nil {
		err = fmt.Errorf("invalid uint: %q", root.Text)
		err = xmlErrWrap(root, err)
	}
	return
}

// decodeUint64 decodes uint64 from the XML attribute
func decodeUint64Attr(attr xmldoc.Attr) (v uint64, err error) {
	v, err = strconv.ParseUint(attr.Value, 10, 64)
	if err != nil {
		err = fmt.Errorf("invalid uint: %q", attr.Value)
		err = xmlErrWrapAttr(attr, err)
	}
	return
}
