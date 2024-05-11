// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Argv parser test

package argv

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// TestNewParserPanic tests panic from newParser()
func TestNewParserPanic(t *testing.T) {
	defer func() {
		v := recover()
		err, ok := v.(error)
		if !ok || err.Error() != "missed command name" {
			panic(v)
		}

	}()

	// It must panic, because empty Command is invalid
	newParser(&Command{}, []string{})
}

// TestParser tests argv parser
func TestParser(t *testing.T) {
	type testData struct {
		argv    []string            // Input
		cmd     Command             // Command description
		err     string              // Expected error, "" if none
		out     map[string][]string // Expected options values
		params  []string            // Expected positional parameters
		subcmd  string              // Expected sub-command
		subargv []string            // Expected sub-command argv
	}

	tests := []testData{
		// Test 0: options on various combinations
		{
			argv: []string{
				"-n", "123",
				"-v456",
				"value1",
				"--long1", "hello",
				"--long2=world",
				"value2",
				"-abc",
				"--",
				"--value3",
			},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "-n",
						Aliases:  []string{"--long-n"},
						Validate: ValidateInt32,
					},

					{
						Name:     "-v",
						Validate: ValidateInt32,
					},

					{
						Name:     "--long1",
						Validate: ValidateAny,
					},

					{
						Name:     "--long2",
						Validate: ValidateAny,
					},

					{Name: "-a"},
					{Name: "-b"},
					{Name: "-c"},
				},

				Parameters: []Parameter{
					{Name: "param1", Validate: ValidateAny},
					{Name: "[param2]"},
					{Name: "[param3]"},
				},
			},
			out: map[string][]string{
				"-a":       {""},
				"-b":       {""},
				"-c":       {""},
				"--long1":  {"hello"},
				"--long2":  {"world"},
				"--long-n": {"123"},
				"-n":       {"123"},
				"-v":       {"456"},
				"[param2]": {"value2"},
				"[param3]": {"--value3"},
				"param1":   {"value1"},
			},
			params: []string{"value1", "value2", "--value3"},
		},

		// Test 1: repeated parameters
		{
			argv: []string{
				"a", "b", "c",
			},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1"},
					{Name: "param2..."},
				},
			},
			out: map[string][]string{
				"param1":    {"a"},
				"param2...": {"b", "c"},
			},
			params: []string{"a", "b", "c"},
		},

		// Test 2: repeated parameters, followed by required parameter
		{
			argv: []string{
				"a", "b", "c",
			},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1..."},
					{Name: "param2"},
				},
			},
			out: map[string][]string{
				"param1...": {"a", "b"},
				"param2":    {"c"},
			},
			params: []string{"a", "b", "c"},
		},

		// Test 3: sub-commands
		{
			argv: []string{
				"sub-2",
			},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "sub-1"},
					{Name: "sub-2"},
					{Name: "sub-3"},
				},
			},
			subcmd:  "sub-2",
			subargv: []string{},
		},

		// Test 4: options and abbreviated sub-command with params
		{
			argv: []string{
				"--long", "l1",
				"-x", "xxx",
				"sub-2", "param1", "param2", "param3",
			},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "-l",
						Aliases:  []string{"--long"},
						Validate: ValidateAny,
					},
					{
						Name:     "-x",
						Aliases:  []string{"--xxl"},
						Validate: ValidateAny,
					},
				},
				SubCommands: []Command{
					{Name: "sub-1-cmd"},
					{Name: "sub-2-cmd"},
					{Name: "sub-3-cmd"},
				},
			},
			out: map[string][]string{
				"--long": {"l1"},
				"--xxl":  {"xxx"},
				"-l":     {"l1"},
				"-x":     {"xxx"},
			},
			subcmd:  "sub-2-cmd",
			subargv: []string{"param1", "param2", "param3"},
		},

		// Test 5: "unexpected parameter"
		{
			argv: []string{"a", "b", "c"},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1"},
					{Name: "param2"},
				},
			},
			err: `unexpected parameter: "c"`,
		},

		// Test 6: "unexpected parameter" with optional parameters
		{
			argv: []string{"a", "b", "c"},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1"},
					{Name: "[param2]"},
				},
			},
			err: `unexpected parameter: "c"`,
		},

		// Test 7: "missed parameter"
		{
			argv: []string{"a"},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1"},
					{Name: "param2"},
				},
			},
			err: `missed parameter: "param2"`,
		},

		// Test 8: "missed sub-comman name"
		{
			argv: []string{"-x", "5"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-x", Validate: ValidateAny},
				},
				SubCommands: []Command{
					{Name: "sub-1"},
					{Name: "sub-2"},
					{Name: "sub-3"},
				},
			},
			err: `missed sub-command name`,
		},

		// Test 9: "unknown option" for short option
		{
			argv: []string{"-x", "5"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-a", Validate: ValidateAny},
					{Name: "-b", Validate: ValidateAny},
					{Name: "-c", Validate: ValidateAny},
				},
			},
			err: `unknown option: "-x"`,
		},

		// Test 10: "unknown option" for long optiob
		{
			argv: []string{"--unknown=5"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-a", Validate: ValidateAny},
					{Name: "-b", Validate: ValidateAny},
					{Name: "-c", Validate: ValidateAny},
				},
			},
			err: `unknown option: "--unknown"`,
		},

		// Test 11: "unknown option" from combined options
		{
			argv: []string{"-abx"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-a"},
					{Name: "-b"},
					{Name: "-c"},
				},
			},
			err: `unknown option: "-x"`,
		},

		// Test 12: "option requires operand"
		{
			argv: []string{"-x"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-x", Validate: ValidateAny},
				},
			},
			err: `option requires operand: "-x"`,
		},

		// Test 13: "option requires operand" from combined options
		{
			argv: []string{"-abc"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-a"},
					{Name: "-b", Validate: ValidateAny},
					{Name: "-c"},
				},
			},
			err: `option requires operand: "-b"`,
		},

		// Test 14: "unknown sub-command"
		{
			argv: []string{
				"unknown",
			},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "sub-1"},
					{Name: "sub-2"},
					{Name: "sub-3"},
				},
			},
			err: `unknown sub-command: "unknown"`,
		},

		// Test 15: "ambiguous sub-command"
		{
			argv: []string{
				"sub",
			},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "sub-1"},
					{Name: "sub-2"},
					{Name: "sub-3"},
				},
			},
			err: `ambiguous sub-command: "sub"`,
		},

		// Test 16: "invalid option value" for short option
		{
			argv: []string{"-b", "hello"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-a", Validate: ValidateInt32},
					{Name: "-b", Validate: ValidateInt32},
					{Name: "-c", Validate: ValidateInt32},
				},
			},
			err: `invalid integer: -b "hello"`,
		},

		// Test 17: "invalid option value" for long option
		{
			argv: []string{"--long-b", "hello"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "--long-a", Validate: ValidateInt32},
					{Name: "--long-b", Validate: ValidateInt32},
					{Name: "--long-c", Validate: ValidateInt32},
				},
			},
			err: `invalid integer: --long-b "hello"`,
		},

		// Test 18: "invalid parameter value"
		{
			argv: []string{"1", "hello", "2"},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "a", Validate: ValidateInt32},
					{Name: "b", Validate: ValidateInt32},
					{Name: "c", Validate: ValidateInt32},
				},
			},
			err: `"b": invalid integer "hello"`,
		},

		// Test 19: "option conflicts with..."
		{
			argv: []string{"-a", "-b", "-c"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-a"},
					{Name: "-b", Conflicts: []string{"-c"}},
					{Name: "-c"},
				},
			},
			err: `option "-c" conflicts with "-b"`,
		},

		// Test 20: "missed option..."
		{
			argv: []string{"-a", "-b"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-a", Requires: []string{"-c"}},
					{Name: "-b"},
					{Name: "-c"},
				},
			},
			err: `missed option "-c", required by "-a"`,
		},

		// Test 21: option required and present
		{
			argv: []string{"-a", "-b", "-c"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-a", Requires: []string{"-c"}},
					{Name: "-b"},
					{Name: "-c"},
				},
			},
			out: map[string][]string{
				"-a": {""},
				"-b": {""},
				"-c": {""},
			},
		},

		// Test 22: parameters after '--'
		{
			argv: []string{
				"a", "--", "-b", "--long",
			},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{Name: "param1..."},
					{Name: "param2"},
				},
			},
			out: map[string][]string{
				"param1...": {"a", "-b"},
				"param2":    {"--long"},
			},
			params: []string{"a", "-b", "--long"},
		},

		// Test 23: '--', followed by sub-commands
		{
			argv: []string{
				"--", "sub-2",
			},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "sub-1"},
					{Name: "sub-2"},
					{Name: "sub-3"},
				},
			},
			subcmd:  "sub-2",
			subargv: []string{},
		},
	}

	for i, test := range tests {
		inv, err := test.cmd.Parse(test.argv)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("[%d]: error mismatch: expected `%s`, present `%s`",
				i, test.err, err)
		} else if test.err == "" {
			diff := testDiffValues(test.out, inv.byName)
			if len(diff) != 0 {
				t.Errorf("[%d]: results mismatch (<<< expected, >>> present):", i)

				for _, s := range diff {
					t.Errorf("  %s", s)
				}
			}

			var params []string
			for n := 0; n < inv.ParamCount(); n++ {
				params = append(params, inv.ParamGet(n))
			}

			if !reflect.DeepEqual(test.params, params) {
				t.Errorf("[%d]: params mismatch:", i)
				t.Errorf("  expected: %#v", test.params)
				t.Errorf("  present:  %#v", params)
			}

			subcmd := ""
			if inv.subcmd != nil {
				subcmd = inv.subcmd.Name
			}

			if subcmd != test.subcmd {
				t.Errorf("[%d]: subcmd mismatch: expected %q present %q",
					i, test.subcmd, subcmd)
			}

			if !reflect.DeepEqual(test.subargv, inv.subargv) {
				t.Errorf("[%d]: subargv mismatch:", i)
				t.Errorf("  expected: %q", test.subargv)
				t.Errorf("  present:  %q", inv.subargv)
			}
		}
	}
}

// TestParserCompletion tests (*parser) complete()
func TestParserCompletion(t *testing.T) {
	type testData struct {
		argv []string // Input
		cmd  Command  // Command description
		out  []string // Expected output
	}

	tests := []testData{
		// Test 0: short option, separate argument
		{
			argv: []string{"-x", "Ro"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "-x",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 1: short option with embedded argument
		{
			argv: []string{"-xRo"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "-x",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 2: short option, missed argument
		{
			argv: []string{"-x"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "-x",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Ro"},
		},

		// Test 3: short option with preceding unknown optipn
		{
			argv: []string{"-a", "-x", "Ro"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "-x",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{},
		},

		// Test 4: short option without value
		{
			argv: []string{"-x"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-x"},
				},
			},
			out: []string{},
		},

		// Test 5: short option without value,
		// then option that needs value
		{
			argv: []string{"-a", "-x"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name: "-a",
					},
					{
						Name:     "-x",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Ro"},
		},

		// Test 6: long option, separate argument
		{
			argv: []string{"--long", "Ro"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "--long",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 7: two options, second completes
		{
			argv: []string{"--first", "1", "--second", "Ro"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "--first",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"1",
								"2",
							},
						),
					},
					{
						Name:     "--second",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 8: long option with embedded argument
		{
			argv: []string{"--long=Ro"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "--long",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 9: long option, missed argument
		{
			argv: []string{"--long"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "--long",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Ro"},
		},

		// Test 10: long option with preceding unknown optipn
		{
			argv: []string{"--unknown", "--long", "Ro"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name:     "--long",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{},
		},

		// Test 11: long option without value
		{
			argv: []string{"--long"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "--long"},
				},
			},
			out: []string{},
		},

		// Test 12: long option without value,
		// then option that needs value
		{
			argv: []string{"--void", "--long"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{
						Name: "--void",
					},
					{
						Name:     "--long",
						Validate: ValidateAny,
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Ro"},
		},

		// Test 13: long option name auto-completion
		{
			argv: []string{"--long"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "--long-1"},
					{Name: "--long-2"},
					{Name: "--other", Aliases: []string{"--long-3"}},
				},
			},
			out: []string{"--long-"},
		},

		// Test 14: sub-commands, successful completion with prefix
		{
			argv: []string{"Ro"},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "Roger"},
					{Name: "Robert"},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 15: a single sub-command, successful completion with prefix
		{
			argv: []string{"Ro"},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "Roger"},
				},
			},
			out: []string{"Roger "},
		},

		// Test 16: sub-commands, successful completion without prefix
		{
			argv: []string{},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "Roger"},
					{Name: "Robert"},
				},
			},
			out: []string{"Ro"},
		},

		// Test 17: option, "--", sub-commands
		{
			argv: []string{"--long", "value", "--", "Ro"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "--long", Validate: ValidateAny},
				},
				SubCommands: []Command{
					{Name: "Roger"},
					{Name: "Robert"},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 18: sub-commands, extra parameter
		{
			argv: []string{"extra", "Ro"},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "Roger"},
					{Name: "Robert"},
				},
			},
			out: []string{},
		},

		// Test 19: parameter completion
		{
			argv: []string{"Ro"},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "param",
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 20: options, '--', parameter
		{
			argv: []string{"-c", "--", "Ro"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-c"},
				},
				Parameters: []Parameter{
					{
						Name: "param",
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 21: parameter completion, extra parameter
		{
			argv: []string{"extra", "Ro"},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "param",
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{},
		},

		// Test 22: parameter completion, repeated first
		{
			argv: []string{"extra", "Ro"},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "param1...",
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
					{
						Name: "param2",
					},
				},
			},
			out: []string{"Robert ", "Roger "},
		},

		// Test 23: parameter completion, repeated last
		{
			argv: []string{"extra", "Ro"},
			cmd: Command{
				Name: "test",
				Parameters: []Parameter{
					{
						Name: "param1...",
					},
					{
						Name: "param1",
						Complete: CompleteStrings(
							[]string{
								"Roger",
								"Robert",
							},
						),
					},
				},
			},
			out: []string{},
		},
	}

	for i, test := range tests {
		out := test.cmd.Complete(test.argv)

		diff := testDiffCompletion(test.out, out)
		if len(diff) != 0 {
			t.Errorf("[%d]: results mismatch (<<< expected, >>> present):", i)

			for _, s := range diff {
				t.Errorf("  %s", s)
			}
		}
	}
}

// testDiffValues compares two maps of named values and returns formatted
// diff as slice of strings
func testDiffValues(m1, m2 map[string][]string) []string {
	type nmval struct {
		name   string
		values []string
	}

	// Convert maps into sorted slices
	s1 := []nmval{}
	for n, v := range m1 {
		s1 = append(s1, nmval{n, v})
	}

	s2 := []nmval{}
	for n, v := range m2 {
		s2 = append(s2, nmval{n, v})
	}

	sort.Slice(s1, func(i, j int) bool { return s1[i].name < s1[j].name })
	sort.Slice(s2, func(i, j int) bool { return s2[i].name < s2[j].name })

	out := []string{}

	// Compare sorted slices
	for len(s1) > 0 && len(s2) > 0 {
		switch {
		case s1[0].name < s2[0].name:
			s := fmt.Sprintf("<<< %s: %q", s1[0].name, s1[0].values)
			out = append(out, s)
			s1 = s1[1:]

		case s1[0].name > s2[0].name:
			s := fmt.Sprintf(">>> %s: %q", s2[0].name, s2[0].values)
			out = append(out, s)
			s2 = s2[1:]

		default:
			if !reflect.DeepEqual(s1[0].values, s2[0].values) {
				s := fmt.Sprintf("<<< %s: %q",
					s1[0].name, s1[0].values)
				out = append(out, s)
				s = fmt.Sprintf(">>> %s: %q",
					s2[0].name, s2[0].values)
				out = append(out, s)
			}

			s1 = s1[1:]
			s2 = s2[1:]
		}
	}

	for i := range s1 {
		s := fmt.Sprintf("<<< %s: %q", s1[i].name, s1[i].values)
		out = append(out, s)
	}

	for i := range s2 {
		s := fmt.Sprintf(">>> %s: %q", s2[i].name, s2[i].values)
		out = append(out, s)
	}

	return out
}

// testDiffCompletion computes a difference between completion results
func testDiffCompletion(expected, received []string) []string {
	expected = testCopySliceOfStrings(expected)
	received = testCopySliceOfStrings(received)

	sort.Strings(expected)
	sort.Strings(received)

	out := []string{}

	if !reflect.DeepEqual(expected, received) {
		out = append(out, fmt.Sprintf("<<< %q", expected))
		out = append(out, fmt.Sprintf(">>> %q", received))
	}

	return out
}

// testCopySliceOfStrings creates a copy of slice of strings
func testCopySliceOfStrings(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)
	return out
}
