// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Test for DeviceSettings
func TestDeviceSettings_XMLRoundTrip(t *testing.T) {
	ds := DeviceSettings{
		AutoExposureSupported:             "true",
		BrightnessSupported:               "true",
		CompressionQualityFactorSupported: CompressionQualityFactorSupported{1, 100},
		ContentTypesSupported:             ContentTypesSupported{Values: []ContentTypeValue{Auto}},
		ContrastSupported:                 "true",
		DocumentSizeAutoDetectSupported:   "true",
		FormatsSupported:                  FormatsSupported{Values: []FormatValue{PNG}},
		RotationsSupported:                RotationsSupported{Values: []RotationValue{Rotation0}},
		ScalingRangeSupported: ScalingRangeSupported{
			ScalingWidth:  ScalingWidth{1, 1000},
			ScalingHeight: ScalingHeight{1, 1000},
		},
	}
	elm := ds.toXML(NsWSCN + ":DeviceSettings")
	parsed, err := decodeDeviceSettings(elm)
	if err != nil {
		t.Errorf("decodeDeviceSettings: input %+v, unexpected error: %v",
			ds, err)
	}
	if !reflect.DeepEqual(parsed, ds) {
		t.Errorf("XML round-trip: expected %+v, got %+v", ds, parsed)
	}
}

func TestDeviceSettings_DecodeErrors(t *testing.T) {
	// Missing AutoExposureSupported
	elm := xmldoc.Element{Name: NsWSCN + ":DeviceSettings"}
	if _, err := decodeDeviceSettings(elm); err == nil {
		t.Error("decodeDeviceSettings: " +
			"expected error for missing AutoExposureSupported, got nil")
	}
}
