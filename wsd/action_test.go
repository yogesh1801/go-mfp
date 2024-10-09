// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Message actions test

package wsd

import "testing"

// TestAction tests operations with action
func TestAction(t *testing.T) {
	type testData struct {
		act  Action
		name string
	}

	tests := []testData{
		{ActUnknown, "Unknown"},
		{ActHello, "Hello"},
		{ActBye, "Bye"},
		{ActProbe, "Probe"},
		{ActProbeMatches, "ProbeMatches"},
		{ActResolve, "Resolve"},
		{ActResolveMatches, "ResolveMatches"},
		{ActGet, "Get"},
		{ActGetResponse, "GetResponse"},
	}

	for _, test := range tests {
		name := test.act.String()
		if name != test.name {
			t.Errorf("action=%d name=%s expected:%s",
				test.act, name, test.name)
		}

		enc := test.act.Encode()
		dec := ActDecode(enc)

		if dec != test.act {
			t.Errorf("action: %s\n"+
				"encoded: %q\n"+
				"decoded: %d\n"+
				"expected: %d\n",
				test.act, enc, dec, test.act)
		}
	}
}
