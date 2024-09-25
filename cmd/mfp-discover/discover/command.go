// MFP - Miulti-Function Printers and scanners toolkit
// The "discover" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description.

package discover

import (
	"context"
	"encoding/json"
	"os"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/discovery"
	"github.com/alexpevzner/mfp/discovery/dnssd"
	"github.com/alexpevzner/mfp/log"
)

// Command is the 'cups' command description
var Command = argv.Command{
	Name: "discover",
	Help: "search for printers and scanners",
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
			Name:    "-p",
			Aliases: []string{"--printers"},
			Help:    "Search for printers",
		},
		argv.Option{
			Name:    "-s",
			Aliases: []string{"--scanners"},
			Help:    "Search for scanners",
		},
		argv.HelpOption,
	},
	Handler: cmdDiscoverHandler,
}

// cmdCupsHandler is the handler for the 'discover' command.
func cmdDiscoverHandler(ctx context.Context, inv *argv.Invocation) error {
	// Setup logging
	_, dbg := inv.Get("-d")
	_, vrb := inv.Get("-v")

	dbg = true // FIXME

	level := log.LevelInfo
	if dbg {
		level = log.LevelDebug
	}
	if vrb {
		level = log.LevelTrace
	}

	logger := log.NewLogger(level, log.Console)
	ctx = log.NewContext(ctx, logger)

	// Execute the command
	clnt := discovery.NewClient(ctx)
	backend, err := dnssd.NewBackend(ctx, "", 0)
	if err != nil {
		return err
	}

	clnt.AddBackend(backend)
	devices, err := clnt.GetDevices(ctx, discovery.ModeNormal)
	if err != nil {
		return err
	}

	//fmt.Printf(">> %#v\n", devices)
	out, _ := json.MarshalIndent(devices, "", " ")
	os.Stdout.Write(out)

	backend.Close()

	return nil
}
