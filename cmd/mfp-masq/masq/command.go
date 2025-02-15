// MFP - Miulti-Function Printers and scanners toolkit
// The "masq" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description.

package masq

import (
	"context"
	"errors"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/log"
)

// Command is the 'masq' command description
var Command = argv.Command{
	Name: "masq",
	Help: "IPP/eSCL/WSD masquerading proxy",
	Options: []argv.Option{
		argv.Option{
			Name: "--escl",
			Help: "Add eSCL proxy (--escl local-port=target-url)",
			Validate: func(s string) error {
				_, err := parseMapping(protoESCL, s)
				if err == nil {
					err = errors.New("not implemented")
				}
				return err
			},
		},
		argv.Option{
			Name: "--ipp",
			Help: "Add IPP proxy (--ipp local-port=target-url)",
			Validate: func(s string) error {
				_, err := parseMapping(protoIPP, s)
				return err
			},
		},
		argv.Option{
			Name: "--wsd",
			Help: "Add WSD proxy (--wsd local-port=target-url)",
			Validate: func(s string) error {
				_, err := parseMapping(protoWSD, s)
				if err == nil {
					err = errors.New("not implemented")
				}
				return err
			},
		},
		argv.Option{
			Name:     "-t",
			Help:     "write trace to file.log and file.tar",
			HelpArg:  "file",
			Validate: argv.ValidateAny,
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

	// Setup trace
	var trace *traceWriter
	if traceName, _ := inv.Get("-t"); traceName != "" {
		var err error
		trace, err = newTraceWriter(ctx, traceName)
		if err != nil {
			return err
		}

		defer trace.Close()
	}

	// Start proxies
	var mappings []mapping
	var proxies []*proxy

	for _, param := range inv.Values("--escl") {
		mappings = append(mappings, mustParseMapping(protoESCL, param))
	}

	for _, param := range inv.Values("--ipp") {
		mappings = append(mappings, mustParseMapping(protoIPP, param))
	}

	for _, param := range inv.Values("--wsd") {
		mappings = append(mappings, mustParseMapping(protoWSD, param))
	}

	for _, m := range mappings {
		p, err := newProxy(ctx, m, trace)
		if err != nil {
			log.Error(ctx, "%s: %s", m.param, err)
			return errors.New("Initialization failure")
		}

		proxies = append(proxies, p)
	}

	if len(proxies) == 0 {
		return errors.New("no proxies configured")
	}

	// Wait for termination signal
	<-ctx.Done()
	log.Info(ctx, "Exiting...")

	// Shutdown all proxies
	for _, p := range proxies {
		p.Shutdown()
	}

	return nil
}
