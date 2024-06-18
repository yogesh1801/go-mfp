// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// (*Command) Run() test

package argv

import (
	"bytes"
	"errors"
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

		HelpOption,
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

Parameters are:
  input-file
  output-file

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
  help, ?           print help page
`

var testCommandWithSubCommandsHelpConnect = `
usage: connect

connect establishes server connection
`

// TestRun is a test for (*Command) Run()
func TestRun(t *testing.T) {
	type testData struct {
		argv []string // Command's arguments
		cmd  *Command // The Command
		hlp  string   // Expected help text
		err  string   // Expected error
	}

	tests := []testData{
		// ----- help system test -----
		{
			argv: []string{"-h"},
			cmd:  &testCommandWithParameters,
			hlp:  testCommandWithParametersHelp[1:],
		},

		{
			argv: []string{"help"},
			cmd:  &testCommandWithSubCommands,
			hlp:  testCommandWithSubCommandsHelp[1:],
		},

		{
			argv: []string{"help", "connect"},
			cmd:  &testCommandWithSubCommands,
			hlp:  testCommandWithSubCommandsHelpConnect[1:],
		},

		{
			argv: []string{"help", "unknown"},
			cmd:  &testCommandWithSubCommands,
			err:  `unknown sub-command: "unknown"`,
		},

		// Miscellaneous tests
		{
			// Attempt to invoke unknown command
			argv: []string{"unknown"},
			cmd:  &testCommandWithSubCommands,
			err:  `unknown sub-command: "unknown"`,
		},

		{
			// Attempt to invoke command without handler
			argv: []string{"connect"},
			cmd:  &testCommandWithSubCommands,
			err:  `unhandled command: connect`,
		},
	}

	saveHelpOutput := HelpOutput
	defer func() { HelpOutput = saveHelpOutput }()

	for i, test := range tests {
		buf := &bytes.Buffer{}
		HelpOutput = buf
		err := test.cmd.Run(test.argv)
		hlp := buf.String()

		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("Test %d: error mismatch", i)
			t.Errorf("expected: `%s`", test.err)
			t.Errorf("received: `%s`", err)
		} else if hlp != test.hlp {
			t.Errorf("Test %d:", i)
			t.Errorf("expected:\n==========\n%s==========", test.hlp)
			t.Errorf("received:\n==========\n%s==========", hlp)
		}
	}
}
