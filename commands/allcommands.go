// MFP - Miulti-Function Printers and scanners toolkit
// Common functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// AllCommands super-command.

package commands

import (
	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/commands/cmdcups"
)

// AllCommands is the argv.Command, that includes all other commands
// as sub-commands.
var AllCommands = &argv.Command{
	Name: "mfp",
	Options: []argv.Option{
		argv.HelpOption,
	},
	SubCommands: []argv.Command{
		cmdcups.Command,
		argv.HelpCommand,
	},
}
