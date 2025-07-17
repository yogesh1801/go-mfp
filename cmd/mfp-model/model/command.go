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
	"fmt"

	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/discovery"
	"github.com/OpenPrinting/go-mfp/discovery/dnssd"
	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/go-mfp/modeling"
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
			Name:      "-D",
			Aliases:   []string{"--dnssd"},
			Help:      "DNS-SD name of the device",
			HelpArg:   "name",
			Conflicts: []string{"-E", "-I", "-W"},
			Singleton: true,
			Validate:  argv.ValidateAny,
			Complete:  dnssd.ArgvCompleter,
		},
		argv.Option{
			Name:      "-E",
			Aliases:   []string{"--escl"},
			Help:      "eSCL scanner URL",
			HelpArg:   "URL",
			Singleton: true,
			Validate:  transport.ValidateAddr,
		},
		argv.Option{
			Name:      "-I",
			Aliases:   []string{"--ipp"},
			Help:      "IPP printer URL",
			HelpArg:   "URL",
			Singleton: true,
			Validate: func(s string) error {
				err := transport.ValidateAddr(s)
				if err == nil {
					err = errors.New("not implemented")
				}
				return err
			},
		},
		argv.Option{
			Name:      "-W",
			Aliases:   []string{"--wsd"},
			Help:      "WSD scanner URL",
			HelpArg:   "URL",
			Singleton: true,
			Validate: func(s string) error {
				err := transport.ValidateAddr(s)
				if err == nil {
					err = errors.New("not implemented")
				}
				return err
			},
		},
		argv.Option{
			Name:      "-m",
			Aliases:   []string{"--model"},
			Help:      "write model to file",
			HelpArg:   "file",
			Required:  true,
			Singleton: true,
			Validate:  argv.ValidateAny,
			Complete:  argv.CompleteOSPath,
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

	// Check options
	optDHSSD, haveDNSSD := inv.Get("--dnssd")
	escl := inv.Values("--escl")
	ipp := inv.Values("--ipp")
	wsd := inv.Values("--wsd")

	if !haveDNSSD && ipp == nil && escl == nil && wsd == nil {
		err := errors.New("at least one option required: --dnssd, --escl, --ipp or --wsd")
		return err
	}

	// Gather endpoints
	if haveDNSSD {
		dev, err := discoverByName(ctx, optDHSSD)
		if err != nil {
			return err
		}

		for _, unit := range dev.PrintUnits {
			if unit.Proto == discovery.ServiceIPP {
				ipp = append(ipp, unit.Endpoints...)
			}
		}

		for _, unit := range dev.ScanUnits {
			switch unit.Proto {
			case discovery.ServiceESCL:
				escl = append(escl, unit.Endpoints...)
			case discovery.ServiceWSD:
				wsd = append(wsd, unit.Endpoints...)
			}
		}

		// Check that something was discovered.
		if ipp == nil && escl == nil && wsd == nil {
			err := errors.New("no eSCL/IPP/WSD endpoints discovered")
			return err
		}
	}

	// Create a model
	model, err := modeling.NewModel()
	if err != nil {
		return err
	}

	// Query printer and scanner capabilities
	esclcaps, err := queryESCLScannerCapabilities(ctx, escl)
	if err != nil {
		err = fmt.Errorf("Can't get eSCL ScannerCapabilities: %s", err)
		return err
	}

	// Save model to file
	file, _ := inv.Get("-m")

	model.SetESCLScanCaps(esclcaps)
	return model.Save(file)
}
