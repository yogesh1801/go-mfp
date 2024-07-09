// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description.

package cmdcups

import (
	"fmt"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/transport"
)

// Command is the 'cups' command description
var Command = argv.Command{
	Name: "cups",
	Help: "CUPS client",
	Options: []argv.Option{
		argv.Option{
			Name:    "-u",
			Aliases: []string{"--cups"},
			Help: "CUPS server address or URL\n" +
				fmt.Sprintf("default: %q",
					transport.DefaultCupsUNIX),
			Validate: transport.ValidateAddr,
		},
		argv.HelpOption,
	},
	SubCommands: []argv.Command{
		cmdGetDefault,
		cmdGetDevices,
		cmdGetPrinters,
		argv.HelpCommand,
	},
}
