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
	"strings"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/discovery"
	"github.com/alexpevzner/mfp/discovery/dnssd"
	"github.com/alexpevzner/mfp/env"
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

	// Perform device discovery
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

	// Format output
	pager := env.NewPager()
	defer pager.Display()

	for _, dev := range devices {
		pager.Printf("================================")

		pager.Printf("  MakeModel:        %q", dev.MakeModel)
		pager.Printf("  DNS-SD name:      %q", dev.DNSSDName)
		pager.Printf("  DNS-SD UUID:      %q", dev.DNSSDUUID)
		pager.Printf("  PPD Manufacturer: %q", dev.PPDManufacturer)
		pager.Printf("  PPD Model:        %q", dev.PPDModel)
		pager.Printf("  USB serial:       %q", dev.USBSerial)

		s := []string{}
		for _, addr := range dev.Addrs {
			s = append(s, addr.String())
		}
		pager.Printf("  IP addresses: %s", strings.Join(s, ", "))

		if len(dev.PrintUnits) != 0 {
			pager.Printf("")
			pager.Printf("  Print units:")
			for _, un := range dev.PrintUnits {
				p := un.Params

				pager.Printf("    Type:       %s printer",
					un.Proto)
				pager.Printf("    Auth:       %s", p.Auth)
				pager.Printf("    Admin URL:  %q", p.AdminURL)
				pager.Printf("    Location:   %q", p.Location)
				pager.Printf("    Paper Size: %s", p.Paper)
				pager.Printf("    Media Type: %s", p.Media)

				pager.Printf("    Flags:      %s", p.Flags())

				pager.Printf("    PSProduct:  %q", p.PSProduct)
				pager.Printf("    PDL:        %s",
					strings.Join(p.PDL, ","))
				pager.Printf("    Priority:   %d", p.Priority)
				pager.Printf("")
			}
		}
	}

	backend.Close()

	return nil
}
