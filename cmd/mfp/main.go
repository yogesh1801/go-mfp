// MFP     - Miulti-Function Printers and scanners toolkit
// cmd/mfp - Universal command that implements all other as sub-commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The main() function.

package main

import (
	"github.com/alexpevzner/mfp/commands"
)

// main function for the mfp command
func main() {
	commands.AllCommands.Main(nil)
}
