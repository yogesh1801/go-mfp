// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Auto-Completion results

package argv

import (
	"fmt"
	"math/bits"
	"strings"
)

// Completion is the output of Command.Complete. It contains suggested
// completion string and [CompletionFlags]
type Completion struct {
	String string          // Suggested completion string
	Flags  CompletionFlags // Completion flags
}

// CompletionFlags returned by [Command.Complete] and by
// [Completer] callback from [Option] a [Parameter] to provide
// additional information hints how caller should interpret returned
// completion candidates.
//
// See each flag's documentation for more details.
type CompletionFlags int

const (
	// CompletionNoSpace indicates that caller should not
	// append a space after completion.
	//
	// This is useful, for example, for file path completion
	// that ends with path separator character (i.e., '/')
	// and user is prompted to continue entering a full
	// file name.
	CompletionNoSpace CompletionFlags = 1 << iota
)

// completionFlagsNames contains names of CompletionFlags,
// for debugging
var completionFlagsNames = map[CompletionFlags]string{}

func init() {
	// Pupulate completionFlagsNames table
	for i := 0; i < bits.UintSize; i++ {
		f := CompletionFlags(1 << i)
		switch f {
		case CompletionNoSpace:
			completionFlagsNames[f] = "CompletionNoSpace"
		default:
			completionFlagsNames[f] = fmt.Sprintf("0x%x", i)
		}
	}
}

// String converts CompletionFlags into string, for debugging.
func (flags CompletionFlags) String() string {
	if flags == 0 {
		return "0"
	}

	var s []string

	for i := CompletionFlags(1); flags != 0; i <<= 1 {
		if flags&i != 0 {
			s = append(s, completionFlagsNames[i])
		}
		flags &^= i
	}

	return strings.Join(s, ",")
}
