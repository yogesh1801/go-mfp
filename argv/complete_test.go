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
		strings []string       // Strings to choose from
		arg     string         // Input string
		out     []string       // Expected output
		flags   CompleterFlags // Expected flags
	}

	tests := []testData{
		{
			strings: []string{"foo", "bar"},
			arg:     "foo",
			out:     []string{},
		},

		{
			strings: []string{"foo-1", "foo-2"},
			arg:     "foo",
			out:     []string{"foo-1", "foo-2"},
		},

		{
			strings: []string{"foo-1", "foo-2"},
			arg:     "",
			out:     []string{"foo-1", "foo-2"},
		},
	}

	for _, test := range tests {
		completer := CompleteStrings(test.strings)
		out, flags := completer(test.arg)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("CompleteStrings(%#v): %q\nexpected: %#v\nreceived: %#v\n",
				test.strings, test.arg, test.out, out)
		}

		if flags != test.flags {
			t.Errorf("CompleteStrings(%#v) %q:"+
				"\nflags expected: %s\nflags received: %s\n",
				test.strings, test.arg, test.flags, flags)
		}
	}
}

// TestCompleteFs tests CompleteFs
func TestCompleteFs(t *testing.T) {
	var testFs fs.FS = fstest.MapFS{
		"etc/hosts":  &fstest.MapFile{Mode: 0755},
		"usr/bin/ls": &fstest.MapFile{Mode: 0755},
		"usr/bin/ps": &fstest.MapFile{Mode: 0755},
	}

	type testData struct {
		getwd func() (string, error) // getwd function
		in    string                 // Input string
		out   []string               // Expected output
		flags CompleterFlags         // Expected flags
	}

	tests := []testData{
		// Simple tests
		{
			in:  "/",
			out: []string{"/etc", "/usr"},
		},

		{
			in:  "./",
			out: []string{"./etc", "./usr"},
		},

		{
			in:  "/usr/bin/ls",
			out: []string{"/usr/bin/ls"},
		},

		{
			in:    "/usr/bin",
			out:   []string{"/usr/bin/"},
			flags: CompleterNoSpace,
		},

		// Tests with getwd()
		{
			getwd: func() (string, error) {
				return "/usr/bin", nil
			},
			in:  "",
			out: []string{"ls", "ps"},
		},

		{
			getwd: func() (string, error) {
				return "/usr/bin", nil
			},
			in:  "",
			out: []string{"ls", "ps"},
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
		out, flags := complete(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("%q:\nexpected: %#v\nreceived: %#v\n",
				test.in, test.out, out)
			//os.Exit(1)
		}

		if flags != test.flags {
			t.Errorf("%q flags:\nexpected: %s\nreceived: %s\n",
				test.in, test.flags, flags)
		}
	}
}

// TestCompleterFlags tests (CompleterFlags) String()
func TestCompleterFlagsString(t *testing.T) {
	type testData struct {
		in  CompleterFlags // Input value
		out string         // Expected output
	}

	tests := []testData{
		{
			in:  0,
			out: "0",
		},

		{
			in:  1 << 0,
			out: "CompleterNoSpace",
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
			out: "CompleterNoSpace,0x3,0x4,0x5,0xc,0xd",
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
