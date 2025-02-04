// MFP - Miulti-Function Printers and scanners toolkit
// The "proxy" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description.

package proxy

import (
	"context"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/log"
)

// Command is the 'proxy' command description
var Command = argv.Command{
	Name: "proxy",
	Help: "IPP/eSCL/WSD proxy",
	Options: []argv.Option{
		argv.Option{
			Name:    "--ipp",
			HelpArg: "local-port=target-url",
			Help: "IPP proxy maps IPP rqquests " +
				"from local-port to target-url",
			Validate: func(s string) error { return nil },
		},
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
	Handler: cmdProxyHandler,
}

// cmdProxyHandler is the top-level handler for the 'proxy' command.
func cmdProxyHandler(ctx context.Context, inv *argv.Invocation) error {
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

	return nil
}
