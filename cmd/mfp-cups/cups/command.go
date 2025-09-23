// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description.

package cups

import (
	"context"
	"fmt"

	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/cups"
	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/transport"
)

// Command is the 'cups' command description
var Command = argv.Command{
	Name: "cups",
	Help: "CUPS client",
	Options: []argv.Option{
		argv.Option{
			Name:    "-d",
			Aliases: []string{"--debug"},
			Help:    "Enable debug output",
		},
		argv.Option{
			Name:    "-v",
			Aliases: []string{"--verbose"},
			Help:    "Enable verbose debug output",
		},
		argv.Option{
			Name:    "-u",
			Aliases: []string{"--cups"},
			Help: "CUPS server address or URL\n" +
				fmt.Sprintf("default: %q", cups.DefaultUNIXURL),
			Validate: transport.ValidateAddr,
		},
		argv.HelpOption,
	},
	SubCommands: []argv.Command{
		cmdDefaultPrinter,
		cmdDetectPrinters,
		cmdGetPPD,
		cmdListPrinters,
		argv.HelpCommand,
	},
	Handler: cmdCupsHandler,
}

// cmdCupsHandler is the top-level handler for the 'cups' command.
func cmdCupsHandler(ctx context.Context, inv *argv.Invocation) error {
	// Setup logging
	_, dbg := inv.Get("-d")
	_, vrb := inv.Get("-v")

	level := log.LevelInfo
	if dbg {
		level = log.LevelDebug
	}
	if vrb {
		level = log.LevelTrace
	}

	logger := log.NewLogger(level, log.Console)
	ctx = log.NewContext(ctx, logger)

	// Execute subcommand
	return argv.DefaultHandler(ctx, inv)
}
