// MFP - Miulti-Function Printers and scanners toolkit
// IPv6 zone suffixes handling
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Test of functions for IPv6 zone suffixes

package zone

import (
	"net"
	"testing"
)

func TestZone(t *testing.T) {
	ift, err := net.Interfaces()
	if err != nil {
		t.Errorf("net.Interfaces: %s", err)
		return
	}

	for _, ifi := range ift {
		n := Name(ifi.Index)
		i := Index(ifi.Name)

		if n != ifi.Name {
			t.Errorf("Name(%d): expected %q, present %q",
				ifi.Index, ifi.Name, n)
		}

		if i != ifi.Index {
			t.Errorf("Index(%q): expected %d, present %d",
				ifi.Name, ifi.Index, i)
		}
	}

	const badIndex = 1234567
	const badName = "1234567"

	n := Name(badIndex)
	i := Index(badName)

	if n != badName {
		t.Errorf("Name(%d): expected %q, present %q",
			badIndex, badName, n)
	}

	if i != badIndex {
		t.Errorf("Index(%q): expected %d, present %d",
			badName, badIndex, i)
	}

	if Name(0) != "" {
		t.Errorf("Name(0): expected %q, present %q", "", Name(0))
	}

	if Index("") != 0 {
		t.Errorf("Index(%q): expected 0, present %d", "", Index(""))
	}
}
