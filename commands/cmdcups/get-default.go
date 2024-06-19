// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Command description.

package cmdcups

import (
	"context"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/cups"
	"github.com/alexpevzner/mfp/env"
	"github.com/alexpevzner/mfp/transport"
)

// cmdGetDefault defines "get-default" command
var cmdGetDefault = argv.Command{
	Name:    "get-default",
	Help:    "Get default printer",
	Handler: cmdGetDefaultHandler,
	Options: []argv.Option{
		argv.HelpOption,
	},
	Parameters: []argv.Parameter{
		{
			Name: "[attrs...]",
			Help: "Requested attributes",
		},
	},
}

// cmdGetDefaultHandler is the "get-default" command handler
func cmdGetDefaultHandler(ctx context.Context, inv *argv.Invocation) error {
	dest := transport.DefaultCupsUNIX

	if addr, ok := inv.Parent().Get("-u"); ok {
		dest = transport.MustParseAddr(addr, "ipp://localhost/")
	}

	attrs := inv.Values("attrs")
	attrs = append(attrs, "printer-name", "printer-uri-supported")

	pager := env.NewPager()

	pager.Printf("CUPS: %s", dest)

	clnt := cups.NewClient(dest, nil)
	prn, err := clnt.CUPSGetDefault(ctx, attrs)
	if err != nil {
		return err
	}

	pager.Printf("Name: %s", prn.PrinterName)
	pager.Printf("URL:  %s", prn.PrinterURISupported)

	pager.Printf("Printer attributes:")

	for _, attr := range prn.Attrs().All() {
		pager.Printf("  %s: %s", attr.Name, attr.Values)
	}

	return pager.Display()
}
