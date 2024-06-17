// MFP      - Miulti-Function Printers and scanners toolkit
// mainfunc - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Main function for the "mfc" command.

package commands

import (
	"github.com/alexpevzner/mfp/argv"
)

// CmdMfp is the argv.Command, that includes all other commands
// as sub-commands.
var CmdMfp = &argv.Command{
	Name: "mfp",
	Options: []argv.Option{
		argv.HelpOption,
	},
	SubCommands: []argv.Command{
		CmdCups,
		argv.HelpCommand,
	},
}
