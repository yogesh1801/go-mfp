// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Tests of interface addresses

package netstate

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddr(t *testing.T) {
	type testCheck struct {
		op  string // Operation being tested ("equal", "less", ...)
		val bool   // Expected answer
	}

	type testData struct {
		a1, a2 Addr        // A couple of addresses
		checks []testCheck // Operations to test
	}

	if0 := NetIf{0, "if0"}
	if1 := NetIf{1, "if1"}

	tests := []testData{
		// Test: equal addresses on a same interface
		{
			a1: testMakeAddr(if0, "192.168.0.1/24"),
			a2: testMakeAddr(if0, "192.168.0.1/24"),
			checks: []testCheck{
				{"equal", true},
				{"sameinterface", true},
				{"less", false},
				{"overlaps", true},
				{"narrower", false},
			},
		},

		// Test: equal addresses on different interfaces
		{
			a1: testMakeAddr(if0, "192.168.0.1/24"),
			a2: testMakeAddr(if1, "192.168.0.1/24"),
			checks: []testCheck{
				{"equal", false},
				{"sameinterface", false},
				{"less", true},
				{"overlaps", false},
				{"narrower", false},
			},
		},

		// Test Addr.Less in various combinations
		{
			// a1 < a2 by IP
			a1:     testMakeAddr(if0, "192.168.0.1/24"),
			a2:     testMakeAddr(if0, "192.168.0.2/24"),
			checks: []testCheck{{"less", true}},
		},

		{
			// a1 > a2 by IP
			a1:     testMakeAddr(if0, "192.168.0.2/24"),
			a2:     testMakeAddr(if0, "192.168.0.1/24"),
			checks: []testCheck{{"less", false}},
		},

		{
			// a1 narrower that a2
			a1: testMakeAddr(if0, "192.168.0.1/32"),
			a2: testMakeAddr(if0, "192.168.0.1/24"),
			checks: []testCheck{
				{"less", true},
				{"narrower", true},
				{"overlaps", true},
			},
		},

		{
			// a1 wider that a2
			a1: testMakeAddr(if0, "192.168.0.1/24"),
			a2: testMakeAddr(if0, "192.168.0.1/32"),
			checks: []testCheck{
				{"less", false},
				{"narrower", false},
				{"overlaps", true},
			},
		},

		{
			// IP6 vs IP4
			a1: testMakeAddr(if0, "::1/24"),
			a2: testMakeAddr(if0, "127.0.0.1/24"),
			checks: []testCheck{
				{"less", false},
				{"narrower", false},
				{"overlaps", false},
				{"sameinterface", true},
			},
		},

		{
			// IP4 vs IP6
			a1: testMakeAddr(if0, "127.0.0.1/24"),
			a2: testMakeAddr(if0, "::1/24"),
			checks: []testCheck{
				{"less", true},
				{"narrower", false},
				{"overlaps", false},
				{"sameinterface", true},
			},
		},

		{
			// IP4 vs IP6-encoded IP4
			a1: testMakeAddr(if0, "127.0.0.1/24"),
			a2: testMakeAddr(if0, "::127.0.0.1/24"),
			checks: []testCheck{
				{"less", true},
				{"narrower", false},
				{"overlaps", false},
				{"equal", false},
			},
		},

		{
			// Interfaces with the same Index and different name
			a1: testMakeAddr(if0, "127.0.0.1/24"),
			a2: testMakeAddr(
				MakeNetIf(if0.Index(), if1.Name()),
				"127.0.0.1/24",
			),
			checks: []testCheck{
				{"less", true},
				{"overlaps", false},
				{"sameinterface", false},
			},
		},
	}

	for _, test := range tests {
		for _, check := range test.checks {
			var val bool
			switch strings.ToLower(check.op) {
			case "equal":
				val = test.a1 == test.a2
			case "sameinterface":
				val = test.a1.SameInterface(test.a2)
			case "less":
				val = test.a1.Less(test.a2)
			case "overlaps":
				val = test.a1.Overlaps(test.a2)
			case "narrower":
				val = test.a1.Narrower(test.a2)
			default:
				panic(fmt.Errorf("invalid op %q", check.op))
			}

			if val != check.val {
				t.Errorf("%s@%s %s %s@%s:\n"+
					"expected %v, present %v",
					test.a1, test.a1.Interface().Name(),
					check.op,
					test.a2, test.a2.Interface().Name(),
					check.val, val)
			}
		}
	}
}
