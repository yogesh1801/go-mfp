// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description.

package cmdcups

import "github.com/alexpevzner/mfp/argv"

// Command is the 'cups' command description
var Command = argv.Command{
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
