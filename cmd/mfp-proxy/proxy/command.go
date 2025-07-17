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

	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/internal/env"
	"github.com/OpenPrinting/go-mfp/log"
)

// description is printed as a command description text
const description = "" +
	"This command runs the IPP/eSCL/WSD proxy\n" +
	"The proxy can be useful for two purposes:\n" +
	"  - to sniff the protocol traffic between the IPP/eSCL/WSD\n" +
	"    client and server\n" +
	"  - to logically bring the device into the different IP address\n" +
	"    or port\n" +
	"\n" +
	"If optional command is specified, the CUPS_SERVER and the\n" +
	"SANE_AIRSCAN_DEVICE environment variables will be set properly\n" +
	"and the command will be executed, The simulator will exit when\n" +
	"the command finished.\n" +
	"\n" +
	"Without that the simulator will run until termination signal\n" +
	"is received.\n"

// Command is the 'proxy' command description
var Command = argv.Command{
	Name:                     "proxy",
	Help:                     "IPP/eSCL/WSD masquerading proxy",
	Description:              description,
	NoOptionsAfterParameters: true,
	Options: []argv.Option{
		argv.Option{
			Name:      "--escl",
			Help:      "Add eSCL proxy (--escl local-port=target-url)",
			Singleton: true,
			Validate: func(s string) error {
				_, err := parseMapping(protoESCL, s)
				if err == nil {
					err = errors.New("not implemented")
				}
				return err
			},
		},
		argv.Option{
			Name:      "--ipp",
			Help:      "Add IPP proxy (--ipp local-port=target-url)",
			Singleton: true,
			Validate: func(s string) error {
				_, err := parseMapping(protoIPP, s)
				return err
			},
		},
		argv.Option{
			Name:      "--wsd",
			Help:      "Add WSD proxy (--wsd local-port=target-url)",
			Singleton: true,
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
			Aliases:  []string{"--trace"},
			Help:     "write trace to file.log and file.tar",
			HelpArg:  "file",
			Validate: argv.ValidateAny,
			Complete: argv.CompleteOSPath,
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
	Parameters: []argv.Parameter{
		{
			Name: "[command]",
			Help: "command to run under the proxy",
		},
		{
			Name: "[args...]",
			Help: "the command's arguments",
		},
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
	var ippPort, esclPort int

	for _, param := range inv.Values("--escl") {
		m := mustParseMapping(protoESCL, param)
		esclPort = m.localPort
		mappings = append(mappings, m)
	}

	for _, param := range inv.Values("--ipp") {
		m := mustParseMapping(protoIPP, param)
		ippPort = m.localPort
		mappings = append(mappings, m)
	}

	for _, param := range inv.Values("--wsd") {
		m := mustParseMapping(protoWSD, param)
		mappings = append(mappings, m)
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

	// Shutdown all proxies at exit
	defer func() {
		for _, p := range proxies {
			p.Shutdown()
		}
	}()

	// Run external program if requested
	if command, ok := inv.Get("command"); ok {
		runner := env.Runner{
			CUPSPort: ippPort,
			ESCLPort: esclPort,
			ESCLName: "Virtual MFP Scanner",
		}

		argv := inv.Values("args")
		return runner.Run(ctx, command, argv...)
	}

	// Wait for termination signal
	<-ctx.Done()
	log.Info(ctx, "Exiting...")

	return nil
}
