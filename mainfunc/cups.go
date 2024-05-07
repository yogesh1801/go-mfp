// MFP      - Miulti-Function Printers and scanners toolkit
// mainfunc - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Main function for the "cups" command.

package mainfunc

import "github.com/alexpevzner/mfp/argv"

var cmdCups = argv.Command{
	Name: "cups",
	Help: "CUPS client",
	Main: MainCups,
}

// MainCups implements the 'main' function for the 'cups' command
func MainCups(argv []string) error {
	return nil
}
