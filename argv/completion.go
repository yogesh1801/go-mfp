// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Auto-Completion results

package argv

// Completion is the output of Command.Complete. It contains suggested
// completion string and [CompletionFlags]
type Completion struct {
	String  string // Suggested completion string
	NoSpace bool   // Don't append space after completion
}
