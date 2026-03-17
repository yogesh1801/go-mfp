// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for MediaSideImageInfo

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestMediaSideImageInfo_RoundTrip_Back verifies that a MediaSideImageInfo
// encoded as MediaBackImageInfo round-trips through XML identically.
func TestMediaSideImageInfo_RoundTrip_Back(t *testing.T) {
	orig := MediaSideImageInfo{
		BytesPerLine:  2550,
		NumberOfLines: 3300,
		PixelsPerLine: 850,
	}
	elm := orig.toXML(NsWSCN + ":MediaBackImageInfo")
	if elm.Name != NsWSCN+":MediaBackImageInfo" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":MediaBackImageInfo", elm.Name)
	}

	parsed, err := decodeMediaSideImageInfo(elm)
	if err != nil {
		t.Fatalf("decodeMediaSideImageInfo returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestMediaSideImageInfo_RoundTrip_Front verifies that a MediaSideImageInfo
// encoded as MediaFrontImageInfo round-trips through XML identically.
func TestMediaSideImageInfo_RoundTrip_Front(t *testing.T) {
	orig := MediaSideImageInfo{
		BytesPerLine:  0,
		NumberOfLines: 1,
		PixelsPerLine: 1,
	}
	elm := orig.toXML(NsWSCN + ":MediaFrontImageInfo")
	if elm.Name != NsWSCN+":MediaFrontImageInfo" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":MediaFrontImageInfo", elm.Name)
	}

	parsed, err := decodeMediaSideImageInfo(elm)
	if err != nil {
		t.Fatalf("decodeMediaSideImageInfo returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestMediaSideImageInfo_MissingRequired verifies that decoding an element
// missing a required child (e.g. PixelsPerLine) returns an error.
func TestMediaSideImageInfo_MissingRequired(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":MediaBackImageInfo",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":BytesPerLine", Text: "2550"},
			{Name: NsWSCN + ":NumberOfLines", Text: "3300"},
			// PixelsPerLine intentionally omitted
		},
	}
	_, err := decodeMediaSideImageInfo(elm)
	if err == nil {
		t.Error("expected error for missing PixelsPerLine, got nil")
	}
}

// TestMediaSideImageInfo_ZeroNumberOfLines verifies that a NumberOfLines
// value of 0 (below the allowed minimum of 1) is rejected.
func TestMediaSideImageInfo_ZeroNumberOfLines(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":MediaBackImageInfo",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":BytesPerLine", Text: "100"},
			{Name: NsWSCN + ":NumberOfLines", Text: "0"},
			{Name: NsWSCN + ":PixelsPerLine", Text: "100"},
		},
	}
	_, err := decodeMediaSideImageInfo(elm)
	if err == nil {
		t.Error("expected error for NumberOfLines=0, got nil")
	}
}

// TestMediaSideImageInfo_ZeroPixelsPerLine verifies that a PixelsPerLine
// value of 0 (below the allowed minimum of 1) is rejected.
func TestMediaSideImageInfo_ZeroPixelsPerLine(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":MediaBackImageInfo",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":BytesPerLine", Text: "100"},
			{Name: NsWSCN + ":NumberOfLines", Text: "100"},
			{Name: NsWSCN + ":PixelsPerLine", Text: "0"},
		},
	}
	_, err := decodeMediaSideImageInfo(elm)
	if err == nil {
		t.Error("expected error for PixelsPerLine=0, got nil")
	}
}
