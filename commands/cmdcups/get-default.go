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
	"fmt"

	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/cups"
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

	fmt.Printf("CUPS: %s\n", dest)

	clnt := cups.NewClient(dest, nil)
	prn, err := clnt.CUPSGetDefault(ctx, attrs)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", prn.PrinterName)
	fmt.Printf("URL:  %s\n", prn.PrinterURISupported)

	fmt.Printf("Printer attributes:\n")

	for _, attr := range prn.Attrs().All() {
		fmt.Printf("  %s: %s\n", attr.Name, attr.Values)
	}

	return nil
}
