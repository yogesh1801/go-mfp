// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for Sharpness

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
)

func TestSharpness_RoundTrip(t *testing.T) {
	orig := Sharpness(
		ValWithOptions[int]{
			Text:        90,
			Override:    optional.New(BooleanElement("1")),
			UsedDefault: optional.New(BooleanElement("false")),
		},
	)

	elm := orig.toXML("wscn:Sharpness")
	decoded, err := decodeSharpness(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}
