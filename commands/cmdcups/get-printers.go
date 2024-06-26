// MFP - Miulti-Function Printers and scanners toolkit
// The "cups" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The "get-printers" command.

package cmdcups

import (
	"context"
	"math"
	"sort"
	"strconv"

	"github.com/OpenPrinting/goipp"
	"github.com/alexpevzner/mfp/argv"
	"github.com/alexpevzner/mfp/cups"
	"github.com/alexpevzner/mfp/env"
)

// cmdGetPrinters defines the "get-printers" sub-command.
var cmdGetPrinters = argv.Command{
	Name:    "get-printers",
	Help:    "Get information on configured printers",
	Handler: cmdGetPrintersHandler,
	Options: []argv.Option{
		{
			Name:     "--attrs",
			Help:     "Additional attributes",
			Validate: argv.ValidateAny,
			Complete: optAttrsComplete,
		},
		{
			Name:     "--id",
			Help:     "Printer ID (1...65535)",
			Validate: argv.ValidateIntRange(0, 1, 65535),
		},
		{
			Name:     "--limit",
			Help:     "Maximum number of printers",
			Validate: argv.ValidateIntRange(0, 1, math.MaxInt32),
		},
		{
			Name: "--location",
			Help: "Printer location " +
				"(i.e., \"Printers at reception\")",
			Validate: argv.ValidateAny,
		},
		argv.HelpOption,
	},
}

// cmdGetPrintersHandler is the "get-printers" command handler
func cmdGetPrintersHandler(ctx context.Context, inv *argv.Invocation) error {
	// Prepare arguments
	dest := optCUPSURL(inv)

	sel := &cups.GetPrintersSelection{}

	if opt, ok := inv.Get("--id"); ok {
		sel.PrinterID, _ = strconv.Atoi(opt)
	}

	if opt, ok := inv.Get("--limit"); ok {
		sel.Limit, _ = strconv.Atoi(opt)
	}

	if opt, ok := inv.Get("--location"); ok {
		sel.PrinterLocation = opt
	}

	attrList := optAttrsGet(inv)
	attrList = append(attrList, "printer-name", "printer-uri-supported")

	// Perform the query
	pager := env.NewPager()

	pager.Printf("CUPS: %s", dest)

	clnt := cups.NewClient(dest, nil)
	printers, err := clnt.CUPSGetPrinters(ctx, sel, attrList)
	if err != nil {
		return err
	}

	for _, prn := range printers {
		pager.Printf("")
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
	}

	return pager.Display()
}
