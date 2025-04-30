// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Maybe type test

package ipp

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/goipp"
)

var (
	_ maybeCodecInterface = &Maybe[string]{}
	_ maybeCodecInterface = &Maybe[int]{}
)

func TestMaybe(t *testing.T) {
	type TestStruct struct {
		AttrHello   Maybe[string] `ipp:"attr-hello"`
		AttrNoValue Maybe[string] `ipp:"attr-no-value"`
		AttrMissed  Maybe[string] `ipp:"attr-missed"`
	}

	type testData struct {
		data  TestStruct
		attrs goipp.Attributes
	}

	codec := ippCodecMustGenerate(reflect.TypeOf(TestStruct{}))

	tests := []testData{
		{
			data: TestStruct{
				AttrHello:   MaybeSet("hello"),
				AttrNoValue: MaybeDel[string](goipp.TagNoValue),
			},

			attrs: goipp.Attributes{
				goipp.MakeAttribute("attr-hello",
					goipp.TagText, goipp.String("hello")),

				goipp.MakeAttribute("attr-no-value",
					goipp.TagNoValue, goipp.Void{}),
			},
		},
	}

	for _, test := range tests {
		attrs := codec.encodeAttrs(&test.data)

		diff := testDiffAttrs(test.attrs, attrs)
		if diff != "" {
			t.Errorf("encode: input/output mismatch:\n%s", diff)
		}

		var data TestStruct
		err := codec.decodeAttrs(&data, attrs)
		if err != nil {
			t.Errorf("decode: %s", err)
			continue
		}

		diff = testDiffStruct(&test.data, &data)
		if diff != "" {
			t.Errorf("decode: input/output mismatch:\n%s", diff)
		}
	}
}
