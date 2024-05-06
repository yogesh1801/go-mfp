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

var testCommandWithSubCommands = Command{
	Name: "test",
	Options: []Option{
		{
			Name:    "-v",
			Aliases: []string{"--verbose"},
			Help:    "enable verbose logging",
		},
		HelpOption,
	},
	SubCommands: []Command{
		{
			Name:        "connect",
			Help:        "connect to the server",
			Description: "connect establishes server connection",
		},
		{
			Name:        "disconnect",
			Help:        "disconnect from the server",
			Description: "disconnect terminates the server connection",
		},

		{
			Name:        "send-files-to-server",
			Help:        "send-files-to-server uploads the files",
			Description: "send-files-to-server uploads files to server",
		},

		HelpCommand,
	},
}

var testCommandWithSubCommandsHelp = `
usage: test [options] command [arguments]

Options are:
  -v, --verbose     enable verbose logging
  -h, --help        print help page

Commands are:
  connect           connect to the server
  disconnect        disconnect from the server
  send-files-to-server
                    send-files-to-server uploads the files
  help              print help page
`

func TestHelp(t *testing.T) {
	type testData struct {
		cmd *Command
		hlp string
	}

	tests := []testData{
		{&testCommandWithParameters, testCommandWithParametersHelp[1:]},
		{&testCommandWithSubCommands, testCommandWithSubCommandsHelp[1:]},
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
