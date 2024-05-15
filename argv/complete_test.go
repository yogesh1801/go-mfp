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
		strings []string // Strings to choose from
		arg     string   // Input string
		out     []string // Expected output
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
		out := completer(test.arg)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("CompleteStrings(%#v): %q\nexpected: %#v\nreceived: %#v\n",
				test.strings, test.arg, test.out, out)
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
	}

	tests := []testData{
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
