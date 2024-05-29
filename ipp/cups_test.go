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
	"errors"
	"reflect"
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
		rq  Request        // Pointer to Request structure
		msg *goipp.Message // Its IPP representation
		err string         // expected decode error
	}

	tests := []testData{
		// ----- CUPSGetDefaultRequest tests -----
		{
			rq: &CUPSGetDefaultRequest{
				Version:   goipp.DefaultVersion,
				RequestID: 12345,

				AttributesCharset:         DefaultCharset,
				AttributesNaturalLanguage: DefaultNaturalLanguage,
				RequestedAttributes:       []string{"all"},
			},

			msg: goipp.NewMessageWithGroups(
				goipp.DefaultVersion,
				goipp.Code(goipp.OpCupsGetDefault),
				12345,
				goipp.Groups{
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
			),
		},

		{
			rq: &CUPSGetDefaultRequest{},

			msg: goipp.NewMessageWithGroups(
				goipp.DefaultVersion,
				goipp.Code(goipp.OpCupsGetDefault),
				12345,

				goipp.Groups{
					{
						Tag: goipp.TagOperationGroup,
						Attrs: []goipp.Attribute{
							goipp.MakeAttribute(
								"attributes-charset",
								goipp.TagInteger,
								goipp.Integer(111)),
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
			),

			err: `IPP decode ipp.CUPSGetDefaultRequest: "attributes-charset": can't convert integer to String`,
		},
	}

	for _, test := range tests {
		// Encode test
		if test.err == "" {
			msg := test.rq.Encode()
			if !msg.Similar(*test.msg) {
				buf := &bytes.Buffer{}

				t.Errorf("Encode error. Message expected:")
				test.msg.Print(buf, true)
				t.Errorf("Message expected:\n%s", buf)

				buf.Reset()
				msg.Print(buf, true)
				t.Errorf("Message received:\n%s", buf)
			}
		}

		// Decode test
		rq := reflect.
			New(reflect.TypeOf(test.rq).Elem()).
			Interface().(Request)

		err := rq.Decode(test.msg)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("Error mismatch:\n  <<< %s\n  >>> %s\n", test.err, err)
		} else if test.err == "" {
			diff := testDiffStruct(test.rq, rq)
			if diff != "" {
				t.Errorf("Decoded data doesn't match:\n%s", diff)
			}
		}
	}
}

// TestCupsRequests tests CUPS responses
func TestCupsRequesponses(t *testing.T) {
	type testData struct {
		rsp Response       // Pointer to Response structure
		msg *goipp.Message // Its IPP representation
		err string         // Expected decode error
	}

	tests := []testData{
		// ----- CUPSGetDefaultResponse tests -----
		{
			rsp: &CUPSGetDefaultResponse{
				Version:   goipp.DefaultVersion,
				Status:    goipp.StatusOk,
				RequestID: 12345,

				AttributesCharset:         DefaultCharset,
				AttributesNaturalLanguage: DefaultNaturalLanguage,
				StatusMessage:             "success",
			},

			msg: goipp.NewMessageWithGroups(
				goipp.DefaultVersion,
				goipp.Code(goipp.StatusOk),
				12345,

				goipp.Groups{
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
			),
		},

		{
			rsp: &CUPSGetDefaultResponse{},

			msg: goipp.NewMessageWithGroups(
				goipp.DefaultVersion,
				goipp.Code(goipp.StatusOk),
				12345,

				goipp.Groups{
					{
						Tag: goipp.TagOperationGroup,
						Attrs: []goipp.Attribute{
							goipp.MakeAttribute(
								"attributes-charset",
								goipp.TagCharset,
								goipp.String(DefaultCharset)),
							goipp.MakeAttribute(
								"attributes-natural-language",
								goipp.TagBoolean,
								goipp.Boolean(true)),
							goipp.MakeAttribute(
								"status-message",
								goipp.TagText,
								goipp.String("success")),
						},
					},
				},
			),

			err: `IPP decode ipp.CUPSGetDefaultResponse: "attributes-natural-language": can't convert boolean to String`,
		},
	}

	for _, test := range tests {
		// Encode test
		if test.err == "" {
			msg := test.rsp.Encode()
			if !msg.Similar(*test.msg) {
				buf := &bytes.Buffer{}

				t.Errorf("Encode error. Message expected:")
				test.msg.Print(buf, false)
				t.Errorf("Message expected:\n%s", buf)

				buf.Reset()
				msg.Print(buf, false)
				t.Errorf("Message received:\n%s", buf)
			}
		}

		// Decode test
		rsp := reflect.
			New(reflect.TypeOf(test.rsp).Elem()).
			Interface().(Response)

		err := rsp.Decode(test.msg)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("Error mismatch:\n  <<< %s\n  >>> %s\n", test.err, err)
		} else if test.err == "" {
			diff := testDiffStruct(test.rsp, rsp)
			if diff != "" {
				t.Errorf("Decoded data doesn't match:\n%s", diff)
			}
		}
	}
}
