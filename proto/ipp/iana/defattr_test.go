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

// TestHasTag tests DefAttr.HasTag method
func TestHasTag(t *testing.T) {
	type testData struct {
		tags     []goipp.Tag
		tag      goipp.Tag
		expected bool
	}

	tests := []testData{
		{
			tags:     []goipp.Tag{goipp.TagInteger},
			tag:      goipp.TagInteger,
			expected: true,
		},

		{
			tags:     []goipp.Tag{goipp.TagKeyword, goipp.TagName},
			tag:      goipp.TagKeyword,
			expected: true,
		},

		{
			tags:     []goipp.Tag{goipp.TagKeyword, goipp.TagName},
			tag:      goipp.TagName,
			expected: true,
		},

		{
			tags:     []goipp.Tag{goipp.TagKeyword, goipp.TagName},
			tag:      goipp.TagEnum,
			expected: false,
		},
	}

	for _, test := range tests {
		def := &DefAttr{Tags: test.tags}
		answer := def.HasTag(test.tag)

		if answer != test.expected {
			t.Errorf("DefAttr.HasTag:\n"+
				"input:    %v vs %#v\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.tag, test.tags, test.expected, answer)
		}
	}
}

// TestString tests DefAttr.String method
func TestString(t *testing.T) {
	type testData struct {
		def *DefAttr
		str string
	}

	tests := []testData{
		{
			def: &DefAttr{
				SetOf: false,
				Min:   MIN,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			str: "integer",
		},

		{
			def: &DefAttr{
				SetOf: false,
				Min:   1,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			str: "keyword",
		},

		{
			def: &DefAttr{
				SetOf: false,
				Min:   MIN,
				Max:   1234,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			str: "integer(1234)",
		},

		{
			def: &DefAttr{
				SetOf: false,
				Min:   -1234,
				Max:   MAX,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			str: "integer(-1234:MAX)",
		},

		{
			def: &DefAttr{
				SetOf: false,
				Min:   5,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagKeyword},
			},
			str: "keyword(5:MAX)",
		},

		{
			def: &DefAttr{
				SetOf: false,
				Min:   -1234,
				Max:   1234,
				Tags:  []goipp.Tag{goipp.TagInteger},
			},
			str: "integer(-1234:1234)",
		},

		{
			def: &DefAttr{
				SetOf: false,
				Min:   1,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNoValue},
			},
			str: "keyword | no-value",
		},

		{
			def: &DefAttr{
				SetOf: true,
				Min:   1,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagNoValue},
			},
			str: "1setOf keyword | no-value",
		},

		{
			def: &DefAttr{
				SetOf: true,
				Min:   1,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagName, goipp.TagNoValue},
			},
			str: "1setOf (keyword | name(1:MAX)) | no-value",
		},

		{
			def: &DefAttr{
				SetOf: true,
				Min:   1,
				Max:   255,
				Tags:  []goipp.Tag{goipp.TagKeyword, goipp.TagName},
			},
			str: "1setOf (keyword | name(1:MAX))",
		},
	}

	for _, test := range tests {
		str := test.def.String()
		if str != test.str {
			t.Errorf("DefAttr.String:\n"+
				"expected: %q\n"+
				"present:  %q\n",
				test.str, str)
		}
	}
}
