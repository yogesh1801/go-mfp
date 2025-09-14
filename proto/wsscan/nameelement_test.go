// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ConditionName

package wsscan

import (
	"errors"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestConditionName(t *testing.T) {
	type testData struct {
		in  string
		err string
	}

	tests := []testData{
		{
			in:  "Calibrating",
			err: "",
		},
		{
			in:  "CoverOpen",
			err: "",
		},
		{
			in:  "InputTrayEmpty",
			err: "",
		},
		{
			in:  "InterlockOpen",
			err: "",
		},
		{
			in:  "InternalStorageFull",
			err: "",
		},
		{
			in:  "LampError",
			err: "",
		},
		{
			in:  "LampWarming",
			err: "",
		},
		{
			in:  "MediaJam",
			err: "",
		},
		{
			in:  "MultipleFeedError",
			err: "",
		},
		{
			in:  "CustomCondition",
			err: "",
		},
		{
			in:  "valid-name",
			err: "",
		},
		{
			in:  "valid.name",
			err: "",
		},
		{
			in:  "valid:name",
			err: "",
		},
		{
			in:  "123valid",
			err: "",
		},
		{
			in:  "",
			err: `/wsscan:ConditionName: invalid ConditionName: ""`,
		},
		{
			in:  "aa bb",
			err: `/wsscan:ConditionName: invalid ConditionName: "aa bb"`,
		},
		{
			in:  "invalid@name",
			err: `/wsscan:ConditionName: invalid ConditionName: "invalid@name"`,
		},
		{
			in:  "invalid name",
			err: `/wsscan:ConditionName: invalid ConditionName: "invalid name"`,
		},
		{
			in:  "invalid#name",
			err: `/wsscan:ConditionName: invalid ConditionName: "invalid#name"`,
		},
		{
			in:  "invalid$name",
			err: `/wsscan:ConditionName: invalid ConditionName: "invalid$name"`,
		},
	}

	for _, test := range tests {
		xml := xmldoc.WithText("wsscan:ConditionName", test.in)
		reason, err := decodeConditionName(xml)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("error mismatch:\n"+
				"input:    %s\n"+
				"expected: %s\n"+
				"present:  %s\n",
				xml.EncodeString(nil), test.err, err)
			continue
		}

		if err.Error() == "" {
			encoded := reason.toXML("wsscan:ConditionName")
			if !encoded.Similar(xml) {
				t.Errorf("encode mismatch:\n"+
					"expected: %s\n"+
					"present:  %s\n",
					xml.EncodeString(nil),
					encoded.EncodeString(nil))
			}
		}
	}
}
