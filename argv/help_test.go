// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Help test

package argv

import (
	"errors"
	"testing"
)

// TestHelpMisuse tests HelpCommand behavior when miss-used
func TestHelpMisuse(t *testing.T) {
	err := HelpCommand.Run(nil)
	if err == nil {
		err = errors.New("")
	}

	expected := `HelpHandler must be used in sub-command`
	if err.Error() != expected {
		t.Errorf("Error mismatch")
		t.Errorf("expected: `%s`", expected)
		t.Errorf("received: `%s`", err)
	}
}

// TestHelpString tests HelpString() function
func TestHelpString(t *testing.T) {
	expected := "usage: help [command]\n"
	received := HelpString(&HelpCommand)
	if expected != received {
		t.Errorf("output mismatch")
		t.Errorf("expected: `%s`", expected)
		t.Errorf("received: `%s`", received)
	}
}

// TestHelpPanic tests that help panics on the invalid Command
func TestHelpPanic(t *testing.T) {
	defer func() {
		v := recover()
		err, ok := v.(error)
		if !ok || err.Error() != "missed command name" {
			panic(v)
		}

	}()

	// It must panic, because empty Command is invalid
	HelpString(&Command{})
}
