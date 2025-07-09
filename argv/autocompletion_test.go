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

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// TestAutoCompletion tests Command.Complete
func TestAutoCompletion(t *testing.T) {
	type testData struct {
		argv []string     // Input
		cmd  Command      // Command description
		out  []Completion // Expected output
	}

	tests := []testData{
		// ----- Option names auto completion -----

		// Test 0: Misc of short/long options
		{
			argv: []string{"-"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-h"},
					{Name: "--long-1"},
					{Name: "--long-2"},
					{Name: "--other", Aliases: []string{"--long-3"}},
				},
			},
			out: []Completion{
				{"-h", false},
				{"--long-1", false},
				{"--long-2", false},
				{"--other", false},
				{"--long-3", false},
			},
		},

		// Test 1: long option name auto-completion
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
			out: []Completion{
				{"--long-", true},
			},
		},

		// ----- Option arguments auto completion ----

		// Test 2: short option, separate argument
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
			out: []Completion{
				{"Robert", false},
				{"Roger", false},
			},
		},

		// Test 3: short option with embedded argument
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
			out: []Completion{
				{"-xRobert", false},
				{"-xRoger", false},
			},
		},

		// Test 4: short option, missed argument
		{
			argv: []string{"-x", ""},
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
			out: []Completion{
				{"Ro", true},
			},
		},

		// Test 5: short option with preceding unknown optipn
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
			out: []Completion{},
		},

		// Test 6: short option without value
		{
			argv: []string{"-x", ""},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-x"},
				},
			},
			out: []Completion{},
		},

		// Test 7: short option without value,
		// then option that needs value
		{
			argv: []string{"-a", "-x", ""},
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
			out: []Completion{
				{"Ro", true},
			},
		},

		// Test 8: long option, separate argument
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
			out: []Completion{
				{"Robert", false},
				{"Roger", false},
			},
		},

		// Test 9: two options, second completes
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
			out: []Completion{
				{"Robert", false},
				{"Roger", false},
			},
		},

		// Test 10: long option with embedded argument
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
			out: []Completion{
				{"--long=Robert", false},
				{"--long=Roger", false},
			},
		},

		// Test 11: long option, missed argument
		{
			argv: []string{"--long", ""},
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
			out: []Completion{
				{"Ro", true},
			},
		},

		// Test 12: long option with preceding unknown optipn
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
			out: []Completion{},
		},

		// Test 13: long option without value
		{
			argv: []string{"--long", ""},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "--long"},
				},
			},
			out: []Completion{},
		},

		// Test 14: long option without value,
		// then option that needs value
		{
			argv: []string{"--void", "--long", ""},
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
			out: []Completion{
				{"Ro", true},
			},
		},

		// Test 15: unknown short option with value
		{
			argv: []string{"-xV"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-o"},
				},
			},
			out: []Completion{},
		},

		// Test 16: unknown long option with value
		{
			argv: []string{"--long=val"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "--long-known"},
				},
			},
			out: []Completion{},
		},

		// ----- Sub-commands auto-completion -----

		// Test 17: sub-commands, successful completion with prefix
		{
			argv: []string{"Ro"},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "Roger"},
					{Name: "Robert"},
				},
			},
			out: []Completion{
				{"Robert", false},
				{"Roger", false},
			},
		},

		// Test 18: a single sub-command, successful completion with prefix
		{
			argv: []string{"Ro"},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "Roger"},
				},
			},
			out: []Completion{
				{"Roger", false},
			},
		},

		// Test 19: sub-commands, successful completion without prefix
		{
			argv: []string{},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "Roger"},
					{Name: "Robert"},
				},
			},
			out: []Completion{
				{"Ro", true},
			},
		},

		// Test 20: option, "--", sub-commands
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
			out: []Completion{
				{"Robert", false},
				{"Roger", false},
			},
		},

		// Test 21: sub-commands, extra parameter
		{
			argv: []string{"extra", "Ro"},
			cmd: Command{
				Name: "test",
				SubCommands: []Command{
					{Name: "Roger"},
					{Name: "Robert"},
				},
			},
			out: []Completion{},
		},

		// ----- Parameters auto-completion -----

		// Test 22: parameter completion
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
			out: []Completion{
				{"Robert", false},
				{"Roger", false},
			},
		},

		// Test 23: options, '--', parameter
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
			out: []Completion{
				{"Robert", false},
				{"Roger", false},
			},
		},

		// Test 24: options, '--', parameter missed
		{
			argv: []string{"-c", "--"},
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
			out: []Completion{
				{"Ro", true},
			},
		},

		// Test 25: parameter completion, extra parameter
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
			out: []Completion{},
		},

		// Test 26: parameter completion, repeated first
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
			out: []Completion{
				{"Robert", false},
				{"Roger", false},
			},
		},

		// Test 27: parameter completion, repeated last
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
			out: []Completion{},
		},

		// Test 28: '--' at the end and there are long options
		{
			argv: []string{"-c", "--"},
			cmd: Command{
				Name: "test",
				Options: []Option{
					{Name: "-c"},
					{Name: "--long-1"},
					{Name: "--long-2"},
				},
			},
			out: []Completion{
				{"--long-", true},
			},
		},

		// ----- Real-life examples -----

		// Test 29: nested sub-commands
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
			out: []Completion{
				{"get-default", false},
			},
		},
	}

	for i, test := range tests {
		out := test.cmd.Complete(test.argv)

		diff := testDiffCompletion(test.out, out)
		if len(diff) != 0 {
			t.Errorf("[%d]: results mismatch:", i)

			for _, s := range diff {
				t.Errorf("  %s", s)
			}
		}
	}
}

// testDiffCompletion computes a difference between completion results
func testDiffCompletion(expected, received []Completion) []string {
	if len(expected) == 0 && len(received) == 0 {
		return []string{}
	}

	expected = generic.CopySlice(expected)
	received = generic.CopySlice(received)

	sort.Slice(expected, func(i, j int) bool {
		return expected[i].String < expected[j].String
	})
	sort.Slice(received, func(i, j int) bool {
		return received[i].String < received[j].String
	})

	out := []string{}

	if !reflect.DeepEqual(expected, received) {
		out = append(out, fmt.Sprintf("<<< %#v", expected))
		out = append(out, fmt.Sprintf(">>> %#v", received))
	}

	return out
}
