// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Job state reason

package escl

import (
	"errors"
	"testing"

	"github.com/alexpevzner/mfp/util/xmldoc"
)

func TestJobStateReason(t *testing.T) {
	type testData struct {
		in  string
		err string
	}

	tests := []testData{
		{
			in:  "JobCanceledByOperator",
			err: "",
		},

		{
			in:  "",
			err: `/pwg:JobStateReason: invalid JobStateReason: ""`,
		},

		{
			in:  "aa bb",
			err: `/pwg:JobStateReason: invalid JobStateReason: "aa bb"`,
		},
	}

	for _, test := range tests {
		xml := xmldoc.WithText(NsPWG+":JobStateReason", test.in)
		reason, err := decodeJobStateReason(xml)
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
			encoded := reason.toXML(NsPWG + ":JobStateReason")
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
