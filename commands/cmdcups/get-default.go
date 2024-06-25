// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The "get-default" command.

package cmdcups

import (
	"context"
	"sort"

	"github.com/OpenPrinting/goipp"
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
		{
			Name:     "-a",
			Aliases:  []string{"--attrs"},
			Help:     "Additional attributes",
			Validate: argv.ValidateAny,
			Complete: optAttrsComplete,
		},
		argv.HelpOption,
	},
}

// cmdGetDefaultHandler is the "get-default" command handler
func cmdGetDefaultHandler(ctx context.Context, inv *argv.Invocation) error {
	dest := transport.DefaultCupsUNIX

	if addr, ok := inv.Parent().Get("-u"); ok {
		dest = transport.MustParseAddr(addr, "ipp://localhost/")
	}

	attrList := optAttrsGet(inv)
	attrList = append(attrList, "printer-name", "printer-uri-supported")

	pager := env.NewPager()

	pager.Printf("CUPS: %s", dest)

	clnt := cups.NewClient(dest, nil)
	prn, err := clnt.CUPSGetDefault(ctx, attrList)
	if err != nil {
		return err
	}

	pager.Printf("Name: %s", prn.PrinterName)
	pager.Printf("URL:  %s", prn.PrinterURISupported)

	pager.Printf("Printer attributes:")

	attrs := prn.Attrs().All().Clone()
	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Name < attrs[j].Name
	})

	f := goipp.NewFormatter()
	f.SetIndent(2)
	f.FmtAttributes(attrs)

	f.WriteTo(pager)

	return pager.Display()
}
