// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Camera capabilities tests.

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/xmldoc"
)

// TestCamera tests [Camera] to/from the XML conversion
func TestCamera(t *testing.T) {
	type testData struct {
		camera Camera
		xml    xmldoc.Element
	}

	tests := []testData{
		{
			camera: Camera{nil},
			xml:    xmldoc.WithChildren(NsScan + ":Camera"),
		},

		{
			camera: Camera{
				CameraInputCaps: optional.New(testInputSourceCaps),
			},
			xml: xmldoc.WithChildren(
				NsScan+":Camera",
				testInputSourceCaps.toXML(NsScan+":CameraInputCaps"),
			),
		},
	}

	for _, test := range tests {
		xml := test.camera.toXML(NsScan + ":Camera")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		camera, err := decodeCamera(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(camera, test.camera) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.camera, camera)
		}
	}
}

// TestCameraDecodeErrors tests [Camera] XML decode
// errors handling
func TestCameraDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		{
			xml: xmldoc.WithChildren(
				NsScan+":Camera",
				xmldoc.WithChildren(NsScan+":CameraInputCaps"),
			),
			err: `/scan:Camera/scan:CameraInputCaps/scan:MinWidth: missed`,
		},
	}

	for _, test := range tests {
		_, err := decodeCamera(test.xml)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("error mismatch:\n"+
				"input:    %s\n"+
				"expected: %q\n"+
				"present:  %q\n",
				test.xml.EncodeString(nil), test.err, err)
		}
	}
}
