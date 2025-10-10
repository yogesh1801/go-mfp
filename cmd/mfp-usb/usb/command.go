// MFP - Miulti-Function Printers and scanners toolkit
// The "usb" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description

package usb

import (
	"context"

	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/log"
)

// Command is the 'usb' command description
var Command = argv.Command{
	Name: "usb",
	Help: "Troubleshooting tool for the USB-connected MFPs",
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
		argv.HelpOption,
	},
	SubCommands: []argv.Command{
		cmdList,
		argv.HelpCommand,
	},
	Handler: cmdUsbHandler,
}

// cmdUsbHandler is the top-level handler for the 'usb' command.
func cmdUsbHandler(ctx context.Context, inv *argv.Invocation) error {
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
