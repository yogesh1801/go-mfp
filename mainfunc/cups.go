// MFP      - Miulti-Function Printers and scanners toolkit
// mainfunc - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Main function for the "cups" command.

package mainfunc

import "github.com/alexpevzner/mfp/argv"

// CmdCups is the 'cups' command description
var CmdCups = argv.Command{
	Name: "cups",
	Help: "CUPS client",
	Options: []argv.Option{
		argv.HelpOption,
	},
	Handler: cupsHandler,
}

// MainCups implements the Handler callback of the 'cups' command
func cupsHandler(*argv.Invocation) error {
	return nil
}
