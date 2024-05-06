// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Help test

package argv

import (
	"testing"
)

var testCommandWithParameters = Command{
	Name: "copy-files",
	Description: "this command copies multiple input files\n" +
		"info the single output.\n" +
		"\n" +
		"optionally, output can be compressed",
	Options: []Option{
		{
			Name:    "-c",
			Aliases: []string{"--compress"},
			Help:    "compress output",
		},

		{
			Name:    "-Z",
			Aliases: []string{"--gzip-compression"},
			Help: "use gzip compression\n" +
				"(slower but compresses better)",
		},

		{
			Name:    "-h",
			Aliases: []string{"--help"},
			Help:    "print help page",
		},
	},
	Parameters: []Parameter{
		{Name: "input-file..."},
		{Name: "output-file"},
	},
}

var testCommandWithParametersHelp = `
usage: copy-files [options] input-file... output-file

Options are:
  -c, --compress    compress output
  -Z, --gzip-compression
                    use gzip compression
                    (slower but compresses better)
  -h, --help        print help page

this command copies multiple input files
info the single output.

optionally, output can be compressed
`

func TestHelp(t *testing.T) {
	type testData struct {
		cmd *Command
		hlp string
	}

	tests := []testData{
		{&testCommandWithParameters, testCommandWithParametersHelp[1:]},
	}

	for i, test := range tests {
		hlp := HelpString(test.cmd)
		if hlp != test.hlp {
			t.Errorf("Test %d:", i)
			t.Errorf("expected:\n==========\n%s==========", test.hlp)
			t.Errorf("received:\n==========\n%s==========", hlp)
		}
	}
}
