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
				compl := member[len(in):]
				out = append(out, compl)
			}
		}
		return out
	}
}

// CompleteFs returns a completer, that performs file name auto-completion
// on a top of fs.GlobFS.
//
// filesystem must implement fs.ReadDirFS or fs.GlobFS interfaces.
func CompleteFs(filesystem fs.FS) func(string) []string {
	return func(in string) []string {
		// Don't allow special characters within the string
		if strings.IndexAny(in, "/?*") >= 0 {
			return nil
		}

		// Lookup matching file names
		names, err := fs.Glob(filesystem, in+"*")
		if err != nil {
			return nil
		}

		out := []string{}
		for _, name := range names {
			if len(in) < len(name) && strings.HasPrefix(name, in) {
				compl := name[:len(in)]
				out = append(out, compl)
			}
		}
		return out
	}
}
