// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CUPS requests and responses

package ipp

import (
	"bytes"
	"testing"

	"github.com/OpenPrinting/goipp"
)

var (
	_ Request  = &CUPSGetDefaultRequest{}
	_ Response = &CUPSGetDefaultResponse{}
)

// TestCupsRequests tests CUPS requests
func TestCupsRequests(t *testing.T) {
	type testData struct {
		rq  Request       // Request structure
		msg goipp.Message // Its IPP representation
	}

	tests := []testData{
		{
			rq: &CUPSGetDefaultRequest{
				Version:   goipp.DefaultVersion,
				RequestID: 12345,

				AttributesCharset:         DefaultCharset,
				AttributesNaturalLanguage: DefaultNaturalLanguage,
				RequestedAttributes:       []string{"all"},
			},

			msg: goipp.Message{
				Version:   goipp.DefaultVersion,
				Code:      goipp.Code(goipp.OpCupsGetDefault),
				RequestID: 12345,

				Groups: []goipp.Group{
					{
						Tag: goipp.TagOperationGroup,
						Attrs: []goipp.Attribute{
							goipp.MakeAttribute(
								"attributes-charset",
								goipp.TagCharset,
								goipp.String(DefaultCharset)),
							goipp.MakeAttribute(
								"attributes-natural-language",
								goipp.TagLanguage,
								goipp.String(DefaultNaturalLanguage)),
							goipp.MakeAttribute(
								"requested-attributes",
								goipp.TagKeyword,
								goipp.String("all")),
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		msg := test.rq.Encode()
		if !msg.Similar(test.msg) {
			buf := &bytes.Buffer{}

			t.Errorf("Encode error. Message expected:")
			test.msg.Print(buf, true)
			t.Errorf("Message expected:\n%s", buf)

			buf.Reset()
			msg.Print(buf, true)
			t.Errorf("Message received:\n%s", buf)
		}
	}
}

// TestCupsRequests tests CUPS responses
func TestCupsRequesponses(t *testing.T) {
	type testData struct {
		rsp Response
		msg goipp.Message // Its IPP representation
	}

	tests := []testData{
		{
			rsp: &CUPSGetDefaultResponse{
				Version:   goipp.DefaultVersion,
				Status:    goipp.StatusOk,
				RequestID: 12345,

				AttributesCharset:         DefaultCharset,
				AttributesNaturalLanguage: DefaultNaturalLanguage,
				StatusMessage:             "success",
			},

			msg: goipp.Message{
				Version:   goipp.DefaultVersion,
				Code:      goipp.Code(goipp.StatusOk),
				RequestID: 12345,

				Groups: []goipp.Group{
					{
						Tag: goipp.TagOperationGroup,
						Attrs: []goipp.Attribute{
							goipp.MakeAttribute(
								"attributes-charset",
								goipp.TagCharset,
								goipp.String(DefaultCharset)),
							goipp.MakeAttribute(
								"attributes-natural-language",
								goipp.TagLanguage,
								goipp.String(DefaultNaturalLanguage)),
							goipp.MakeAttribute(
								"status-message",
								goipp.TagText,
								goipp.String("success")),
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		msg := test.rsp.Encode()
		if !msg.Similar(test.msg) {
			buf := &bytes.Buffer{}

			t.Errorf("Encode error. Message expected:")
			test.msg.Print(buf, false)
			t.Errorf("Message expected:\n%s", buf)

			buf.Reset()
			msg.Print(buf, false)
			t.Errorf("Message received:\n%s", buf)
		}
	}
}
