// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan intent test

package escl

import "testing"

// TestIntentString tests Intent.String
func TestIntentString(t *testing.T) {
	type testData struct {
		intent Intent
		s      string
	}

	tests := []testData{
		{IntentDocument, "Document"},
		{IntentTextAndGraphic, "TextAndGraphic"},
		{IntentPhoto, "Photo"},
		{IntentPreview, "Preview"},
		{IntentObject, "Object"},
		{IntentBusinessCard, "BusinessCard"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.intent.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.intent), test.s, s)
		}
	}
}

// TestDecodeIntent tests DecodeIntent
func TestDecodeIntent(t *testing.T) {
	type testData struct {
		intent Intent
		s      string
	}

	tests := []testData{
		{IntentDocument, "Document"},
		{IntentTextAndGraphic, "TextAndGraphic"},
		{IntentPhoto, "Photo"},
		{IntentPreview, "Preview"},
		{IntentObject, "Object"},
		{IntentBusinessCard, "BusinessCard"},
		{IntentUnknown, "XXX"},
	}

	for _, test := range tests {
		intent := DecodeIntent(test.s)
		if intent != test.intent {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.intent, intent)
		}
	}
}
