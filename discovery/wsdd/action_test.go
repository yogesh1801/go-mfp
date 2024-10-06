// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Message actions test

package wsdd

import "testing"

// TestAction tests operations with action
func TestAction(t *testing.T) {
	type testData struct {
		act  action
		name string
	}

	tests := []testData{
		{actUnknown, "Unknown"},
		{actHello, "Hello"},
		{actBye, "Bye"},
		{actProbe, "Probe"},
		{actProbeMatches, "ProbeMatches"},
		{actResolve, "Resolve"},
		{actResolveMatches, "ResolveMatches"},
		{actGet, "Get"},
		{actGetResponse, "GetResponse"},
	}

	for _, test := range tests {
		name := test.act.String()
		if name != test.name {
			t.Errorf("action=%d name=%s expected:%s",
				test.act, name, test.name)
		}

		enc := test.act.Encode()
		dec := actDecode(enc)

		if dec != test.act {
			t.Errorf("action: %s\n"+
				"encoded: %q\n"+
				"decoded: %d\n"+
				"expected: %d\n",
				test.act, enc, dec, test.act)
		}
	}
}
