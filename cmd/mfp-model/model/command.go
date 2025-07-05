// MFP - Miulti-Function Printers and scanners toolkit
// The "model" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description.

package model

import (
	"context"
	"errors"

	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/transport"
)

// DefaultTCPPort is the default TCP port for the MFP simulator
const DefaultTCPPort = 50000

// description is printed as a command description text
const description = "" +
	"This command generates models for the MFP simulator\n" +
	""

// Command is the 'model' command description
var Command = argv.Command{
	Name:                     "model",
	Help:                     "Model generator for MFP simulator",
	Description:              description,
	NoOptionsAfterParameters: true,
	Options: []argv.Option{
		argv.Option{
			Name:     "-D",
			Aliases:  []string{"--dnssd"},
			Help:     "DNS-SD name of the device",
			HelpArg:  "name",
			Validate: argv.ValidateAny,
		},
		argv.Option{
			Name:     "-E",
			Aliases:  []string{"--escl"},
			Help:     "eSCL scanner URL",
			HelpArg:  "URL",
			Validate: transport.ValidateURL,
		},
		argv.Option{
			Name:     "-I",
			Aliases:  []string{"--ipp"},
			Help:     "IPP printer URL",
			HelpArg:  "URL",
			Validate: transport.ValidateURL,
		},
		argv.Option{
			Name:     "-W",
			Aliases:  []string{"--wsd"},
			Help:     "WSD scanner URL",
			HelpArg:  "URL",
			Validate: transport.ValidateURL,
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
	Handler: cmdModelHandler,
}

// cmdModelHandler is the top-level handler for the 'model' command.
func cmdModelHandler(ctx context.Context, inv *argv.Invocation) error {
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

	_ = ctx

	return errors.New("not implemented")
}
