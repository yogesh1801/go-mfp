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
	"errors"

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
			Validate: func(s string) error {
				_, err := parseMapping("--ipp", s)
				return err
			},
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

	// Start proxies
	var proxies []*proxy

	for _, param := range inv.Values("--ipp") {
		m, err := parseMapping("--ipp", param)
		if err != nil {
			// It MUST NOT happen, because parameters already
			// validated at this point.
			panic(err)
		}

		p, err := newProxy(ctx, m)
		if err != nil {
			log.Error(ctx, "%s: %s", param, err)
			return errors.New("Initialization failure")
		}

		proxies = append(proxies, p)
	}

	<-ctx.Done()
	log.Info(ctx, "Exiting...")

	return nil
}
