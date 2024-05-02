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

// TestParser tests argv parser
func TestParser(t *testing.T) {
	type testData struct {
		argv    []string            // Input
		cmd     Command             // Command description
		err     string              // Expected error, "" if none
		out     map[string][]string // Expected options values
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

		// Test 3: options and abbreviated sub-command with params
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
	}

	for i, test := range tests {
		prs := newParser(&test.cmd, test.argv)
		err := prs.parse()
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("[%d}: error mismatch: expected `%s`, present `%s`",
				i, test.err, err)
		} else if test.err == "" {
			diff := testDiffValues(test.out, prs.byName)
			if len(diff) != 0 {
				t.Errorf("[%d}: results mismatch (<<< expected, >>> present):", i)

				for _, s := range diff {
					t.Errorf("  %s", s)
				}
			}

			subcmd := ""
			if prs.subcmd != nil {
				subcmd = prs.subcmd.Name
			}

			if subcmd != test.subcmd {
				t.Errorf("[%d}: subcmd mismatch: expected %q present %q",
					i, test.subcmd, subcmd)
			}

			if !reflect.DeepEqual(test.subargv, prs.subargv) {
				t.Errorf("[%d}: subargv mismatch:", i)
				t.Errorf("  expected: %q", test.subargv)
				t.Errorf("  present:  %q", prs.subargv)
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
