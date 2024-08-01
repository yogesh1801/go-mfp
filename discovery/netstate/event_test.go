// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Events tests

package netstate

import (
	"errors"
	"testing"
)

// TestEvent tests EventXXX
func TestEvent(t *testing.T) {
	type testData struct {
		evnt Event  // Event being tested
		str  string // Expected Event.String output
	}

	if0 := NetIf{1, "net1"}
	a4 := testMakeAddr(if0, "127.0.0.1/24")
	a6 := testMakeAddr(if0, "::1/24")

	tests := []testData{
		{
			evnt: EventAddInterface{if0},
			str:  `add-interface: net1(#1)`,
		},
		{
			evnt: EventDelInterface{if0},
			str:  `del-interface: net1(#1)`,
		},

		{
			evnt: EventAddAddress{a4},
			str:  `add-address: 127.0.0.1/24%net1(#1)`,
		},
		{
			evnt: EventAddAddress{a6},
			str:  `add-address: ::1/24%net1(#1)`,
		},

		{
			evnt: EventDelAddress{a4},
			str:  `del-address: 127.0.0.1/24%net1(#1)`,
		},
		{
			evnt: EventDelAddress{a6},
			str:  `del-address: ::1/24%net1(#1)`,
		},

		{
			evnt: EventAddPrimaryAddress{a4},
			str:  `add-primary: 127.0.0.1/24%net1(#1)`,
		},
		{
			evnt: EventAddPrimaryAddress{a6},
			str:  `add-primary: ::1/24%net1(#1)`,
		},

		{
			evnt: EventDelPrimaryAddress{a4},
			str:  `del-primary: 127.0.0.1/24%net1(#1)`,
		},
		{
			evnt: EventDelPrimaryAddress{a6},
			str:  `del-primary: ::1/24%net1(#1)`,
		},

		{
			evnt: EventError{errors.New("very fatal error")},
			str:  `error: very fatal error`,
		},
	}

	for _, test := range tests {
		test.evnt.event() // Does nothing; just for coverage
		s := test.evnt.String()
		if s != test.str {
			t.Errorf("Event.String mismatch:\n"+
				"expected: %s\n"+
				" present: %s\n",
				test.str, s)
		}
	}
}
