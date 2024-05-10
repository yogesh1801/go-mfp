// MFP     - Miulti-Function Printers and scanners toolkit
// cmd/mfp - Universal command that implements all other as sub-commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Test of main() function

package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/alexpevzner/mfp/argv"
)

func TestMain(t *testing.T) {
	saveHelpOutput := argv.HelpOutput
	defer func() { argv.HelpOutput = saveHelpOutput }()

	buf := &bytes.Buffer{}
	argv.HelpOutput = buf

	os.Args = []string{os.Args[0], "-h"}
	main()

	if !strings.HasPrefix(buf.String(), "usage:") {
		t.Errorf("Option -h not properly handled")
	}
}
