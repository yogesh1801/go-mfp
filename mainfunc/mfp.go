// MFP      - Miulti-Function Printers and scanners toolkit
// mainfunc - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Main function for the "mfc" command.

package mainfunc

import (
	"fmt"

	"github.com/alexpevzner/mfp/argv"
)

// AllCommands is the argv.Command, that includes all other commands
// as sub-commands.
var cmdMfp = argv.Command{
	Name: "mfp",
	SubCommands: []argv.Command{
		cmdCups,
	},
	Main: MainMfp,
}

// MainMfp implements the 'main' function for the 'mfp' command
func MainMfp(argv []string) error {
	fmt.Printf("%s\n", argv)
	return nil
}
