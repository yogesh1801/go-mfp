// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Value completers test.

package argv

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

// TestCompleteStrings tests CompleteStrings
func TestCompleteStrings(t *testing.T) {
	type testData struct {
		strings []string     // Strings to choose from
		arg     string       // Input string
		out     []Completion // Expected output
	}

	tests := []testData{
		{
			strings: []string{"foo", "bar"},
			arg:     "foo",
			out:     []Completion{},
		},

		{
			strings: []string{"foo-1", "foo-2"},
			arg:     "foo",
			out: []Completion{
				{"foo-1", 0},
				{"foo-2", 0},
			},
		},

		{
			strings: []string{"foo-1", "foo-2"},
			arg:     "",
			out: []Completion{
				{"foo-1", 0},
				{"foo-2", 0},
			},
		},
	}

	for _, test := range tests {
		completer := CompleteStrings(test.strings)
		out := completer(test.arg)

		if (len(out) != 0 || len(test.out) != 0) &&
			!reflect.DeepEqual(out, test.out) {

			t.Errorf("CompleteStrings(%#v): %q\n"+
				"expected: %#v\nreceived: %#v\n",
				test.strings, test.arg, test.out, out)
		}
	}
}

// TestCompleteFs tests CompleteFs
func TestCompleteFs(t *testing.T) {
	var testFs fs.FS = fstest.MapFS{
		"etc/hosts":        &fstest.MapFile{Mode: 0755},
		"usr/bin/ls":       &fstest.MapFile{Mode: 0755},
		"usr/bin/ps":       &fstest.MapFile{Mode: 0755},
		"test/subdir/file": &fstest.MapFile{Mode: 0755},
		"test/subdirfile":  &fstest.MapFile{Mode: 0755},
	}

	type testData struct {
		getwd func() (string, error) // getwd function
		in    string                 // Input string
		out   []Completion           // Expected output
	}

	tests := []testData{
		// Simple tests
		{
			in: "/",
			out: []Completion{
				{"/etc", 0},
				{"/test", 0},
				{"/usr", 0},
			},
		},

		{
			in: "./",
			out: []Completion{
				{"./etc", 0},
				{"./test", 0},
				{"./usr", 0},
			},
		},

		{
			in: "/usr/bin/ls",
			out: []Completion{
				{"/usr/bin/ls", 0},
			},
		},

		{
			in: "/usr/bin",
			out: []Completion{
				{"/usr/bin/", CompletionNoSpace},
			},
		},

		{
			in: "test/subdir",
			out: []Completion{
				{"test/subdir", 0},
				{"test/subdirfile", 0},
			},
		},

		// Tests with getwd()
		{
			getwd: func() (string, error) {
				return "/usr/bin", nil
			},
			in: "",
			out: []Completion{
				{"ls", 0},
				{"ps", 0},
			},
		},

		{
			getwd: func() (string, error) {
				return "/usr/bin", nil
			},
			in: "",
			out: []Completion{
				{"ls", 0},
				{"ps", 0},
			},
		},

		// Errors handling
		{
			in:  "/path/not/exist",
			out: nil,
		},

		{
			in: "/usr/bin/ls",
			getwd: func() (string, error) {
				return "", errors.New("oops")
			},
			out: nil,
		},
	}

	for _, test := range tests {
		complete := CompleteFs(testFs, test.getwd)
		out := complete(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("%q:\nexpected: %#v\nreceived: %#v\n",
				test.in, test.out, out)
		}
	}
}

// TestCompleterFlags tests (CompletionrFlags) String()
func TestCompleterFlagsString(t *testing.T) {
	type testData struct {
		in  CompletionFlags // Input value
		out string          // Expected output
	}

	tests := []testData{
		{
			in:  0,
			out: "0",
		},

		{
			in:  1 << 0,
			out: "CompletionNoSpace",
		},

		{
			in:  1 << 2,
			out: "0x2",
		},

		{
			in:  1 << 4,
			out: "0x4",
		},

		{
			in:  12345,
			out: "CompletionNoSpace,0x3,0x4,0x5,0xc,0xd",
		},
	}

	for _, test := range tests {
		out := test.in.String()
		if out != test.out {
			t.Errorf("%x:\nexpected: %s\nreceived: %s",
				uint(test.in), test.out, out)
		}
	}
}
