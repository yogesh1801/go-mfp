// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CUPS requests and responses

package ipp

import (
	"os"
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
		rq     Request      // Request structure
		groups goipp.Groups // Its IPP representation
	}

	tests := []testData{
		{
			rq: &CUPSGetDefaultRequest{
				AttributesCharset:         DefaultCharset,
				AttributesNaturalLanguage: DefaultNaturalLanguage,
				RequestedAttributes:       []string{"all"},
			},

			groups: []goipp.Group{
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
	}

	for _, test := range tests {
		msg := test.rq.Encode()
		msg.Print(os.Stdout, true)
	}
}

func TestCupsRequesponses(t *testing.T) {
	type testData struct {
		rsp Response
	}

	tests := []testData{
		{
			rsp: &CUPSGetDefaultResponse{
				AttributesCharset:         DefaultCharset,
				AttributesNaturalLanguage: DefaultNaturalLanguage,
				StatusMessage:             "success",
			},
		},
	}

	for _, test := range tests {
		msg := test.rsp.Encode()

		msg.Print(os.Stdout, false)
	}
}
