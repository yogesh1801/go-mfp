// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for Contrast

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
)

func TestContrast_RoundTrip(t *testing.T) {
	orig := Contrast(
		ValWithOptions[int]{
			Text:        75,
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
	)

	elm := orig.toXML("wscn:Contrast")
	decoded, err := decodeContrast(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}
