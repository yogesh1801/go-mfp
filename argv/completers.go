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
	"os"
	"strings"
)

// Completer is a callback called for auto-completion
//
// Any [Option] or [Parameter] may have its own Completer.
//
// It receives the Option's value prefix, already typed
// by user, and must return a slice of completion candidates
// that match the prefix.
//
// For example, if possible Option or Parameter values are "Richard",
// "Roger" and  "Robert", then, depending of supplied prefix, the following
// output is expected:
//
//	"R"   -> ["Richard", "Roger", "Robert"]
//	"Ro"  -> ["Roger", "Robert"]
//	"Rog" -> ["Roger"]
//	"Rol" -> []
type Completer func(string) []Completion

// CompleteStrings returns a [Completer], that performs auto-completion,
// choosing from a set of supplied strings.
func CompleteStrings(s []string) Completer {
	// Create a copy of input, to protect from callers
	// that may change the slice after the call.
	set := make([]string, len(s))
	copy(set, s)

	// Create completer
	return func(in string) (compl []Completion) {
		for _, member := range set {
			if len(in) < len(member) &&
				strings.HasPrefix(member, in) {
				compl = append(compl, Completion{member, 0})
			}
		}
		return compl
	}
}

// CompletePath is the [Completer] that completes file system paths.
func CompletePath(s string) []Completion {
	return completePath(s)
}

var completePath = CompleteFs(os.DirFS("/"), os.Getwd)

// CompleteFs returns a [Completer], that performs file name auto-completion
// on a top of a virtual (or real) filesystem, represented as fs.FS,
//
// getwd callback returns a current directory within that file system.
// It's signature is compatible with os.Getwd(), so this function can
// be used directly.
//
// If getwd is nil, current directory assumed to be "/"
func CompleteFs(fsys fs.FS, getwd func() (string, error)) Completer {
	fscompl := newFscompleter(fsys, getwd)
	return func(arg string) []Completion {
		return fscompl.complete(arg)
	}
}
