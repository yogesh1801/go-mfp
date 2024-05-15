// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Value completers.

package argv

import (
	"io/fs"
	"strings"
)

// CompleteStrings returns a completer, that performs auto-completion,
// choosing from a set of supplied strings.
func CompleteStrings(s []string) func(string) []string {
	// Create a copy of input, to protect from callers
	// that may change the slice after the call.
	set := make([]string, len(s))
	copy(set, s)

	// Create completer
	return func(in string) []string {
		out := []string{}
		for _, member := range set {
			if len(in) < len(member) &&
				strings.HasPrefix(member, in) {
				out = append(out, member)
			}
		}
		return out
	}
}

// CompleteFs returns a completer, that performs file name auto-completion
// on a top of a virtual (or real) filesystem, represented as fs.FS,
//
// filesystem must implement fs.ReadDirFS or fs.GlobFS interfaces.
//
// getwd callback returns a current directory within that file system.
// It's signature is compatible with os.Getwd(), so this function can
// be used directly.
//
// If getwd is nil, current directory assumed to be "/"
func CompleteFs(fsys fs.FS,
	getwd func() (string, error)) func(string) []string {

	fscompl := newFscompleter(fsys, getwd)
	return func(arg string) []string {
		return fscompl.complete(arg)
	}
}
