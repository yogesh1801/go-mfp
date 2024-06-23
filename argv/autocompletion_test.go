// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Argv parser test

package argv

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// TestAutoCompletion tests Command.Complete
func TestAutoCompletion(t *testing.T) {
	type testData struct {
		argv  []string       // Input
		cmd   Command        // Command description
		out   []string       // Expected output
		flags CompleterFlags // Expected flags
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
			out:   []string{"Robert", "Roger"},
			flags: 0,
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
			out:   []string{"Robert", "Roger"},
			flags: 0,
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
			out:   []string{"Ro"},
			flags: CompleterNoSpace,
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
			out:   []string{"Ro"},
			flags: CompleterNoSpace,
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
			out: []string{"Robert", "Roger"},
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
			out: []string{"Robert", "Roger"},
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
			out: []string{"Robert", "Roger"},
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
			out:   []string{"Ro"},
			flags: CompleterNoSpace,
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
			out:   []string{"Ro"},
			flags: CompleterNoSpace,
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
			out:   []string{"--long-"},
			flags: CompleterNoSpace,
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
			out: []string{"Robert", "Roger"},
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
			out: []string{"Roger"},
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
			out:   []string{"Ro"},
			flags: CompleterNoSpace,
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
			out: []string{"Robert", "Roger"},
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
			out: []string{"Robert", "Roger"},
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
			out: []string{"Robert", "Roger"},
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
			out: []string{"Robert", "Roger"},
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

		// Test 24: nested sub-commands
		{
			argv: []string{"cups", "ge"},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{
						Name: "cups",
						SubCommands: []Command{
							{
								Name: "get-default",
							},
						},
					},
				},
			},
			out: []string{"get-default"},
		},
	}

	for i, test := range tests {
		out, flags := test.cmd.Complete(test.argv)

		diff := testDiffCompletion(test.out, out)
		if len(diff) != 0 {
			t.Errorf("[%d]: results mismatch (<<< expected, >>> present):", i)

			for _, s := range diff {
				t.Errorf("  %s", s)
			}
		}

		if flags != test.flags {
			t.Errorf("[%d]: flags mismatch\n"+
				"extected: %s\nreceived: %s",
				i, test.flags, flags)
		}
	}
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
