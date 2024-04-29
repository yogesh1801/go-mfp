// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Commands test

package argv

import (
	"errors"
	"testing"
)

// TestCommandVerify is a test for (*Command) Verify()
func TestCommandVerify(t *testing.T) {
	// testData contains a single-test data
	type testData struct {
		cmd *Command
		err string
	}

	tests := []testData{
		{
			cmd: &Command{},
			err: `missed command name`,
		},

		// Tests for malformed Options
		{
			cmd: &Command{
				Name: "test",
				Options: []Option{
					{
						Name: "",
					},
				},
			},
			err: `test: option must have a name`,
		},

		{
			cmd: &Command{
				Name: "test",
				Options: []Option{
					{
						Name: "opt",
					},
				},
			},
			err: `test: option must start with dash (-): "opt"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Options: []Option{
					{
						Name: "-",
					},
				},
			},
			err: `test: empty option name: "-"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Options: []Option{
					{
						Name: "--",
					},
				},
			},
			err: `test: empty option name: "--"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Options: []Option{
					{
						Name: "---long",
					},
				},
			},
			err: `test: invalid char '-' in option: "---long"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Options: []Option{
					{
						Name: "-long123@",
					},
				},
			},
			err: `test: invalid char '@' in option: "-long123@"`,
		},

		// Tests for duplicated options
		{
			cmd: &Command{
				Name: "test",
				Options: []Option{
					{
						Name: "-c",
					},
					{
						Name: "-c",
					},
				},
			},
			err: `test: duplicated option "-c"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Options: []Option{
					{
						Name: "-c",
					},
					{
						Name:    "-v",
						Aliases: []string{"-c"},
					},
				},
			},
			err: `test: duplicated option "-c"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Options: []Option{
					{
						Name:    "-c",
						Aliases: []string{"-help"},
					},
					{
						Name:    "-v",
						Aliases: []string{"-help"},
					},
				},
			},
			err: `test: duplicated option "-help"`,
		},

		// Tests for malformed Parameters
		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "param1",
					},
				},
				SubCommands: []Command{
					{
						Name: "param1",
					},
				},
			},
			err: `test: Parameters and SubCommands are mutually exclusive`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "param1",
					},
					{
						Name: "param1",
					},
				},
			},
			err: `test: duplicated parameter "param1"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "",
					},
				},
			},
			err: `test: parameter must have a name`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "[]",
					},
				},
			},
			err: `test: parameter name is empty: "[]"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "...",
					},
				},
			},
			err: `test: parameter name is empty: "..."`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "[...]",
					},
				},
			},
			err: `test: parameter name is empty: "[...]"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "[param",
					},
				},
			},
			err: `test: missed closing ']' character in parameter "[param"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "-param",
					},
				},
			},
			err: `test: invalid char '-' in parameter: "-param"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "pa-ram",
					},
				},
			},
			err: "",
		},

		// Tests for Parameters disposition
		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1"},
					{Name: "param2"},
					{Name: "param3"},
				},
			},
			err: "",
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1"},
					{Name: "[param2]"},
					{Name: "param3"},
				},
			},
			err: `test: required parameter "param3" used after optional "[param2]"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1"},
					{Name: "param2..."},
					{Name: "[param3]"},
				},
			},
			err: `test: optional parameter "[param3]" used after repeated "param2..."`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1"},
					{Name: "[param2...]"},
					{Name: "[param3]"},
				},
			},
			err: `test: optional parameter "[param3]" used after repeated "[param2...]"`,
		},

		{
			cmd: &Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1"},
					{Name: "param2..."},
					{Name: "param3..."},
				},
			},
			err: `test: repeated parameter used twice ("param2..." and "param3...")`,
		},

		// Tests SubCommands
		{
			cmd: &Command{
				Name: "test",
				SubCommands: []Command{
					{
						Name: "subcmd",
						Parameters: []Parameter{
							{Name: "param1"},
							{Name: "param2..."},
							{Name: "param3..."},
						},
					},
				},
			},
			err: `test.subcmd: repeated parameter used twice ("param2..." and "param3...")`,
		},

		{
			cmd: &Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "subcmd"},
					{Name: "subcmd"},
				},
			},
			err: `test.duplicated subcommand "subcmd"`,
		},
	}

	for _, test := range tests {
		err := test.cmd.Verify()
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("(*Command) Verify(): expected %q, present %q",
				test.err, err)
		}
	}
}
