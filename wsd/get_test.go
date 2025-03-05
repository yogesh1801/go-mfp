// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Get test

package wsd

import (
	"testing"

	"github.com/alexpevzner/mfp/util/xmldoc"
)

// TestGet tests Get encoding and decoding
func TestGet(t *testing.T) {
	xml := Get{}.ToXML()
	if !xml.IsZero() {
		t.Errorf("Get.ToXML: unexpected output: %#v", xml)
	}

	_, err := DecodeGet(xmldoc.Element{})
	if err != nil {
		t.Errorf("DecodeGet: %s", err)
	}
}
