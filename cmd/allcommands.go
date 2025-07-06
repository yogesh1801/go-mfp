// MFP - Miulti-Function Printers and scanners toolkit
// Common functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// AllCommands super-command.

package cmd

import (
	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/cmd/mfp-cups/cups"
	"github.com/OpenPrinting/go-mfp/cmd/mfp-discover/discover"
	"github.com/OpenPrinting/go-mfp/cmd/mfp-proxy/proxy"
)

// AllCommands is the argv.Command, that includes all other commands
// as sub-commands.
var AllCommands = &argv.Command{
	Name: "mfp",
	Options: []argv.Option{
		argv.HelpOption,
	},
	SubCommands: []argv.Command{
		cups.Command,
		proxy.Command,
		discover.Command,
		argv.HelpCommand,
	},
}
