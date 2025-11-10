// MFP           - Miulti-Function Printers and scanners toolkit
// cmd/mfp-ppd   - Utility for PPD files
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Test of main() function test

package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/OpenPrinting/go-mfp/argv"
)

func TestMain(t *testing.T) {
	saveHelpOutput := argv.HelpOutput
	defer func() { argv.HelpOutput = saveHelpOutput }()

	buf := &bytes.Buffer{}
	argv.HelpOutput = buf

	saveArgs := os.Args
	defer func() { os.Args = saveArgs }()

	os.Args = []string{os.Args[0], "-h"}
	main()

	if !strings.HasPrefix(buf.String(), "usage:") {
		t.Errorf("Option -h not properly handled")
	}
}
