// MFP - Miulti-Function Printers and scanners toolkit
// Network state monitoring
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Event queue test

package netstate

import (
	"errors"
	"reflect"
	"testing"
)

func TestEventQueue(t *testing.T) {
	netifmaker := testNewNetIfMaker()
	if0 := netifmaker.new()
	a4 := testMakeAddr(if0, "127.0.0.1/24")
	a6 := testMakeAddr(if0, "::1/24")

	events := []Event{
		EventAddInterface{if0},
		EventAddAddress{a4},
		EventAddAddress{a6},
		EventError{errors.New("very serious error")},
		EventDelAddress{a6},
		EventDelAddress{a4},
		EventDelInterface{if0},
	}

	var q eventqueue
	for _, evnt := range events {
		q.push(evnt)
	}

	for _, expected := range append(events, nil) {
		evnt := q.pull()
		if !reflect.DeepEqual(evnt, expected) {
			t.Errorf("expected: %v, present: %v", expected, evnt)
		}
	}
}
