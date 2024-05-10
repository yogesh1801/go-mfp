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
