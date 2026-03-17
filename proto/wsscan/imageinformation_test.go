// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ImageInformation

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
)

func createValidMediaSideImageInfo() MediaSideImageInfo {
	return MediaSideImageInfo{
		BytesPerLine:  2550,
		NumberOfLines: 3300,
		PixelsPerLine: 850,
	}
}

// TestImageInformation_RoundTrip_Both verifies that an ImageInformation with
// both MediaBackImageInfo and MediaFrontImageInfo present encodes to XML and
// decodes back to an identical value.
func TestImageInformation_RoundTrip_Both(t *testing.T) {
	orig := ImageInformation{
		MediaBackImageInfo:  optional.New(createValidMediaSideImageInfo()),
		MediaFrontImageInfo: optional.New(createValidMediaSideImageInfo()),
	}
	elm := orig.toXML(NsWSCN + ":ImageInformation")
	if elm.Name != NsWSCN+":ImageInformation" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":ImageInformation", elm.Name)
	}

	parsed, err := decodeImageInformation(elm)
	if err != nil {
		t.Fatalf("decodeImageInformation returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestImageInformation_RoundTrip_BackOnly verifies that an ImageInformation
// with only MediaBackImageInfo present encodes and decodes correctly.
func TestImageInformation_RoundTrip_BackOnly(t *testing.T) {
	orig := ImageInformation{
		MediaBackImageInfo: optional.New(createValidMediaSideImageInfo()),
	}
	elm := orig.toXML(NsWSCN + ":ImageInformation")

	parsed, err := decodeImageInformation(elm)
	if err != nil {
		t.Fatalf("decodeImageInformation returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestImageInformation_RoundTrip_FrontOnly verifies that an ImageInformation
// with only MediaFrontImageInfo present encodes and decodes correctly.
func TestImageInformation_RoundTrip_FrontOnly(t *testing.T) {
	orig := ImageInformation{
		MediaFrontImageInfo: optional.New(createValidMediaSideImageInfo()),
	}
	elm := orig.toXML(NsWSCN + ":ImageInformation")

	parsed, err := decodeImageInformation(elm)
	if err != nil {
		t.Fatalf("decodeImageInformation returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestImageInformation_RoundTrip_Empty verifies that an ImageInformation with
// neither side present (no optional children) encodes and decodes correctly.
func TestImageInformation_RoundTrip_Empty(t *testing.T) {
	orig := ImageInformation{}
	elm := orig.toXML(NsWSCN + ":ImageInformation")

	parsed, err := decodeImageInformation(elm)
	if err != nil {
		t.Fatalf("decodeImageInformation returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}
