// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// DefAttr methods test

package iana

import (
	"testing"

	"github.com/OpenPrinting/goipp"
)

// TestOOBTag tests DefAttr.OOBTag method
func TestOOBTag(t *testing.T) {
	type testData struct {
		tags     []goipp.Tag
		expected goipp.Tag
	}

	tests := []testData{
		{
			tags:     []goipp.Tag{goipp.TagName, goipp.TagUnsupportedValue},
			expected: goipp.TagUnsupportedValue,
		},
		{
			tags:     []goipp.Tag{goipp.TagText, goipp.TagDefault},
			expected: goipp.TagDefault,
		},
		{
			tags:     []goipp.Tag{goipp.TagKeyword, goipp.TagUnknown},
			expected: goipp.TagUnknown,
		},
		{
			tags:     []goipp.Tag{goipp.TagURI, goipp.TagNoValue},
			expected: goipp.TagNoValue,
		},
		{
			tags:     []goipp.Tag{goipp.TagLanguage, goipp.TagNotSettable},
			expected: goipp.TagNotSettable,
		},
		{
			tags:     []goipp.Tag{goipp.TagMimeType, goipp.TagDeleteAttr},
			expected: goipp.TagDeleteAttr,
		},
		{
			tags:     []goipp.Tag{goipp.TagDateTime, goipp.TagAdminDefine},
			expected: goipp.TagAdminDefine,
		},
		{
			tags:     []goipp.Tag{goipp.TagInteger},
			expected: goipp.TagZero,
		},
	}

	for _, test := range tests {
		def := &DefAttr{Tags: test.tags}
		oob := def.OOBTag()

		if oob != test.expected {
			t.Errorf("DefAttr.OOBTag:\n"+
				"input:    %#v\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.tags, test.expected, oob)
		}
	}
}
